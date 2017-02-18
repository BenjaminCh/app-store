package usecases

import (
	"../domain"
	"../interfaces"
)

type AppInteractor struct {
	Repository    string
	AppRepository interfaces.IRepository
}

// NewAppInteractor allows to create new AppInteractor object.
func NewAppInteractor(appRepository interfaces.IRepository, repository string) AppInteractor {
	return AppInteractor{
		Repository:    repository,
		AppRepository: appRepository,
	}
}

// GetApps allows to retrieve App Objects from a query.
// It returns matching App objects or an error if the identifier does not exist.
// Implements IAppInteractor interface.
func (ai *AppInteractor) GetApps(id string) ([]domain.App, error) {
	return ai.AppRepository.SearchApps(id, ai.Repository)
}

// PersistApps allows to persist an App Object.
// It returns the persisted App Object if App was persisted properly.
// It returns an if something went wrong during the persist operation.
// Implements IAppInteractor interface.
func (ai *AppInteractor) PersistApps(newApps []domain.App) ([]domain.App, error) {
	return ai.AppRepository.AddApps(newApps, ai.Repository)
}
