package celscript

type lexer struct {
	In chan string
}
func NewLexer(in chan string) *lexer {
	return &lexer{
		In: in,
	}
}
