package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/loickcherimont/confirm-mail/internal/handlers"
)

func init() { // To load sender data
	err := godotenv.Load(".env")

	fmt.Println(os.Getenv("SENDER_ADDRESS"))
	fmt.Println(os.Getenv("SENDER_PASSWORD"))

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	var port = "3000"

	http.HandleFunc("/signup", handlers.Signup)

	fmt.Printf("Server is up and listening.\n Try http://localhost:%s/signup.\n Stop the server with  [Ctrl + C]", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
