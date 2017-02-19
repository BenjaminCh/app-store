package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/BenjaminCh/app-store/backend/handlers"
	"github.com/BenjaminCh/app-store/backend/infrastructure"
	"github.com/BenjaminCh/app-store/backend/interfaces"
	"github.com/BenjaminCh/app-store/backend/repositories"
	"github.com/BenjaminCh/app-store/backend/usecases"
)

func main() {

	// Create the router
	router := mux.NewRouter()

	// Create an set configuration manager
	configInteractor := interfaces.ConfigurationManager{}
	configInteractor.ConfigurationInteractor = infrastructure.NewViperConfig()

	// TODO : Add a logger class

	// Create the controller (aka webservice)
	webserviceHandler := handlers.AppWebserviceHandler{}
	// Attach the App Model interactor
	webserviceHandler.AppInteractor = usecases.NewAppInteractor(
		repositories.NewAlgoliaRepository(
			configInteractor.GetConfigString("algolia.applicationID"),
			configInteractor.GetConfigString("algolia.apiKey"),
			configInteractor.GetConfigString("algolia.indexes.apps"),
		),
	)

	// TODO : Move the router to an external class
	// Route app get (GET apps/:id)
	router.
		Methods("GET").
		Path("/api/1/apps/{id:[0-9]*}").
		HandlerFunc(
			func(res http.ResponseWriter, req *http.Request) {
				// Call the webservice handler injecting the configuration interactor as well.
				webserviceHandler.Get(configInteractor, res, req)
			},
		)
	// Route app creation (POST /apps)
	router.
		Methods("POST").
		Path("/api/1/apps").
		HandlerFunc(
			func(res http.ResponseWriter, req *http.Request) {
				// Call the webservice handler injecting the configuration interactor as well.
				webserviceHandler.Create(configInteractor, res, req)
			},
		)
	// Route app deletion (DELETE /apps)
	router.
		Methods("DELETE").
		Path("/api/1/apps/{id:[0-9]*}").
		HandlerFunc(
			func(res http.ResponseWriter, req *http.Request) {
				// Call the webservice handler injecting the configuration interactor as well.
				webserviceHandler.Delete(configInteractor, res, req)
			},
		)

	// Launch the server
	fmt.Println("Server launching port : ", configInteractor.GetConfigString("server.port"))
	port := configInteractor.GetConfigString("server.port")
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	log.Fatal(
		http.ListenAndServe(
			":"+port,
			router,
		),
	)

}
