package database

import (
	"os"

	"github.com/ostafen/clover"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// Global db conn

var Conn *clover.DB

func first_run() {
	if db, err := clover.Open("./data/data.json"); err != nil {
		log.Fatal().Msgf("Error opening database: %s", err.Error())
	} else {
		// Default user is admin/adminadmin
		db.CreateCollection("users")

		commit := clover.NewDocument()
		commit.Set("username", "admin")
		password, _ := bcrypt.GenerateFromPassword([]byte("adminadmin"), bcrypt.DefaultCost)
		commit.Set("password", password)
		db.InsertOne("users", commit)

		// Default collections
		db.CreateCollection("torrents")
		db.CreateCollection("rss")
		db.Close()
	}
}

func Init() {
	if _, err := os.Stat("./data/data.json"); err != nil {
		first_run()
	}

	var err error
	Conn, err = clover.Open("./data/data.json")
	if err != nil {
		log.Fatal().Msgf("Error opening database: %s", err.Error())
	}
}
