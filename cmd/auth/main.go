package main

import (
	"fmt"

	"github.com/SergeyGolang/grpc-auth/internal/config"
)

func main() {
	// TODO: init config
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO: init logger

	// TODO: init app

	// TODO: start gRPC-server app
}
