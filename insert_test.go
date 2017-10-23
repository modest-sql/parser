package parser

import (
	"os"
	"testing"
)

func TestInsert(t *testing.T) {
	sqlFile, err := os.Open("samples/insert.sql")
	if err != nil {
		t.Fatal(err)
	}

	if err := Parse(sqlFile); err != nil {
		t.Error(err)
	}
}
