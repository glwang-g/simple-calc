// Code generated from SimpleCalc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SimpleCalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SimpleCalcLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func simplecalclexerLexerInit() {
	staticData := &SimpleCalcLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"LET", "FUNCTION", "IF", "ELSE", "RETURN", "TRUE", "FALSE", "IDENT",
		"INT", "NOT_EQ", "EQ", "ASSIGN", "PLUS", "MINUS", "BANG", "MULTIPLY",
		"DIVIDE", "GT", "LT", "COMMA", "LPAREN", "RPAREN", "LBRACE", "RBRACE",
		"SEMICOLON", "WS", "ILLEGAL",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 27, 143, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 5, 7, 91, 8, 7, 10, 7, 12, 7, 94, 9, 7, 1, 8, 4, 8, 97, 8, 8,
		11, 8, 12, 8, 98, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 4, 25, 136, 8, 25, 11, 25, 12,
		25, 137, 1, 25, 1, 25, 1, 26, 1, 26, 0, 0, 27, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 1, 0, 4, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13,
		32, 32, 145, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 1, 55, 1, 0, 0, 0, 3, 59, 1, 0, 0, 0, 5, 62, 1, 0, 0, 0, 7,
		65, 1, 0, 0, 0, 9, 70, 1, 0, 0, 0, 11, 77, 1, 0, 0, 0, 13, 82, 1, 0, 0,
		0, 15, 88, 1, 0, 0, 0, 17, 96, 1, 0, 0, 0, 19, 100, 1, 0, 0, 0, 21, 103,
		1, 0, 0, 0, 23, 106, 1, 0, 0, 0, 25, 108, 1, 0, 0, 0, 27, 110, 1, 0, 0,
		0, 29, 112, 1, 0, 0, 0, 31, 114, 1, 0, 0, 0, 33, 116, 1, 0, 0, 0, 35, 118,
		1, 0, 0, 0, 37, 120, 1, 0, 0, 0, 39, 122, 1, 0, 0, 0, 41, 124, 1, 0, 0,
		0, 43, 126, 1, 0, 0, 0, 45, 128, 1, 0, 0, 0, 47, 130, 1, 0, 0, 0, 49, 132,
		1, 0, 0, 0, 51, 135, 1, 0, 0, 0, 53, 141, 1, 0, 0, 0, 55, 56, 5, 108, 0,
		0, 56, 57, 5, 101, 0, 0, 57, 58, 5, 116, 0, 0, 58, 2, 1, 0, 0, 0, 59, 60,
		5, 102, 0, 0, 60, 61, 5, 110, 0, 0, 61, 4, 1, 0, 0, 0, 62, 63, 5, 105,
		0, 0, 63, 64, 5, 102, 0, 0, 64, 6, 1, 0, 0, 0, 65, 66, 5, 101, 0, 0, 66,
		67, 5, 108, 0, 0, 67, 68, 5, 115, 0, 0, 68, 69, 5, 101, 0, 0, 69, 8, 1,
		0, 0, 0, 70, 71, 5, 114, 0, 0, 71, 72, 5, 101, 0, 0, 72, 73, 5, 116, 0,
		0, 73, 74, 5, 117, 0, 0, 74, 75, 5, 114, 0, 0, 75, 76, 5, 110, 0, 0, 76,
		10, 1, 0, 0, 0, 77, 78, 5, 116, 0, 0, 78, 79, 5, 114, 0, 0, 79, 80, 5,
		117, 0, 0, 80, 81, 5, 101, 0, 0, 81, 12, 1, 0, 0, 0, 82, 83, 5, 102, 0,
		0, 83, 84, 5, 97, 0, 0, 84, 85, 5, 108, 0, 0, 85, 86, 5, 115, 0, 0, 86,
		87, 5, 101, 0, 0, 87, 14, 1, 0, 0, 0, 88, 92, 7, 0, 0, 0, 89, 91, 7, 1,
		0, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93,
		1, 0, 0, 0, 93, 16, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 97, 7, 2, 0, 0,
		96, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1,
		0, 0, 0, 99, 18, 1, 0, 0, 0, 100, 101, 5, 33, 0, 0, 101, 102, 5, 61, 0,
		0, 102, 20, 1, 0, 0, 0, 103, 104, 5, 61, 0, 0, 104, 105, 5, 61, 0, 0, 105,
		22, 1, 0, 0, 0, 106, 107, 5, 61, 0, 0, 107, 24, 1, 0, 0, 0, 108, 109, 5,
		43, 0, 0, 109, 26, 1, 0, 0, 0, 110, 111, 5, 45, 0, 0, 111, 28, 1, 0, 0,
		0, 112, 113, 5, 33, 0, 0, 113, 30, 1, 0, 0, 0, 114, 115, 5, 42, 0, 0, 115,
		32, 1, 0, 0, 0, 116, 117, 5, 47, 0, 0, 117, 34, 1, 0, 0, 0, 118, 119, 5,
		62, 0, 0, 119, 36, 1, 0, 0, 0, 120, 121, 5, 60, 0, 0, 121, 38, 1, 0, 0,
		0, 122, 123, 5, 44, 0, 0, 123, 40, 1, 0, 0, 0, 124, 125, 5, 40, 0, 0, 125,
		42, 1, 0, 0, 0, 126, 127, 5, 41, 0, 0, 127, 44, 1, 0, 0, 0, 128, 129, 5,
		123, 0, 0, 129, 46, 1, 0, 0, 0, 130, 131, 5, 125, 0, 0, 131, 48, 1, 0,
		0, 0, 132, 133, 5, 59, 0, 0, 133, 50, 1, 0, 0, 0, 134, 136, 7, 3, 0, 0,
		135, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137,
		138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 140, 6, 25, 0, 0, 140, 52,
		1, 0, 0, 0, 141, 142, 9, 0, 0, 0, 142, 54, 1, 0, 0, 0, 4, 0, 92, 98, 137,
		1, 6, 0, 0,
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

// SimpleCalcLexerInit initializes any static state used to implement SimpleCalcLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSimpleCalcLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SimpleCalcLexerInit() {
	staticData := &SimpleCalcLexerLexerStaticData
	staticData.once.Do(simplecalclexerLexerInit)
}

// NewSimpleCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSimpleCalcLexer(input antlr.CharStream) *SimpleCalcLexer {
	SimpleCalcLexerInit()
	l := new(SimpleCalcLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SimpleCalcLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "SimpleCalc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimpleCalcLexer tokens.
const (
	SimpleCalcLexerLET       = 1
	SimpleCalcLexerFUNCTION  = 2
	SimpleCalcLexerIF        = 3
	SimpleCalcLexerELSE      = 4
	SimpleCalcLexerRETURN    = 5
	SimpleCalcLexerTRUE      = 6
	SimpleCalcLexerFALSE     = 7
	SimpleCalcLexerIDENT     = 8
	SimpleCalcLexerINT       = 9
	SimpleCalcLexerNOT_EQ    = 10
	SimpleCalcLexerEQ        = 11
	SimpleCalcLexerASSIGN    = 12
	SimpleCalcLexerPLUS      = 13
	SimpleCalcLexerMINUS     = 14
	SimpleCalcLexerBANG      = 15
	SimpleCalcLexerMULTIPLY  = 16
	SimpleCalcLexerDIVIDE    = 17
	SimpleCalcLexerGT        = 18
	SimpleCalcLexerLT        = 19
	SimpleCalcLexerCOMMA     = 20
	SimpleCalcLexerLPAREN    = 21
	SimpleCalcLexerRPAREN    = 22
	SimpleCalcLexerLBRACE    = 23
	SimpleCalcLexerRBRACE    = 24
	SimpleCalcLexerSEMICOLON = 25
	SimpleCalcLexerWS        = 26
	SimpleCalcLexerILLEGAL   = 27
)
