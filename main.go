package main

import (
	"log"
	"os"
	"simple-cli/app"
)

func main() {
	app := app.Generate()

	error := app.Run(os.Args)
	if error != nil {
		log.Fatal(error)
	}
}
