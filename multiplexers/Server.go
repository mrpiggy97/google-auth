package multiplexers

import (
	"fmt"
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
	AllowedAppHost     string
}

func (serverInstance *Server) Throttle(writer http.ResponseWriter, req *http.Request) {
	//this function will thottle, add cors headers with the allowed methods the app
	//will handle and the cross site origin allowed, and
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
		var currentHost string = fmt.Sprintf("this is the host being used %v", req.Host)
		fmt.Println(currentHost)
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
	var allowedHost string = "localhost:8080"
	var muxServer Server = Server{
		Router:             router,
		AllowedMethods:     allowedMethods,
		AllowedCrossOrigin: allowedOrigin,
		AllowedAppHost:     allowedHost,
	}

	var muxServerPointer *Server = &muxServer
	return muxServerPointer
}
