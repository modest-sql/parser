package parser

import (
	"os"
	"testing"
)

func TestUpdate(t *testing.T) {
	sqlFile, err := os.Open("samples/update.sql")
	if err != nil {
		t.Fatal(err)
	}

	if err := Parse(sqlFile); err != nil {
		t.Error(err)
	}
}
