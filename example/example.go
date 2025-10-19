package main

import (
	"fmt"
	"log"
	"os"

	"github.com/soumayg9673/dotenv"
)

func main() {
	// Setting required key
	dotenv.AddRqdKey("HELLO", true)
	dotenv.AddRqdKey("AU", true)

	if err := dotenv.LoadEnvFile("./example/test.env"); err != nil {
		log.Println(err)
	}

	fmt.Println(os.Getenv("AUTHOR"))

	if err := dotenv.ValidateRqdEnv(); err != nil {
		fmt.Println(err)
	}
}
