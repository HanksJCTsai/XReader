package logger

import (
	"github.com/HanksJCTsai/XReader/internal/config"
	"log/slog"
	"os"
)

var Log *slog.Logger

func InitLogger() {
	var level slog.Level
	switch config.Cfg.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	Log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	slog.SetDefault(Log)
}
