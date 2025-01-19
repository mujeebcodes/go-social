package main

import (
	"github.com/mujeebcodes/go-social/internal/db"
	"github.com/mujeebcodes/go-social/internal/env"
	"github.com/mujeebcodes/go-social/internal/store"
	"log"
)

func main() {
	addr := env.GetString("DB_ADDR", "")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
