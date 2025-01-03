package main

import (
	"github.com/mujeebcodes/go-social/internal/db"
	"github.com/mujeebcodes/go-social/internal/env"
	"github.com/mujeebcodes/go-social/internal/store"
	"log"
)

const version = "0.0.1"

//	@title			Go-Social API
//	@description	API for Go-Social, a social network
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/v1
//
//@securityDefinitions.apikey ApiKeyAuth
//@in 	header
//@name Authorization
//@description

func main() {

	cfg := config{
		addr:   env.GetString("PORT", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),

		db: dbConfig{
			addr:         env.GetString("DB_ADDR", ""),
			maxOpenConns: env.GetInt("MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
