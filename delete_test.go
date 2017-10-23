package parser_test

import (
	"os"
	"testing"

	"github.com/modest-sql/parser"
)

func TestDelete(t *testing.T) {
	sqlFile, err := os.Open("samples/delete.sql")
	if err != nil {
		t.Fatal(err)
	}

	if err := parser.Parse(sqlFile); err != nil {
		t.Error(err)
	}
}
