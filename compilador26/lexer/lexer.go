package lexer

import (
	"compilador26/token"
	"fmt"
	"unicode"
)

// LexError representa um erro léxico encontrado
type LexError struct {
	Message string
	Line    int
}

func (e LexError) String() string {
	return fmt.Sprintf("[Linha %d] %s", e.Line, e.Message)
}

// Lexer implementa o analisador léxico baseado no autômato finito
type Lexer struct {
	source []rune
	pos    int
	line   int
	tokens []token.Token
	errors []LexError
}

func New(source string) *Lexer {
	return &Lexer{
		source: []rune(source),
		pos:    0,
		line:   1,
	}
}

// Scan executa a análise léxica completa e retorna tokens e erros
func (l *Lexer) Scan() ([]token.Token, []LexError) {
	for !l.isAtEnd() {
		l.skipWhitespace()
		if l.isAtEnd() {
			break
		}
		l.scanToken()
	}
	return l.tokens, l.errors
}

func (l *Lexer) scanToken() {
	ch := l.peek()

	switch {
	case isLetter(ch) || ch == '_':
		l.scanIdentifier()
	case isDigit(ch):
		l.scanNumber()
	case ch == '=':
		l.advance()
		if !l.isAtEnd() && l.peek() == '=' {
			l.advance()
			l.addToken(token.EQ, "==")
		} else {
			l.addToken(token.ASSIGN, "=")
		}
	case ch == '!':
		l.advance()
		if !l.isAtEnd() && l.peek() == '=' {
			l.advance()
			l.addToken(token.NEQ, "!=")
		} else {
			l.addError("Caractere invalido '!'")
		}
	case ch == '<':
		l.advance()
		if !l.isAtEnd() && l.peek() == '=' {
			l.advance()
			l.addToken(token.LTE, "<=")
		} else {
			l.addToken(token.LT, "<")
		}
	case ch == '>':
		l.advance()
		if !l.isAtEnd() && l.peek() == '=' {
			l.advance()
			l.addToken(token.GTE, ">=")
		} else {
			l.addToken(token.GT, ">")
		}
	case ch == '/':
		l.scanSlash()
	case ch == '+':
		l.advance()
		l.addToken(token.PLUS, "+")
	case ch == '-':
		l.advance()
		l.addToken(token.MINUS, "-")
	case ch == '*':
		l.advance()
		l.addToken(token.MULT, "*")
	case ch == '(':
		l.advance()
		l.addToken(token.LPAREN, "(")
	case ch == ')':
		l.advance()
		l.addToken(token.RPAREN, ")")
	case ch == '{':
		l.advance()
		l.addToken(token.LBRACE, "{")
	case ch == '}':
		l.advance()
		l.addToken(token.RBRACE, "}")
	case ch == ';':
		l.advance()
		l.addToken(token.SEMI, ";")
	case ch == ',':
		l.advance()
		l.addToken(token.COMMA, ",")
	default:
		l.addError(fmt.Sprintf("Caractere invalido '%c'", ch))
		l.advance()
	}
}

// scanIdentifier reconhece identificadores e palavras reservadas
func (l *Lexer) scanIdentifier() {
	start := l.pos
	startLine := l.line

	for !l.isAtEnd() && (isLetter(l.peek()) || isDigit(l.peek()) || l.peek() == '_') {
		l.advance()
	}

	lexeme := string(l.source[start:l.pos])

	if len(lexeme) > 30 {
		l.errors = append(l.errors, LexError{
			Message: fmt.Sprintf("Identificador acima do tamanho maximo (30 caracteres): '%s'", lexeme),
			Line:    startLine,
		})
		return
	}

	if code, ok := token.Keywords[lexeme]; ok {
		l.tokens = append(l.tokens, token.Token{Code: code, Lexeme: lexeme, Line: startLine})
	} else {
		l.tokens = append(l.tokens, token.Token{Code: token.ID, Lexeme: lexeme, Line: startLine})
	}
}

// scanNumber reconhece números inteiros e reais
func (l *Lexer) scanNumber() {
	start := l.pos
	startLine := l.line

	// Lê dígitos da parte inteira
	for !l.isAtEnd() && isDigit(l.peek()) {
		l.advance()
	}

	// Verifica se é número real (ponto seguido de dígito)
	if !l.isAtEnd() && l.peek() == '.' {
		l.advance() // consome o ponto
		if !l.isAtEnd() && isDigit(l.peek()) {
			for !l.isAtEnd() && isDigit(l.peek()) {
				l.advance()
			}
		} else {
			// Número mal formado: "12." sem dígito após o ponto
			lexeme := string(l.source[start:l.pos])
			l.errors = append(l.errors, LexError{
				Message: fmt.Sprintf("Numero mal formado: '%s'", lexeme),
				Line:    startLine,
			})
			return
		}
	}

	// Verifica se há letra colada ao número: "123abc"
	if !l.isAtEnd() && (isLetter(l.peek()) || l.peek() == '_') {
		for !l.isAtEnd() && (isLetter(l.peek()) || isDigit(l.peek()) || l.peek() == '_') {
			l.advance()
		}
		lexeme := string(l.source[start:l.pos])
		l.errors = append(l.errors, LexError{
			Message: fmt.Sprintf("Numero mal formado: '%s'", lexeme),
			Line:    startLine,
		})
		return
	}

	lexeme := string(l.source[start:l.pos])
	l.tokens = append(l.tokens, token.Token{Code: token.NUM, Lexeme: lexeme, Line: startLine})
}

// scanSlash trata / (divisão), // (comentário de linha) e /* (comentário de bloco)
func (l *Lexer) scanSlash() {
	l.advance() // consome '/'

	if l.isAtEnd() {
		l.addToken(token.DIV, "/")
		return
	}

	switch l.peek() {
	case '/':
		// Comentário de linha: consome até o fim da linha
		l.advance()
		for !l.isAtEnd() && l.peek() != '\n' {
			l.advance()
		}
	case '*':
		// Comentário de bloco: consome até */
		startLine := l.line
		l.advance() // consome '*'
		closed := false
		for !l.isAtEnd() {
			if l.peek() == '*' {
				l.advance()
				if !l.isAtEnd() && l.peek() == '/' {
					l.advance()
					closed = true
					break
				}
			} else {
				if l.peek() == '\n' {
					l.line++
				}
				l.advance()
			}
		}
		if !closed {
			l.errors = append(l.errors, LexError{
				Message: "Comentario de bloco nao fechado",
				Line:    startLine,
			})
		}
	default:
		l.addToken(token.DIV, "/")
	}
}

// skipWhitespace pula espaços, tabs e quebras de linha
func (l *Lexer) skipWhitespace() {
	for !l.isAtEnd() {
		ch := l.peek()
		if ch == '\n' {
			l.line++
			l.pos++
		} else if ch == ' ' || ch == '\t' || ch == '\r' {
			l.pos++
		} else {
			break
		}
	}
}

func (l *Lexer) peek() rune {
	if l.pos < len(l.source) {
		return l.source[l.pos]
	}
	return 0
}

func (l *Lexer) advance() {
	l.pos++
}

func (l *Lexer) isAtEnd() bool {
	return l.pos >= len(l.source)
}

func (l *Lexer) addToken(code int, lexeme string) {
	l.tokens = append(l.tokens, token.Token{Code: code, Lexeme: lexeme, Line: l.line})
}

func (l *Lexer) addError(msg string) {
	l.errors = append(l.errors, LexError{Message: msg, Line: l.line})
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
