package main

import (
	"github.com/joho/godotenv"
	"github.com/mrpiggy97/google-auth/multiplexers"
)

func main() {
	var err error = godotenv.Load()

	if err != nil {
		panic("loading of environment variables failed to load")
	}
	multiplexers.Runserver(8080)
}
