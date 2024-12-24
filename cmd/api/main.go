package main

import (
	"github.com/mujeebcodes/go-social/internal/env"
	"github.com/mujeebcodes/go-social/internal/store"
	"log"
)

func main() {

	cfg := config{
		addr: env.GetString("ADDRESS", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
