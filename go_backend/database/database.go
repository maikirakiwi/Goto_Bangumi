package database

import (
	"os"
	"time"

	"github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

// Global db conn

var Conn *clover.DB
var Cache *cache.Cache

func firstRun() {
	err := os.Mkdir("./data", 0755)
	if err != nil {
		log.Fatal().Msgf("Error creating data directory: %s", err.Error())
	}

	db, err := clover.Open("./data")
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	// Default user is admin/adminadmin
	db.CreateCollection("users")

	commit := document.NewDocument()
	commit.Set("username", "admin")
	password, _ := bcrypt.GenerateFromPassword([]byte("adminadmin"), bcrypt.DefaultCost)
	commit.Set("password", password)
	db.InsertOne("users", commit)

	// Default collections
	db.CreateCollection("sessions")
	db.CreateCollection("torrents")
	db.CreateCollection("rss")
	db.Close()

}

func Init() {
	if _, err := os.Stat("./data"); err != nil {
		firstRun()
	}

	var err error
	Conn, err = clover.Open("./data")
	if err != nil {
		log.Fatal().Msgf("Error opening database: %s", err.Error())
	}

	// Have to be ran every time
	Cache = cache.New(7*24*time.Hour, 10*time.Minute)
	Cache.SetDefault("activeUser", "")
}

func Teardown() {
	Conn.Close()
	Cache.Flush()
}

func FindOne(collection string, field string, equ string) (*document.Document, error) {
	doc, err := Conn.FindFirst(query.NewQuery(collection).Where(query.Field(field).Eq(equ)))
	if err != nil {
		return nil, err
	} else {
		return doc, nil
	}
}

func InsertOne(collection string, field string, value string, ttl time.Duration) (string, error) {
	doc := document.NewDocument()
	doc.Set(field, value)
	// ttl = -1 means no ttl
	if ttl != -1 {
		doc.SetExpiresAt(time.Now().Add(ttl))
	}
	return Conn.InsertOne(collection, doc)
}
