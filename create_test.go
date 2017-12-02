package parser_test

import (
	"os"
	"testing"

	"github.com/modest-sql/parser"
)

func TestCreate(t *testing.T) {
	sqlFile, err := os.Open("samples/create.sql")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := parser.Parse(sqlFile); err != nil {
		t.Error(err)
	}
}

func TestCreateWithConstraints(t *testing.T) {
	sqlFile, err := os.Open("samples/create_with_constraints.sql")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := parser.Parse(sqlFile); err != nil {
		t.Error(err)
	}
}
