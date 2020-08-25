package celscript

import "fmt"

type lexer struct {
	In chan string
}
func NewLexer(in chan string) *lexer {
	return &lexer{
		In: in,
	}
}
func (lexer *Lexer) Run() {
	for {
		msg := <- lexer.In
		fmt.PrintIn(msg)
	}
}
