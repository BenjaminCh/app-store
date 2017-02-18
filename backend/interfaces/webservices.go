package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	app, err = handler.AppInteractor.GetApp(appID)
	if err != nil {
		// An error occured while trying to retrieve the app
		// Return a HTTP 500 (Internal Server Error)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(app)
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
