grammar SimpleCalc;
//语法规则
prog: ( funcStatement | letStatement )* EOF;

funcStatement: FUNCTION IDENT LPAREN (IDENT (COMMA IDENT)*)? RPAREN LBRACE letStatement* RBRACE;
letStatement: LET IDENT ASSIGN INT SEMICOLON;

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
NOT_EQ: '!=';
EQ: '==';

ASSIGN: '=';
PLUS: '+';
MINUS: '-';
BANG: '!';
MULTIPLY: '*';
DIVIDE: '/';
GT: '>';
LT: '<';

COMMA: ',';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
SEMICOLON: ';';

// 忽略空白字符
WS: [ \t\r\n]+ -> skip;

// 处理非法字符/25
ILLEGAL: .;