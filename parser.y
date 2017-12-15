/* Based on https://ronsavage.github.io/SQL/sql-2003-2.bnf.html */

%{
package parser

import (
    "io"
    "sync"
)

var statements statementList
var lock sync.Mutex

%}

%union {
    int64_t int64
    string_t string
    float64_t float64

    expr_t expression

    stmt_list_t statementList
    stmt_t statement
    columnSpec_t columnSpec
    columnSpec_list_t []columnSpec
    joinSpec_list_t []joinSpec
    joinSpec_t joinSpec
    col_list_t columnDefinitions
    col_t *columnDefinition
    assignment_t assignment
    data_t dataType
    assignments_list []assignment
    obj_list_t []interface{}
    string_list_t []string
    group_by_t GroupBySpec
    group_by_list_t []GroupBySpec
    obj_t interface{}
}

%token TK_PLUS TK_MINUS TK_STAR TK_DIV TK_LT TK_GT TK_GTE TK_LTE TK_EQ TK_NE
%token TK_LEFT_PAR TK_RIGHT_PAR TK_COMMA TK_DOT TK_SEMICOLON
%token KW_OR KW_AND KW_NOT KW_INTEGER KW_FLOAT KW_CHAR KW_BOOLEAN KW_DATETIME
%token KW_PRIMARY KW_FOREIGN KW_KEY KW_BY
%token KW_CREATE KW_TABLE KW_DELETE KW_INSERT
%token KW_INTO KW_SELECT KW_WHERE KW_FROM KW_UPDATE KW_SET
%token KW_ALTER KW_VALUES KW_BETWEEN KW_LIKE KW_INNER
%token KW_HAVING KW_SUM KW_COUNT KW_AVG KW_MIN KW_MAX
%token KW_NULL KW_IN KW_IS KW_AUTO_INCREMENT KW_JOIN KW_ON KW_GROUP KW_BY KW_DROP KW_DEFAULT
%token KW_TRUE KW_FALSE KW_AS KW_ADD KW_COLUMN

%token<int64_t> INT_LIT
%token<float64_t> FLOAT_LIT
%token<string_t> TK_ID STR_LIT
%type<string_t> multipart_id_suffix alias_spec


%type<stmt_list_t> statements_list
%type<stmt_t> statement data_statement schema_statement create_statement alter_statement drop_statement
%type<stmt_t> select_statement insert_statement delete_statement update_statement

%type<col_list_t> table_element_list
%type<col_t> table_element column_definition

%type<obj_list_t> column_constraint_list values_list values_tuple values_tuples_list
%type<obj_t> column_constraint constr_not_null alter_instruction key_constraint

%type<data_t> data_type
%type<assignments_list>set_assignments_list set_list
%type<assignment_t>set_assignment
%type<obj_t> value_literal
%type<expr_t> addi_factor relational_factor relational_term truth_value between_term relational_expression
%type<expr_t> boolean_factor boolean_term boolean_value_expression opt_where_clause search_condition
%type<string_list_t>column_names_list 
%type<columnSpec_t>select_col
%type<columnSpec_list_t>select_col_list 
%type<joinSpec_list_t> opt_joins_list join_list
%type<joinSpec_t> inner_join
%type<group_by_t>op_groupBy 
%type<group_by_list_t>op_groupBy_List
%%

input: statements_list { lock.Lock(); statements = $1; }
    | { }
;

statements_list: statements_list statement TK_SEMICOLON { $$ = $1; $$ = append($$, $2) }
    | statement TK_SEMICOLON { $$ = append($$, $1) }
    | statement { $$ = append($$, $1) }
;

statement: data_statement { $$ = $1}
    | schema_statement { $$ = $1}
;

schema_statement: create_statement { $$ = $1 }
    | alter_statement { $$ = $1 }
    | drop_statement { $$ = $1 }
;

data_statement: select_statement { $$ = $1 }
    | insert_statement { $$ = $1  }
    | delete_statement { $$ = $1 }
    | update_statement { $$ = $1 }
;

create_statement: KW_CREATE KW_TABLE TK_ID TK_LEFT_PAR table_element_list TK_RIGHT_PAR { $$ = &createStatement{$3, $5} }
;

table_element_list: table_element_list TK_COMMA table_element { $$ = $1; $$ = append($$, $3) }
    | table_element { $$ = append($$, $1) }
;

table_element: column_definition {$$ = $1}
;

column_definition: TK_ID data_type column_constraint_list { $$ = &columnDefinition{$1, $2, $3} }
    | TK_ID data_type { $$ = &columnDefinition{$1, $2, nil } }
;

data_type: KW_CHAR TK_LEFT_PAR INT_LIT TK_RIGHT_PAR { $$ = &charType{$3} }
    | KW_INTEGER { $$ = &integerType{} }
    | KW_BOOLEAN { $$ = &booleanType{} }
    | KW_DATETIME { $$ = &datetimeType{} }
    | KW_FLOAT { $$ = &floatType{} }
;

column_constraint_list: column_constraint_list column_constraint { $$ = $1; $$ = append($$, $2) }
    | column_constraint { $$ = append($$, $1) }
;

column_constraint: constr_not_null { $$  = $1 }
;

constr_not_null: KW_NOT KW_NULL { $$ = &notNullConstraint{ } }
    | KW_DEFAULT value_literal { $$ = &defaultConstraint{ $2 } }
    | KW_AUTO_INCREMENT { $$ = &autoincrementConstraint{ } }
    | key_constraint { $$ = $1 }
;

key_constraint: KW_PRIMARY KW_KEY { $$ = &primaryKeyConstraint{ } }
    | KW_FOREIGN KW_KEY { $$ = &foreignKeyConstraint{ } }
;

alter_statement: KW_ALTER KW_TABLE TK_ID alter_instruction { $$ = &alterStatement{$3,$4}  }
;

alter_instruction: KW_DROP KW_COLUMN TK_ID {$$ = &alterDrop{ $3 } }
    | KW_ADD TK_ID data_type column_constraint_list { $$ = &alterAdd{ $2,$3,$4} }
    | KW_ADD TK_ID data_type { $$ = &alterAdd{ $2,$3, nil} }
    | KW_ADD KW_COLUMN TK_ID data_type column_constraint_list { $$ = &alterAdd{ $3,$4,$5} }
    | KW_ADD KW_COLUMN TK_ID data_type { $$ = &alterAdd{ $3,$4, nil} }
    | KW_ALTER KW_COLUMN TK_ID data_type { $$ = &alterModify{ $3,$4, nil} }
    | KW_ALTER KW_COLUMN TK_ID data_type column_constraint_list { $$ = &alterModify{ $3,$4, $5} }
;

drop_statement: KW_DROP KW_TABLE TK_ID { $$ = &dropStatement{ $3 } }
;

select_statement: KW_SELECT select_col_list KW_FROM TK_ID alias_spec opt_joins_list opt_where_clause op_groupBy_List { $$ = &selectStatement{$2,$4,$5,$6, $7,$8} }
    | KW_SELECT select_col_list KW_FROM TK_ID opt_joins_list opt_where_clause op_groupBy_List { $$ = &selectStatement{$2,$4,"", $5, $6,$7} }
;

op_groupBy_List:op_groupBy_List TK_COMMA  op_groupBy  { $$ = $1 ; $$ = append($$,$3) }
                |op_groupBy { $$ = append($$,$1) }
                | { $$ = nil }
;

op_groupBy:KW_GROUP KW_BY TK_ID { $$ = GroupBySpec{ "",$3} }
        | KW_GROUP KW_BY TK_ID multipart_id_suffix {  $$ = GroupBySpec{ $3,$4} }
        

;
select_col_list: select_col_list TK_COMMA select_col { $$ = $1 ; $$ = append($$,$3) }
    | select_col { $$ = append($$,$1) }
;

select_col: TK_STAR { $$ = columnSpec{true,"","","",nil} }
    | TK_ID alias_spec {$$ = columnSpec{false,$1,"",$2,nil} }
    | TK_ID { $$ = columnSpec{false,$1,"","",nil}}
    | TK_ID multipart_id_suffix alias_spec { $$ = columnSpec{false,$1,$2,$3,nil} }
    | TK_ID multipart_id_suffix { $$ = columnSpec{false,$1,$2,"",nil} }
    | KW_SUM TK_LEFT_PAR TK_ID TK_RIGHT_PAR { $$ = columnSpec{false,$3,"","",&functionSum{}}}
    | KW_SUM TK_LEFT_PAR TK_ID multipart_id_suffix TK_RIGHT_PAR { $$ = columnSpec{false,$3,$4,"",&functionSum{}}}
    | KW_COUNT TK_LEFT_PAR TK_ID TK_RIGHT_PAR { $$ = columnSpec{false,$3,"","",&functionCount{}}}
    | KW_COUNT TK_LEFT_PAR TK_ID multipart_id_suffix TK_RIGHT_PAR { $$ = columnSpec{false,$3,$4,"",&functionCount{}}}
    | KW_AVG TK_LEFT_PAR TK_ID TK_RIGHT_PAR { $$ = columnSpec{false,$3,"","",&functionAvg{}}}
    | KW_AVG TK_LEFT_PAR TK_ID multipart_id_suffix TK_RIGHT_PAR { $$ = columnSpec{false,$3,$4,"",&functionAvg{}}}
    | KW_MIN TK_LEFT_PAR TK_ID TK_RIGHT_PAR { $$ = columnSpec{false,$3,"","",&functionMin{}}}
    | KW_MIN TK_LEFT_PAR TK_ID multipart_id_suffix TK_RIGHT_PAR { $$ = columnSpec{false,$3,$4,"",&functionMin{}}}
    | KW_MAX TK_LEFT_PAR TK_ID TK_RIGHT_PAR { $$ = columnSpec{false,$3,"","",&functionMax{}}}
    | KW_MAX TK_LEFT_PAR TK_ID multipart_id_suffix TK_RIGHT_PAR { $$ = columnSpec{false,$3,$4,"",&functionMax{}}}

insert_statement: KW_INSERT KW_INTO TK_ID TK_LEFT_PAR column_names_list TK_RIGHT_PAR KW_VALUES values_tuples_list {  $$ = &insertStatement{$3,$5,$8} }
    | KW_INSERT KW_INTO TK_ID KW_VALUES values_tuples_list {  $$ = &insertStatement{$3, nil, $5} }
;

column_names_list: column_names_list TK_COMMA TK_ID { $$ = $1; $$ = append($$, $3) }
    | TK_ID { $$ = append($$,$1) }
;
 
values_tuples_list: values_tuples_list TK_COMMA values_tuple { $$ = $1 ; $$ = append($$,$3) }
    | values_tuple { $$ = append($1)}
;

values_tuple: TK_LEFT_PAR values_list TK_RIGHT_PAR {$$ = $2}
;

values_list: values_list TK_COMMA value_literal { $$ = $1; $$ = append($$,$3)}
    | value_literal {$$ = append($$,$1) }
;

value_literal: STR_LIT { $$ = $1 }
    | INT_LIT { $$ = $1 }
    | truth_value { }
    | FLOAT_LIT { }
;

opt_joins_list: join_list { $$ = $1 }
    | { $$ = nil }
;

join_list: join_list inner_join { $$ = $1; $$ = append($$, $2); }
    | inner_join { $$ = append($$, $1) }
;

inner_join: KW_INNER KW_JOIN TK_ID alias_spec KW_ON search_condition { $$ = joinSpec{ $3, $4, $6 } }
    | KW_INNER KW_JOIN TK_ID KW_ON search_condition { $$ = joinSpec{ $3, "", $5 } }
;

delete_statement: KW_DELETE TK_ID alias_spec opt_where_clause { $$ = &deleteStatement{$2,$3,$4} }
    | KW_DELETE KW_FROM TK_ID alias_spec opt_where_clause { $$ = &deleteStatement{$3,$4,$5} }
    | KW_DELETE TK_ID opt_where_clause { $$ = &deleteStatement{$2,"",$3} }
    | KW_DELETE KW_FROM TK_ID opt_where_clause { $$ = &deleteStatement{$3,"",$4} }
;

alias_spec: KW_AS TK_ID { $$ = $2 }
    | TK_ID { $$ = $1  }
;

opt_where_clause: KW_WHERE search_condition { $$ = $2 }
    | { $$ = nil }
;

search_condition: boolean_value_expression {$$ = $1}
;

update_statement: KW_UPDATE TK_ID set_list opt_where_clause { $$ = &updateStatement{$2,$3,$4} }
;

set_list: KW_SET set_assignments_list { $$ = $2 }
;

set_assignments_list: set_assignments_list TK_COMMA set_assignment { $$ = $1; $$ = append($1, $3) }
    | set_assignment {  $$ = append($$, $1) }
;

set_assignment: TK_ID TK_EQ relational_term { $$ = assignment{ $1 , $3 } }
;

boolean_value_expression: boolean_value_expression KW_OR boolean_term { $$ = &orExpression{ $1 , $3 } }
    | boolean_term { $$ = $1 }
;

boolean_term: boolean_term KW_AND boolean_factor { $$ = &andExpression{ $1 , $3 } }
    | boolean_factor { $$ = $1  }
;

boolean_factor: KW_NOT relational_expression { $$ = &notExpression{ $2 } }
    | relational_expression { $$ = $1 }
;

relational_expression: relational_expression TK_LT relational_term { $$ = &ltExpression{ $1 , $3 } }
    | relational_expression TK_GT relational_term { $$ = &gtExpression{ $1 , $3 } }
    | relational_expression TK_EQ relational_term { $$ = &eqExpression{ $1 , $3 } }
    | relational_expression TK_LTE relational_term { $$ = &lteExpression{ $1 , $3 } }
    | relational_expression TK_GTE relational_term { $$ = &gteExpression{ $1 , $3 } }
    | relational_expression TK_NE relational_term { $$ = &neExpression{ $1 , $3 } }
    | relational_expression KW_LIKE relational_term { $$ = &likeExpression{ $1 , $3 } }
    | relational_expression KW_BETWEEN between_term { $$ = &betweenExpression{ $1 , $3} }
    | relational_term { $$ = $1 }
;

between_term: INT_LIT KW_AND INT_LIT { $$ = &betweenExpression{ &intExpression{$1} , &intExpression{$3} } }
;

relational_term: relational_term TK_PLUS relational_factor { $$ = &sumExpression{ $1 , $3 } }
    | relational_term TK_MINUS relational_factor { $$ = &subExpression{ $1 , $3 } }
    | relational_factor { $$ = $1 }
;

relational_factor: relational_factor TK_STAR addi_factor { $$ = &multExpression{ $1 , $3 } }
    | relational_factor TK_DIV addi_factor { $$ = &divExpression{ $1 , $3 } }
    | addi_factor { $$ = $1  }
;

addi_factor: INT_LIT { $$ = &intExpression{ $1 } }
    | truth_value { $$ = $1 }
    | STR_LIT { $$ = &stringExpression{ $1 } }
    | TK_ID { $$ = &idExpression{ $1 , "" } }
    | TK_ID multipart_id_suffix { $$ = &idExpression{ $1 , $2 }  }
    | TK_LEFT_PAR relational_expression TK_RIGHT_PAR { $$ = $2 }
;

multipart_id_suffix: TK_DOT TK_ID { $$ = $2 }
;

truth_value: KW_TRUE { $$ = &trueExpression{} }
    | KW_FALSE { $$ = &falseExpression{} }
    | KW_NULL { $$ = &nullExpression{} }
;

%%

func init() {
    yyErrorVerbose = true
}

func (l *Lexer) Error(s string) {
	panic(Error{
        line: l.Line() + 1,
        column: l.Column() + 1,
        message: s,
    })
}

func Parse(in io.Reader) (commands []interface{}, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = r.(error)
        }else{
            lock.Unlock()
        }
    }()    

    yyParse(NewLexer(in))

    return statements.convert(), nil
}