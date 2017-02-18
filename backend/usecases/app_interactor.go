package usecases

import "../domain"

type AppInteractor struct {
	AppRepository IRepository
}

// NewAppInteractor allows to create new AppInteractor object.
func NewAppInteractor(appRepository IRepository) *AppInteractor {
	return &AppInteractor{
		AppRepository: appRepository,
	}
}

// GetApps allows to retrieve App Objects from a query.
// It returns matching App objects or an error.
// Implements IAppInteractor interface.
func (ai *AppInteractor) GetApps(id string) ([]domain.App, error) {
	return ai.AppRepository.SearchApps(id)
}

// GetApp allows to retrieve App Object from an identifier.
// It returns matching App object or an error if the identifier does not exist.
// Implements IAppInteractor interface.
func (ai *AppInteractor) GetApp(id string) (domain.App, error) {
	return ai.AppRepository.SearchApp(id)
}

// PersistApps allows to persist multiple App Objects.
// It returns the persisted objects identifiers if Apps were persisted properly.
// It returns an if something went wrong during the persist operation.
// Implements IAppInteractor interface.
func (ai *AppInteractor) PersistApps(apps []domain.App) ([]string, error) {
	return ai.AppRepository.AddApps(apps)
}

// PersistApp allows to persist an App Object.
// It returns the persisted App Object identifier if App was persisted properly.
// It returns an if something went wrong during the persist operation.
// Implements IAppInteractor interface.
func (ai *AppInteractor) PersistApp(app domain.App) (string, error) {
	return ai.AppRepository.AddApp(app)
}
