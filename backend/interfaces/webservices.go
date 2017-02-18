package interfaces

import "net/http"
import "fmt"

type AppsInteractor interface{}

type WebserviceHandler struct {
	AppsInteractor AppsInteractor
	// Helper                                   *WebserviceHelper
}

func (handler WebserviceHandler) AppGet(config ConfigurationManager, res http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello Algolia !")
	res.WriteHeader(200)
}
