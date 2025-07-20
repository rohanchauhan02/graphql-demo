package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/rohanchauhan02/graphql-demo/internal/user/delivery/graph"
	rest "github.com/rohanchauhan02/graphql-demo/internal/user/delivery/https"
	grpcserver "github.com/rohanchauhan02/graphql-demo/internal/user/delivery/rpc"
	"github.com/rohanchauhan02/graphql-demo/internal/user/repository"
	"github.com/rohanchauhan02/graphql-demo/internal/user/usecase"
	cm "github.com/rohanchauhan02/graphql-demo/middleware"
	"github.com/rohanchauhan02/graphql-demo/models"
	userpb "github.com/rohanchauhan02/graphql-demo/proto/userpb"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found; using system environment variables")
	}

	// Configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "host=localhost user=root password=root dbname=internal_transfer_local port=5432 sslmode=disable TimeZone=UTC"
	}

	// Initialize Echo
	e := echo.New()
	e.Use(cm.RequestIDMiddleware())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
	})

	// Connect to DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Auto-migrate models
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("❌ Auto migration failed: %v", err)
	}

	// Initialize application layers
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// GraphQL setup
	gqlHandler := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.NewResolver(userUsecase),
	}))
	e.POST("/query", echo.WrapHandler(gqlHandler))
	e.GET("/", echo.WrapHandler(playground.Handler("GraphQL Playground", "/query")))

	// REST API
	api := e.Group("/api/v1")
	rest.NewUserHandler(api, userUsecase)

	// Start gRPC server
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Failed to listen on gRPC port: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, grpcserver.NewUserGRPCServer(userUsecase))

	go func() {
		log.Println("🚀 gRPC server listening on :50051")
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatalf("❌ gRPC server failed: %v", err)
		}
	}()

	// Start HTTP server in a goroutine
	go func() {
		log.Printf("🚀 HTTP server started at http://localhost:%s", port)
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ HTTP server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("🛑 Shutdown signal received, cleaning up...")

	// Stop gRPC server
	grpcServer.GracefulStop()

	// Stop HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("❌ HTTP server forced to shutdown: %v", err)
	}

	log.Println("✅ Server shut down gracefully")
}
