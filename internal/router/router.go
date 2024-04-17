package router

import (
	"github.com/alexlzrv/gokeeper/internal/handlers"
	mw "github.com/alexlzrv/gokeeper/internal/middleware/auth"
	"github.com/gin-gonic/gin"
)

func Router(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	auth := r.Group("/")
	{
		auth.Use(mw.Auth())
		auth.POST("/api/secret/create", h.CreateSecret)
		auth.POST("/api/secret/read", h.ReadSecret)
	}
	r.POST("/api/user/register", h.RegisterUser)
	r.POST("/api/user/login", h.LoginUser)

	return r
}
