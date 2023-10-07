package main

import (
	"fmt"
	"starter/bookapp"
	"starter/connectmongo"
	"starter/server"
)

func main() {
	var takeInput string

	fmt.Printf("Which app would you like to run? (bookapp, simple_server, templated_server, static_server, momngo_users): ")

	fmt.Scanln(&takeInput)

	switch takeInput {
	case "bookapp":
		fmt.Println("Running bookapp...")
		bookapp.RunBookApp()
	case "simple_server":
		fmt.Println("Running server...")
		server.RunSimpleHelloWorldServer()

	case "templated_server":
		fmt.Println("Running template server...")
		server.RunWithTemplates()

	case "static_server":
		fmt.Println("Running static server...")
		server.RunStaticServer()

	case "mongo_users":
		fmt.Println("Running mongo users...")
		connectmongo.MongoConnect()
	default:
		fmt.Println("No app selected.")
		fmt.Println("Exiting...")
	}
}
