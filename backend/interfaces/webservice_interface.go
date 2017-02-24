package interfaces

import "net/http"

type IWebservice interface {
	Default(config ConfigurationManager, res http.ResponseWriter, req *http.Request)
	Get(config ConfigurationManager, res http.ResponseWriter, req *http.Request)
	Create(config ConfigurationManager, res http.ResponseWriter, req *http.Request)
	Delete(config ConfigurationManager, res http.ResponseWriter, req *http.Request)
}
