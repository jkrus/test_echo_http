package handler

import (
	"github.com/jkrus/test_echo_http/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()
	user := router.Group("/user")
	{
		user.POST("/create", h.createUser)
		user.GET("/all", h.getAllUsers)
		user.GET("/:id", h.getById)
		user.PUT("/:id", h.updateUser)
		user.DELETE("/:id", h.deleteUser)
	}
	return router
}
