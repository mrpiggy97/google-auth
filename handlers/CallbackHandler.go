package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

func CallbackHandler(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var data map[string]string
	var queries url.Values = req.URL.Query()
	var code string = queries.Get("code")
	data["code"] = code

	jsonData, _ := json.Marshal(&data)

	var statusCode int = http.StatusAccepted
	writer.WriteHeader(statusCode)
	writer.Header().Add("Content-Type", "application/json")
	writer.Write(jsonData)
}
