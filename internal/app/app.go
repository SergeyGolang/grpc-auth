package app

import (
	"log/slog"
	"time"

	grpcApp "github.com/SergeyGolang/grpc-auth/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcApp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: init storage

	// TODO: init auth service

	grpcApp := grpcApp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}

}
