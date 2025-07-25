package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/SergeyGolang/grpc-auth/internal/app"
	"github.com/SergeyGolang/grpc-auth/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config
	cfg := config.MustLoad()

	// TODO: init logger
	log := setupLogger(cfg.Env)

	log.Info("starting applications!")

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go application.GRPCSrv.MustRun()

	// TODO: init app

	// TODO: start gRPC-server app

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// Здесь мы ждём пока в нашу горутину не придёт сигнал и если сигнал придёт, то мы делаем Graceful shutdown
	sign := <-stop

	log.Info("stopping application", slog.String("signal", sign.String()))

	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}),
		)
	}
	return log
}
