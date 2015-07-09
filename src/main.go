package main

import (
	"log"
	"os"
	"web/app"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("no arg 'port'")
	}

	app.Run(os.Args[1])
}
