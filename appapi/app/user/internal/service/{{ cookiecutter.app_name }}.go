package service

import (
	"context"

	{{ cookiecutter.app_name }} "{{ cookiecutter.__project_name_snake }}/api/{{ cookiecutter.app_name }}"
	"{{ cookiecutter.__project_name_snake }}/app/user/internal/biz"
)

// {{ cookiecutter.app_name_camel }}Service is a {{ cookiecutter.app_name_camel }} service.
type {{ cookiecutter.app_name_camel }}Service struct {
	{{ cookiecutter.app_name }}.Unimplemented{{ cookiecutter.app_name_camel }}Server

	uc *biz.{{ cookiecutter.app_name_camel }}Usecase
}

// New{{ cookiecutter.app_name_camel }}Service new a {{ cookiecutter.app_name_camel }} service.
func New{{ cookiecutter.app_name_camel }}Service(uc *biz.{{ cookiecutter.app_name_camel }}Usecase) *{{ cookiecutter.app_name_camel }}Service {
	return &{{ cookiecutter.app_name_camel }}Service{uc: uc}
}

// SayHello implements helloworld.{{ cookiecutter.app_name_camel }}Server.
func (s *{{ cookiecutter.app_name_camel }}Service) SayHello(ctx context.Context, in *{{ cookiecutter.app_name }}.HelloRequest) (*{{ cookiecutter.app_name }}.HelloReply, error) {
	g, err := s.uc.Create{{ cookiecutter.app_name_camel }}(ctx, &biz.{{ cookiecutter.app_name_camel }}{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &{{ cookiecutter.app_name }}.HelloReply{Message: "Hello " + g.Hello}, nil
}
