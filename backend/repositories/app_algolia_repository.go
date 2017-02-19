package repositories

import (
	"encoding/json"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"

	"github.com/BenjaminCh/algolia-app-store/backend/domain"
)

// AlgoliaRepository allows to interact with Algolia backend.
type AlgoliaRepository struct {
	client     algoliasearch.Client
	repository string
}

// NewAlgoliaRepository creates a new AlgoliaRepository object.
func NewAlgoliaRepository(applicationID string, apiKey string, repository string) *AlgoliaRepository {
	return &AlgoliaRepository{
		client:     algoliasearch.NewClient(applicationID, apiKey),
		repository: repository,
	}
}

// AppToObject allows to convert a domain.App object to an Algolia Object.
func (ar *AlgoliaRepository) AppToObject(app domain.App) (algoliasearch.Object, error) {
	var err error
	var object algoliasearch.Object
	appJSON, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(appJSON, &object); err != nil {
		return nil, err
	}
	return object, err
}

// ObjectToApp allows to convert a Algolia Map object to an domain.App.
func (ar *AlgoliaRepository) HitToApp(object algoliasearch.Map) (domain.App, error) {
	var err error
	var app domain.App

	app = domain.NewApp(
		object["name"].(string),
		object["image"].(string),
		object["link"].(string),
		object["category"].(string),
		object["rank"].(float64),
	)

	return app, err
}

// AppsToObjects allows to convert multiple domain.App objects to Algolia Objects.
func (ar *AlgoliaRepository) AppsToObjects(app []domain.App) ([]algoliasearch.Object, error) {
	var objects []algoliasearch.Object
	var object algoliasearch.Object
	var err error

	for index := range app {
		object, err = ar.AppToObject(app[index])
		if err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}

	return objects, err
}

// ObjectToApp allows to convert a Algolia Object to an domain.App.
func (ar *AlgoliaRepository) ObjectToApp(object algoliasearch.Object) (domain.App, error) {
	var err error
	var app domain.App

	app = domain.NewApp(
		object["name"].(string),
		object["image"].(string),
		object["link"].(string),
		object["category"].(string),
		object["rank"].(float64),
	)

	return app, err
}

// ObjectsToApps allows to convert multiple Algolia Maps to domain.App objects.
func (ar *AlgoliaRepository) HitsToApps(object []algoliasearch.Map) ([]domain.App, error) {
	var apps []domain.App
	var app domain.App
	var err error

	for index := range object {
		app, err = ar.HitToApp(object[index])
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}

	return apps, err
}

// Create allows to add app into the app index.
// Returns the id of added app or an error.
// Implements IRepository interface.
func (ar *AlgoliaRepository) Create(newApp domain.App) (string, error) {
	var err error
	index := ar.client.InitIndex(ar.repository)

	// Convert app objects to algolia objects
	objects, err := ar.AppToObject(newApp)
	if err != nil {
		return "", err
	}

	// Add the apps to algolia index
	res, err := index.AddObject(objects)
	if err != nil {
		return "", err
	}

	return res.ObjectID, err
}

// CreateBatch allows to add apps into the app index.
// Returns the ids of added apps or an error.
// Implements IRepository interface.
func (ar *AlgoliaRepository) CreateBatch(newApps []domain.App) ([]string, error) {
	var err error
	index := ar.client.InitIndex(ar.repository)

	// Convert app objects to algolia objects
	objects, err := ar.AppsToObjects(newApps)
	if err != nil {
		return nil, err
	}

	// Add the apps to algolia index
	res, err := index.AddObjects(objects)
	if err != nil {
		return nil, err
	}

	return res.ObjectIDs, err
}

// Search allows to search apps in the app index.
// Returns a list of Apps Objects or an error.
// Implements IRepository interface.
func (ar *AlgoliaRepository) Search(query string) ([]domain.App, error) {
	var err error
	var apps []domain.App

	index := ar.client.InitIndex(ar.repository)

	// Get objects from Index
	objects, err := index.Search(query, nil)
	if err != nil {
		return nil, err
	}

	// Convert those objects back to App Objects
	apps, err = ar.HitsToApps(objects.Hits)
	if err != nil {
		return nil, err
	}

	return apps, err
}

// Get allows to search an app in the app index.
// Returns the matching App Object or an error.
// Implements IRepository interface.
func (ar *AlgoliaRepository) Get(id string) (domain.App, error) {
	var err error
	var app domain.App

	index := ar.client.InitIndex(ar.repository)

	// Get objects from Index
	object, err := index.GetObject(id, nil)
	if err != nil {
		return app, err
	}

	// Convert those objects back to App Objects
	app, err = ar.ObjectToApp(object)
	if err != nil {
		return app, err
	}

	return app, err
}

// Delete allows to delete an app from the app index.
// Returns the id of deleted app or an error.
// Implements IRepository interface.
func (ar *AlgoliaRepository) Delete(id string) (string, error) {
	var err error
	index := ar.client.InitIndex(ar.repository)

	// Add the apps to algolia index
	res, err := index.DeleteObject(id)
	if err != nil {
		return "", err
	}

	return res.DeletedAt, err // TODO : No really the deleted ID but, well, to be improved in my workflow :)
}

// Delete allows to delete multiple apps from the app index.
// Returns the ids of deleted apps or an error.
// Implements IRepository interface.
func (ar *AlgoliaRepository) DeleteBatch(ids []string) ([]string, error) {
	var err error
	index := ar.client.InitIndex(ar.repository)

	// Delete the apps from algolia index
	res, err := index.DeleteObjects(ids)
	if err != nil {
		return nil, err
	}

	return res.ObjectIDs, err
}
