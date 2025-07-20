package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id := uuid.New().String()
			c.Request().Header.Set("X-Request-ID", id)
			c.Response().Header().Set("X-Request-ID", id)
			c.Set("X-Request-ID", id)
			return next(c)
		}
	}
}
