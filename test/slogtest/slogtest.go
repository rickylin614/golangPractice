package slogtest

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
)

var zlog *zap.Logger

var zlogsugar *zap.SugaredLogger

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	zlog, _ = zap.NewProduction()
	zlogsugar = zlog.Sugar()
}

func SlogPrint() {
	for i := 0; i < 10; i++ {
		slog.Info("test %s", slog.Int("test", i))
	}
}

func ZlogPrint() {
	for i := 0; i < 10; i++ {
		zlog.Info("test", zap.Int("test", i))
	}
}

func ZlogSugarPrint() {
	for i := 0; i < 10; i++ {
		zlogsugar.Info("test", i)
	}
}
