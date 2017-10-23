package parser

import "fmt"

type parserError struct {
	line    int
	column  int
	message string
}

func (e parserError) Line() int {
	return e.line
}

func (e parserError) Column() int {
	return e.column
}

func (e parserError) Message() string {
	return e.message
}

func (e parserError) Error() string {
	return fmt.Sprintf("%s at Line %d, Column %d", e.message, e.line, e.column)
}
