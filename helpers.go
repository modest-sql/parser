package parser

type node interface {
	line() int
	column() int
}

type notNullConstraint struct {
}

type autoincrementConstraint struct {
}

type defaultConstraint struct {
	value interface{}
}

type columnDefinitions []*columnDefinition

type columnDefinition struct {
	identifier        string
	dataType          dataType
	columnConstraints []interface{}
}

type starSelectColumn struct {
}

type selectColumn struct {
	identifier string
	alias      string
	source     string
}

type assignment struct {
	identifier string
	value      interface{}
}
