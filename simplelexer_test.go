package main

import (
	"log"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

func TestSimpleCalcLexer(t *testing.T) {
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
	`
	input := antlr.NewInputStream(inputstr)
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
		log.Printf("Token: {%d, %s, %q}", token.GetTokenType(), lexer.SymbolicNames[token.GetTokenType()], token.GetText())
	}
}

func TestSimpleCalcParser(t *testing.T) {
	inputstr := `
		let b = 3;
		fn add(x, y, z) {
			let xyz = 321;
			return xyz;
		}
		let a = 10;
		let ten = 10 + five;

		let add2 = fn (x, y) {
		  let xyz = x + y;
		  return 123+123;
		};

		let result = add(five, ten);
	`
	input := antlr.NewInputStream(inputstr)
	lexer := NewSimpleCalcLexer(input)

	// 创建并添加错误监听器
	errorListener := NewSimpleCalcErrorListener()
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)


	parser := NewSimpleCalcParser(antlr.NewCommonTokenStream(lexer, 0))
	if parser == nil {
		t.Fatal("Failed to create parser")
		return
	}

	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)

	listener := &BaseSimpleCalcListener{}
	parser.AddParseListener(listener)

	log.Printf("Parsing...")

	prog := parser.Prog()
	if prog == nil {
		t.Fatal("Failed to parse program")
	}

	errorListener.PrintErrors()

	log.Printf("post Parsing...")
}
