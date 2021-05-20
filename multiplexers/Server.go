package multiplexers

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/time/rate"
)

var limiter *rate.Limiter = rate.NewLimiter(1, 3)

type Server struct {
	Router             *httprouter.Router
	AllowedMethods     [5]string
	AllowedCrossOrigin string
}

func (serverInstance *Server) Throttle(writer http.ResponseWriter, req *http.Request) {
	//this function will thottle and add cors headers with the allowed methods the app
	//will handle and the cross site origin allowed
	if !limiter.Allow() {

		var messageCode int = http.StatusTooManyRequests
		var message string = http.StatusText(messageCode)
		http.Error(writer, message, messageCode)
	} else {
		var allowedMethods []string = serverInstance.AllowedMethods[0:]
		var allowedMethodString string = strings.Join(allowedMethods, ",")
		writer.Header().Set("Access-Content-Allow-Methods", allowedMethodString)
		writer.Header().Set("Access-Content-Allow-Origin", serverInstance.AllowedCrossOrigin)
		serverInstance.Router.ServeHTTP(writer, req)
	}
}

func (serverInstance *Server) ServeAndThrottle() http.HandlerFunc {
	//return a handler
	return http.HandlerFunc(serverInstance.Throttle)
}

func NewServer() *Server {
	var router *httprouter.Router = httprouter.New()
	var allowedMethods [5]string = [5]string{"GET"}
	var allowedOrigin string = "http://localhost:3000"
	var muxServer Server = Server{
		Router:             router,
		AllowedMethods:     allowedMethods,
		AllowedCrossOrigin: allowedOrigin,
	}

	var muxServerPointer *Server = &muxServer
	return muxServerPointer
}
