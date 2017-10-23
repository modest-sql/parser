package parser

import "fmt"

//Error provides information about errors that occurred during parsing or execution.
type Error struct {
	line    int
	column  int
	message string
}

//Line returns the line number in which the error occurred.
func (e Error) Line() int {
	return e.line
}

//Column returns the column number in which the error occurred.
func (e Error) Column() int {
	return e.column
}

//Message returns a description of the error that occurred.
func (e Error) Message() string {
	return e.message
}

func (e Error) Error() string {
	return fmt.Sprintf("%s at Line %d, Column %d", e.message, e.line, e.column)
}
