package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	 _ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", "postgresql://root:root@localhost:5432/cv?sslmode=disable")

	if err != nil {
		fmt.Println("Failed to connect Database")
	}

	testQueries = New(testDB)
	os.Exit(m.Run())	

}
