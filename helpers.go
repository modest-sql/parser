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

type primaryKeyConstraint struct {
}

type foreignKeyConstraint struct {
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
		return common.NewBooleanTableColumn(c.identifier, c.defaultValue(), c.nullable(), c.autoincrementable(), c.isPrimaryKey(), c.isForeignKey())
	case *integerType:
		return common.NewIntegerTableColumn(c.identifier, c.defaultValue(), c.nullable(), c.autoincrementable(), c.isPrimaryKey(), c.isForeignKey())
	case *floatType:
		return common.NewFloatTableColumn(c.identifier, c.defaultValue(), c.nullable(), c.autoincrementable(), c.isPrimaryKey(), c.isForeignKey())
	case *datetimeType:
		return common.NewDatetimeTableColumn(c.identifier, c.defaultValue(), c.nullable(), c.autoincrementable(), c.isPrimaryKey(), c.isForeignKey())
	case *charType:
		return common.NewCharTableColumn(c.identifier, c.defaultValue(), c.nullable(), c.autoincrementable(), c.isPrimaryKey(), c.isForeignKey(), uint16(v.length))
	}

	return nil
}

func (c *columnDefinition) nullable() bool {
	for _, constraint := range c.columnConstraints {
		if _, ok := constraint.(*notNullConstraint); ok {
			return false
		}
	}

	return true
}

func (c *columnDefinition) autoincrementable() bool {
	for _, constraint := range c.columnConstraints {
		if _, ok := constraint.(*autoincrementConstraint); ok {
			return true
		}
	}

	return false
}

func (c *columnDefinition) isPrimaryKey() bool {
	for _, constraint := range c.columnConstraints {
		if _, ok := constraint.(*primaryKeyConstraint); ok {
			return true
		}
	}

	return false
}

func (c *columnDefinition) isForeignKey() bool {
	for _, constraint := range c.columnConstraints {
		if _, ok := constraint.(*foreignKeyConstraint); ok {
			return true
		}
	}

	return false
}

func (c *columnDefinition) defaultValue() interface{} {
	for _, constraint := range c.columnConstraints {
		if defaultConstraint, ok := constraint.(*defaultConstraint); ok {
			return defaultConstraint.value
		}
	}

	return nil
}
