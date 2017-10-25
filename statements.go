package parser

type statement interface {
	execute() error
}

type createStatement struct {
	identifier        string
	columnDefinitions []columnDefinition
}

type dropStatement struct {
	identifier string
}

type insertStatement struct {
	table       string
	columnNames []string
	values      []interface{}
}

type updateStatement struct {
	table           string
	assignments     []assignment
	whereExpression expression
}

type deleteStatement struct {
	table           string
	alias           string
	whereExpression expression
}

type selectStatement struct {
	table           string
	alias           string
	selectColumns   []interface{}
	whereExpression expression
}
