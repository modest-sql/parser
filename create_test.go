package parser

import (
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	sqlFile, err := os.Open("samples/create.sql")
	if err != nil {
		t.Fatal(err)
	}

	if err := Parse(sqlFile); err != nil {
		t.Error(err)
	}
}
