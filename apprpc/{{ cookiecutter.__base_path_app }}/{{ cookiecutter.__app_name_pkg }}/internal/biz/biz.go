package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(New{{cookiecutter.__app_name_camel}}Usecase)
