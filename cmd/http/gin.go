package http

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-ddd/internal/infra/cfg"
	"go-ddd/internal/infra/http/gin"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
)

func RunGinServer(cfg cfg.Config, store *sqlx.DB, log *zap.SugaredLogger) {
	server := gin.NewServer(cfg, store, log)

	_ = server.Start(cfg.HTTPServerAddress)
}
