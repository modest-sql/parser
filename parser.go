//line parser.y:4
package parser

import __yyfmt__ "fmt"

//line parser.y:4
import "io"

/*import "github.com/modest-sql/common"*/

var statements statementList

//line parser.y:13
type yySymType struct {
	yys       int
	int64_t   int64
	string_t  string
	float64_t float64

	expr_t expression

	stmt_list_t       statementList
	stmt_t            statement
	columnSpec_t      columnSpec
	columnSpec_list_t []columnSpec
	joinSpec_list_t   []joinSpec
	joinSpec_t        joinSpec
	col_list_t        columnDefinitions
	col_t             *columnDefinition
	assignment_t      assignment
	data_t            dataType
	assignments_list  []assignment
	obj_list_t        []interface{}
	string_list_t     []string

	obj_t interface{}
}
