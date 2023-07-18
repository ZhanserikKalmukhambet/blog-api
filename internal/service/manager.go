package service

import (
	"github.com/ZhanserikKalmukhambet/blog-api/internal/config"
	"github.com/ZhanserikKalmukhambet/blog-api/internal/repository"
)

type Manager struct {
	Repository repository.Repository
	Config     *config.Config
}

func New(repo repository.Repository, conf *config.Config) *Manager {
	return &Manager{
		Repository: repo,
		Config:     conf,
	}
}
