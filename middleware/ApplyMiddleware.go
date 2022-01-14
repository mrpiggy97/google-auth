package middleware

import "net/http"

type middlewareFunc func(request *http.Request) *http.Request

func ApplyMiddleware(request *http.Request) *http.Request {
	//every middleware function will modify request and return it
	var middlewareFunctions []middlewareFunc = []middlewareFunc{
		MessageMiddleware,
		UUIDMiddleware,
	}
	for _, middleware := range middlewareFunctions {
		request = middleware(request)
	}

	return request
}
