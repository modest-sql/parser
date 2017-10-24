package parser

const (
	booleanSize = 1
	charSize    = 1
	integerSize = 4
	floatSize   = 8
)

type node interface {
	line() int
	column() int
}

type expression interface {
	evaluate() (interface{}, error)
}

type statement interface {
	execute() error
}

type dataType interface {
	size() int
}

type booleanType struct {
}

func (t *booleanType) size() int {
	return booleanSize
}

type charType struct {
	length int
}

func (t *charType) size() int {
	return t.length * charSize
}

type integerType struct {
}

func (t *integerType) size() int {
	return integerSize
}

type floatType struct {
}

func (t *floatType) size() int {
	return floatSize
}

type columnDefinition struct {
	identifier string
	dataType   dataType
}

type createStatement struct {
	identifier        string
	columnDefinitions []columnDefinition
}
