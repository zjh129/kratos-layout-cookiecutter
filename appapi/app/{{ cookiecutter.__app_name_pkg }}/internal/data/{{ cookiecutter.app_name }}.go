package data

import (
	"context"

	"{{ cookiecutter.__project_name_camel }}/app/{{ cookiecutter.__app_name_pkg }}/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type {{ cookiecutter.app_name_camel }}Repo struct {
	data *Data
	log  *log.Helper
}

// New{{ cookiecutter.app_name_camel }}Repo .
func New{{ cookiecutter.app_name_camel }}Repo(data *Data, logger log.Logger) biz.{{ cookiecutter.app_name_camel }}Repo {
	return &{{ cookiecutter.app_name_camel }}Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *{{ cookiecutter.app_name_camel }}Repo) Save(ctx context.Context, g *biz.{{ cookiecutter.app_name_camel }}) (*biz.{{ cookiecutter.app_name_camel }}, error) {
	return g, nil
}

func (r *{{ cookiecutter.app_name_camel }}Repo) Update(ctx context.Context, g *biz.{{ cookiecutter.app_name_camel }}) (*biz.{{ cookiecutter.app_name_camel }}, error) {
	return g, nil
}

func (r *{{ cookiecutter.app_name_camel }}Repo) FindByID(context.Context, int64) (*biz.{{ cookiecutter.app_name_camel }}, error) {
	return nil, nil
}

func (r *{{ cookiecutter.app_name_camel }}Repo) ListByHello(context.Context, string) ([]*biz.{{ cookiecutter.app_name_camel }}, error) {
	return nil, nil
}

func (r *{{ cookiecutter.app_name_camel }}Repo) ListAll(context.Context) ([]*biz.{{ cookiecutter.app_name_camel }}, error) {
	return nil, nil
}
