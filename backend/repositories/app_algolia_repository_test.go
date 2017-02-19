package repositories_test

import (
	"testing"

	"github.com/BenjaminCh/app-store/backend/domain"
	"github.com/BenjaminCh/app-store/backend/infrastructure"
	"github.com/BenjaminCh/app-store/backend/interfaces"
	"github.com/BenjaminCh/app-store/backend/usecases"
)

// TestAddApp, test if the Algolia app interactor does work properly adding one app.
func TestAddApp(t *testing.T) {

	var err error

	// Use the configuration manager to get Algolia's keys to mock the app interactor
	configInteractor := interfaces.ConfigurationManager{}
	configInteractor.ConfigurationInteractor = infrastructure.NewViperConfig()

	// Instanciate the App interactor
	appInteractor := usecases.NewAppInteractor(
		interfaces.NewAlgoliaRepository(
			configInteractor.GetConfigString("algolia.applicationID"),
			configInteractor.GetConfigString("algolia.apiKey"),
			configInteractor.GetConfigString("algolia.indexes.apps"),
		),
	)

	// Single addition
	// Create a random app
	testApp := domain.NewApp(
		"Unit testing app interactor",
		"static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
		"http://whatthetvshow.com/fe/snapshots/2205",
		"Quiz unit tests Benjamin",
		223,
	)
	// Try to persist it
	res, err := appInteractor.Create(testApp)

	// Testing returns
	if err != nil {
		// Error raised during the creation
		t.Error("The app was not properly added : ", err)
		return
	}
	if res == "" {
		// No object created
		t.Error("The app was not properly added : no identifier returned")
		return
	}

	_, _ = appInteractor.Delete(res)

	t.Log("TestAddApp: Test Clear")
	return
}

// TestAddApps, test if the Algolia app interactor does work properly adding apps.
func TestAddApps(t *testing.T) {

	var err error

	// Use the configuration manager to get Algolia's keys to mock the app interactor
	configInteractor := interfaces.ConfigurationManager{}
	configInteractor.ConfigurationInteractor = infrastructure.NewViperConfig()

	// Instanciate the App interactor
	appInteractor := usecases.NewAppInteractor(
		interfaces.NewAlgoliaRepository(
			configInteractor.GetConfigString("algolia.applicationID"),
			configInteractor.GetConfigString("algolia.apiKey"),
			configInteractor.GetConfigString("algolia.indexes.apps"),
		),
	)

	// Batch addition
	// Create a random app
	testApps := []domain.App{
		domain.NewApp(
			"Unit testing app interactor",
			"static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
			"http://whatthetvshow.com/fe/snapshots/2205",
			"Quiz unit tests Benjamin",
			223,
		),
		domain.NewApp(
			"Unit testing app interactor 2",
			"static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
			"http://whatthetvshow.com/fe/snapshots/2205",
			"Quiz unit tests Benjamin",
			224,
		),
		domain.NewApp(
			"Unit testing app interactor 3",
			"static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
			"http://whatthetvshow.com/fe/snapshots/2205",
			"Quiz unit tests Benjamin",
			225,
		),
		domain.NewApp(
			"Unit testing app interactor 4",
			"static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
			"http://whatthetvshow.com/fe/snapshots/2205",
			"Quiz unit tests Benjamin",
			226,
		),
	}
	// Try to persist it
	res, err := appInteractor.CreateBatch(testApps)

	// Testing returns
	if err != nil {
		// Error raised during the creation
		t.Error("Apps were not properly added : ", err)
		return
	}
	if len(res) == 0 {
		// None of apps were created properly
		t.Error("Apps were not properly added : no identifiers returned")
		return
	}
	for i := range testApps {
		if len(res) > 0 && res[i] == "" {
			// No object created
			t.Error("App number ", i, " '", testApps[i].Name, " was not properly added : no identifier returned")
		}
	}

	_, _ = appInteractor.DeleteBatch(res)

	t.Log("TestAddApp: Test Clear")

	t.Log("TestAddApps: Test Clear")
	return
}

// TestDeleteApp, test if the Algolia app interactor does work properly deleting an app.
func TestDeleteApp(t *testing.T) {

	var err error

	// Use the configuration manager to get Algolia's keys to mock the app interactor
	configInteractor := interfaces.ConfigurationManager{}
	configInteractor.ConfigurationInteractor = infrastructure.NewViperConfig()

	// Instanciate the App interactor
	appInteractor := usecases.NewAppInteractor(
		interfaces.NewAlgoliaRepository(
			configInteractor.GetConfigString("algolia.applicationID"),
			configInteractor.GetConfigString("algolia.apiKey"),
			configInteractor.GetConfigString("algolia.indexes.apps"),
		),
	)

	// Single addition
	// Create a random app
	testApp := domain.NewApp(
		"Unit testing app interactor",
		"static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
		"http://whatthetvshow.com/fe/snapshots/2205",
		"Quiz unit tests Benjamin",
		223,
	)
	// Persist it
	res, err := appInteractor.Create(testApp)

	// Testing returns
	// No need to test returns here since they are already handled in TestAddApp test

	// Try to delete it
	res, err = appInteractor.Delete(res)
	if err != nil {
		// Error raised during the deletion
		t.Error("App was not properly delete : ", err)
		return
	}
	if res == "" {
		// No identifier was passed to confirm deletion
		t.Error("App was not properly deleted : no identifier returned")
		return
	}

	t.Log("TestDeleteApp: Test Clear")
	return
}

// TestDeleteApps, test if the Algolia app interactor does work properly deleting apps.
func TestDeleteApps(t *testing.T) {

	var err error

	// Use the configuration manager to get Algolia's keys to mock the app interactor
	configInteractor := interfaces.ConfigurationManager{}
	configInteractor.ConfigurationInteractor = infrastructure.NewViperConfig()

	// Instanciate the App interactor
	appInteractor := usecases.NewAppInteractor(
		interfaces.NewAlgoliaRepository(
			configInteractor.GetConfigString("algolia.applicationID"),
			configInteractor.GetConfigString("algolia.apiKey"),
			configInteractor.GetConfigString("algolia.indexes.apps"),
		),
	)

	// Batch addition
	// Create a random app
	testApps := []domain.App{
		domain.NewApp(
			"Unit testing app interactor",
			"static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
			"http://whatthetvshow.com/fe/snapshots/2205",
			"Quiz unit tests Benjamin",
			223,
		),
		domain.NewApp(
			"Unit testing app interactor 2",
			"static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
			"http://whatthetvshow.com/fe/snapshots/2205",
			"Quiz unit tests Benjamin",
			224,
		),
		domain.NewApp(
			"Unit testing app interactor 3",
			"static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
			"http://whatthetvshow.com/fe/snapshots/2205",
			"Quiz unit tests Benjamin",
			225,
		),
		domain.NewApp(
			"Unit testing app interactor 4",
			"static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
			"http://whatthetvshow.com/fe/snapshots/2205",
			"Quiz unit tests Benjamin",
			226,
		),
	}
	// Try to persist it
	res, err := appInteractor.CreateBatch(testApps)

	// Testing returns
	// No need to test returns here since they are already handled in TestAddApps test
	res, err = appInteractor.DeleteBatch(res)

	// Try to delete those apps
	if err != nil {
		// Error raised during the deletion
		t.Error("Apps were not properly delete : ", err)
		return
	}
	if len(res) == 0 {
		// None of apps were deleted properly
		t.Error("Apps were not properly deleted : no identifiers returned")
		return
	}
	for i := range testApps {
		if len(res) > 0 && res[i] == "" {
			// No object created
			t.Error("App number ", i, " '", testApps[i].Name, " was not properly deleted : no identifier returned")
		}
	}

	t.Log("TestDeleteApps: Test Clear")
	return
}

// TestAddApp, test if the Algolia app interactor does work properly getting one app from its identifier.
// TODO : Debug this not working why, I don't understand why it returns {"message":"ObjectID does not exist","status":404}
// While I am able to find the app by its id in the UI from the printed id.
/*func TestGetApp(t *testing.T) {

	var err error

	// Use the configuration manager to get Algolia's keys to mock the app interactor
	configInteractor := interfaces.ConfigurationManager{}
	configInteractor.ConfigurationInteractor = infrastructure.NewViperConfig()

	// Instanciate the App interactor
	appInteractor := usecases.NewAppInteractor(
		interfaces.NewAlgoliaRepository(
			configInteractor.GetConfigString("algolia.applicationID"),
			configInteractor.GetConfigString("algolia.apiKey"),
			configInteractor.GetConfigString("algolia.indexes.apps"),
		),
	)

	// Single addition
	// Create a random app
	tracker := time.Now().Unix()
	testApp := domain.NewApp(
		"Unit testing app interactor",
		"static.whatthetvshow.com:9000/media/snapshots/c4c3021b168ba93572d402e313f0f884_medium.png",
		"http://whatthetvshow.com/fe/snapshots/2205",
		fmt.Sprintf("Quiz unit tests Benjamin %d", tracker), // Used to be sure we will be able to be sure this is the good object once returned
		223,
	)
	// Try to persist it
	res, err := appInteractor.Create(testApp)

	// Testing returns
	// No need to test returns here since they are already handled in TestAddApp test

	// Try to get the app from its identifier
	app, err := appInteractor.Get(res)

	if err != nil {
		// Error raised during the get operation
		t.Error("App was not found : ", err)
		return
	}
	if !strings.Contains(app.Category, strconv.FormatInt(tracker, 10)) {
		// No identifier was passed to confirm deletion
		t.Error("The returned app is not the one we just created.")
		return
	}

	_, _ = appInteractor.Delete(res)

	t.Log("TestGetApp: Test Clear")
	return
}*/
