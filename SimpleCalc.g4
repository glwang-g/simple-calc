grammar SimpleCalc;
//语法规则
prog: ( funcStatement | letStatement )* EOF;

funcStatement: FUNCTION IDENT LPAREN (IDENT (COMMA IDENT)*)? RPAREN LBRACE letStatement* RBRACE;
letStatement: LET IDENT ASSIGN (callFunc | expression) SEMICOLON;

callFunc: IDENT LPAREN (IDENT (COMMA IDENT)*)? RPAREN;

expression:
    | expression PLUS expression
    | expression MINUS expression
    | expression MULTIPLY expression
    | expression DIVIDE expression
    | INT
    | FLOAT
    | IDENT
    | LPAREN expression RPAREN;


// 词法规则
LET: 'let';
FUNCTION: 'fn';
IF: 'if';
ELSE: 'else';
RETURN: 'return';
TRUE: 'true';
FALSE: 'false';

IDENT: [a-zA-Z_][a-zA-Z0-9_]*;
INT: [0-9]+;
FLOAT: [0-9]+'.'[0-9]+;
NOT_EQ: '!=';
EQ: '==';
AND: '&&';
OR: '||';


ASSIGN: '=';
PLUS: '+';
MINUS: '-';
BANG: '!';
MULTIPLY: '*';
DIVIDE: '/';
GT: '>';
LT: '<';
GTE: '>=';
LTE: '<=';

COMMA: ',';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
SEMICOLON: ';';

// 忽略空白字符
WHITESPACE: [ \t\r\n]+ -> skip;

// 处理非法字符
ILLEGAL: .;
