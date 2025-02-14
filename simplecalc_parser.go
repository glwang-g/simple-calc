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
		"", "", "'!='", "'=='", "'='", "'+'", "'-'", "'!'", "'*'", "'/'", "'>'",
		"'<'", "','", "'('", "')'", "'{'", "'}'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "LET", "FUNCTION", "IF", "ELSE", "RETURN", "TRUE", "FALSE", "IDENT",
		"INT", "NOT_EQ", "EQ", "ASSIGN", "PLUS", "MINUS", "BANG", "MULTIPLY",
		"DIVIDE", "GT", "LT", "COMMA", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"SEMICOLON", "WS", "ILLEGAL",
	}
	staticData.RuleNames = []string{
		"prog", "funcStatement", "letStatement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 27, 45, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 5, 0, 9,
		8, 0, 10, 0, 12, 0, 12, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 5, 1, 22, 8, 1, 10, 1, 12, 1, 25, 9, 1, 3, 1, 27, 8, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 32, 8, 1, 10, 1, 12, 1, 35, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 0, 0, 3, 0, 2, 4, 0, 0, 46, 0, 10, 1, 0,
		0, 0, 2, 15, 1, 0, 0, 0, 4, 38, 1, 0, 0, 0, 6, 9, 3, 2, 1, 0, 7, 9, 3,
		4, 2, 0, 8, 6, 1, 0, 0, 0, 8, 7, 1, 0, 0, 0, 9, 12, 1, 0, 0, 0, 10, 8,
		1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11, 13, 1, 0, 0, 0, 12, 10, 1, 0, 0, 0,
		13, 14, 5, 0, 0, 1, 14, 1, 1, 0, 0, 0, 15, 16, 5, 2, 0, 0, 16, 17, 5, 8,
		0, 0, 17, 26, 5, 21, 0, 0, 18, 23, 5, 8, 0, 0, 19, 20, 5, 20, 0, 0, 20,
		22, 5, 8, 0, 0, 21, 19, 1, 0, 0, 0, 22, 25, 1, 0, 0, 0, 23, 21, 1, 0, 0,
		0, 23, 24, 1, 0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 26, 18,
		1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 29, 5, 22, 0, 0,
		29, 33, 5, 23, 0, 0, 30, 32, 3, 4, 2, 0, 31, 30, 1, 0, 0, 0, 32, 35, 1,
		0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 36, 1, 0, 0, 0, 35,
		33, 1, 0, 0, 0, 36, 37, 5, 24, 0, 0, 37, 3, 1, 0, 0, 0, 38, 39, 5, 1, 0,
		0, 39, 40, 5, 8, 0, 0, 40, 41, 5, 12, 0, 0, 41, 42, 5, 9, 0, 0, 42, 43,
		5, 25, 0, 0, 43, 5, 1, 0, 0, 0, 5, 8, 10, 23, 26, 33,
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
	SimpleCalcParserEOF       = antlr.TokenEOF
	SimpleCalcParserLET       = 1
	SimpleCalcParserFUNCTION  = 2
	SimpleCalcParserIF        = 3
	SimpleCalcParserELSE      = 4
	SimpleCalcParserRETURN    = 5
	SimpleCalcParserTRUE      = 6
	SimpleCalcParserFALSE     = 7
	SimpleCalcParserIDENT     = 8
	SimpleCalcParserINT       = 9
	SimpleCalcParserNOT_EQ    = 10
	SimpleCalcParserEQ        = 11
	SimpleCalcParserASSIGN    = 12
	SimpleCalcParserPLUS      = 13
	SimpleCalcParserMINUS     = 14
	SimpleCalcParserBANG      = 15
	SimpleCalcParserMULTIPLY  = 16
	SimpleCalcParserDIVIDE    = 17
	SimpleCalcParserGT        = 18
	SimpleCalcParserLT        = 19
	SimpleCalcParserCOMMA     = 20
	SimpleCalcParserLPAREN    = 21
	SimpleCalcParserRPAREN    = 22
	SimpleCalcParserLBRACE    = 23
	SimpleCalcParserRBRACE    = 24
	SimpleCalcParserSEMICOLON = 25
	SimpleCalcParserWS        = 26
	SimpleCalcParserILLEGAL   = 27
)

// SimpleCalcParser rules.
const (
	SimpleCalcParserRULE_prog          = 0
	SimpleCalcParserRULE_funcStatement = 1
	SimpleCalcParserRULE_letStatement  = 2
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
	p.SetState(10)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SimpleCalcParserLET || _la == SimpleCalcParserFUNCTION {
		p.SetState(8)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SimpleCalcParserFUNCTION:
			{
				p.SetState(6)
				p.FuncStatement()
			}

		case SimpleCalcParserLET:
			{
				p.SetState(7)
				p.LetStatement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(12)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(13)
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
		p.SetState(15)
		p.Match(SimpleCalcParserFUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(16)
		p.Match(SimpleCalcParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(17)
		p.Match(SimpleCalcParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SimpleCalcParserIDENT {
		{
			p.SetState(18)
			p.Match(SimpleCalcParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SimpleCalcParserCOMMA {
			{
				p.SetState(19)
				p.Match(SimpleCalcParserCOMMA)
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

			p.SetState(25)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(28)
		p.Match(SimpleCalcParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Match(SimpleCalcParserLBRACE)
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

	for _la == SimpleCalcParserLET {
		{
			p.SetState(30)
			p.LetStatement()
		}

		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(36)
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
	INT() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

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

func (s *LetStatementContext) INT() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserINT, 0)
}

func (s *LetStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(SimpleCalcParserSEMICOLON, 0)
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
		p.SetState(38)
		p.Match(SimpleCalcParserLET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(SimpleCalcParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(SimpleCalcParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Match(SimpleCalcParserINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
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
