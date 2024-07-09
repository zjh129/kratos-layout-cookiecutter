//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"{{ cookiecutter.__app_name_camel }}/app/{{ cookiecutter.__app_name_pkg }}/internal/biz"
	"{{ cookiecutter.__app_name_camel }}/app/{{ cookiecutter.__app_name_pkg }}/internal/conf"
	"{{ cookiecutter.__app_name_camel }}/app/{{ cookiecutter.__app_name_pkg }}/internal/data"
	"{{ cookiecutter.__app_name_camel }}/app/{{ cookiecutter.__app_name_pkg }}/internal/server"
	"{{ cookiecutter.__app_name_camel }}/app/{{ cookiecutter.__app_name_pkg }}/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
