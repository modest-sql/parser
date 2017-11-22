package parser

import "github.com/modest-sql/common"

type expression interface {
	convert() interface{}
}

type assignment struct {
	identifier string
	value      expression
}

func (s *assignment) convert() interface{} {
	return common.NewAssignmentCommon(s.identifier, s.value.convert())
}

type idExpression struct {
	name  string
	alias string
}

func (s *idExpression) convert() interface{} {
	return common.NewIdCommon(s.name, s.alias)
}

type intExpression struct {
	value int64
}

func (s *intExpression) convert() interface{} {
	return common.NewIntCommon(s.value)
}

type boolExpression struct {
	value bool
}

func (s *boolExpression) convert() interface{} {
	return common.NewBoolCommon(s.value)
}

type floatExpression struct {
	value float64
}

func (s *floatExpression) convert() interface{} {
	return common.NewFloatCommon(s.value)
}

type stringExpression struct {
	value string
}

func (s *stringExpression) convert() interface{} {
	return common.NewStringCommon(s.value)
}

type sumExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *sumExpression) convert() interface{} {

	return common.NewSumCommon(s.rightValue.convert(), s.leftValue.convert())
}

type subExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *subExpression) convert() interface{} {
	return common.NewSubCommon(s.rightValue.convert(), s.leftValue.convert())
}

type multExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *multExpression) convert() interface{} {
	return common.NewMultCommon(s.rightValue.convert(), s.leftValue.convert())
}

type divExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *divExpression) convert() interface{} {
	return common.NewDivCommon(s.rightValue.convert(), s.leftValue.convert())
}

type eqExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *eqExpression) convert() interface{} {
	return common.NewEqCommon(s.rightValue.convert(), s.leftValue.convert())
}

type neExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *neExpression) convert() interface{} {
	return common.NewNeCommon(s.rightValue.convert(), s.leftValue.convert())
}

type ltExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *ltExpression) convert() interface{} {
	return common.NewLtCommon(s.rightValue.convert(), s.leftValue.convert())
}

type gtExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *gtExpression) convert() interface{} {
	return common.NewGtCommon(s.rightValue.convert(), s.leftValue.convert())
}

type lteExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *lteExpression) convert() interface{} {
	return common.NewLteCommon(s.rightValue.convert(), s.leftValue.convert())
}

type gteExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *gteExpression) convert() interface{} {
	return common.NewGteCommon(s.rightValue.convert(), s.leftValue.convert())
}

type betweenExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *betweenExpression) convert() interface{} {
	return common.NewBetweenCommon(s.rightValue.convert(), s.leftValue.convert())
}

type likeExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *likeExpression) convert() interface{} {
	return common.NewLikeCommon(s.rightValue.convert(), s.leftValue.convert())
}

type notExpression struct {
	not expression
}

func (s *notExpression) convert() interface{} {
	return common.NewNotCommon(s.not.convert())
}

type andExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *andExpression) convert() interface{} {
	return common.NewAndCommon(s.rightValue.convert(), s.leftValue.convert())
}

type orExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *orExpression) convert() interface{} {
	return common.NewOrCommon(s.rightValue.convert(), s.leftValue.convert())
}

type nullExpression struct {
}

func (s *nullExpression) convert() interface{} {
	return common.NewNullCommon()
}

type falseExpression struct {
}

func (s *falseExpression) convert() interface{} {
	return common.NewFalseCommon()
}

type trueExpression struct {
}

func (s *trueExpression) convert() interface{} {
	return common.NewTrueCommon()
}
