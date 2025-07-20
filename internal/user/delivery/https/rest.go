package rest

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rohanchauhan02/graphql-demo/internal/user"
)

type UserHandler struct {
	usecase user.Usecase
}

func NewUserHandler(group *echo.Group, uc user.Usecase) {
	h := &UserHandler{usecase: uc}

	group.GET("/users", h.ListUsers)
}

func (h *UserHandler) ListUsers(c echo.Context) error {
	users, err := h.usecase.GetUser(context.TODO(), 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}
