/* Based on https://ronsavage.github.io/SQL/sql-2003-2.bnf.html */

%{
    package main
    import "fmt"
%}

%union {
    int_t int
    string_t string
}

%token '+' '-' '*' '/' '(' ')' ',' '.' ';' '=' '<' '>' TK_GTE TK_LTE TK_NE KW_OR KW_AND KW_NOT KW_INTEGER KW_CHAR 
%token KW_CREATE KW_TABLE KW_DELETE KW_INSERT KW_INTO KW_SELECT KW_WHERE KW_FROM KW_UPDATE KW_SET TK_WORD
%token KW_ALTER KW_VALUES KW_BETWEEN KW_LIKE KW_INNER  KW_HAVING KW_SUM KW_COUNT KW_AVG KW_MIN KW_MAX
%token KW_NULL KW_IN  KW_IS TK_QUOTES KW_AUTO_INCREMENT KW_JOIN KW_DROP KW_DEFAULT

%token<int_t> NUM
%token<string_t> TK_WORD TK_ID

%%

input: statements_list {  }
    | /* empty */
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

create_statement: KW_CREATE KW_TABLE TK_ID '(' table_element_list ')' ';'  {  }
;

table_element_list: table_element_list ',' table_element {  }
    | table_element {  }
;

table_element: column_definition {  }
;

column_definition: TK_ID data_type column_constraint_list {  }
;

data_type: KW_CHAR '(' NUM ')' { }
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
;

alter_statement: KW_ALTER KW_TABLE { }
;

drop_statement: KW_DROP KW_TABLE TK_ID { }
;

select_statement: KW_SELECT  {  }
;

insert_statement: KW_INSERT KW_INTO TK_ID '(' column_names_list ')' KW_VALUES values_tuples_list ';' { }
;

column_names_list: column_names_list ',' TK_ID { }
    | TK_ID { }
;

values_tuples_list: values_tuples_list ',' values_tuple { }
    | values_tuple { }
;

values_tuple: '(' values_list ')'
;

values_list: values_list ',' value_literal
    | value_literal
;

value_literal: TK_WORD
    | NUM
;

delete_statement: KW_DELETE KW_TABLE TK_ID { }
;

update_statement: KW_UPDATE { }
;

%%

func (l *Lexer) Error(s string) {
	fmt.Printf("Syntax error at Ln %d Col %d: %s with input %s\n", l.Line(), l.Column(), s, l.Text())
}