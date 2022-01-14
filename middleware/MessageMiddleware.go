package middleware

import (
	"context"
	"net/http"
)

type MessageKeyType string

func MessageMiddleware(req *http.Request) *http.Request {
	var messageKey MessageKeyType = "message"
	var messageValue string = "this is the feeing message"
	var previousContext context.Context = req.Context()
	var newContext context.Context = context.WithValue(previousContext, messageKey, messageValue)
	req = req.WithContext(newContext)
	return req
}
