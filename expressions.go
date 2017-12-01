package parser

import "github.com/modest-sql/common"

type expression interface {
	convert() interface{}
	sintactico() string
}

type assignment struct {
	identifier string
	value      expression
}

func (s *assignment) convert() interface{} {
	return common.NewAssignmentCommon(s.identifier, s.value.convert())
}

func (s *assignment) sintactico() string {
	s.value.sintactico()
	return "assignment"
}

type idExpression struct {
	name  string
	alias string
}

func (s *idExpression) convert() interface{} {
	return common.NewIdCommon(s.name, s.alias)
}

func (s *idExpression) sintactico() string {
	return "id"
}
type intExpression struct {
	value int64
}

func (s *intExpression) convert() interface{} {
	return common.NewIntCommon(s.value)
}

func (s *intExpression) sintactico() string {
	return "int"
}
type boolExpression struct {
	value bool
}

func (s *boolExpression) convert() interface{} {
	return common.NewBoolCommon(s.value)
}

func (s *boolExpression) sintactico() string {
	return "bool"
}

type floatExpression struct {
	value float64
}

func (s *floatExpression) convert() interface{} {
	return common.NewFloatCommon(s.value)
}

func (s *floatExpression) sintactico() string {
	return "float"
}

type stringExpression struct {
	value string
}

func (s *stringExpression) convert() interface{} {
	return common.NewStringCommon(s.value)
}

func (s *stringExpression) sintactico() string {
	return "string"
}

type sumExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *sumExpression) convert() interface{} {

	switch v1 := s.leftValue.(type) {
	case *intExpression:
		switch v2 := s.rightValue.(type) {
		case *intExpression:
			value := intExpression{v1.value + v2.value}
			return value.convert()
		case *floatExpression:
			value := floatExpression{float64(v1.value) + v2.value}
			return value.convert()
		}
	case *floatExpression:
		switch v2 := s.rightValue.(type) {
		case *floatExpression:
			value := floatExpression{v1.value + v2.value}
			return value.convert()
		case *intExpression:
			value := floatExpression{v1.value + float64(v2.value)}
			return value.convert()
		}
	case *stringExpression:
		switch v2 := s.rightValue.(type) {
		case *stringExpression:
			value := stringExpression{v1.value + v2.value}
			return value.convert()
		}
	}
	return common.NewSumCommon(s.rightValue.convert(), s.leftValue.convert())
}
func (s *sumExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	if v1=="string"||v2=="string" {
		return "string"
	}else if (v1=="float"&&v2=="float")||(v1=="int"&&v2=="float")||(v1=="float"&&v2=="int"){
		return "float"
	}else if((v1=="int"&&v2=="bool")||(v1=="int"&&v2=="bool")||(v1=="int"&&v2=="int")){
		return "int"
	}
	panic(true)
}

type subExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *subExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	 if (v1=="float"&&v2=="float")||(v1=="int"&&v2=="float")||(v1=="float"&&v2=="int"){
		return "float"
	}else if((v1=="int"&&v2=="bool")||(v1=="int"&&v2=="bool")||(v1=="int"&&v2=="int")){
		return "int"
	}
	panic(true)
}

func (s *subExpression) convert() interface{} {
	switch v1 := s.leftValue.(type) {
	case *intExpression:
		switch v2 := s.rightValue.(type) {
		case *intExpression:
			value := intExpression{v1.value - v2.value}
			return value.convert()
		case *floatExpression:
			value := floatExpression{float64(v1.value) - v2.value}
			return value.convert()
		}
	case *floatExpression:
		switch v2 := s.rightValue.(type) {
		case *floatExpression:
			value := floatExpression{v1.value - v2.value}
			return value.convert()
		case *intExpression:
			value := floatExpression{v1.value - float64(v2.value)}
			return value.convert()
		}
	}
	return common.NewSubCommon(s.rightValue.convert(), s.leftValue.convert())
}

type multExpression struct {
	rightValue expression
	leftValue  expression
}


func (s *multExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	 if ((v1=="float"&&v2=="float")||(v1=="int"&&v2=="float")||(v1=="float"&&v2=="int")){
		return "float"
	}else if((v1=="int"&&v2=="bool")||(v1=="int"&&v2=="bool")||(v1=="int"&&v2=="int")){
		return "int"
	}
	panic(true)
}
func (s *multExpression) convert() interface{} {
	switch v1 := s.leftValue.(type) {
	case *intExpression:
		switch v2 := s.rightValue.(type) {
		case *intExpression:
			value := intExpression{v1.value * v2.value}
			return value.convert()
		case *floatExpression:
			value := floatExpression{float64(v1.value) * v2.value}
			return value.convert()
		}
	case *floatExpression:
		switch v2 := s.rightValue.(type) {
		case *floatExpression:
			value := floatExpression{v1.value * v2.value}
			return value.convert()
		case *intExpression:
			value := floatExpression{v1.value * float64(v2.value)}
			return value.convert()
		}
	}
	return common.NewMultCommon(s.rightValue.convert(), s.leftValue.convert())
}

type divExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *divExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	 if ((v1=="float"&&v2=="float")||(v1=="int"&&v2=="float")||(v1=="float"&&v2=="int")){
		return "float"
	 }else if((v1=="int"&&v2=="bool")||(v1=="int"&&v2=="bool")||(v1=="int"&&v2=="int")){
		return "int"
	 }

	panic(true)
}

func (s *divExpression) convert() interface{} {
	switch v1 := s.leftValue.(type) {
	case *intExpression:
		switch v2 := s.rightValue.(type) {
		case *intExpression:
			value := intExpression{v1.value / v2.value}
			return value.convert()
		case *floatExpression:
			value := floatExpression{float64(v1.value) / v2.value}
			return value.convert()
		}
	case *floatExpression:
		switch v2 := s.rightValue.(type) {
		case *floatExpression:
			value := floatExpression{v1.value / v2.value}
			return value.convert()
		case *intExpression:
			value := floatExpression{v1.value / float64(v2.value)}
			return value.convert()
		}
	}
	return common.NewDivCommon(s.rightValue.convert(), s.leftValue.convert())
}

type eqExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *eqExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	return ""
}
func (s *eqExpression) convert() interface{} {
	return common.NewEqCommon(s.rightValue.convert(), s.leftValue.convert())
}

type neExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *neExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	return ""
}

func (s *neExpression) convert() interface{} {
	return common.NewNeCommon(s.rightValue.convert(), s.leftValue.convert())
}

type ltExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *ltExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	return ""
}

func (s *ltExpression) convert() interface{} {
	return common.NewLtCommon(s.rightValue.convert(), s.leftValue.convert())
}

type gtExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *gtExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	return ""
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
func (s *lteExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	return ""
}
type gteExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *gteExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	return ""
}
func (s *gteExpression) convert() interface{} {
	return common.NewGteCommon(s.rightValue.convert(), s.leftValue.convert())
}

type betweenExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *betweenExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	return ""
}
func (s *betweenExpression) convert() interface{} {
	return common.NewBetweenCommon(s.rightValue.convert(), s.leftValue.convert())
}

type likeExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *likeExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	return ""
}
func (s *likeExpression) convert() interface{} {
	return common.NewLikeCommon(s.rightValue.convert(), s.leftValue.convert())
}

type notExpression struct {
	not expression
}
func (s *notExpression) sintactico() string {
	v1 := s.not.sintactico()
	return ""
	
}
func (s *notExpression) convert() interface{} {
	return common.NewNotCommon(s.not.convert())
}

type andExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *andExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	return ""
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
func (s *orExpression) sintactico() string {
	v1 := s.leftValue.sintactico()
	v2 := s.rightValue.sintactico()
	return ""
}
type nullExpression struct {

}
func (s *nullExpression) sintactico() string {
	return ""
}
func (s *nullExpression) convert() interface{} {
	return common.NewNullCommon()
}

type falseExpression struct {
}
func (s *falseExpression) sintactico() string {
	return ""
}
func (s *falseExpression) convert() interface{} {
	return common.NewFalseCommon()
}

type trueExpression struct {
}

func (s *trueExpression) convert() interface{} {
	return common.NewTrueCommon()
}
func (s *trueExpression) sintactico() string {
	return ""
}