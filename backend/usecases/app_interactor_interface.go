package usecases

import (
	"../domain"
)

type IAppInteractor interface {
	Get(id int64) (*domain.App, error)
	Persist(newApp *domain.App) (*domain.App, error)
}
