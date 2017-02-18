package usecases

import (
	"../domain"
)

type IRepository interface {
	AddApps(apps []domain.App) ([]string, error)
	AddApp(app domain.App) (string, error)
	SearchApps(query string) ([]domain.App, error)
	SearchApp(id string) (domain.App, error)
}
