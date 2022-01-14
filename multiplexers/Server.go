package multiplexers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/google-auth/middleware"
	"golang.org/x/time/rate"
)

var limiter *rate.Limiter = rate.NewLimiter(1, 3)

type Server struct {
	Router             *httprouter.Router
	AllowedMethods     [5]string
	AllowedCrossOrigin string
	AllowedAppHost     string
}

func (serverInstance *Server) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	//this function will thottle and add cors headers
	if !limiter.Allow() {
		var messageCode int = http.StatusTooManyRequests
		var message string = http.StatusText(messageCode)
		http.Error(writer, message, messageCode)

	} else if serverInstance.AllowedAppHost != req.Host {

		fmt.Println("wrong host")
		var messageCode int = http.StatusForbidden
		var message string = http.StatusText(messageCode)
		http.Error(writer, message, messageCode)

	} else {

		var allowedMethods []string = serverInstance.AllowedMethods[0:]
		var allowedMethodString string = strings.Join(allowedMethods, ",")
		writer.Header().Set("Access-Control-Allow-Methods", allowedMethodString)
		writer.Header().Set("Access-Control-Allow-Origin", serverInstance.AllowedCrossOrigin)
		req = middleware.ApplyMiddleware(req)
		serverInstance.Router.ServeHTTP(writer, req)
	}
}

func NewServer() *Server {
	var router *httprouter.Router = httprouter.New()
	var allowedMethods [5]string = [5]string{"GET"}
	var allowedCrossOrigin string = os.Getenv("ALLOWED_CROSS_ORIGIN")
	var allowedAppHost string = os.Getenv("ALLOWED_APP_HOST")

	var muxServer *Server = &Server{
		Router:             router,
		AllowedMethods:     allowedMethods,
		AllowedCrossOrigin: allowedCrossOrigin,
		AllowedAppHost:     allowedAppHost,
	}

	return muxServer
}
