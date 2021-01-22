// main.go

package main

import (
	"fmt"
	"os"

	"github.com/godemo/context"
)

func main() {
	a := context.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	port := os.Getenv("PORT")

	// Default port in development
	if port == "" {
		port = "5000"
	}
	fmt.Printf("Server started listening at http://localhost:%s \n", port)
	a.Run(":" + port)
}
