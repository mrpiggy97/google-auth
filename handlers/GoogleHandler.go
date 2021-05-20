package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GoogleHandler(writer http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	//writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	defer req.Body.Close()
	var response map[string][]int = map[string][]int{
		"numbers": {1, 2, 3, 3, 4},
	}

	responseJson, _ := json.Marshal(&response)
	fmt.Println(req.Header.Get("Origin"))
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(responseJson)
	fmt.Println("it gets here")
}
