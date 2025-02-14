// Code generated from SimpleCalc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main // SimpleCalc
import "github.com/antlr4-go/antlr/v4"

// SimpleCalcListener is a complete listener for a parse tree produced by SimpleCalcParser.
type SimpleCalcListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterFuncStatement is called when entering the funcStatement production.
	EnterFuncStatement(c *FuncStatementContext)

	// EnterLetStatement is called when entering the letStatement production.
	EnterLetStatement(c *LetStatementContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitFuncStatement is called when exiting the funcStatement production.
	ExitFuncStatement(c *FuncStatementContext)

	// ExitLetStatement is called when exiting the letStatement production.
	ExitLetStatement(c *LetStatementContext)
}
