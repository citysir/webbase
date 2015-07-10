package main

import (
	"log"
	"os"
	"web/app"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("not enough args.")
	}

	app.Run(os.Args[1], os.Args[2])
}
