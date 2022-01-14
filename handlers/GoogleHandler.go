package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/google-auth/loggers"
)

type HomeController struct {
	handler httprouter.Handle
}

func GoogleHandler(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	loggers.HandlerLogger("GoogleHandler", req)

	//this will be the response for the client
	var response map[string][]int = map[string][]int{
		"numbers": {1, 2, 3, 3, 4},
	}

	jsonDataFromBody, _ := io.ReadAll(req.Body)

	var data map[string]string
	var err error = json.Unmarshal(jsonDataFromBody, &data)

	if err != nil {
		panic("error during unmarshaling of jsonDataFromBody")
	} else {
		fmt.Println(data)
	}

	var querys url.Values = req.URL.Query()
	var name string = querys.Get("name")
	var code string = querys.Get("code")

	var queryParams string = fmt.Sprintf("%v %v", name, code)
	fmt.Println(queryParams)

	fmt.Println(data["message"])

	responseJson, _ := json.Marshal(&response)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(responseJson)
}

func NewHomeController() HomeController {
	var controller HomeController = HomeController{handler: GoogleHandler}
	return controller
}
