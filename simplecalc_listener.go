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

	// EnterCallFunc is called when entering the callFunc production.
	EnterCallFunc(c *CallFuncContext)

	// EnterAnonFunction is called when entering the anonFunction production.
	EnterAnonFunction(c *AnonFunctionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitFuncStatement is called when exiting the funcStatement production.
	ExitFuncStatement(c *FuncStatementContext)

	// ExitLetStatement is called when exiting the letStatement production.
	ExitLetStatement(c *LetStatementContext)

	// ExitCallFunc is called when exiting the callFunc production.
	ExitCallFunc(c *CallFuncContext)

	// ExitAnonFunction is called when exiting the anonFunction production.
	ExitAnonFunction(c *AnonFunctionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)
}
