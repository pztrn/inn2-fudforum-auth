package main

import (
	// stdlib
	"bufio"
	"crypto/sha1"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	// local
	"develop.pztrn.name/pztrn/inn2-fudforum-auth/configuration"
	"develop.pztrn.name/pztrn/inn2-fudforum-auth/database"
)

// This structure holds neccessary data after parsing stdin or passed
// by flags.
type passedParameters struct {
	Username string
	Password string
}

// This structure holds data for database requests.
type databaseParams struct {
	Username string `db:"login"`
	Password string `db:"passwd"`
	Salt     string `db:"salt"`
}

func main() {
	configuration.Initialize()

	params := &passedParameters{}

	flag.StringVar(&params.Username, "user", "", "Username to authenticate. You can use this flag for debugging.")
	flag.StringVar(&params.Username, "password", "", "Password to use. You can use this flag for debugging.")

	flag.Parse()

	configuration.Cfg.Initialize()
	database.Initialize()

	if configuration.Cfg.Debug {
		log.Println("Starting inn2-fudforum-auth in debug mode...")
	}

	// Check our running mode. If params structure isn't filled - then
	// we should read from stdin.
	if params.Username != "" || params.Password != "" {
		if configuration.Cfg.Debug {
			log.Println("-user or -password passed, will use these parameters values for authentication.")
		}
		// We should have both fields filled.
		if params.Username == "" {
			log.Fatalln("You should provide -user parameter.")
		}

		if params.Password == "" {
			log.Fatalln("You should provide -password parameter.")
		}
	} else {
		if configuration.Cfg.Debug {
			log.Println("-user or -password WASN'T passed, will use stdin for authentication.")
		}

		input := bufio.NewScanner(os.Stdin)

		for {
			input.Scan()
			// That means inn2 stopped sending data to stdin due to
			// error or timeout.
			if err := input.Err(); err != nil {
				break
			}

			// If we gathered all needed data - stop iterating.
			if params.Username != "" && params.Password != "" {
				break
			}

			inputDataRaw := input.Text()
			if strings.Contains(inputDataRaw, "ClientAuthname: ") {
				inputData := strings.Split(inputDataRaw, ": ")
				if len(inputData) == 1 {
					log.Fatalln("Empty auth name (login) passed.")
				}
				params.Username = inputData[1]
				if configuration.Cfg.Debug {
					log.Println("Username gathered: " + params.Username)
				}
			}

			if strings.Contains(inputDataRaw, "ClientPassword: ") {
				inputData := strings.Split(inputDataRaw, ": ")
				if len(inputData) == 1 {
					log.Fatalln("Empty password passed.")
				}
				params.Password = inputData[1]
				if configuration.Cfg.Debug {
					log.Println("Password gathered: " + params.Password)
				}
			}
		}
	}

	if configuration.Cfg.Debug {
		log.Printf("Got authentication data: %+v\n", params)
	}

	// Get data from FUDForum's database.
	dbData := &databaseParams{}
	err := database.Conn.Get(dbData, "SELECT login, passwd, salt FROM "+configuration.Cfg.Database.Prefix+"users WHERE login=$1 OR alias=$2", params.Username, params.Username)
	if err != nil {
		log.Fatalln("Failed to get data from FUDForum database: " + err.Error())
	}

	if configuration.Cfg.Debug {
		log.Printf("Got data from FUDForum database: %+v\n", dbData)
	}

	// FUDForum uses sha1(salt + sha1(password)).
	passHashRaw := sha1.New()
	_, _ = passHashRaw.Write([]byte(params.Password))
	passHash := fmt.Sprintf("%x", passHashRaw.Sum(nil))

	saltedPassHashRaw := sha1.New()
	_, _ = saltedPassHashRaw.Write([]byte(dbData.Salt + passHash))
	saltedPassHash := fmt.Sprintf("%x", saltedPassHashRaw.Sum(nil))

	if configuration.Cfg.Debug {
		log.Printf("Password stored in database: %s, we hashed: %s (pre: %s)\n", dbData.Password, saltedPassHash, passHash)
	}

	// Check groups mapping.
	// This is temporary, in future versions all groups memberships
	// should be managed on FUDForum side.
	var group string
	for _, groupMapping := range configuration.Cfg.Groups.Groups {
		var userFound bool
		for _, user := range groupMapping.Users {
			if user == params.Username {
				group = groupMapping.Group
				userFound = true
				break
			}
		}

		if userFound {
			break
		}
	}

	fmt.Printf("User:%s@%s\r\n", params.Username, group)

	database.Shutdown()
}
