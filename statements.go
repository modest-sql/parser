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
	return common.NewDropCommand(s.identifier)
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
	values := map[string]interface{}{}
	if len(s.columnNames) != len(s.values) {
		panic(true)
	}
	for i, columnName := range s.columnNames {
		values[columnName] = s.values[i]
	}

	return common.NewInsertCommand(s.table, values)
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
	function interface{}
}

func (s *columnSpec) convert() interface{} {
	return common.NewTableColumnSelector(s.isStar,s.table,s.column,s.alias,s.function)
}
type functionSum struct{

}
type functionCount struct{
	
}
type functionAvg struct{
	
}
type functionMin struct{
	
}
type functionMax struct{
	
}
type GroupBySpec struct{
	table string
	column  string
}

func (s *GroupBySpec) convert() interface{} {
	return common.NewGroupBySelect(s.table,s.column)
}
type selectStatement struct {
	selectColumns   []columnSpec
	mainTable string
	mainAlias string
	joinList []joinSpec
	whereExpression expression
	groupBy  []GroupBySpec
}


type joinSpec struct {
	targetTable string
	targetAlias string
	filterCriteria expression
}

func (s *joinSpec) convert() interface{} {
	return common.NewJoinSelect(s.targetTable,s.targetAlias,s.filterCriteria.convert())
}
func (s *selectStatement) convert() interface{} {
	var tablecolumm common.TableColumnSelectors
	for _, column := range s.selectColumns {
		tablecolumm = append(tablecolumm,column.convert())
	}

	var tablejoin []common.JoinSelect
	for _, join := range s.joinList {
		s := join.convert().(common.JoinSelect)
		tablejoin = append(tablejoin,s)
	}

	var tablegroupBy []common.GroupBySelect
	for _, group := range s.groupBy {
		s := group.convert().(common.GroupBySelect)
		tablegroupBy = append(tablegroupBy,s)
	}
	return common.NewSelectTableCommand(s.mainTable,s.mainAlias,tablecolumm,tablejoin,s.whereExpression.convert(),tablegroupBy)
}

func (s *selectStatement) execute() error {
	return nil
}

type alterStatement struct {
	table       string
	instruction interface{}
}

func (s *alterStatement) convert() interface{} {
	switch v := s.instruction.(type) {
	case *alterDrop:
		return common.NewAlterCommand(s.table, v.convert())
	case *alterAdd:
		return common.NewAlterCommand(s.table, v.convert())
	case *alterModify:
		return common.NewAlterCommand(s.table, v.convert())
	}
	return common.NewAlterCommand(s.table, nil)
}
func (s *alterStatement) execute() error {
	return nil
}

type alterDrop struct {
	table string
}

func (s *alterDrop) convert() interface{} {
	return common.NewAlterDropInst(s.table)
}
func (s *alterDrop) execute() error {
	return nil
}

type alterAdd struct {
	table             string
	dataType          dataType
	columnConstraints []interface{}
}

func (s *alterAdd) convert() interface{} {
	column := columnDefinition{s.table, s.dataType, s.columnConstraints}
	return common.NewAlterAddInst(column.convert())
}
func (s *alterAdd) execute() error {
	return nil
}

type alterModify struct {
	table             string
	dataType          dataType
	columnConstraints []interface{}
}

func (s *alterModify) convert() interface{} {
	column := columnDefinition{s.table, s.dataType, s.columnConstraints}
	return common.NewAlterModifyInst(column.convert())
}
func (s *alterModify) execute() error {
	return nil
}
