package parser

type expression interface {
	evaluate() (interface{}, error)
}

type idExpression struct {
	name  string
	alias string
}

func (e *idExpression) evaluate() (interface{}, error) {
	return &idExpression{}, nil
}

type intExpression struct {
	value int
}

func (e *intExpression) evaluate() (interface{}, error) {
	return &intExpression{}, nil
}

type boolExpression struct {
	value bool
}

func (e *boolExpression) evaluate() (interface{}, error) {
	return &boolExpression{}, nil
}

type floatExpression struct {
	value float64
}

func (e *floatExpression) evaluate() (interface{}, error) {
	return &floatExpression{}, nil
}

type stringExpression struct {
	value string
}

func (e *stringExpression) evaluate() (interface{}, error) {
	return &stringExpression{}, nil
}

type sumExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *sumExpression) evaluate() (interface{}, error) {
	return &sumExpression{}, nil
}

type subExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *subExpression) evaluate() (interface{}, error) {
	return &subExpression{}, nil
}

type multExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *multExpression) evaluate() (interface{}, error) {
	return &multExpression{}, nil
}

type divExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *divExpression) evaluate() (interface{}, error) {
	return &divExpression{}, nil
}

type eqExpression struct {
	rightValue expression
	leftValue  expression
}
type neExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *neExpression) evaluate() (interface{}, error) {
	return &neExpression{}, nil
}

func (e *eqExpression) evaluate() (interface{}, error) {
	return &divExpression{}, nil
}

type ltExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *ltExpression) evaluate() (interface{}, error) {
	return &ltExpression{}, nil
}

type gtExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *gtExpression) evaluate() (interface{}, error) {
	return &gtExpression{}, nil
}

type lteExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *lteExpression) evaluate() (interface{}, error) {
	return &lteExpression{}, nil
}

type gteExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *gteExpression) evaluate() (interface{}, error) {
	return &gteExpression{}, nil
}

type betweenExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *betweenExpression) evaluate() (interface{}, error) {
	return &betweenExpression{}, nil
}

type likeExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *likeExpression) evaluate() (interface{}, error) {
	return &likeExpression{}, nil
}

type notExpression struct {
	not expression
}

func (e *notExpression) evaluate() (interface{}, error) {
	return &notExpression{}, nil
}

type andExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *andExpression) evaluate() (interface{}, error) {
	return &andExpression{}, nil
}

type orExpression struct {
	rightValue expression
	leftValue  expression
}

func (e *orExpression) evaluate() (interface{}, error) {
	return &orExpression{}, nil
}

type nullExpression struct {
}
type falseExpression struct {
}
type trueExpression struct {
}
