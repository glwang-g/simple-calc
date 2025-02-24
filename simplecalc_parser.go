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
		"prog", "funcStatement", "letStatement", "callFunc", "anonFunction",
		"expression", "returnStatement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 129, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 5, 0, 17, 8, 0, 10, 0, 12, 0, 20,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 30, 8, 1, 10,
		1, 12, 1, 33, 9, 1, 3, 1, 35, 8, 1, 1, 1, 1, 1, 1, 1, 5, 1, 40, 8, 1, 10,
		1, 12, 1, 43, 9, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 54, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 63,
		8, 3, 10, 3, 12, 3, 66, 9, 3, 3, 3, 68, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 5, 4, 77, 8, 4, 10, 4, 12, 4, 80, 9, 4, 3, 4, 82, 8, 4,
		1, 4, 1, 4, 1, 4, 5, 4, 87, 8, 4, 10, 4, 12, 4, 90, 9, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 103, 8, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5,
		5, 117, 8, 5, 10, 5, 12, 5, 120, 9, 5, 1, 6, 1, 6, 1, 6, 3, 6, 125, 8,
		6, 1, 6, 1, 6, 1, 6, 0, 1, 10, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 142, 0,
		18, 1, 0, 0, 0, 2, 23, 1, 0, 0, 0, 4, 47, 1, 0, 0, 0, 6, 57, 1, 0, 0, 0,
		8, 71, 1, 0, 0, 0, 10, 102, 1, 0, 0, 0, 12, 121, 1, 0, 0, 0, 14, 17, 3,
		2, 1, 0, 15, 17, 3, 4, 2, 0, 16, 14, 1, 0, 0, 0, 16, 15, 1, 0, 0, 0, 17,
		20, 1, 0, 0, 0, 18, 16, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 21, 1, 0, 0,
		0, 20, 18, 1, 0, 0, 0, 21, 22, 5, 0, 0, 1, 22, 1, 1, 0, 0, 0, 23, 24, 5,
		2, 0, 0, 24, 25, 5, 8, 0, 0, 25, 34, 5, 26, 0, 0, 26, 31, 5, 8, 0, 0, 27,
		28, 5, 25, 0, 0, 28, 30, 5, 8, 0, 0, 29, 27, 1, 0, 0, 0, 30, 33, 1, 0,
		0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31,
		1, 0, 0, 0, 34, 26, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0,
		36, 37, 5, 27, 0, 0, 37, 41, 5, 28, 0, 0, 38, 40, 3, 4, 2, 0, 39, 38, 1,
		0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42,
		44, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 45, 3, 12, 6, 0, 45, 46, 5, 29,
		0, 0, 46, 3, 1, 0, 0, 0, 47, 48, 5, 1, 0, 0, 48, 49, 5, 8, 0, 0, 49, 53,
		5, 15, 0, 0, 50, 54, 3, 8, 4, 0, 51, 54, 3, 6, 3, 0, 52, 54, 3, 10, 5,
		0, 53, 50, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 52, 1, 0, 0, 0, 54, 55,
		1, 0, 0, 0, 55, 56, 5, 30, 0, 0, 56, 5, 1, 0, 0, 0, 57, 58, 5, 8, 0, 0,
		58, 67, 5, 26, 0, 0, 59, 64, 5, 8, 0, 0, 60, 61, 5, 25, 0, 0, 61, 63, 5,
		8, 0, 0, 62, 60, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64,
		65, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 59, 1, 0, 0,
		0, 67, 68, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 5, 27, 0, 0, 70, 7,
		1, 0, 0, 0, 71, 72, 5, 2, 0, 0, 72, 81, 5, 26, 0, 0, 73, 78, 5, 8, 0, 0,
		74, 75, 5, 25, 0, 0, 75, 77, 5, 8, 0, 0, 76, 74, 1, 0, 0, 0, 77, 80, 1,
		0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80,
		78, 1, 0, 0, 0, 81, 73, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 1, 0, 0,
		0, 83, 84, 5, 27, 0, 0, 84, 88, 5, 28, 0, 0, 85, 87, 3, 4, 2, 0, 86, 85,
		1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0,
		89, 91, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 92, 3, 12, 6, 0, 92, 93, 5,
		29, 0, 0, 93, 9, 1, 0, 0, 0, 94, 103, 6, 5, -1, 0, 95, 103, 5, 9, 0, 0,
		96, 103, 5, 10, 0, 0, 97, 103, 5, 8, 0, 0, 98, 99, 5, 26, 0, 0, 99, 100,
		3, 10, 5, 0, 100, 101, 5, 27, 0, 0, 101, 103, 1, 0, 0, 0, 102, 94, 1, 0,
		0, 0, 102, 95, 1, 0, 0, 0, 102, 96, 1, 0, 0, 0, 102, 97, 1, 0, 0, 0, 102,
		98, 1, 0, 0, 0, 103, 118, 1, 0, 0, 0, 104, 105, 10, 8, 0, 0, 105, 106,
		5, 16, 0, 0, 106, 117, 3, 10, 5, 9, 107, 108, 10, 7, 0, 0, 108, 109, 5,
		17, 0, 0, 109, 117, 3, 10, 5, 8, 110, 111, 10, 6, 0, 0, 111, 112, 5, 19,
		0, 0, 112, 117, 3, 10, 5, 7, 113, 114, 10, 5, 0, 0, 114, 115, 5, 20, 0,
		0, 115, 117, 3, 10, 5, 6, 116, 104, 1, 0, 0, 0, 116, 107, 1, 0, 0, 0, 116,
		110, 1, 0, 0, 0, 116, 113, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116,
		1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 11, 1, 0, 0, 0, 120, 118, 1, 0,
		0, 0, 121, 124, 5, 5, 0, 0, 122, 125, 3, 10, 5, 0, 123, 125, 3, 8, 4, 0,
		124, 122, 1, 0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126,
		127, 5, 30, 0, 0, 127, 13, 1, 0, 0, 0, 15, 16, 18, 31, 34, 41, 53, 64,
		67, 78, 81, 88, 102, 116, 118, 124,
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
	SimpleCalcParserRULE_prog            = 0
	SimpleCalcParserRULE_funcStatement   = 1
	SimpleCalcParserRULE_letStatement    = 2
	SimpleCalcParserRULE_callFunc        = 3
	SimpleCalcParserRULE_anonFunction    = 4
	SimpleCalcParserRULE_expression      = 5
	SimpleCalcParserRULE_returnStatement = 6
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
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleCalcParserLET || _la == SimpleCalcParserFUNCTION {
		p.SetState(16)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SimpleCalcParserFUNCTION:
			{
				p.SetState(14)
				p.FuncStatement()
			}

		case SimpleCalcParserLET:
			{
				p.SetState(15)
				p.LetStatement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(21)
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
	ReturnStatement() IReturnStatementContext
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

func (s *FuncStatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
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
		p.SetState(23)
		p.Match(SimpleCalcParserFUNCTION)
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
	{
		p.SetState(25)
		p.Match(SimpleCalcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleCalcParserIDENT {
		{
			p.SetState(26)
			p.Match(SimpleCalcParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleCalcParserCOMMA {
			{
				p.SetState(27)
				p.Match(SimpleCalcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(28)
				p.Match(SimpleCalcParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(33)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(36)
		p.Match(SimpleCalcParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		p.Match(SimpleCalcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleCalcParserLET {
		{
			p.SetState(38)
			p.LetStatement()
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
		p.ReturnStatement()
	}
	{
		p.SetState(45)
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
	AnonFunction() IAnonFunctionContext
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

func (s *LetStatementContext) AnonFunction() IAnonFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonFunctionContext)
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
		p.SetState(47)
		p.Match(SimpleCalcParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.Match(SimpleCalcParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(SimpleCalcParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(50)
			p.AnonFunction()
		}

	case 2:
		{
			p.SetState(51)
			p.CallFunc()
		}

	case 3:
		{
			p.SetState(52)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(55)
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
		p.SetState(57)
		p.Match(SimpleCalcParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(SimpleCalcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleCalcParserIDENT {
		{
			p.SetState(59)
			p.Match(SimpleCalcParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleCalcParserCOMMA {
			{
				p.SetState(60)
				p.Match(SimpleCalcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(61)
				p.Match(SimpleCalcParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(69)
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

// IAnonFunctionContext is an interface to support dynamic dispatch.
type IAnonFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	ReturnStatement() IReturnStatementContext
	RBRACE() antlr.TerminalNode
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode
	AllLetStatement() []ILetStatementContext
	LetStatement(i int) ILetStatementContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAnonFunctionContext differentiates from other interfaces.
	IsAnonFunctionContext()
}

type AnonFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnonFunctionContext() *AnonFunctionContext {
	var p = new(AnonFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_anonFunction
	return p
}

func InitEmptyAnonFunctionContext(p *AnonFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_anonFunction
}

func (*AnonFunctionContext) IsAnonFunctionContext() {}

func NewAnonFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnonFunctionContext {
	var p = new(AnonFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleCalcParserRULE_anonFunction

	return p
}

func (s *AnonFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AnonFunctionContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserFUNCTION, 0)
}

func (s *AnonFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserLPAREN, 0)
}

func (s *AnonFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserRPAREN, 0)
}

func (s *AnonFunctionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserLBRACE, 0)
}

func (s *AnonFunctionContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *AnonFunctionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserRBRACE, 0)
}

func (s *AnonFunctionContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SimpleCalcParserIDENT)
}

func (s *AnonFunctionContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserIDENT, i)
}

func (s *AnonFunctionContext) AllLetStatement() []ILetStatementContext {
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

func (s *AnonFunctionContext) LetStatement(i int) ILetStatementContext {
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

func (s *AnonFunctionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SimpleCalcParserCOMMA)
}

func (s *AnonFunctionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserCOMMA, i)
}

func (s *AnonFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnonFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnonFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.EnterAnonFunction(s)
	}
}

func (s *AnonFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.ExitAnonFunction(s)
	}
}

func (p *SimpleCalcParser) AnonFunction() (localctx IAnonFunctionContext) {
	localctx = NewAnonFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SimpleCalcParserRULE_anonFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(SimpleCalcParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Match(SimpleCalcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleCalcParserIDENT {
		{
			p.SetState(73)
			p.Match(SimpleCalcParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleCalcParserCOMMA {
			{
				p.SetState(74)
				p.Match(SimpleCalcParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(75)
				p.Match(SimpleCalcParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(83)
		p.Match(SimpleCalcParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(SimpleCalcParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleCalcParserLET {
		{
			p.SetState(85)
			p.LetStatement()
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
		p.ReturnStatement()
	}
	{
		p.SetState(92)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, SimpleCalcParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:

	case 2:
		{
			p.SetState(95)
			p.Match(SimpleCalcParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(96)
			p.Match(SimpleCalcParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(97)
			p.Match(SimpleCalcParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(98)
			p.Match(SimpleCalcParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.expression(0)
		}
		{
			p.SetState(100)
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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(116)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleCalcParserRULE_expression)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(105)
					p.Match(SimpleCalcParserPLUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(106)
					p.expression(9)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleCalcParserRULE_expression)
				p.SetState(107)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(108)
					p.Match(SimpleCalcParserMINUS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(109)
					p.expression(8)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleCalcParserRULE_expression)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(111)
					p.Match(SimpleCalcParserMULTIPLY)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(112)
					p.expression(7)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SimpleCalcParserRULE_expression)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(114)
					p.Match(SimpleCalcParserDIVIDE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(115)
					p.expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
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

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	Expression() IExpressionContext
	AnonFunction() IAnonFunctionContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SimpleCalcParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleCalcParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserRETURN, 0)
}

func (s *ReturnStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserSEMICOLON, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
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

func (s *ReturnStatementContext) AnonFunction() IAnonFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonFunctionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleCalcListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *SimpleCalcParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SimpleCalcParserRULE_returnStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(SimpleCalcParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(122)
			p.expression(0)
		}

	case 2:
		{
			p.SetState(123)
			p.AnonFunction()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(126)
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

func (p *SimpleCalcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
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
