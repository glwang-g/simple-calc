package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// SimpleCalcErrorListener 是自定义的错误监听器
type SimpleCalcErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

// NewSimpleCalcErrorListener 创建一个新的错误监听器实例
func NewSimpleCalcErrorListener() *SimpleCalcErrorListener {
	return &SimpleCalcErrorListener{
		Errors: make([]string, 0),
	}
}

// SyntaxError 处理语法错误
func (l *SimpleCalcErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{},
	line, column int, msg string, e antlr.RecognitionException) {
	err := fmt.Sprintf("语法错误 [%d:%d]: %s", line, column, msg)
	l.Errors = append(l.Errors, err)
}

// HasErrors 返回是否有错误
func (l *SimpleCalcErrorListener) HasErrors() bool {
	return len(l.Errors) > 0
}

// PrintErrors 打印所有错误信息
func (l *SimpleCalcErrorListener) PrintErrors() {
	if len(l.Errors) == 0 {
		fmt.Println("没有错误")
		return
	}
	for i, err := range l.Errors {
		fmt.Printf("错误 %d: %s\n", i+1, err)
	}
}