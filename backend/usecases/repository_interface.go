package usecases

import (
	"../domain"
)

type IRepository interface {
	CreateBatch(apps []domain.App) ([]string, error)
	Create(app domain.App) (string, error)
	Search(query string) ([]domain.App, error)
	Get(id string) (domain.App, error)
	Delete(id string) (string, error)
	DeleteBatch(ids []string) ([]string, error)
}
