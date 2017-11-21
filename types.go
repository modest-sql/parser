package parser

const (
	booleanSize  = 1
	charSize     = 1
	integerSize  = 8
	floatSize    = 8
	datetimeSize = 8
)

type dataType interface {
	size() int
}

type booleanType struct {
}

func (t *booleanType) size() int {
	return booleanSize
}

type charType struct {
	length int64
}

func (t *charType) size() int {
	return int(t.length) * charSize
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

type datetimeType struct {
}

func (t *datetimeType) size() int {
	return datetimeSize
}
