package parser

import (
	"os"
	"testing"
)

func TestDelete(t *testing.T) {
	sqlFile, err := os.Open("samples/delete.sql")
	if err != nil {
		t.Fatal(err)
	}

	if err := Parse(sqlFile); err != nil {
		t.Error(err)
	}
}
