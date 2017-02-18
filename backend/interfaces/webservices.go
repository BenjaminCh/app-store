package interfaces

import (
	"fmt"
	"net/http"
)

type WebserviceHandler struct {
	AppInteractor IAppInteractor
	// Helper                                   *WebserviceHelper
}

func (handler WebserviceHandler) AppGet(config ConfigurationManager, res http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello Algolia !")

	fmt.Println(handler.AppInteractor.GetApps("test"))

	res.WriteHeader(200)
}
