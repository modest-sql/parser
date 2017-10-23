TARGET=parser
PARSER_SRC=parser.go
LEXER_SRC=lexer.go

.PHONY: clean

$(TARGET): $(PARSER_SRC) $(LEXER_SRC)
	go build -o $(TARGET)

$(PARSER_SRC): parser.y
	goyacc -o $@ $<

$(LEXER_SRC): parser.nex
	nex -o $@ $<

run: $(TARGET)
	./$< < input.txt

clean:
	rm -f $(TARGET) $(PARSER_SRC) $(LEXER_SRC) *.output