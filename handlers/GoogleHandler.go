package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/google-auth/loggers"
)

func GoogleHandler(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	loggers.Logger("GoogleHandler", req)
	var response map[string][]int = map[string][]int{
		"numbers": {1, 2, 3, 3, 4},
	}

	responseJson, _ := json.Marshal(&response)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(responseJson)
}
