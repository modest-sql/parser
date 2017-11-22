package parser

type expression interface {
	convert() interface{}
}

type idExpression struct {
	name  string
	alias string
}

func (s *idExpression) convert() interface{} {
	return nil
}

type intExpression struct {
	value int64
}

func (s *intExpression) convert() interface{} {
	return nil
}

type boolExpression struct {
	value bool
}

func (s *boolExpression) convert() interface{} {
	return nil
}

type floatExpression struct {
	value float64
}

func (s *floatExpression) convert() interface{} {
	return nil
}

type stringExpression struct {
	value string
}

func (s *stringExpression) convert() interface{} {
	return nil
}

type sumExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *sumExpression) convert() interface{} {
	return nil
}

type subExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *subExpression) convert() interface{} {
	return nil
}

type multExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *multExpression) convert() interface{} {
	return nil
}

type divExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *divExpression) convert() interface{} {
	return nil
}

type eqExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *eqExpression) convert() interface{} {
	return nil
}

type neExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *neExpression) convert() interface{} {
	return nil
}

type ltExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *ltExpression) convert() interface{} {
	return nil
}

type gtExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *gtExpression) convert() interface{} {
	return nil
}

type lteExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *lteExpression) convert() interface{} {
	return nil
}

type gteExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *gteExpression) convert() interface{} {
	return nil
}

type betweenExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *betweenExpression) convert() interface{} {
	return nil
}

type likeExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *likeExpression) convert() interface{} {
	return nil
}

type notExpression struct {
	not expression
}

func (s *notExpression) convert() interface{} {
	return nil
}

type andExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *andExpression) convert() interface{} {
	return nil
}

type orExpression struct {
	rightValue expression
	leftValue  expression
}

func (s *orExpression) convert() interface{} {
	return nil
}

type nullExpression struct {
}

func (s *nullExpression) convert() interface{} {
	return nil
}

type falseExpression struct {
}

func (s *falseExpression) convert() interface{} {
	return nil
}

type trueExpression struct {
}

func (s *trueExpression) convert() interface{} {
	return nil
}
