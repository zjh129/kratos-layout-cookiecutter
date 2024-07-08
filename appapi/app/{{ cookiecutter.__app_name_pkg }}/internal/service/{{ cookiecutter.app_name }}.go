package service

import (
	"context"

	{{ cookiecutter.__app_name_camel }} "{{ cookiecutter.__project_name_snake }}/api/{{ cookiecutter.__app_name_camel }}"
	"{{ cookiecutter.__project_name_snake }}/app/{{ cookiecutter.__app_name_pkg }}/internal/biz"
)

// {{ cookiecutter.app_name_camel }}Service is a {{ cookiecutter.app_name_camel }} service.
type {{ cookiecutter.app_name_camel }}Service struct {
	{{ cookiecutter.__app_name_camel }}.Unimplemented{{ cookiecutter.app_name_camel }}Server

	uc *biz.{{ cookiecutter.app_name_camel }}Usecase
}

// New{{ cookiecutter.app_name_camel }}Service new a {{ cookiecutter.app_name_camel }} service.
func New{{ cookiecutter.app_name_camel }}Service(uc *biz.{{ cookiecutter.app_name_camel }}Usecase) *{{ cookiecutter.app_name_camel }}Service {
	return &{{ cookiecutter.app_name_camel }}Service{uc: uc}
}

// SayHello implements helloworld.{{ cookiecutter.app_name_camel }}Server.
func (s *{{ cookiecutter.app_name_camel }}Service) SayHello(ctx context.Context, in *{{ cookiecutter.__app_name_camel }}.HelloRequest) (*{{ cookiecutter.__app_name_camel }}.HelloReply, error) {
	g, err := s.uc.Create{{ cookiecutter.app_name_camel }}(ctx, &biz.{{ cookiecutter.app_name_camel }}{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &{{ cookiecutter.__app_name_camel }}.HelloReply{Message: "Hello " + g.Hello}, nil
}
