package service

import (
	"context"

	{{ cookiecutter.__app_name_pkg }}api "{{ cookiecutter.__project_name_snake }}/api/{{ cookiecutter.__app_name_pkg }}"
	"{{ cookiecutter.__project_name_snake }}/app/{{ cookiecutter.__app_name_pkg }}/internal/biz"
)

// {{ cookiecutter.__app_name_camel }}Service is a {{ cookiecutter.__app_name_camel }} service.
type {{ cookiecutter.__app_name_camel }}Service struct {
	{{ cookiecutter.__app_name_pkg }}api.Unimplemented{{ cookiecutter.__app_name_camel }}Server

	uc *biz.{{ cookiecutter.__app_name_camel }}Usecase
}

// New{{ cookiecutter.__app_name_camel }}Service new a {{ cookiecutter.__app_name_camel }} service.
func New{{ cookiecutter.__app_name_camel }}Service(uc *biz.{{ cookiecutter.__app_name_camel }}Usecase) *{{ cookiecutter.__app_name_camel }}Service {
	return &{{ cookiecutter.__app_name_camel }}Service{uc: uc}
}

// SayHello implements helloworld.{{ cookiecutter.__app_name_camel }}Server.
func (s *{{ cookiecutter.__app_name_camel }}Service) SayHello(ctx context.Context, in *{{ cookiecutter.__app_name_pkg }}_api.HelloRequest) (*{{ cookiecutter.__app_name_pkg }}api.HelloReply, error) {
	g, err := s.uc.Create{{ cookiecutter.__app_name_camel }}(ctx, &biz.{{ cookiecutter.__app_name_camel }}{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &{{ cookiecutter.__app_name_pkg }}api.HelloReply{Message: "Hello " + g.Hello}, nil
}
