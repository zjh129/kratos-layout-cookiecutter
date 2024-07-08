package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(New{{ cookiecutter.__app_name_camel }}Service)
