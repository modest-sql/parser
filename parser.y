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
%token KW_ALTER KW_VALUE KW_BETWEEN KW_LIKE KW_INNER_JOIN KW_HAVING KW_SUM KW_COUNT KW_AVG KW_MIN KW_MAX
%token KW_NULL KW_IN  KW_IS TK_QUOTES KW_AUTO_INCREMENT TK_ID

%type<int_t> expression term factor
%token<int_t> NUM
%token<string_t> TK_WORD

%%

input: statements_list {  }
    | /* empty */
;

statements_list: statements_list statement { fmt.Println("Hey there, I'm a statements_list!"); }
    | statement { fmt.Println("Here's a statment"); }
;

statement: expression { fmt.Printf("Result: %d\n", $1); }
    | KW_SELECT { fmt.Println("Heres a SELECT statement"); }
;

expression: expression '+' term { $$ = $1 + $3; }
    | expression '-' term { $$ = $1 - $3; }
    | term
;

term: term '*' factor { $$ = $1 * $3; }
    | term '/' factor { $$ = $1 / $3; }
    | factor
;

factor: NUM
    | '(' expression ')' { $$ = $2; }  
    | TK_ID {fmt.Print("es un id");}
    | TK_WORD {fmt.Print("es un palabra");} 
;