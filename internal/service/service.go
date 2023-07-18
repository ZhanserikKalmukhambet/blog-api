package service

import (
	"context"
	"github.com/ZhanserikKalmukhambet/blog-api/internal/model"
)

type Service interface {
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, user *model.User) error

	// other methods
}
