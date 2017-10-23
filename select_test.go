package parser

import (
	"os"
	"testing"
)

func TestSelect(t *testing.T) {
	sqlFile, err := os.Open("samples/select.sql")
	if err != nil {
		t.Fatal(err)
	}

	if err := Parse(sqlFile); err != nil {
		t.Error(err)
	}
}
