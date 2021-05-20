package main

import (
	"fmt"
	"net/http"

	"github.com/mrpiggy97/google-auth/handlers"
	"github.com/mrpiggy97/google-auth/multiplexers"
)

func main() {
	var basicServer *multiplexers.Server = multiplexers.NewServer()
	basicServer.Router.GET("/", handlers.GoogleHandler)
	fmt.Println("listening at port 8080")
	http.ListenAndServe("0.0.0.0:8080", basicServer.ServeAndThrottle())
}
