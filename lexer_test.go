package parser

import (
	"os"
	"testing"
)

//DELETE THESE CONSTANTS!!!!
const (
	INT_LIT     = 0
	FLOAT_LIT   = 0
	STR_LIT     = 0
	KW_ON       = 0
	KW_GROUP    = 0
	KW_BY       = 0
	KW_AS       = 0
	KW_FLOAT    = 0
	KW_BOOLEAN  = 0
	KW_DATETIME = 0
)

func TestLexer(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()

	tokensFile, err := os.Open("samples/tokens.sql")
	if err != nil {
		t.Fatal(err)
	}

	lexer := NewLexer(tokensFile)
	lval := &yySymType{}

	tokenCount := 0
	expectedTokenCount := 62

	var integers []int
	var floats []float64
	var strings []string
	var identifiers []string
	var keywords []int

	for token := lexer.Lex(lval); token != 0; token = lexer.Lex(lval) {
		tokenCount++

		if token == INT_LIT {
			integers = append(integers, lval.int_t)
		} else if token == FLOAT_LIT {
			floats = append(floats, lval.float_t)
		} else if token == STR_LIT {
			strings = append(strings, lval.string_t)
		} else if token == TK_ID {
			identifiers = append(identifiers, lval.string_t)
		} else if token >= KW_OR && token <= KW_FALSE {
			keywords = append(keywords, token)
		}

		lval = &yySymType{}
	}

	if tokenCount != expectedTokenCount {
		t.Errorf("Lexical analysis was incorrect, got: %d tokens, want: %d tokens", tokenCount, expectedTokenCount)
	}

	t.Run("TestIntegerLiterals", func(t *testing.T) {
		expectedValue := 27819
		expectedIntCount := 1
		intCount := len(integers)

		if intCount != expectedIntCount {
			t.Fatalf("Expected 1 integer literal, got %d", intCount)
		}

		if integers[0] != expectedValue {
			t.Errorf("Expected %d, got %d", expectedValue, integers[0])
		}
	})

	t.Run("TestFloatLiterals", func(t *testing.T) {
		expectedValue := 3.1416
		expectedFloatCount := 1
		floatCount := len(floats)

		if floatCount != expectedFloatCount {
			t.Fatalf("Expected 1 float literal, got %d", floatCount)
		}

		if floats[0] != expectedValue {
			t.Errorf("Expected %f, got %f", expectedValue, floats[0])
		}
	})

	t.Run("TestStringLiterals", func(t *testing.T) {
		expectedStringCount := 2
		expectedFirstValue, expectedSecondValue := "Hello", "Multi line\nstring! Why not?"
		stringCount := len(strings)

		if stringCount != expectedStringCount {
			t.Fatalf("Expected 2 string literals, got %d", stringCount)
		}

		if strings[0] != expectedFirstValue {
			t.Errorf("Expected `%s', got `%s'", expectedFirstValue, strings[0])
		}

		if strings[0] != expectedSecondValue {
			t.Errorf("Expected `%s', got `%s'", expectedSecondValue, strings[1])
		}
	})

	t.Run("TestIdentifiers", func(t *testing.T) {
		expectedIdentifierCount := 4
		expectedIdentifiers := []string{"movies", "movies", "title", "movies2"}
		identifierCount := len(identifiers)

		if identifierCount != expectedIdentifierCount {
			t.Fatalf("Expected 4 identifiers, got %d", identifierCount)
		}

		for index, identifier := range identifiers {
			if identifier != expectedIdentifiers[index] {
				t.Errorf("Expected `%s', got `%s'", expectedIdentifiers[index], identifier)
			}
		}
	})

	t.Run("TestKeywords", func(t *testing.T) {
		expectedKeywords := []int{
			KW_BETWEEN,
			KW_LIKE,
			KW_IS,
			KW_AND,
			KW_OR,
			KW_NULL,
			KW_TRUE,
			KW_FALSE,
			KW_SELECT,
			KW_FROM,
			KW_INNER,
			KW_JOIN,
			KW_ON,
			KW_WHERE,
			KW_GROUP,
			KW_BY,
			KW_NOT,
			KW_IN,
			KW_HAVING,
			KW_AS,
			KW_INTO,
			KW_SET,
			KW_SUM,
			KW_COUNT,
			KW_AVG,
			KW_MIN,
			KW_MAX,
			KW_CREATE,
			KW_ALTER,
			KW_DROP,
			KW_INSERT,
			KW_UPDATE,
			KW_DELETE,
			KW_INTEGER,
			KW_FLOAT,
			KW_BOOLEAN,
			KW_CHAR,
			KW_DATETIME,
			KW_DEFAULT,
			KW_AUTO_INCREMENT,
		}

		expectedKeywordCount := len(expectedKeywords)
		keywordCount := len(keywords)

		if keywordCount != expectedKeywordCount {
			t.Fatalf("Expected %d keywords, got %d", expectedKeywordCount, keywordCount)
		}

		for index, keyword := range keywords {
			if keyword != expectedKeywords[index] {
				t.Errorf("Expected token %d, got %d", expectedKeywords[index], keyword)
			}
		}
	})
}
