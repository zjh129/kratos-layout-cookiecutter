package server

import (
	{{ cookiecutter.app_name }} "{{ cookiecutter.app_name_camel }}/api/{{ cookiecutter.app_name }}"
	"{{ cookiecutter.app_name_camel }}/app/user/internal/conf"
	"{{ cookiecutter.app_name_camel }}/app/user/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, {{ cookiecutter.app_name }} *service.{{ cookiecutter.app_name_camel }}Service, logger log.Logger) *http.Server {
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
	{{ cookiecutter.app_name }}.Register{{ cookiecutter.app_name_camel }}HTTPServer(srv, {{ cookiecutter.app_name }})
	return srv
}
