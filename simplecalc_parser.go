// Code generated from SimpleCalc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main // SimpleCalc
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SimpleCalcParser struct {
	*antlr.BaseParser
}

var SimpleCalcParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplecalcParserInit() {
	staticData := &SimpleCalcParserStaticData
	staticData.LiteralNames = []string{
		"", "'let'", "'fn'", "'if'", "'else'", "'return'", "'true'", "'false'",
		"", "", "", "'!='", "'=='", "'&&'", "'||'", "'='", "'+'", "'-'", "'!'",
		"'*'", "'/'", "'>'", "'<'", "'>='", "'<='", "','", "'('", "')'", "'{'",
		"'}'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "LET", "FUNCTION", "IF", "ELSE", "RETURN", "TRUE", "FALSE", "IDENT",
		"INT", "FLOAT", "NOT_EQ", "EQ", "AND", "OR", "ASSIGN", "PLUS", "MINUS",
		"BANG", "MULTIPLY", "DIVIDE", "GT", "LT", "GTE", "LTE", "COMMA", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "SEMICOLON", "WHITESPACE", "ILLEGAL",
	}
	staticData.RuleNames = []string{
		"prog", "funcStatement", "letStatement", "callFunc", "expression",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 93, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 5, 0, 13, 8, 0, 10, 0, 12, 0, 16, 9, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 26, 8, 1, 10, 1, 12, 1, 29, 9, 1, 3,
		1, 31, 8, 1, 1, 1, 1, 1, 1, 1, 5, 1, 36, 8, 1, 10, 1, 12, 1, 39, 9, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 48, 8, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 57, 8, 3, 10, 3, 12, 3, 60, 9, 3, 3, 3,
		62, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3,
		4, 74, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 5, 4, 88, 8, 4, 10, 4, 12, 4, 91, 9, 4, 1, 4, 0, 1, 8, 5, 0,
		2, 4, 6, 8, 0, 0, 103, 0, 14, 1, 0, 0, 0, 2, 19, 1, 0, 0, 0, 4, 42, 1,
		0, 0, 0, 6, 51, 1, 0, 0, 0, 8, 73, 1, 0, 0, 0, 10, 13, 3, 2, 1, 0, 11,
		13, 3, 4, 2, 0, 12, 10, 1, 0, 0, 0, 12, 11, 1, 0, 0, 0, 13, 16, 1, 0, 0,
		0, 14, 12, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 17, 1, 0, 0, 0, 16, 14,
		1, 0, 0, 0, 17, 18, 5, 0, 0, 1, 18, 1, 1, 0, 0, 0, 19, 20, 5, 2, 0, 0,
		20, 21, 5, 8, 0, 0, 21, 30, 5, 26, 0, 0, 22, 27, 5, 8, 0, 0, 23, 24, 5,
		25, 0, 0, 24, 26, 5, 8, 0, 0, 25, 23, 1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27,
		25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 27, 1, 0, 0,
		0, 30, 22, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 33,
		5, 27, 0, 0, 33, 37, 5, 28, 0, 0, 34, 36, 3, 4, 2, 0, 35, 34, 1, 0, 0,
		0, 36, 39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40,
		1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 40, 41, 5, 29, 0, 0, 41, 3, 1, 0, 0, 0,
		42, 43, 5, 1, 0, 0, 43, 44, 5, 8, 0, 0, 44, 47, 5, 15, 0, 0, 45, 48, 3,
		6, 3, 0, 46, 48, 3, 8, 4, 0, 47, 45, 1, 0, 0, 0, 47, 46, 1, 0, 0, 0, 48,
		49, 1, 0, 0, 0, 49, 50, 5, 30, 0, 0, 50, 5, 1, 0, 0, 0, 51, 52, 5, 8, 0,
		0, 52, 61, 5, 26, 0, 0, 53, 58, 5, 8, 0, 0, 54, 55, 5, 25, 0, 0, 55, 57,
		5, 8, 0, 0, 56, 54, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0,
		58, 59, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 53, 1,
		0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 64, 5, 27, 0, 0, 64,
		7, 1, 0, 0, 0, 65, 74, 6, 4, -1, 0, 66, 74, 5, 9, 0, 0, 67, 74, 5, 10,
		0, 0, 68, 74, 5, 8, 0, 0, 69, 70, 5, 26, 0, 0, 70, 71, 3, 8, 4, 0, 71,
		72, 5, 27, 0, 0, 72, 74, 1, 0, 0, 0, 73, 65, 1, 0, 0, 0, 73, 66, 1, 0,
		0, 0, 73, 67, 1, 0, 0, 0, 73, 68, 1, 0, 0, 0, 73, 69, 1, 0, 0, 0, 74, 89,
		1, 0, 0, 0, 75, 76, 10, 8, 0, 0, 76, 77, 5, 16, 0, 0, 77, 88, 3, 8, 4,
		9, 78, 79, 10, 7, 0, 0, 79, 80, 5, 17, 0, 0, 80, 88, 3, 8, 4, 8, 81, 82,
		10, 6, 0, 0, 82, 83, 5, 19, 0, 0, 83, 88, 3, 8, 4, 7, 84, 85, 10, 5, 0,
		0, 85, 86, 5, 20, 0, 0, 86, 88, 3, 8, 4, 6, 87, 75, 1, 0, 0, 0, 87, 78,
		1, 0, 0, 0, 87, 81, 1, 0, 0, 0, 87, 84, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0,
		89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 9, 1, 0, 0, 0, 91, 89, 1, 0,
		0, 0, 11, 12, 14, 27, 30, 37, 47, 58, 61, 73, 87, 89,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SimpleCalcParserInit initializes any static state used to implement SimpleCalcParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSimpleCalcParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleCalcParserInit() {
	staticData := &SimpleCalcParserStaticData
	staticData.once.Do(simplecalcParserInit)
}

// NewSimpleCalcParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSimpleCalcParser(input antlr.TokenStream) *SimpleCalcParser {
	SimpleCalcParserInit()
	this := new(SimpleCalcParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SimpleCalcParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SimpleCalc.g4"

	return this
}

// SimpleCalcParser tokens.
const (
	SimpleCalcParserEOF        = antlr.TokenEOF
	SimpleCalcParserLET        = 1
	SimpleCalcParserFUNCTION   = 2
	SimpleCalcParserIF         = 3
	SimpleCalcParserELSE       = 4
	SimpleCalcParserRETURN     = 5
	SimpleCalcParserTRUE       = 6
	SimpleCalcParserFALSE      = 7
	SimpleCalcParserIDENT      = 8
	SimpleCalcParserINT        = 9
	SimpleCalcParserFLOAT      = 10
	SimpleCalcParserNOT_EQ     = 11
	SimpleCalcParserEQ         = 12
	SimpleCalcParserAND        = 13
	SimpleCalcParserOR         = 14
	SimpleCalcParserASSIGN     = 15
	SimpleCalcParserPLUS       = 16
	SimpleCalcParserMINUS      = 17
	SimpleCalcParserBANG       = 18
	SimpleCalcParserMULTIPLY   = 19
	SimpleCalcParserDIVIDE     = 20
	SimpleCalcParserGT         = 21
	SimpleCalcParserLT         = 22
	SimpleCalcParserGTE        = 23
	SimpleCalcParserLTE        = 24
	SimpleCalcParserCOMMA      = 25
	SimpleCalcParserLPAREN     = 26
	SimpleCalcParserRPAREN     = 27
	SimpleCalcParserLBRACE     = 28
	SimpleCalcParserRBRACE     = 29
	SimpleCalcParserSEMICOLON  = 30
	SimpleCalcParserWHITESPACE = 31
	SimpleCalcParserILLEGAL    = 32
)

// SimpleCalcParser rules.
const (
	SimpleCalcParserRULE_prog          = 0
	SimpleCalcParserRULE_funcStatement = 1
	SimpleCalcParserRULE_letStatement  = 2
	SimpleCalcParserRULE_callFunc      = 3
	SimpleCalcParserRULE_expression    = 4
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllFuncStatement() []IFuncStatementContext
	FuncStatement(i int) IFuncStatementContext
	AllLetStatement() []ILetStatementContext
	LetStatement(i int) ILetStatementContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleCalcParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserEOF, 0)
}

func (s *ProgContext) AllFuncStatement() []IFuncStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncStatementContext); ok {
			len++
		}
	}

	tst := make([]IFuncStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncStatementContext); ok {
			tst[i] = t.(IFuncStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) FuncStatement(i int) IFuncStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncStatementContext)
}

func (s *ProgContext) AllLetStatement() []ILetStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILetStatementContext); ok {
			len++
		}
	}

	tst := make([]ILetStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILetStatementContext); ok {
			tst[i] = t.(ILetStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) LetStatement(i int) ILetStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetStatementContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *SimpleCalcParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimpleCalcParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleCalcParserLET || _la == SimpleCalcParserFUNCTION {
		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SimpleCalcParserFUNCTION:
			{
				p.SetState(10)
				p.FuncStatement()
			}

		case SimpleCalcParserLET:
			{
				p.SetState(11)
				p.LetStatement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(16)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(17)
		p.Match(SimpleCalcParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncStatementContext is an interface to support dynamic dispatch.
type IFuncStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllLetStatement() []ILetStatementContext
	LetStatement(i int) ILetStatementContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncStatementContext differentiates from other interfaces.
	IsFuncStatementContext()
}

type FuncStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncStatementContext() *FuncStatementContext {
	var p = new(FuncStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_funcStatement
	return p
}

func InitEmptyFuncStatementContext(p *FuncStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_funcStatement
}

func (*FuncStatementContext) IsFuncStatementContext() {}

func NewFuncStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncStatementContext {
	var p = new(FuncStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleCalcParserRULE_funcStatement

	return p
}

func (s *FuncStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncStatementContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserFUNCTION, 0)
}

func (s *FuncStatementContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SimpleCalcParserIDENT)
}

func (s *FuncStatementContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserIDENT, i)
}

func (s *FuncStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserLPAREN, 0)
}

func (s *FuncStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserRPAREN, 0)
}

func (s *FuncStatementContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserLBRACE, 0)
}

func (s *FuncStatementContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserRBRACE, 0)
}

func (s *FuncStatementContext) AllLetStatement() []ILetStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILetStatementContext); ok {
			len++
		}
	}

	tst := make([]ILetStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILetStatementContext); ok {
			tst[i] = t.(ILetStatementContext)
			i++
		}
	}

	return tst
}

func (s *FuncStatementContext) LetStatement(i int) ILetStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetStatementContext)
}

func (s *FuncStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SimpleCalcParserCOMMA)
}

func (s *FuncStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserCOMMA, i)
}

func (s *FuncStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.EnterFuncStatement(s)
	}
}

func (s *FuncStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.ExitFuncStatement(s)
	}
}

func (p *SimpleCalcParser) FuncStatement() (localctx IFuncStatementContext) {
	localctx = NewFuncStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SimpleCalcParserRULE_funcStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Match(SimpleCalcParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(20)
		p.Match(SimpleCalcParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(21)
		p.Match(SimpleCalcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleCalcParserIDENT {
		{
			p.SetState(22)
			p.Match(SimpleCalcParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleCalcParserCOMMA {
			{
				p.SetState(23)
				p.Match(SimpleCalcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(24)
				p.Match(SimpleCalcParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(32)
		p.Match(SimpleCalcParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(SimpleCalcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleCalcParserLET {
		{
			p.SetState(34)
			p.LetStatement()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(SimpleCalcParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILetStatementContext is an interface to support dynamic dispatch.
type ILetStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LET() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	CallFunc() ICallFuncContext
	Expression() IExpressionContext

	// IsLetStatementContext differentiates from other interfaces.
	IsLetStatementContext()
}

type LetStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetStatementContext() *LetStatementContext {
	var p = new(LetStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_letStatement
	return p
}

func InitEmptyLetStatementContext(p *LetStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_letStatement
}

func (*LetStatementContext) IsLetStatementContext() {}

func NewLetStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetStatementContext {
	var p = new(LetStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleCalcParserRULE_letStatement

	return p
}

func (s *LetStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LetStatementContext) LET() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserLET, 0)
}

func (s *LetStatementContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserIDENT, 0)
}

func (s *LetStatementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserASSIGN, 0)
}

func (s *LetStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserSEMICOLON, 0)
}

func (s *LetStatementContext) CallFunc() ICallFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFuncContext)
}

func (s *LetStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LetStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.EnterLetStatement(s)
	}
}

func (s *LetStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.ExitLetStatement(s)
	}
}

func (p *SimpleCalcParser) LetStatement() (localctx ILetStatementContext) {
	localctx = NewLetStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SimpleCalcParserRULE_letStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(SimpleCalcParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(SimpleCalcParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(SimpleCalcParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(45)
			p.CallFunc()
		}

	case 2:
		{
			p.SetState(46)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(49)
		p.Match(SimpleCalcParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallFuncContext is an interface to support dynamic dispatch.
type ICallFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCallFuncContext differentiates from other interfaces.
	IsCallFuncContext()
}

type CallFuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallFuncContext() *CallFuncContext {
	var p = new(CallFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_callFunc
	return p
}

func InitEmptyCallFuncContext(p *CallFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_callFunc
}

func (*CallFuncContext) IsCallFuncContext() {}

func NewCallFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFuncContext {
	var p = new(CallFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleCalcParserRULE_callFunc

	return p
}

func (s *CallFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFuncContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SimpleCalcParserIDENT)
}

func (s *CallFuncContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserIDENT, i)
}

func (s *CallFuncContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserLPAREN, 0)
}

func (s *CallFuncContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserRPAREN, 0)
}

func (s *CallFuncContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SimpleCalcParserCOMMA)
}

func (s *CallFuncContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserCOMMA, i)
}

func (s *CallFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.EnterCallFunc(s)
	}
}

func (s *CallFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.ExitCallFunc(s)
	}
}

func (p *SimpleCalcParser) CallFunc() (localctx ICallFuncContext) {
	localctx = NewCallFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SimpleCalcParserRULE_callFunc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(SimpleCalcParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(SimpleCalcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleCalcParserIDENT {
		{
			p.SetState(53)
			p.Match(SimpleCalcParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleCalcParserCOMMA {
			{
				p.SetState(54)
				p.Match(SimpleCalcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(55)
				p.Match(SimpleCalcParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(60)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(63)
		p.Match(SimpleCalcParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	IDENT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	RPAREN() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	MULTIPLY() antlr.TerminalNode
	DIVIDE() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleCalcParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) INT() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserINT, 0)
}

func (s *ExpressionContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserFLOAT, 0)
}

func (s *ExpressionContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserIDENT, 0)
}

func (s *ExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserLPAREN, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserRPAREN, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserMINUS, 0)
}

func (s *ExpressionContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserMULTIPLY, 0)
}

func (s *ExpressionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserDIVIDE, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SimpleCalcParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *SimpleCalcParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, SimpleCalcParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:

	case 2:
		{
			p.SetState(66)
			p.Match(SimpleCalcParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(67)
			p.Match(SimpleCalcParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(68)
			p.Match(SimpleCalcParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(69)
			p.Match(SimpleCalcParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.expression(0)
		}
		{
			p.SetState(71)
			p.Match(SimpleCalcParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleCalcParserRULE_expression)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(76)
					p.Match(SimpleCalcParserPLUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(77)
					p.expression(9)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleCalcParserRULE_expression)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(79)
					p.Match(SimpleCalcParserMINUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.expression(8)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleCalcParserRULE_expression)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(82)
					p.Match(SimpleCalcParserMULTIPLY)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(83)
					p.expression(7)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleCalcParserRULE_expression)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(85)
					p.Match(SimpleCalcParserDIVIDE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(86)
					p.expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *SimpleCalcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SimpleCalcParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
