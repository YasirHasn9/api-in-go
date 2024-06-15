// here where we organize the entry points
package main

import (
	"log"

	"github.com/YasirHasn9/api-in-go/config"
	"github.com/YasirHasn9/api-in-go/internal/db"
)

func main() {

	cfg := config.LoadConfig()
	db := db.InitDb(cfg)

	defer db.Close()

	log.Println("Application started successfully!")
}
