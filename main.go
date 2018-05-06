package main

import (
	"log"
	"os"
)

func main() {
	// configure cli app
	app := configureApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
