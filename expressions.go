package parser

import "github.com/modest-sql/common"

const constfloat = "float"
const constint = "int"
const conststring = "string"
const constid = "id"
const constbool = "bool"

type expression interface {
	convert() interface{}
	evaluateType() string
}

type assignment struct {
	identifier string
	value      expression
}

func (s *assignment) convert() interface{} {
	return common.NewAssignmentCommon(s.identifier, s.value.convert())
}

func (s *assignment) evaluateType() string {
	s.value.evaluateType()
	return "assignment"
}

type idExpression struct {
	name  string
	alias string
}

func (s *idExpression) convert() interface{} {
	return common.NewIdCommon(s.name, s.alias)
}

func (s *idExpression) evaluateType() string {
	return constid
}
type intExpression struct {
	value int64
}

func (s *intExpression) convert() interface{} {
	return common.NewIntCommon(s.value)
}

func (s *intExpression) evaluateType() string {
	return constint
}
type boolExpression struct {
	value bool
}

func (s *boolExpression) convert() interface{} {
	return common.NewBoolCommon(s.value)
}

func (s *boolExpression) evaluateType() string {
	return constbool
}

type floatExpression struct {
	value float64
}

func (s *floatExpression) convert() interface{} {
	return common.NewFloatCommon(s.value)
}

func (s *floatExpression) evaluateType() string {
	return constfloat
}

type stringExpression struct {
	value string
}

func (s *stringExpression) convert() interface{} {
	return common.NewStringCommon(s.value)
}

func (s *stringExpression) evaluateType() string {
	return conststring
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
func (s *sumExpression) evaluateType() string {
	v1 := s.leftValue.evaluateType()
	v2 := s.rightValue.evaluateType()
	if v1==conststring||v2==conststring {
		return conststring
	}else if (v1==constfloat&&v2==constfloat)||(v1==constint&&v2==constfloat)||(v1==constfloat&&v2==constint){
		return constfloat
	}else if((v1==constint&&v2==constbool)||(v1==constint&&v2==constbool)||(v1==constint&&v2==constint)){
		return constint
	}else if(v1==constid||v2==constid){
		return constid
	}
	panic("incompatible datatype")
}

type subExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *subExpression) evaluateType() string {
	v1 := s.leftValue.evaluateType()
	v2 := s.rightValue.evaluateType()
	if (v1== constfloat && v2==constfloat)||(v1==constint&&v2==constfloat)||(v1==constfloat&&v2==constint){
		return constfloat
	}else if((v1==constint&&v2==constbool)||(v1==constint&&v2==constbool)||(v1==constint&&v2==constint)){
		return constint
	}else if(v1==constid||v2==constid){
		return constid
	}
	panic("incompatible datatype")
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


func (s *multExpression) evaluateType() string {
	v1 := s.leftValue.evaluateType()
	v2 := s.rightValue.evaluateType()
	 if ((v1==constfloat&&v2==constfloat)||(v1==constint&&v2==constfloat)||(v1==constfloat&&v2==constint)){
		return constfloat
	}else if((v1==constint&&v2==constbool)||(v1==constint&&v2==constbool)||(v1==constint&&v2==constint)){
		return constint
	}else if(v1==constid||v2==constid){
		return constid
	}
	panic("incompatible datatype")
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
func (s *divExpression) evaluateType() string {
	v1 := s.leftValue.evaluateType()
	v2 := s.rightValue.evaluateType()
	 if ((v1==constfloat&&v2==constfloat)||(v1==constint&&v2==constfloat)||(v1==constfloat&&v2==constint)){
		return constfloat
	 }else if((v1==constint&&v2==constbool)||(v1==constint&&v2==constbool)||(v1==constint&&v2==constint)){
		return constint
	 }else if(v1==constid||v2==constid){
		return constid
	}

	panic("incompatible datatype")
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
func (s *eqExpression) evaluateType() string {
	 s.leftValue.evaluateType()
	 s.rightValue.evaluateType()
	return ""
}
func (s *eqExpression) convert() interface{} {
	return common.NewEqCommon(s.rightValue.convert(), s.leftValue.convert())
}

type neExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *neExpression) evaluateType() string {
	 s.leftValue.evaluateType()
	 s.rightValue.evaluateType()
	return ""
}

func (s *neExpression) convert() interface{} {
	return common.NewNeCommon(s.rightValue.convert(), s.leftValue.convert())
}

type ltExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *ltExpression) evaluateType() string {
	 s.leftValue.evaluateType()
	 s.rightValue.evaluateType()
	return ""
}

func (s *ltExpression) convert() interface{} {
	return common.NewLtCommon(s.rightValue.convert(), s.leftValue.convert())
}

type gtExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *gtExpression) evaluateType() string {
	 s.leftValue.evaluateType()
	 s.rightValue.evaluateType()
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
func (s *lteExpression) evaluateType() string {
	 s.leftValue.evaluateType()
	 s.rightValue.evaluateType()
	return ""
}
type gteExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *gteExpression) evaluateType() string {
	 s.leftValue.evaluateType()
	 s.rightValue.evaluateType()
	return ""
}
func (s *gteExpression) convert() interface{} {
	return common.NewGteCommon(s.rightValue.convert(), s.leftValue.convert())
}

type betweenExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *betweenExpression) evaluateType() string {
	 s.leftValue.evaluateType()
	 s.rightValue.evaluateType()
	return ""
}
func (s *betweenExpression) convert() interface{} {
	return common.NewBetweenCommon(s.rightValue.convert(), s.leftValue.convert())
}

type likeExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *likeExpression) evaluateType() string {
	 s.leftValue.evaluateType()
	 s.rightValue.evaluateType()
	return ""
}
func (s *likeExpression) convert() interface{} {
	return common.NewLikeCommon(s.rightValue.convert(), s.leftValue.convert())
}

type notExpression struct {
	not expression
}
func (s *notExpression) evaluateType() string {
	 s.not.evaluateType()
	return ""
	
}
func (s *notExpression) convert() interface{} {
	return common.NewNotCommon(s.not.convert())
}

type andExpression struct {
	rightValue expression
	leftValue  expression
}
func (s *andExpression) evaluateType() string {
	 s.leftValue.evaluateType()
	 s.rightValue.evaluateType()
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
func (s *orExpression) evaluateType() string {
	 s.leftValue.evaluateType()
	 s.rightValue.evaluateType()
	return ""
}
type nullExpression struct {

}
func (s *nullExpression) evaluateType() string {
	return ""
}
func (s *nullExpression) convert() interface{} {
	return common.NewNullCommon()
}

type falseExpression struct {
}
func (s *falseExpression) evaluateType() string {
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
func (s *trueExpression) evaluateType() string {
	return ""
}