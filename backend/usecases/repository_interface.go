package usecases

import (
	"../domain"
)

type IRepository interface {
	AddApps(apps []domain.App, repository string) ([]domain.App, error)
	SearchApps(query string, repository string) ([]domain.App, error)
}
