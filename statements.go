package parser

import (
	"github.com/modest-sql/common"
)

type statement interface {
	execute() error
	convert() interface{}
}

type statementList []statement

func (sl statementList) convert() (commands []interface{}) {
	for _, statement := range sl {
		command := statement.convert()

		if command != nil {
			commands = append(commands, command)
		}
	}
	return commands
}

func (sl statementList) execute() error {
	for _, statement := range sl {
		if err := statement.execute(); err != nil {
			panic(err)
		}
	}

	return nil
}

type createStatement struct {
	identifier        string
	columnDefinitions columnDefinitions
}

func (s *createStatement) convert() interface{} {
	return common.NewCreateTableCommand(s.identifier, s.columnDefinitions.convert())
}

func (s *createStatement) execute() error {
	return nil
}

type dropStatement struct {
	identifier string
}

func (s *dropStatement) convert() interface{} {
	return nil
}

func (s *dropStatement) execute() error {
	return nil
}

type insertStatement struct {
	table       string
	columnNames []string
	values      []interface{}
}

func (s *insertStatement) convert() interface{} {
	return nil
}

func (s *insertStatement) execute() error {
	return nil
}

type updateStatement struct {
	table           string
	assignments     []assignment
	whereExpression expression
}

func (s *updateStatement) convert() interface{} {
	return nil
}

func (s *updateStatement) execute() error {
	return nil
}

type deleteStatement struct {
	table           string
	alias           string
	whereExpression expression
}

func (s *deleteStatement) convert() interface{} {
	return nil
}

func (s *deleteStatement) execute() error {
	return nil
}

type columnSpec struct {
	isStar bool
	table  string
	column string
	alias  string
}

type selectStatement struct {
	table string /* discouraged, use mainTable from source instead. Will disappear */
	alias string /* discoraged, use mainAlias from source instead. Will disappear */
	mainTable string
	mainAlias string
	joinList []joinSpec
	selectColumns   []columnSpec
	whereExpression expression
}

type joinSpec struct {
	targetTable string
	targetAlias string
	filterCriteria expression
}
func (s *selectStatement) convert() interface{} {
	return nil
}

func (s *selectStatement) execute() error {
	return nil
}

type alterStatement struct {
	table       string
	instruction interface{}
}

func (s *alterStatement) convert() interface{} {
	return nil
}

func (s *alterStatement) execute() error {
	return nil
}

type alterDrop struct {
	table string
}
type alterAdd struct {
	table             string
	dataType          dataType
	columnConstraints []interface{}
}
