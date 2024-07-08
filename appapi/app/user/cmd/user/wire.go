//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"{{ cookiecutter.app_name }}/app/user/internal/biz"
	"{{ cookiecutter.app_name }}/app/user/internal/conf"
	"{{ cookiecutter.app_name }}/app/user/internal/data"
	"{{ cookiecutter.app_name }}/app/user/internal/server"
	"{{ cookiecutter.app_name }}/app/user/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
