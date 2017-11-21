package parser

import (
	"github.com/modest-sql/common"
)

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

func (cd columnDefinitions) convert() (tableColumns common.TableColumnDefiners) {
	for _, columnDefinition := range cd {
		tableColumn := columnDefinition.convert()

		if tableColumn != nil {
			tableColumns = append(tableColumns, tableColumn)
		}
	}
	return tableColumns
}

type columnDefinition struct {
	identifier        string
	dataType          dataType
	columnConstraints []interface{}
}

func (c *columnDefinition) convert() common.TableColumnDefiner {
	switch v := c.dataType.(type) {
	case *booleanType:
		return common.NewBooleanTableColumn(c.identifier, c.defaultValue(), c.nullable(), c.autoincrementable())
	case *integerType:
		return common.NewIntegerTableColumn(c.identifier, c.defaultValue(), c.nullable(), c.autoincrementable())
	case *floatType:
		return common.NewFloatTableColumn(c.identifier, c.defaultValue(), c.nullable(), c.autoincrementable())
	case *datetimeType:
		return common.NewDatetimeTableColumn(c.identifier, c.defaultValue(), c.nullable(), c.autoincrementable())
	case *charType:
		return common.NewCharTableColumn(c.identifier, c.defaultValue(), c.nullable(), c.autoincrementable(), uint32(v.length))
	}

	return nil
}

func (c *columnDefinition) nullable() bool {
	for _, constraint := range c.columnConstraints {
		if _, ok := constraint.(notNullConstraint); ok {
			return false
		}
	}

	return true
}

func (c *columnDefinition) autoincrementable() bool {
	for _, constraint := range c.columnConstraints {
		if _, ok := constraint.(autoincrementConstraint); ok {
			return true
		}
	}

	return false
}

func (c *columnDefinition) defaultValue() interface{} {
	for _, constraint := range c.columnConstraints {
		if defaultConstraint, ok := constraint.(defaultConstraint); ok {
			return defaultConstraint.value
		}
	}

	return nil
}

type assignment struct {
	identifier string
	value      expression
}
