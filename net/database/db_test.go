package database

import (
	"fmt"
	"testing"
)

type testLogger struct {
}

func (t *testLogger)Handler(msg string) bool {
	fmt.Print(msg)
	return true
}

func TestDataBase_Connect(t *testing.T) {
	DB := DataBase{
		Type: "sqlite3",
		Path: "./testDB.db",
		Logger: &testLogger{},
	}
	ok := DB.Connect()
	if ok != true {
		t.Error("Database init error")
	}
}