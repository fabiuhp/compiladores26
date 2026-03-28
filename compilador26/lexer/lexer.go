package lexer

import (
	"compilador26/token"
	"fmt"
	"regexp"
	"strings"
)

type LexError struct {
	Message string
	Line    int
}

func (e LexError) String() string {
	return fmt.Sprintf("[Linha %d] %s", e.Line, e.Message)
}

const (
	skip         = -1
	blockComment = -2
	errNumber    = -3
	errChar      = -4
)

type lexRule struct {
	pattern *regexp.Regexp
	code    int
}

var rules = []lexRule{
	{regexp.MustCompile(`^[ \t\r\n]+`), skip},
	{regexp.MustCompile(`^//[^\n]*`), skip},
	{regexp.MustCompile(`^/\*`), blockComment},

	{regexp.MustCompile(`^[0-9]+\.[0-9]+[A-Za-z_][A-Za-z0-9_]*`), errNumber},
	{regexp.MustCompile(`^[0-9]+\.[0-9]+`), token.NUM},
	{regexp.MustCompile(`^[0-9]+\.[A-Za-z_][A-Za-z0-9_]*`), errNumber},
	{regexp.MustCompile(`^[0-9]+\.`), errNumber},
	{regexp.MustCompile(`^[0-9]+[A-Za-z_][A-Za-z0-9_]*`), errNumber},
	{regexp.MustCompile(`^[0-9]+`), token.NUM},

	{regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*`), token.ID},

	{regexp.MustCompile(`^==`), token.EQ},
	{regexp.MustCompile(`^!=`), token.NEQ},
	{regexp.MustCompile(`^<=`), token.LTE},
	{regexp.MustCompile(`^>=`), token.GTE},

	{regexp.MustCompile(`^=`), token.ASSIGN},
	{regexp.MustCompile(`^\+`), token.PLUS},
	{regexp.MustCompile(`^-`), token.MINUS},
	{regexp.MustCompile(`^\*`), token.MULT},
	{regexp.MustCompile(`^/`), token.DIV},
	{regexp.MustCompile(`^<`), token.LT},
	{regexp.MustCompile(`^>`), token.GT},

	{regexp.MustCompile(`^\(`), token.LPAREN},
	{regexp.MustCompile(`^\)`), token.RPAREN},
	{regexp.MustCompile(`^\{`), token.LBRACE},
	{regexp.MustCompile(`^\}`), token.RBRACE},
	{regexp.MustCompile(`^;`), token.SEMI},
	{regexp.MustCompile(`^,`), token.COMMA},

	{regexp.MustCompile(`^!`), errChar},
	{regexp.MustCompile(`^.`), errChar},
}

type Lexer struct {
	source string
	pos    int
	line   int
	tokens []token.Token
	errors []LexError
}

func New(source string) *Lexer {
	return &Lexer{source: source, line: 1}
}

func (l *Lexer) Scan() ([]token.Token, []LexError) {
	for l.pos < len(l.source) {
		remaining := l.source[l.pos:]

		for _, r := range rules {
			m := r.pattern.FindString(remaining)
			if m == "" {
				continue
			}

			startLine := l.line

			switch r.code {
			case skip:
				l.line += strings.Count(m, "\n")
				l.pos += len(m)

			case blockComment:
				l.pos += len(m)
				l.handleBlockComment(startLine)

			case errNumber:
				l.errors = append(l.errors, LexError{fmt.Sprintf("Numero mal formado: '%s'", m), startLine})
				l.pos += len(m)

			case errChar:
				l.errors = append(l.errors, LexError{fmt.Sprintf("Caractere invalido '%s'", m), startLine})
				l.pos += len(m)

			case token.ID:
				l.pos += len(m)
				if len(m) > 30 {
					l.errors = append(l.errors, LexError{fmt.Sprintf("Identificador acima do tamanho maximo (30 caracteres): '%s'", m), startLine})
				} else if kwCode, ok := token.Keywords[m]; ok {
					l.tokens = append(l.tokens, token.Token{Code: kwCode, Lexeme: m, Line: startLine})
				} else {
					l.tokens = append(l.tokens, token.Token{Code: token.ID, Lexeme: m, Line: startLine})
				}

			default:
				l.tokens = append(l.tokens, token.Token{Code: r.code, Lexeme: m, Line: startLine})
				l.pos += len(m)
			}

			break
		}
	}

	return l.tokens, l.errors
}

func (l *Lexer) handleBlockComment(startLine int) {
	rest := l.source[l.pos:]
	closeIdx := strings.Index(rest, "*/")

	if closeIdx == -1 {
		l.line += strings.Count(rest, "\n")
		l.pos = len(l.source)
		l.errors = append(l.errors, LexError{"Comentario de bloco nao fechado", startLine})
	} else {
		content := rest[:closeIdx+2]
		l.line += strings.Count(content, "\n")
		l.pos += closeIdx + 2
	}
}
