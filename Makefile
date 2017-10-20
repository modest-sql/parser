TARGET=parser
PARSER_SRC=y.go
LEXER_SRC=parser.nn.go

.PHONY: clean

$(TARGET): $(PARSER_SRC) $(LEXER_SRC)
	go build -o $(TARGET)

$(PARSER_SRC): parser.y
	goyacc $<

$(LEXER_SRC): parser.nex
	nex $<

run: $(TARGET)
	./$< < input.txt

clean:
	rm -f $(TARGET) $(PARSER_SRC) $(LEXER_SRC) *.output