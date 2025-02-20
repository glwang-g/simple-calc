// Code generated from SimpleCalc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main // SimpleCalc
import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// BaseSimpleCalcListener is a complete listener for a parse tree produced by SimpleCalcParser.
type BaseSimpleCalcListener struct{}

var _ SimpleCalcListener = &BaseSimpleCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("Enter %d, %s, %s \n", ctx.GetRuleIndex(), SimpleCalcParserStaticData.RuleNames[ctx.GetRuleIndex()], ctx.GetText())
}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("Exit %d, %s, %s \n", ctx.GetRuleIndex(), SimpleCalcParserStaticData.RuleNames[ctx.GetRuleIndex()], ctx.GetText())
}

// EnterProg is called when production prog is entered.
func (s *BaseSimpleCalcListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseSimpleCalcListener) ExitProg(ctx *ProgContext) {}

// EnterFuncStatement is called when production funcStatement is entered.
func (s *BaseSimpleCalcListener) EnterFuncStatement(ctx *FuncStatementContext) {}

// ExitFuncStatement is called when production funcStatement is exited.
func (s *BaseSimpleCalcListener) ExitFuncStatement(ctx *FuncStatementContext) {}

// EnterLetStatement is called when production letStatement is entered.
func (s *BaseSimpleCalcListener) EnterLetStatement(ctx *LetStatementContext) {}

// ExitLetStatement is called when production letStatement is exited.
func (s *BaseSimpleCalcListener) ExitLetStatement(ctx *LetStatementContext) {}

// EnterCallFunc is called when production callFunc is entered.
func (s *BaseSimpleCalcListener) EnterCallFunc(ctx *CallFuncContext) {}

// ExitCallFunc is called when production callFunc is exited.
func (s *BaseSimpleCalcListener) ExitCallFunc(ctx *CallFuncContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSimpleCalcListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSimpleCalcListener) ExitExpression(ctx *ExpressionContext) {}
