package main

import (
	"fmt"
	"strings"

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
	parser := recognizer.(antlr.Parser)
    stack := parser.GetRuleInvocationStack(nil)
    
    // 构建详细的错误信息
    var errBuilder strings.Builder
    errBuilder.WriteString(fmt.Sprintf("语法错误 [%d:%d]: %s\n", line, column, msg))

    errBuilder.WriteString("规则调用栈:\n")
	for i := len(stack) - 1; i >= 0; i-- {
		errBuilder.WriteString(fmt.Sprintf("  %d: %s\n", len(stack)-i, stack[i]))
	}
    
    // 如果有导致错误的符号,添加相关信息
    if offendingSymbol != nil {
        if token, ok := offendingSymbol.(antlr.Token); ok {
            errBuilder.WriteString(fmt.Sprintf("出错的符号: %s\n", token.GetText()))
        }
    }
    
    l.Errors = append(l.Errors, errBuilder.String())
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