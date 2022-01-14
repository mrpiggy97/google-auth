package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type UUIDKeyType string

func UUIDMiddleware(request *http.Request) *http.Request {
	var UUIDKey UUIDKeyType = "uuid"
	var UUIDValue uuid.UUID = uuid.New()
	var previousContext context.Context = request.Context()
	var newContext context.Context = context.WithValue(previousContext, UUIDKey, UUIDValue)
	request = request.WithContext(newContext)
	return request
}
