package handlers

import (
	"context"
	"github.com/alexlzrv/gokeeper/cmd/server/config"
	"github.com/alexlzrv/gokeeper/internal/services/auth"
	"github.com/alexlzrv/gokeeper/internal/services/keeper"
	"github.com/alexlzrv/gokeeper/internal/storage"
)

type Handler struct {
	Config        *config.Config
	Storage       *storage.Storage
	AuthService   *auth.Auth
	KeeperService *keeper.Keeper
	Ctx           context.Context
}

func NewHandler(config *config.Config, ctx context.Context) *Handler {
	store := storage.NewStorage(config.DatabaseURI)

	return &Handler{
		Config:        config,
		Storage:       store,
		AuthService:   auth.NewAuth(store, config.TokenTTL),
		KeeperService: keeper.NewKeeper(store),
		Ctx:           ctx,
	}
}
