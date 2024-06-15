// we are gonna create a package to handle the db connection
package db

import (
	"database/sql" // standard package for SQL database
	"fmt"          // used for formatting and connect string
	"log"

	"github.com/YasirHasn9/api-in-go/config" // importing the config package I built

	// Go uses _ before importing a package solely
	_ "github.com/lib/pq" // postgresSQL driver
)

// * Note
// when the function in Go starts with capital letter, then it means it can be called from different place
//  (F)unctionName in capital means it exported

func InitDb(cfg *config.Config) *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Failed to open database: %q", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping database %q", err)
	}

	log.Println("Successfully connected to the database!")
	return db
}
