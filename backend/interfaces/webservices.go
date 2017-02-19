package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"../domain"
)

// WebserviceHandler : REST API Handler
type WebserviceHandler struct {
	AppInteractor IAppInteractor
}

// AppGet is called on get app HTTP request and returns the given app if found.
func (handler WebserviceHandler) AppGet(config ConfigurationManager, res http.ResponseWriter, req *http.Request) {
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

// AppDelete is called on delete app HTTP request and delete the given app from the index.
func (handler WebserviceHandler) AppDelete(config ConfigurationManager, res http.ResponseWriter, req *http.Request) {
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

// AppCreate is called on create app HTTP request and create the given app to the index.
func (handler WebserviceHandler) AppCreate(config ConfigurationManager, res http.ResponseWriter, req *http.Request) {
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
	vars := mux.Vars(req)

	rank, err := strconv.ParseFloat(vars["rank"], 16)
	if err != nil {
		rank = -1
	}
	app = domain.NewApp(
		vars["name"],
		vars["image"],
		vars["link"],
		vars["category"],
		rank,
	)

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

	// Return result
	// App was successfully deleted
	// Returns a 204
	res.WriteHeader(204)
	return
}
