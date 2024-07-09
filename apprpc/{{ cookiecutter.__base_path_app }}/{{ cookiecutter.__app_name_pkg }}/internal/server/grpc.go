package server

import (
	{{ cookiecutter.__app_name_pkg }}_api "{{ cookiecutter.__project_name_snake }}/api/{{ cookiecutter.__app_name_pkg }}"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos-layout/internal/service"

	

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, {{ cookiecutter.__app_name_pkg }} *service.{{ cookiecutter.__app_name_pkg }}Service, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	{{ cookiecutter.__app_name_pkg }}_api.Register{{ cookiecutter.__app_name_pkg }}Server(srv, {{ cookiecutter.__app_name_pkg }})
	return srv
}
