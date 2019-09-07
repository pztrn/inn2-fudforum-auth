package database

import (
	// stdlib
	"log"

	// local
	"develop.pztrn.name/pztrn/inn2-fudforum-auth/configuration"

	// other
	"github.com/jmoiron/sqlx"
	// postgres driver
	_ "github.com/lib/pq"
)

var Conn *sqlx.DB

// Initialize initializes package.
func Initialize() {
	if configuration.Cfg.Debug {
		log.Println("Initializing database connection...")
	}

	dsn := configuration.Cfg.Database.DSN
	if configuration.Cfg.Database.Parameters != "" {
		dsn = configuration.Cfg.Database.DSN + "?" + configuration.Cfg.Database.Parameters
	}
	conn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("Can't connect to fudforum database: " + err.Error())
	}

	if configuration.Cfg.Debug {
		log.Println("Database connection established.")
	}

	Conn = conn
}

// Shutdown closes database connection.
func Shutdown() {
	if Conn != nil {
		if configuration.Cfg.Debug {
			log.Println("Closing database connection...")
		}
		err := Conn.Close()
		if err != nil {
			log.Fatalln("Failed to close database connection: " + err.Error())
		} else {
			if configuration.Cfg.Debug {
				log.Println("Database connection closed.")
			}
		}
	}
}
