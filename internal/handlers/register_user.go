package handlers

import (
	"github.com/alexlzrv/gokeeper/internal/logger"
	"github.com/alexlzrv/gokeeper/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterUser - register new user
func (h *Handler) RegisterUser(ctx *gin.Context) {
	var auth *models.Auth
	if err := ctx.ShouldBindJSON(&auth); err != nil {
		logger.Log.Info("Dont read json")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	err := h.AuthService.CreateUser(auth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "the new user has been successfully registered!"})
}
