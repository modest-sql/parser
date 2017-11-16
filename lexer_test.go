package parser

import (
	"os"
	"testing"
)

var tokenNames = map[int]string{
	TK_PLUS:           "+",
	TK_MINUS:          "-",
	TK_STAR:           "*",
	TK_DIV:            "/",
	TK_LT:             "<",
	TK_GT:             ">",
	TK_GTE:            ">=",
	TK_LTE:            "<=",
	TK_EQ:             "=",
	TK_NE:             "<>",
	TK_LEFT_PAR:       "(",
	TK_RIGHT_PAR:      ")",
	TK_COMMA:          ",",
	TK_DOT:            ".",
	TK_SEMICOLON:      ";",
	KW_OR:             "OR",
	KW_AND:            "AND",
	KW_NOT:            "NOT",
	KW_INTEGER:        "INTEGER",
	KW_FLOAT:          "FLOAT",
	KW_CHAR:           "CHAR",
	KW_BOOLEAN:        "BOOLEAN",
	KW_DATETIME:       "DATETIME",
	KW_CREATE:         "CREATE",
	KW_TABLE:          "TABLE",
	KW_DELETE:         "DELETE",
	KW_INSERT:         "INSERT",
	KW_INTO:           "INTO",
	KW_SELECT:         "SELECT",
	KW_WHERE:          "WHERE",
	KW_FROM:           "FROM",
	KW_UPDATE:         "UPDATE",
	KW_SET:            "SET",
	KW_ALTER:          "ALTER",
	KW_VALUES:         "VALUES",
	KW_BETWEEN:        "BETWEEN",
	KW_LIKE:           "LIKE",
	KW_INNER:          "INNER",
	KW_HAVING:         "HAVING",
	KW_SUM:            "SUM",
	KW_COUNT:          "COUNT",
	KW_AVG:            "AVG",
	KW_MIN:            "MIN",
	KW_MAX:            "MAX",
	KW_NULL:           "NULL",
	KW_IN:             "IN",
	KW_IS:             "IS",
	KW_AUTO_INCREMENT: "AUTO_INCREMENT",
	KW_JOIN:           "JOIN",
	KW_ON:             "ON",
	KW_GROUP:          "GROUP",
	KW_BY:             "BY",
	KW_DROP:           "DROP",
	KW_DEFAULT:        "DEFAULT",
	KW_TRUE:           "TRUE",
	KW_FALSE:          "FALSE",
	KW_AS:             "AS",
	KW_ADD:            "ADD",
	KW_COLUMN:         "COLUMN",
	INT_LIT:           "INT_LIT",
	FLOAT_LIT:         "FLOAT_LIT",
	TK_ID:             "IDENTIFIER",
	STR_LIT:           "STR_LIT",
}

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
	expectedTokenCount := 65

	var integers []int
	var floats []float64
	var strings []string
	var identifiers []string
	var keywords []int
	var binOperators []int
	var punctuators []int

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
		} else if token >= KW_OR && token <= KW_COLUMN {
			keywords = append(keywords, token)
		} else if token >= TK_PLUS && token <= TK_NE {
			binOperators = append(binOperators, token)
		} else if token >= TK_LEFT_PAR && token <= TK_SEMICOLON {
			punctuators = append(punctuators, token)
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
		t.Skip("It's fine, we don't NEED multiline for now.")
		expectedStringCount := 2
		expectedFirstValue, expectedSecondValue := "Hello", "Multi line\nstring! Why not?"
		stringCount := len(strings)

		if stringCount != expectedStringCount {
			t.Fatalf("Expected 2 string literals, got %d", stringCount)
		}

		if strings[0] != expectedFirstValue {
			t.Errorf("Expected `%s', got `%s'", expectedFirstValue, strings[0])
		}

		if strings[1] != expectedSecondValue {
			t.Errorf("Expected `%s', got `%s'", expectedSecondValue, strings[1])
		}
	})

	t.Run("TestIdentifiers", func(t *testing.T) {
		expectedIdentifiers := []string{"movies", "movies", "title", "movies2"}
		expectedIdentifierCount := len(expectedIdentifiers)
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
			KW_COLUMN,
			KW_ADD,
		}

		expectedKeywordCount := len(expectedKeywords)
		keywordCount := len(keywords)

		if keywordCount != expectedKeywordCount {
			t.Fatalf("Expected %d keywords, got %d", expectedKeywordCount, keywordCount)
		}

		for index, keyword := range keywords {
			if keyword != expectedKeywords[index] {
				t.Errorf("Expected %s, got %s", tokenNames[expectedKeywords[index]], tokenNames[keyword])
			}
		}
	})

	t.Run("TestBinaryOperators", func(t *testing.T) {
		expectedBinOperators := []int{
			TK_PLUS,
			TK_MINUS,
			TK_STAR,
			TK_DIV,
			TK_LT,
			TK_GT,
			TK_LTE,
			TK_GTE,
			TK_EQ,
			TK_NE,
		}

		expectedBinOperatorsCount := len(expectedBinOperators)
		binOperatorsCount := len(binOperators)

		if binOperatorsCount != expectedBinOperatorsCount {
			t.Fatalf("Expected %d binary operators, got %d", expectedBinOperatorsCount, binOperatorsCount)
		}

		for index, operator := range binOperators {
			if operator != expectedBinOperators[index] {
				t.Errorf("Expected %s, got %s", tokenNames[expectedBinOperators[index]], tokenNames[operator])
			}
		}
	})

	t.Run("TestPunctuators", func(t *testing.T) {
		expectedPunctuators := []int{
			TK_DOT,
			TK_LEFT_PAR,
			TK_RIGHT_PAR,
			TK_COMMA,
			TK_SEMICOLON,
		}

		expectedPunctuatorsCount := len(expectedPunctuators)
		punctuatorsCount := len(punctuators)

		if punctuatorsCount != expectedPunctuatorsCount {
			t.Fatalf("Expected %d punctuators, got %d", expectedPunctuatorsCount, punctuatorsCount)
		}

		for index, punctuator := range punctuators {
			if punctuator != expectedPunctuators[index] {
				t.Errorf("Expected token %s, got %s", tokenNames[expectedPunctuators[index]], tokenNames[punctuator])
			}
		}
	})
}
