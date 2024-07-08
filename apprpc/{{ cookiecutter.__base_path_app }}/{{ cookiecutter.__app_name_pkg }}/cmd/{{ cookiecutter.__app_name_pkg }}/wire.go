//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"{{ cookiecutter.__app_name_camel }}/app/user/internal/biz"
	"{{ cookiecutter.__app_name_camel }}/app/user/internal/conf"
	"{{ cookiecutter.__app_name_camel }}/app/user/internal/data"
	"{{ cookiecutter.__app_name_camel }}/app/user/internal/server"
	"{{ cookiecutter.__app_name_camel }}/app/user/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
