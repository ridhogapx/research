package db

import (
	"testing"
	"context"
)

func TestCreateUser(t *testing.T) {
	err := testQueries.CreateUser(context.Background(), "foo@mai.com")
	if err != nil {
		t.Error(err)
	}
}
