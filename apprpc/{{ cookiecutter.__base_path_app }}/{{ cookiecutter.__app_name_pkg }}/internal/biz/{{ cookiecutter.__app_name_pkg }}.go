package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

// {{ cookiecutter.__app_name_camel }} is a {{ cookiecutter.__app_name_camel }} model.
type {{ cookiecutter.__app_name_camel }} struct {
	Hello string
}

// {{ cookiecutter.__app_name_camel }}Repo is a Greater repo.
type {{ cookiecutter.__app_name_camel }}Repo interface {
	Save(context.Context, *{{ cookiecutter.__app_name_camel }}) (*{{ cookiecutter.__app_name_camel }}, error)
	Update(context.Context, *{{ cookiecutter.__app_name_camel }}) (*{{ cookiecutter.__app_name_camel }}, error)
	FindByID(context.Context, int64) (*{{ cookiecutter.__app_name_camel }}, error)
	ListByHello(context.Context, string) ([]*{{ cookiecutter.__app_name_camel }}, error)
	ListAll(context.Context) ([]*{{ cookiecutter.__app_name_camel }}, error)
}

// {{ cookiecutter.__app_name_camel }}Usecase is a {{ cookiecutter.__app_name_camel }} usecase.
type {{ cookiecutter.__app_name_camel }}Usecase struct {
	repo {{ cookiecutter.__app_name_camel }}Repo
	log  *log.Helper
}

// New{{ cookiecutter.__app_name_camel }}Usecase new a {{ cookiecutter.__app_name_camel }} usecase.
func New{{ cookiecutter.__app_name_camel }}Usecase(repo {{ cookiecutter.__app_name_camel }}Repo, logger log.Logger) *{{ cookiecutter.__app_name_camel }}Usecase {
	return &{{ cookiecutter.__app_name_camel }}Usecase{repo: repo, log: log.NewHelper(logger)}
}

// Create{{ cookiecutter.__app_name_camel }} creates a {{ cookiecutter.__app_name_camel }}, and returns the new {{ cookiecutter.__app_name_camel }}.
func (uc *{{ cookiecutter.__app_name_camel }}Usecase) Create{{ cookiecutter.__app_name_camel }}(ctx context.Context, g *{{ cookiecutter.__app_name_camel }}) (*{{ cookiecutter.__app_name_camel }}, error) {
	uc.log.WithContext(ctx).Infof("Create{{ cookiecutter.__app_name_camel }}: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
