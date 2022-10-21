package main

import (
	"finder-rest-api/app/route"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	port := os.Getenv("PORT")
	r := route.New()

	fmt.Printf("Starting server on %v\n", port)
	http.ListenAndServe(port, r.Router)

}
