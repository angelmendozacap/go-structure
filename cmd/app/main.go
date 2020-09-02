package main

import "log"

func main() {
	e := initRoutes()

	log.Fatal(e.Start(":3030"))
}
