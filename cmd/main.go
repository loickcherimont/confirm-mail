package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/loickcherimont/confirm-mail/internal/handlers"
)

func init() { // To load sender data
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	const port = "3000"

	http.HandleFunc("/signup", handlers.Signup)

	fmt.Printf(`
Server ready and listening on port %s
➜  Try:   http://localhost:%s/signup
➜  Press [Ctrl + C] to stop it`, port, port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
