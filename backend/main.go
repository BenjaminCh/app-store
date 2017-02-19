package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"fmt"

	"./infrastructure"
	"./interfaces"
	"./usecases"
)

func main() {

	// Create the router
	router := mux.NewRouter()

	// Create an set configuration manager
	configInteractor := interfaces.ConfigurationManager{}
	configInteractor.ConfigurationInteractor = infrastructure.NewViperConfig()

	// TODO : Add a logger class

	// Create the controller (aka webservice)
	webserviceHandler := interfaces.WebserviceHandler{}
	// Attach the App Model interactor
	webserviceHandler.AppInteractor = usecases.NewAppInteractor(
		interfaces.NewAlgoliaRepository(
			configInteractor.GetConfigString("algolia.applicationID"),
			configInteractor.GetConfigString("algolia.apiKey"),
			configInteractor.GetConfigString("algolia.indexes.apps"),
		),
	)
	// Attach the webservice helper
	//webserviceHandler.Helper = interfaces.NewWebserviceHelper()

	// TODO : Move the router to an external class
	// Route app get (apps/:id)
	router.
		Methods("GET").
		Path("/api/1/apps/{id:[0-9]*}").
		HandlerFunc(
			func(res http.ResponseWriter, req *http.Request) {
				// Call the webservice handler injecting the congiguration interactor as well.
				webserviceHandler.AppGet(configInteractor, res, req)
			},
		)
	// Route app creation (apps/:app_id)
	router.
		Methods("POST").
		Path("/api/1/apps").
		HandlerFunc(
			func(res http.ResponseWriter, req *http.Request) {
				webserviceHandler.AppCreate(configInteractor, res, req)
			},
		)
	// Route app creation (apps/:app_id)
	router.
		Methods("DELETE").
		Path("/api/1/apps/{id:[0-9]*}").
		HandlerFunc(
			func(res http.ResponseWriter, req *http.Request) {
				webserviceHandler.AppDelete(configInteractor, res, req)
			},
		)

	// Launch the server
	fmt.Println("Server launching port : ", configInteractor.GetConfigString("server.port"))
	http.ListenAndServe(
		":"+configInteractor.GetConfigString("server.port"),
		router,
	)

}
