package parser

type expression interface {
	evaluate() (interface{}, error)
}
