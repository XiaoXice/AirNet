package database

import (
	"fmt"
	"testing"
)

type testLogger struct {
}

func (t *testLogger) Handler(msg string) bool {
	fmt.Print(msg)
	return true
}

var DB = DataBase{
	Type:   "sqlite3",
	Path:   "./testDB.db",
	Logger: &testLogger{},
}

func TestDataBase_Connect(t *testing.T) {
	ok := DB.Connect()
	if ok != true {
		t.Error("Database init error")
	}
}

func TestDataBase_Select(t *testing.T) {
	if !DB.CanUse() {
		TestDataBase_Connect(t)
	}
	//rows, rowCount, err := DB.Select(NodeListNew{},"SELECT * FROM NodeList;")
	//if err != nil && len(rows) != rowCount {
	//	t.Error("Something error.")
	//}
}
