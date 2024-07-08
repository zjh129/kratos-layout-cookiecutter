package biz

import (
	"context"

	v1 "{{ cookiecutter.__project_name_snake }}/api/{{ cookiecutter.__app_name_pkg }}"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// {{ cookiecutter.app_name_camel }} is a {{ cookiecutter.app_name_camel }} model.
type {{ cookiecutter.app_name_camel }} struct {
	Hello string
}

// {{ cookiecutter.app_name_camel }}Repo is a Greater repo.
type {{ cookiecutter.app_name_camel }}Repo interface {
	Save(context.Context, *{{ cookiecutter.app_name_camel }}) (*{{ cookiecutter.app_name_camel }}, error)
	Update(context.Context, *{{ cookiecutter.app_name_camel }}) (*{{ cookiecutter.app_name_camel }}, error)
	FindByID(context.Context, int64) (*{{ cookiecutter.app_name_camel }}, error)
	ListByHello(context.Context, string) ([]*{{ cookiecutter.app_name_camel }}, error)
	ListAll(context.Context) ([]*{{ cookiecutter.app_name_camel }}, error)
}

// {{ cookiecutter.app_name_camel }}Usecase is a {{ cookiecutter.app_name_camel }} usecase.
type {{ cookiecutter.app_name_camel }}Usecase struct {
	repo {{ cookiecutter.app_name_camel }}Repo
	log  *log.Helper
}

// New{{ cookiecutter.app_name_camel }}Usecase new a {{ cookiecutter.app_name_camel }} usecase.
func New{{ cookiecutter.app_name_camel }}Usecase(repo {{ cookiecutter.app_name_camel }}Repo, logger log.Logger) *{{ cookiecutter.app_name_camel }}Usecase {
	return &{{ cookiecutter.app_name_camel }}Usecase{repo: repo, log: log.NewHelper(logger)}
}

// Create{{ cookiecutter.app_name_camel }} creates a {{ cookiecutter.app_name_camel }}, and returns the new {{ cookiecutter.app_name_camel }}.
func (uc *{{ cookiecutter.app_name_camel }}Usecase) Create{{ cookiecutter.app_name_camel }}(ctx context.Context, g *{{ cookiecutter.app_name_camel }}) (*{{ cookiecutter.app_name_camel }}, error) {
	uc.log.WithContext(ctx).Infof("Create{{ cookiecutter.app_name_camel }}: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
