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
		"LET", "FUNCTION", "IF", "ELSE", "RETURN", "TRUE", "FALSE", "IDENT",
		"INT", "FLOAT", "NOT_EQ", "EQ", "AND", "OR", "ASSIGN", "PLUS", "MINUS",
		"BANG", "MULTIPLY", "DIVIDE", "GT", "LT", "GTE", "LTE", "COMMA", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "SEMICOLON", "WHITESPACE", "ILLEGAL",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 32, 176, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 5, 7, 101, 8, 7, 10, 7, 12, 7, 104, 9, 7, 1, 8, 4, 8, 107, 8, 8,
		11, 8, 12, 8, 108, 1, 9, 4, 9, 112, 8, 9, 11, 9, 12, 9, 113, 1, 9, 1, 9,
		4, 9, 118, 8, 9, 11, 9, 12, 9, 119, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30,
		4, 30, 169, 8, 30, 11, 30, 12, 30, 170, 1, 30, 1, 30, 1, 31, 1, 31, 0,
		0, 32, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 1, 0, 4, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13,
		32, 32, 180, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0,
		61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 1, 65, 1, 0, 0, 0, 3, 69, 1, 0, 0, 0,
		5, 72, 1, 0, 0, 0, 7, 75, 1, 0, 0, 0, 9, 80, 1, 0, 0, 0, 11, 87, 1, 0,
		0, 0, 13, 92, 1, 0, 0, 0, 15, 98, 1, 0, 0, 0, 17, 106, 1, 0, 0, 0, 19,
		111, 1, 0, 0, 0, 21, 121, 1, 0, 0, 0, 23, 124, 1, 0, 0, 0, 25, 127, 1,
		0, 0, 0, 27, 130, 1, 0, 0, 0, 29, 133, 1, 0, 0, 0, 31, 135, 1, 0, 0, 0,
		33, 137, 1, 0, 0, 0, 35, 139, 1, 0, 0, 0, 37, 141, 1, 0, 0, 0, 39, 143,
		1, 0, 0, 0, 41, 145, 1, 0, 0, 0, 43, 147, 1, 0, 0, 0, 45, 149, 1, 0, 0,
		0, 47, 152, 1, 0, 0, 0, 49, 155, 1, 0, 0, 0, 51, 157, 1, 0, 0, 0, 53, 159,
		1, 0, 0, 0, 55, 161, 1, 0, 0, 0, 57, 163, 1, 0, 0, 0, 59, 165, 1, 0, 0,
		0, 61, 168, 1, 0, 0, 0, 63, 174, 1, 0, 0, 0, 65, 66, 5, 108, 0, 0, 66,
		67, 5, 101, 0, 0, 67, 68, 5, 116, 0, 0, 68, 2, 1, 0, 0, 0, 69, 70, 5, 102,
		0, 0, 70, 71, 5, 110, 0, 0, 71, 4, 1, 0, 0, 0, 72, 73, 5, 105, 0, 0, 73,
		74, 5, 102, 0, 0, 74, 6, 1, 0, 0, 0, 75, 76, 5, 101, 0, 0, 76, 77, 5, 108,
		0, 0, 77, 78, 5, 115, 0, 0, 78, 79, 5, 101, 0, 0, 79, 8, 1, 0, 0, 0, 80,
		81, 5, 114, 0, 0, 81, 82, 5, 101, 0, 0, 82, 83, 5, 116, 0, 0, 83, 84, 5,
		117, 0, 0, 84, 85, 5, 114, 0, 0, 85, 86, 5, 110, 0, 0, 86, 10, 1, 0, 0,
		0, 87, 88, 5, 116, 0, 0, 88, 89, 5, 114, 0, 0, 89, 90, 5, 117, 0, 0, 90,
		91, 5, 101, 0, 0, 91, 12, 1, 0, 0, 0, 92, 93, 5, 102, 0, 0, 93, 94, 5,
		97, 0, 0, 94, 95, 5, 108, 0, 0, 95, 96, 5, 115, 0, 0, 96, 97, 5, 101, 0,
		0, 97, 14, 1, 0, 0, 0, 98, 102, 7, 0, 0, 0, 99, 101, 7, 1, 0, 0, 100, 99,
		1, 0, 0, 0, 101, 104, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0,
		0, 0, 103, 16, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 107, 7, 2, 0, 0,
		106, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108,
		109, 1, 0, 0, 0, 109, 18, 1, 0, 0, 0, 110, 112, 7, 2, 0, 0, 111, 110, 1,
		0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0,
		0, 114, 115, 1, 0, 0, 0, 115, 117, 5, 46, 0, 0, 116, 118, 7, 2, 0, 0, 117,
		116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120,
		1, 0, 0, 0, 120, 20, 1, 0, 0, 0, 121, 122, 5, 33, 0, 0, 122, 123, 5, 61,
		0, 0, 123, 22, 1, 0, 0, 0, 124, 125, 5, 61, 0, 0, 125, 126, 5, 61, 0, 0,
		126, 24, 1, 0, 0, 0, 127, 128, 5, 38, 0, 0, 128, 129, 5, 38, 0, 0, 129,
		26, 1, 0, 0, 0, 130, 131, 5, 124, 0, 0, 131, 132, 5, 124, 0, 0, 132, 28,
		1, 0, 0, 0, 133, 134, 5, 61, 0, 0, 134, 30, 1, 0, 0, 0, 135, 136, 5, 43,
		0, 0, 136, 32, 1, 0, 0, 0, 137, 138, 5, 45, 0, 0, 138, 34, 1, 0, 0, 0,
		139, 140, 5, 33, 0, 0, 140, 36, 1, 0, 0, 0, 141, 142, 5, 42, 0, 0, 142,
		38, 1, 0, 0, 0, 143, 144, 5, 47, 0, 0, 144, 40, 1, 0, 0, 0, 145, 146, 5,
		62, 0, 0, 146, 42, 1, 0, 0, 0, 147, 148, 5, 60, 0, 0, 148, 44, 1, 0, 0,
		0, 149, 150, 5, 62, 0, 0, 150, 151, 5, 61, 0, 0, 151, 46, 1, 0, 0, 0, 152,
		153, 5, 60, 0, 0, 153, 154, 5, 61, 0, 0, 154, 48, 1, 0, 0, 0, 155, 156,
		5, 44, 0, 0, 156, 50, 1, 0, 0, 0, 157, 158, 5, 40, 0, 0, 158, 52, 1, 0,
		0, 0, 159, 160, 5, 41, 0, 0, 160, 54, 1, 0, 0, 0, 161, 162, 5, 123, 0,
		0, 162, 56, 1, 0, 0, 0, 163, 164, 5, 125, 0, 0, 164, 58, 1, 0, 0, 0, 165,
		166, 5, 59, 0, 0, 166, 60, 1, 0, 0, 0, 167, 169, 7, 3, 0, 0, 168, 167,
		1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0,
		0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 6, 30, 0, 0, 173, 62, 1, 0, 0, 0,
		174, 175, 9, 0, 0, 0, 175, 64, 1, 0, 0, 0, 6, 0, 102, 108, 113, 119, 170,
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
	SimpleCalcLexerLET        = 1
	SimpleCalcLexerFUNCTION   = 2
	SimpleCalcLexerIF         = 3
	SimpleCalcLexerELSE       = 4
	SimpleCalcLexerRETURN     = 5
	SimpleCalcLexerTRUE       = 6
	SimpleCalcLexerFALSE      = 7
	SimpleCalcLexerIDENT      = 8
	SimpleCalcLexerINT        = 9
	SimpleCalcLexerFLOAT      = 10
	SimpleCalcLexerNOT_EQ     = 11
	SimpleCalcLexerEQ         = 12
	SimpleCalcLexerAND        = 13
	SimpleCalcLexerOR         = 14
	SimpleCalcLexerASSIGN     = 15
	SimpleCalcLexerPLUS       = 16
	SimpleCalcLexerMINUS      = 17
	SimpleCalcLexerBANG       = 18
	SimpleCalcLexerMULTIPLY   = 19
	SimpleCalcLexerDIVIDE     = 20
	SimpleCalcLexerGT         = 21
	SimpleCalcLexerLT         = 22
	SimpleCalcLexerGTE        = 23
	SimpleCalcLexerLTE        = 24
	SimpleCalcLexerCOMMA      = 25
	SimpleCalcLexerLPAREN     = 26
	SimpleCalcLexerRPAREN     = 27
	SimpleCalcLexerLBRACE     = 28
	SimpleCalcLexerRBRACE     = 29
	SimpleCalcLexerSEMICOLON  = 30
	SimpleCalcLexerWHITESPACE = 31
	SimpleCalcLexerILLEGAL    = 32
)
