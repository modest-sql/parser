%{
package main
import "fmt"
%}

%union {
    int_t int
    string_t string
}

%type<int_t> expression term factor

%token '+' '-' '*' '/' '(' ')'

%token<int_t> NUM
%token<string_t> KW_SELECT KW_FROM KW_WHERE

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
;