package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

var fs embed.FS

func main() {
	d, err := iofs.New(fs, "data/")
	if err != nil {
		log.Fatal(err)
	}

	m, errD := migrate.NewWithSourceInstance("iofs", d, "postgres://root:root@localhost/cvz_auth?sslmode=disable")

	if errD != nil {
		log.Fatal(err)
	}

	errS := m.Up()

	if errS != nil {
		log.Fatal(errS)
	}

	fmt.Println("Migrating...")
}
