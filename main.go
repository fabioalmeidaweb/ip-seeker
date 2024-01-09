package main

import (
	"ip-seeker/app"
	"log"
	"os"
)

// main is the entry point of the program.
//
// It generates an application and runs it with the given command-line arguments.
func main() {
	application := app.Generate()
	err := application.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
