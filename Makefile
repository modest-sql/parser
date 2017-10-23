TARGET=parser
PARSER_SRC=parser.go
LEXER_SRC=lexer.go

.PHONY: clean

$(TARGET): $(PARSER_SRC) $(LEXER_SRC) ast.go parser_error.go
	go build -o $(TARGET)

$(PARSER_SRC): parser.y
	goyacc -o $@ $<

$(LEXER_SRC): parser.nex
	nex -e -o $@ $<

run: $(TARGET)
	./$< < input.txt

clean:
	rm -f $(TARGET) $(PARSER_SRC) $(LEXER_SRC) *.output