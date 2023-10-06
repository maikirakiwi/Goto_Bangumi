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

	"Auto_Bangumi/v2/models"
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
	db.CreateCollection("config")
	db.CreateCollection("sessions")
	db.CreateCollection("bangumi")
	db.CreateCollection("torrents")
	db.CreateCollection("rss")

	cfg := document.NewDocument()
	cfg.Set("backend", "go")
	cfg.Set("content", models.InitConfigModel())
	db.InsertOne("config", cfg)

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

	cfg, err := Conn.FindFirst(query.NewQuery("config").Where(query.Field("backend").Eq("go")))
	if err != nil {
		log.Fatal().Msgf("Error getting config [Init]: %s", err.Error())
	}
	Cache.SetDefault("config", cfg.Get("content"))
}

func Teardown() {
	Conn.Close()
	Cache.Flush()
}

func FindOne(collection string, key string, value string) (*document.Document, error) {
	doc, err := Conn.FindFirst(query.NewQuery(collection).Where(query.Field(key).Eq(value)))
	if err != nil {
		return nil, err
	} else {
		return doc, nil
	}
}

func UpdateOne(collection string, key string, value string, changingKey string, changingValue interface{}) error {
	return Conn.Update(query.NewQuery(collection).Where(query.Field(key).Eq(value)), map[string]interface{}{changingKey: changingValue})
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
