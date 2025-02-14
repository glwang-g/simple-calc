package main

import (
	"log"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

func TestNewSimpleCalc(t *testing.T) {
	inputstr := `
		let five = 5;
		let add = fn(x, y) {
		x + y;
		};

		let result = add(five, ten);
		!-/*5;
		5 < 10 > 5;
		if (5 < 10) {
			return true;
		} else {
			return false;
		}
		10 == 10;
		10 != 9;
	`;
	input := antlr.NewInputStream(inputstr);
	lexer := NewSimpleCalcLexer(input)
	if lexer == nil {
		t.Error("NewSimpleLexer returned nil")
	}
	// 循环读取每个 token，直到遇到 EOF
	for {
		token := lexer.NextToken()
		if token.GetTokenType() == antlr.TokenEOF {
			break
		}
		if token.GetTokenType() == SimpleCalcLexerILLEGAL {
			t.Errorf("Illegal token: %q", token.GetText())
		}
		// 打印 token 的类型和文本
		log.SetFlags(0)
		log.Printf("Token: {%s, %q}", lexer.SymbolicNames[token.GetTokenType()], token.GetText())
	}	
}


func TestNewSimpleCalcParser(t *testing.T) {
	inputstr := `
		let b = 3;
		fn add(x, y, z) {
			let xyz = 321;
		}
		let a = 10;
	`;
	/**
		let ten = 10 + five;


		let result = add(five, ten);
		!-/*5;
		5 < 10 > 5;
		if (5 < 10) {
			return true;
		} else {
			return false;
		}
		10 == 10;
		10 != 9;
	*/
	input := antlr.NewInputStream(inputstr)
	lexer := NewSimpleCalcLexer(input)
	parser := NewSimpleCalcParser(antlr.NewCommonTokenStream(lexer, 0))
	if parser == nil {
		t.Errorf("Expected parser to be non-nil")
	}

	 listener := &BaseSimpleCalcListener{}
	 parser.AddParseListener(listener);
	 log.Printf("Prog: %s", parser.Prog().FuncStatement(0).GetText());
}
