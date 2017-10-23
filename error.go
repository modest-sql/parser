package parser

import "fmt"

type error struct {
	line    int
	column  int
	message string
}

func (e *error) Line() int {
	return e.line
}

func (e *error) Column() int {
	return e.column
}

func (e *error) Message() string {
	return e.message
}

func (e *error) Error() string {
	return fmt.Sprintf("%s at Line %d, Column %d", e.message, e.line, e.column)
}
