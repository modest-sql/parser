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
