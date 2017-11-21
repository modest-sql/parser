package parser_test

import (
	"os"
	"testing"

	"github.com/modest-sql/parser"
)

func TestAlterAdd(t *testing.T) {
	sqlFile, err := os.Open("samples/alter_add.sql")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := parser.Parse(sqlFile); err != nil {
		t.Error(err)
	}
}

func TestAlterDrop(t *testing.T) {
	sqlFile, err := os.Open("samples/alter_drop.sql")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := parser.Parse(sqlFile); err != nil {
		t.Error(err)
	}
}

func TestAlterMod(t *testing.T) {
	sqlFile, err := os.Open("samples/alter_mod.sql")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := parser.Parse(sqlFile); err != nil {
		t.Error(err)
	}
}