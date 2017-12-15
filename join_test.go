package parser_test

import (
	"os"
	"testing"

	"github.com/modest-sql/parser"
)

func TestJoinWithoutAlias(t *testing.T) {
	sqlFile, err := os.Open("samples/inner_join.sql")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := parser.Parse(sqlFile); err != nil {
		t.Error(err)
	}
}