package server

import (
	{{ cookiecutter.__app_name_camel }} "{{ cookiecutter.__project_name_camel }}/api/{{ cookiecutter.__app_name_pkg }}"
	"{{ cookiecutter.__app_name_camel }}/app/{{ cookiecutter.__app_name_pkg }}/internal/conf"
	"{{ cookiecutter.__app_name_camel }}/app/{{ cookiecutter.__app_name_pkg }}/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, {{ cookiecutter.__app_name_camel }} *service.{{ cookiecutter.__app_name_camel }}Service, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	{{ cookiecutter.__app_name_camel }}.Register{{ cookiecutter.__app_name_camel }}HTTPServer(srv, {{ cookiecutter.__app_name_camel }})
	return srv
}
