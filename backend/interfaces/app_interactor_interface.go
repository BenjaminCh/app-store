package interfaces

import (
	"github.com/BenjaminCh/app-store/backend/domain"
)

type IAppInteractor interface {
	Search(query string) ([]domain.App, error)
	Get(id string) (domain.App, error)
	CreateBatch(newApps []domain.App) ([]string, error)
	Create(newApp domain.App) (string, error)
	Delete(id string) (string, error)
	DeleteBatch(ids []string) ([]string, error)
}
