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

type selectStatement struct {
	identifier string
	alias      string
	whereClause
}
