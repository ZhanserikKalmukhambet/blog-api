package service

import (
	"context"
	"github.com/ZhanserikKalmukhambet/blog-api/internal/model"
	"github.com/ZhanserikKalmukhambet/blog-api/pkg/util"
)

func (m *Manager) CreateUser(ctx context.Context, user *model.User) error {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	err = m.Repository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) UpdateUser(ctx context.Context, user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (m *Manager) DeleteUser(ctx context.Context, user *model.User) error {
	//TODO implement me
	panic("implement me")
}
