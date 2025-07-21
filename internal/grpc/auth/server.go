package auth

import (
	"context"

	ssov1 "github.com/SergeyGolang/protos/gen/go/auth"
	"google.golang.org/grpc"
)

// Заглушка для как будто имплементации методов
type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

// Регистрируем обработчик
func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	panic("implement me!")
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	panic("implement me!")
}
func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("implement me!")
}
