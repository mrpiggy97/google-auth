package multiplexers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mrpiggy97/google-auth/handlers"
)

func Runserver(port int) {
	var currentTime time.Time = time.Now()
	var address string = fmt.Sprintf("0.0.0.0:%v", port)
	var message1 string = fmt.Sprintf("server started at %v", currentTime)
	var message2 string = fmt.Sprintf("server listening at address %v", address)
	fmt.Println(message1)
	fmt.Println(message2)
	var basicServer *Server = NewServer()
	basicServer.Router.POST("/test-json", handlers.GoogleHandler)
	basicServer.Router.GET("/", handlers.HomeHandler)
	http.ListenAndServe(address, basicServer)
}
