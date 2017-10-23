/* Based on https://ronsavage.github.io/SQL/sql-2003-2.bnf.html */

%{
package parser

import (
    "fmt"
    "io"
)

%}

%union {
    int_t int
    string_t string
    float_t float64
    expr_t expression
    stmt_t statement
}

%token TK_PLUS TK_MINUS TK_STAR TK_DIV TK_LT TK_GT TK_GTE TK_LTE TK_EQ TK_NE
%token TK_LEFT_PAR TK_RIGHT_PAR TK_COMMA TK_DOT TK_SEMICOLON
%token KW_OR KW_AND KW_NOT KW_INTEGER KW_FLOAT KW_CHAR KW_BOOLEAN KW_DATETIME
%token KW_CREATE KW_TABLE KW_DELETE KW_INSERT
%token KW_INTO KW_SELECT KW_WHERE KW_FROM KW_UPDATE KW_SET
%token KW_ALTER KW_VALUES KW_BETWEEN KW_LIKE KW_INNER
%token KW_HAVING KW_SUM KW_COUNT KW_AVG KW_MIN KW_MAX
%token KW_NULL KW_IN KW_IS KW_AUTO_INCREMENT KW_JOIN KW_ON KW_GROUP KW_BY KW_DROP KW_DEFAULT
%token KW_TRUE KW_FALSE KW_AS KW_ADD KW_COLUMN

%token<int_t> INT_LIT
%token<float_t> FLOAT_LIT
%token<string_t> TK_ID STR_LIT

%%

input: statements_list {  }
    | { }
;

statements_list: statements_list statement { fmt.Println("Hey there, I'm a statements_list!\n"); }
    | statement { }
;

statement: data_statement { fmt.Printf("Data access statement found\n"); }
    | schema_statement { fmt.Printf("Schema Definition/Manipulation statement found\n"); }
;

schema_statement: create_statement {  }
    | alter_statement { }
    | drop_statement { }
;

data_statement: select_statement { }
    | insert_statement { }
    | delete_statement { }
    | update_statement { }
;

create_statement: KW_CREATE KW_TABLE TK_ID TK_LEFT_PAR table_element_list TK_RIGHT_PAR TK_SEMICOLON  {  }
;

table_element_list: table_element_list TK_COMMA table_element {  }
    | table_element {  }
;

table_element: column_definition {  }
;

column_definition: TK_ID data_type column_constraint_list {  }
;

data_type: KW_CHAR TK_LEFT_PAR INT_LIT TK_RIGHT_PAR { }
    | KW_INTEGER { }
;

column_constraint_list: column_constraint_list column_constraint { }
    | column_constraint { }
    | { }
;

column_constraint: constr_not_null { }
;

constr_not_null: KW_NOT KW_NULL { }
    | KW_DEFAULT value_literal { }
    | KW_AUTO_INCREMENT { }
;

alter_statement: KW_ALTER KW_TABLE TK_ID alter_instruction TK_SEMICOLON { }
;

alter_instruction: KW_ADD TK_ID { }
    | KW_DROP KW_COLUMN TK_ID data_type column_constraint_list { }
;

drop_statement: KW_DROP KW_TABLE TK_ID { }
;

select_statement: KW_SELECT select_col_list KW_FROM TK_ID opt_alias_spec where_clause TK_SEMICOLON {  }
;

select_col_list: select_col_list TK_COMMA select_col { }
    | select_col { }
;

select_col: TK_STAR { }
    | TK_ID opt_multipart_id_suffix opt_alias_spec { }
    | value_literal opt_alias_spec { }
    | truth_value opt_alias_spec
;

insert_statement: KW_INSERT KW_INTO TK_ID TK_LEFT_PAR column_names_list TK_RIGHT_PAR KW_VALUES values_tuples_list TK_SEMICOLON { }
;

column_names_list: column_names_list TK_COMMA TK_ID { }
    | TK_ID { }
;

values_tuples_list: values_tuples_list TK_COMMA values_tuple { }
    | values_tuple { }
;

values_tuple: TK_LEFT_PAR values_list TK_RIGHT_PAR
;

values_list: values_list TK_COMMA value_literal
    | value_literal
;

value_literal: STR_LIT
    | INT_LIT
;

delete_statement: KW_DELETE TK_ID opt_alias_spec where_clause TK_SEMICOLON { }
;

opt_alias_spec: KW_AS TK_ID { }
    | TK_ID { }
    | { }
;

where_clause: KW_WHERE search_condition { }
;

search_condition: boolean_value_expression
;

update_statement: KW_UPDATE TK_ID set_list where_clause TK_SEMICOLON { }
;

set_list: KW_SET set_assignments_list { }
;

set_assignments_list: set_assignments_list TK_COMMA set_assignment { }
    | set_assignment { }
;

set_assignment: TK_ID TK_EQ relational_term { }
;

boolean_value_expression: boolean_value_expression KW_OR boolean_term { }
    | boolean_term { }
;

boolean_term: boolean_term KW_AND boolean_factor { }
    | boolean_factor { }
;

boolean_factor: KW_NOT relational_expression { }
    | relational_expression { }
;

relational_expression: relational_expression TK_LT relational_term { }
    | relational_expression TK_GT relational_term { }
    | relational_expression TK_EQ relational_term { }
    | relational_expression TK_LTE relational_term { }
    | relational_expression TK_GTE relational_term { }
    | relational_expression TK_NE relational_term { }
    | relational_expression KW_LIKE relational_term { }
    | relational_expression KW_BETWEEN between_term { }
    | relational_term { }
;

between_term: INT_LIT KW_AND INT_LIT { }
;

relational_term: relational_term TK_PLUS relational_factor { }
    | relational_term TK_MINUS relational_factor { }
    | relational_factor { }
;

relational_factor: relational_factor TK_STAR addi_factor { }
    | relational_factor TK_DIV addi_factor { }
    | addi_factor { }
;

addi_factor: INT_LIT { }
    | truth_value { }
    | STR_LIT { }
    | TK_ID opt_multipart_id_suffix { }
    | TK_LEFT_PAR relational_expression TK_RIGHT_PAR { }
;

opt_multipart_id_suffix: TK_DOT TK_ID
    | { }
;

truth_value: KW_TRUE { }
    | KW_FALSE { }
    | KW_NULL
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

func Parse(in io.Reader) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = r.(error)
        }
    }()    

    yyParse(NewLexer(in))

    return nil
}