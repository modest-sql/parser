package parser

import "fmt"

type statement interface {
	execute() error
}

type statementList []statement

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

func (s *createStatement) execute() error {
	fmt.Println("TO DO: Create statement execution")
	fmt.Printf("%+v\n", *(s.columnDefinitions[0]))
	return nil
}

type dropStatement struct {
	identifier string
}

func (s *dropStatement) execute() error {
	fmt.Println("TO DO: Drop statement execution")
	return nil
}

type insertStatement struct {
	table       string
	columnNames []string
	values      []interface{}
}

func (s *insertStatement) execute() error {
	fmt.Println("TO DO: Insert statement execution")
	return nil
}

type updateStatement struct {
	table           string
	assignments     []assignment
	whereExpression expression
}

func (s *updateStatement) execute() error {
	fmt.Println("TO DO: Update statement execution")
	return nil
}

type deleteStatement struct {
	table           string
	alias           string
	whereExpression expression
}

func (s *deleteStatement) execute() error {
	fmt.Println("TO DO: Delete statement execution")
	return nil
}

type selectStatement struct {
	table           string
	alias           string
	selectColumns   []interface{}
	whereExpression expression
}

func (s *selectStatement) execute() error {
	fmt.Println("TO DO: Select statement execution")
	return nil
}
