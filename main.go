package main

import (
	"final-project-3/handlers"
)

const port = ":8000"

func main() {
	r := handlers.StartApp()

	r.Run(port)
}
