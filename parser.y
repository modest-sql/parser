%{
package main
import "fmt"
%}

%union {
    int_t int
}

%type<int_t> expression term factor

%token '+' '-' '*' '/' '(' ')'

%token<int_t> NUM

%%

input: expression { fmt.Printf("Result: %d\n", $1); }
    | /* empty */
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