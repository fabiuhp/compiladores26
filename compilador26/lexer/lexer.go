package lexer

import (
	"compilador26/token"
	"fmt"
	"regexp"
	"strings"
)

const maxIdentifierLength = 30

type LexicalError struct {
	Line    int
	Message string
}

func (e LexicalError) String() string {
	return fmt.Sprintf("[Linha %d] %s", e.Line, e.Message)
}

type matchAction int

const (
	emitToken matchAction = iota
	skipMatch
	enterBlockComment
	reportMalformedNumber
	reportInvalidCharacter
)

type lexicalRule struct {
	pattern *regexp.Regexp
	action  matchAction
	code    int
}

func emit(pattern string, code int) lexicalRule {
	return lexicalRule{regexp.MustCompile(pattern), emitToken, code}
}

func skip(pattern string) lexicalRule {
	return lexicalRule{regexp.MustCompile(pattern), skipMatch, 0}
}

func startBlockComment(pattern string) lexicalRule {
	return lexicalRule{regexp.MustCompile(pattern), enterBlockComment, 0}
}

func malformedNumber(pattern string) lexicalRule {
	return lexicalRule{regexp.MustCompile(pattern), reportMalformedNumber, 0}
}

func invalidCharacter(pattern string) lexicalRule {
	return lexicalRule{regexp.MustCompile(pattern), reportInvalidCharacter, 0}
}

var rules = []lexicalRule{
	skip(`^[ \t\r\n]+`),
	skip(`^//[^\n]*`),
	startBlockComment(`^/\*`),

	malformedNumber(`^[0-9]+\.[0-9]+[A-Za-z_][A-Za-z0-9_]*`),
	emit(`^[0-9]+\.[0-9]+`, token.NUM),
	malformedNumber(`^[0-9]+\.[A-Za-z_][A-Za-z0-9_]*`),
	malformedNumber(`^[0-9]+\.`),
	malformedNumber(`^[0-9]+[A-Za-z_][A-Za-z0-9_]*`),
	emit(`^[0-9]+`, token.NUM),

	emit(`^[A-Za-z_][A-Za-z0-9_]*`, token.ID),

	emit(`^==`, token.EQ),
	emit(`^!=`, token.NEQ),
	emit(`^<=`, token.LTE),
	emit(`^>=`, token.GTE),

	emit(`^=`, token.ASSIGN),
	emit(`^\+`, token.PLUS),
	emit(`^-`, token.MINUS),
	emit(`^\*`, token.MULT),
	emit(`^/`, token.DIV),
	emit(`^<`, token.LT),
	emit(`^>`, token.GT),

	emit(`^\(`, token.LPAREN),
	emit(`^\)`, token.RPAREN),
	emit(`^\{`, token.LBRACE),
	emit(`^\}`, token.RBRACE),
	emit(`^;`, token.SEMI),
	emit(`^,`, token.COMMA),

	invalidCharacter(`^!`),
	invalidCharacter(`^.`),
}

type Lexer struct {
	source       string
	position     int
	currentLine  int
	tokens       []token.Token
	errors       []LexicalError
}

func New(source string) *Lexer {
	return &Lexer{source: source, currentLine: 1}
}

func (l *Lexer) Scan() ([]token.Token, []LexicalError) {
	for l.hasMoreInput() {
		l.processNextMatch()
	}
	return l.tokens, l.errors
}

func (l *Lexer) hasMoreInput() bool {
	return l.position < len(l.source)
}

func (l *Lexer) processNextMatch() {
	remaining := l.source[l.position:]

	for _, rule := range rules {
		match := rule.pattern.FindString(remaining)
		if match == "" {
			continue
		}
		l.applyAction(rule, match)
		return
	}
}

func (l *Lexer) applyAction(rule lexicalRule, match string) {
	startLine := l.currentLine

	switch rule.action {
	case skipMatch:
		l.advanceCountingNewlines(match)

	case enterBlockComment:
		l.position += len(match)
		l.consumeBlockComment(startLine)

	case reportMalformedNumber:
		l.recordError(startLine, fmt.Sprintf("Numero mal formado: '%s'", match))
		l.position += len(match)

	case reportInvalidCharacter:
		l.recordError(startLine, fmt.Sprintf("Caractere invalido '%s'", match))
		l.position += len(match)

	case emitToken:
		l.position += len(match)
		if rule.code == token.ID {
			l.emitIdentifierOrKeyword(match, startLine)
		} else {
			l.addToken(rule.code, match, startLine)
		}
	}
}

func (l *Lexer) advanceCountingNewlines(text string) {
	l.currentLine += strings.Count(text, "\n")
	l.position += len(text)
}

func (l *Lexer) emitIdentifierOrKeyword(lexeme string, line int) {
	if len(lexeme) > maxIdentifierLength {
		l.recordError(line, fmt.Sprintf(
			"Identificador acima do tamanho maximo (%d caracteres): '%s'",
			maxIdentifierLength, lexeme))
		return
	}
	if keywordCode, isKeyword := token.Keywords[lexeme]; isKeyword {
		l.addToken(keywordCode, lexeme, line)
		return
	}
	l.addToken(token.ID, lexeme, line)
}

func (l *Lexer) addToken(code int, lexeme string, line int) {
	l.tokens = append(l.tokens, token.Token{Code: code, Lexeme: lexeme, Line: line})
}

func (l *Lexer) recordError(line int, message string) {
	l.errors = append(l.errors, LexicalError{Line: line, Message: message})
}

func (l *Lexer) consumeBlockComment(startLine int) {
	rest := l.source[l.position:]
	closeIndex := strings.Index(rest, "*/")

	if closeIndex == -1 {
		l.advanceCountingNewlines(rest)
		l.recordError(startLine, "Comentario de bloco nao fechado")
		return
	}

	commentBody := rest[:closeIndex+2]
	l.advanceCountingNewlines(commentBody)
}
