package parser

type node interface {
	line() int
	column() int
}

type columnDefinition struct {
	identifier        string
	nullable          bool
	autoincrementable bool
	defaultValue      interface{}
	dataType
}
type whereClause struct {
}

type starSelectColumn struct {
}

type identifierSelectColumn struct {
}
