package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"../domain"
	"../interfaces"
)

// AppWebserviceHandler : REST API Handler
type AppWebserviceHandler struct {
	AppInteractor interfaces.IAppInteractor
}

// Get is called on get app HTTP request and returns the given app if found.
// Implements IWebservice
func (handler AppWebserviceHandler) Get(config interfaces.ConfigurationManager, res http.ResponseWriter, req *http.Request) {
	var err error
	var app domain.App
	var appID string

	defer req.Body.Close()

	// Reject every requests that are not HTTP GET
	if req.Method != "GET" {
		// Return a HTTP 403 (UnAuthorized)
		http.Error(res, err.Error(), http.StatusUnauthorized)
		return
	}

	// Parse params to extract app identifier to be retrieved
	vars := mux.Vars(req)
	appID = vars["id"]
	if appID == "" {
		// No app identifier passed via the HTTP query
		// Return a HTTP 404 (Not Found)
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	// Search for app
	app, err = handler.AppInteractor.Get(appID)
	if err != nil {
		// An error occured while trying to retrieve the app
		// Return a HTTP 500 (Internal Server Error)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	appJSON, err := json.Marshal(app)
	if err != nil {
		// An error occured while trying to serialize the app object to JSON
		// Return a HTTP 500 (Internal Server Error)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return result
	res.Header().Set("Content-Type", "application/json")
	res.Write(appJSON)
	return
}

// Delete is called on delete app HTTP request and delete the given app from the index.
// Implements IWebservice
func (handler AppWebserviceHandler) Delete(config interfaces.ConfigurationManager, res http.ResponseWriter, req *http.Request) {
	var err error
	var appID string

	defer req.Body.Close()

	// Reject every requests that are not HTTP DELETE
	if req.Method != "DELETE" {
		// Return a HTTP 403 (UnAuthorized)
		http.Error(res, err.Error(), http.StatusUnauthorized)
		return
	}

	// Parse params to extract app identifier to be retrieved
	vars := mux.Vars(req)
	appID = vars["id"]
	if appID == "" {
		// No app identifier passed via the HTTP query
		// Return a HTTP 404 (Not Found)
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	// Delete the app
	deleteResult, err := handler.AppInteractor.Delete(appID)
	if err != nil {
		// An error occured while trying to retrieve the app
		// Return a HTTP 500 (Internal Server Error)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(deleteResult) == 0 {
		// If no app was deleted
		// Return a HTTP 404 (Not Found)
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	// Return result
	// App was successfully deleted
	// Returns a 200
	res.WriteHeader(200)
	return
}

// Create is called on create app HTTP request and create the given app to the index.
// Implements IWebservice
func (handler AppWebserviceHandler) Create(config interfaces.ConfigurationManager, res http.ResponseWriter, req *http.Request) {
	var err error
	var app domain.App

	defer req.Body.Close()

	// Reject every requests that are not HTTP POST
	if req.Method != "POST" {
		// Return a HTTP 403 (UnAuthorized)
		http.Error(res, err.Error(), http.StatusUnauthorized)
		return
	}

	// Parse params to extract all params needed to create an app
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&app)
	if err != nil {
		// An error occured while trying to decode POST params
		// Return a HTTP 500 (Internal Server Error)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(app)

	// Create the app
	createResult, err := handler.AppInteractor.Create(app)
	if err != nil {
		// An error occured while trying to retrieve the app
		// Return a HTTP 500 (Internal Server Error)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(createResult) == 0 {
		// If app wasn't created
		// Return a HTTP 400 (StatusBadRequest)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	resultJSON, err := json.Marshal(createResult)
	if err != nil {
		// An error occured while trying to serialize the result object to JSON
		// Return a HTTP 500 (Internal Server Error)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return result
	// App was successfully deleted
	// Returns a 204
	res.WriteHeader(204)
	res.Write(resultJSON)
	return
}
