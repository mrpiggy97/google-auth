package loggers

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(handlerName string, request *http.Request) {
	var logDate time.Time = time.Now()
	var origin string = request.Header.Get("Origin")
	var message string = fmt.Sprintf(
		"handler %v called at %v from origin %v",
		handlerName, logDate, origin,
	)
	fmt.Println(message)
}
