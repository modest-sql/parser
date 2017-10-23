package parser

type node interface {
	line() int
	column() int
}

type expression interface {
	evaluate() (interface{}, *error)
}

type statement interface {
	execute() *error
}

type tableElement struct {
	identifier string
}

type createStatement struct {
	identifier    string
	tableElements []tableElement
}
