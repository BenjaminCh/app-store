package interfaces

import (
	"../domain"
)

type IAppInteractor interface {
	GetApps(query string) ([]domain.App, error)
	GetApp(id string) (domain.App, error)
	PersistApps(newApps []domain.App) ([]string, error)
	PersistApp(newApp domain.App) (string, error)
}
