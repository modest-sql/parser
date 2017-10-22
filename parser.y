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
%token KW_CREATE
%token KW_TABLE KW_DELETE KW_INSERT KW_INTO KW_SELECT KW_WHERE KW_FROM KW_UPDATE KW_SET TK_WORD
%token KW_ALTER KW_VALUE KW_BETWEEN KW_LIKE KW_INNER  KW_HAVING KW_SUM KW_COUNT KW_AVG KW_MIN KW_MAX
%token KW_NULL KW_IN  KW_IS TK_QUOTES KW_AUTO_INCREMENT KW_JOIN KW_DROP

%token<int_t> NUM
%token<string_t> TK_WORD TK_ID

%%

input: statements_list {  }
    | /* empty */
;

statements_list: statements_list statement {  }
    | statement { }
;

statement: data_statement {  }
    | schema_statement { }
;

schema_statement: create_statement { }
    | alter_statement { }
    | drop_statement { }
;

data_statement: select_statement { }
    | insert_statement { }
    | delete_statement { }
    | update_statement { }
;

create_statement: KW_CREATE KW_TABLE { fmt.Printf("Reading CREATE TABLE statement found"); }
;

alter_statement: KW_ALTER { }
;

drop_statement: KW_DROP { }
;

select_statement: KW_SELECT { fmt.Println("Heres a SELECT statement"); }
;

insert_statement: KW_INSERT { }
;

delete_statement: KW_DELETE { }
;

update_statement: KW_UPDATE { }
;


%%

func (l *Lexer) Error(s string) {
	fmt.Printf("Syntax error Ln %d Col %d with: %s Error: %s\n", l.Line(), l.Column(), l.Text(), s)
}