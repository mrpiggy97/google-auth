package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/google-auth/loggers"
	"github.com/mrpiggy97/google-auth/middleware"
)

func HomeHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loggers.HandlerLogger("HomeHandler", request)
	var requestContext context.Context = request.Context()
	var uuidKey middleware.UUIDKeyType = "uuid"
	var messageKey middleware.MessageKeyType = "message"
	var data map[string]interface{} = map[string]interface{}{
		"message": requestContext.Value(messageKey),
		"uuid":    requestContext.Value(uuidKey),
	}
	jsonData, _ := json.Marshal(data)
	var statusCode int = http.StatusOK
	response.WriteHeader(statusCode)
	response.Write(jsonData)
}
