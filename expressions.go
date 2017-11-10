package parser

type expression interface {
	
}

type idExpression struct {
	name  string
	alias string
}

type intExpression struct {
	value int
}

type boolExpression struct {
	value bool
}

type floatExpression struct {
	value float64
}

type stringExpression struct {
	value string
}

type sumExpression struct {
	rightValue expression
	leftValue  expression
}

type subExpression struct {
	rightValue expression
	leftValue  expression
}

type multExpression struct {
	rightValue expression
	leftValue  expression
}

type divExpression struct {
	rightValue expression
	leftValue  expression
}

type eqExpression struct {
	rightValue expression
	leftValue  expression
}
type neExpression struct {
	rightValue expression
	leftValue  expression
}

type ltExpression struct {
	rightValue expression
	leftValue  expression
}

type gtExpression struct {
	rightValue expression
	leftValue  expression
}

type lteExpression struct {
	rightValue expression
	leftValue  expression
}
type gteExpression struct {
	rightValue expression
	leftValue  expression
}


type betweenExpression struct {
	rightValue expression
	leftValue  expression
}


type likeExpression struct {
	rightValue expression
	leftValue  expression
}


type notExpression struct {
	not expression
}
type andExpression struct {
	rightValue expression
	leftValue  expression
}

type orExpression struct {
	rightValue expression
	leftValue  expression
}
type nullExpression struct {
}
type falseExpression struct {
}
type trueExpression struct {
}
