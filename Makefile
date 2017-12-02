TARGET=parser
PARSER_SRC=parser.go
LEXER_SRC=lexer.go

.PHONY: clean

$(TARGET): $(PARSER_SRC) $(LEXER_SRC) helpers.go types.go expressions.go statements.go error.go
	go build -o $(TARGET)

$(PARSER_SRC): parser.y
	goyacc -o $@ $<

$(LEXER_SRC): parser.nex
	nex -e -o $@ $<

clean:
	rm -f $(TARGET) $(PARSER_SRC) $(LEXER_SRC) *.output