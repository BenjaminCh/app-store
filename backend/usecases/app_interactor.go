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

// Search allows to retrieve App Objects from a query.
// It returns matching App objects or an error.
// Implements IAppInteractor interface.
func (ai *AppInteractor) Search(id string) ([]domain.App, error) {
	return ai.AppRepository.Search(id)
}

// Get allows to retrieve App Object from an identifier.
// It returns matching App object or an error if the identifier does not exist.
// Implements IAppInteractor interface.
func (ai *AppInteractor) Get(id string) (domain.App, error) {
	return ai.AppRepository.Get(id)
}

// CreateBatch allows to persist multiple App Objects.
// It returns the persisted objects identifiers if Apps were persisted properly.
// It returns an error if something went wrong during the persist operation.
// Implements IAppInteractor interface.
func (ai *AppInteractor) CreateBatch(apps []domain.App) ([]string, error) {
	return ai.AppRepository.CreateBatch(apps)
}

// Create allows to persist an App Object.
// It returns the persisted App Object identifier if App was persisted properly.
// It returns an error if something went wrong during the persist operation.
// Implements IAppInteractor interface.
func (ai *AppInteractor) Create(app domain.App) (string, error) {
	return ai.AppRepository.Create(app)
}

// Delete allows to delete an App Object.
// It returns the deleted App Object identifier if App was deleted properly.
// It returns an error if something went wrong during the delete operation.
// Implements IAppInteractor interface.
func (ai *AppInteractor) Delete(id string) (string, error) {
	return ai.AppRepository.Delete(id)
}

// DeleteBatch allows to delete multiple App Objects.
// It returns the deleted App Object identifiers if App were deleted properly.
// It returns an error if something went wrong during the delete operation.
// Implements IAppInteractor interface.
func (ai *AppInteractor) DeleteBatch(ids []string) ([]string, error) {
	return ai.AppRepository.DeleteBatch(ids)
}
