package external

import (
	"context"
	"time"

	"github.com/alexlast/bunzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"datink/config"
	"datink/pkg/database"
)

func NewDatabase(cfg *config.Config, logger *zap.Logger) *database.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	db := database.NewBunDB(ctx, &database.Config{
		Database: cfg.DBName,
		User:     cfg.DBUsername,
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		SSLMode:  cfg.DBSSLMode,
		Password: cfg.DBPassword,
	})

	logOption := bunzap.QueryHookOptions{
		Logger:       logger,
		SlowDuration: 200 * time.Millisecond,
	}

	//show all query of database
	lvl, _ := zapcore.ParseLevel(cfg.AppLogLevel)
	if lvl == zapcore.DebugLevel {
		logOption.SlowDuration = 0
	}

	db.AddQueryHook(bunzap.NewQueryHook(logOption))

	return db
}
