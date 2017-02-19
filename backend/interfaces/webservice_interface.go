package interfaces

import "net/http"

type IWebservice interface {
	Get(config ConfigurationManager, res http.ResponseWriter, req *http.Request)
	Create(config ConfigurationManager, res http.ResponseWriter, req *http.Request)
	Delete(config ConfigurationManager, res http.ResponseWriter, req *http.Request)
}
