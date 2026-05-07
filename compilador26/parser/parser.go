package parser

import (
	"compilador26/token"
	"fmt"
	"sort"
	"strings"
)

const maxErrorsBeforeAbort = 50

type SyntaxError struct {
	Line    int
	Message string
}

func (e SyntaxError) String() string {
	if e.Line > 0 {
		return fmt.Sprintf("[Linha %d] %s", e.Line, e.Message)
	}
	return e.Message
}

type Parser struct {
	tokens     []token.Token
	tokenIndex int
	stack      []int
	stepNumber int
	errors     []SyntaxError
}

func New(tokens []token.Token) *Parser {
	tokensWithEndMarker := append([]token.Token{}, tokens...)
	tokensWithEndMarker = append(tokensWithEndMarker,
		token.Token{Code: token.EOF, Lexeme: "$"})

	return &Parser{
		tokens: tokensWithEndMarker,
		stack:  []int{token.EOF, token.NTProgram},
	}
}

func (p *Parser) Parse() []SyntaxError {
	p.printTraceHeader()

	for !p.tooManyErrors() {
		topSymbol := p.topOfStack()
		currentToken := p.currentToken()

		if topSymbol == token.EOF {
			if currentToken.Code == token.EOF {
				p.logStep(currentToken, "ACEITA")
				return p.errors
			}
			p.handleExtraTokensAfterEnd(currentToken)
			continue
		}

		if token.IsTerminal(topSymbol) {
			if topSymbol == currentToken.Code {
				p.matchTerminal(currentToken)
			} else {
				p.handleTerminalMismatch(topSymbol, currentToken)
			}
			continue
		}

		if rule, found := lookupParseRule(topSymbol, currentToken.Code); found {
			p.applyRule(rule, currentToken)
			continue
		}

		p.handleMissingRule(topSymbol, currentToken)
	}

	p.errors = append(p.errors, SyntaxError{
		Message: fmt.Sprintf("limite de %d erros atingido, analise abortada", maxErrorsBeforeAbort),
	})
	return p.errors
}

func (p *Parser) tooManyErrors() bool {
	return len(p.errors) >= maxErrorsBeforeAbort
}

func (p *Parser) topOfStack() int {
	return p.stack[len(p.stack)-1]
}

func (p *Parser) currentToken() token.Token {
	return p.tokens[p.tokenIndex]
}

func (p *Parser) popStack() {
	p.stack = p.stack[:len(p.stack)-1]
}

func (p *Parser) advanceInput() {
	p.tokenIndex++
}

func (p *Parser) matchTerminal(t token.Token) {
	p.popStack()
	p.logStep(t, fmt.Sprintf("Consome: %s", t.Lexeme))
	p.advanceInput()
}

func (p *Parser) applyRule(ruleNumber int, currentToken token.Token) {
	p.popStack()
	p.pushProductionInReverse(productionRightHandSide[ruleNumber])
	p.logStep(currentToken,
		fmt.Sprintf("R%-2d: %s", ruleNumber, ruleDescription[ruleNumber]))
}

func (p *Parser) pushProductionInReverse(rightHandSide []int) {
	for i := len(rightHandSide) - 1; i >= 0; i-- {
		p.stack = append(p.stack, rightHandSide[i])
	}
}

func (p *Parser) handleTerminalMismatch(expectedTerminal int, foundToken token.Token) {
	message := fmt.Sprintf("esperado '%s', encontrou '%s'",
		token.TerminalName(expectedTerminal), foundToken.Lexeme)

	p.recordError(foundToken.Line, message)
	p.logStep(foundToken, "ERRO: "+message)

	p.popStack()
	p.logStep(foundToken, fmt.Sprintf(
		"RECUPERA: pop terminal '%s' da pilha", token.TerminalName(expectedTerminal)))
}

func (p *Parser) handleMissingRule(nonTerminal int, foundToken token.Token) {
	message := p.describeExpectedTokens(nonTerminal, foundToken)
	p.recordError(foundToken.Line, message)
	p.logStep(foundToken, fmt.Sprintf(
		"ERRO em %s: token '%s' inesperado",
		token.NonTerminalName(nonTerminal), foundToken.Lexeme))

	p.synchronizeUsingFollowSet(nonTerminal)
}

func (p *Parser) handleExtraTokensAfterEnd(extraToken token.Token) {
	p.recordError(extraToken.Line, fmt.Sprintf(
		"token inesperado apos fim do programa: '%s'", extraToken.Lexeme))
	p.logStep(extraToken, fmt.Sprintf(
		"ERRO: token '%s' apos fim do programa (descarta)", extraToken.Lexeme))
	p.advanceInput()
}

func (p *Parser) synchronizeUsingFollowSet(nonTerminal int) {
	follow := followSet[nonTerminal]

	for !p.currentTokenSynchronizes(follow) {
		discarded := p.currentToken()
		p.logStep(discarded, fmt.Sprintf(
			"RECUPERA: descarta '%s' (modo panico)", discarded.Lexeme))
		p.advanceInput()
	}

	p.popStack()
	p.logStep(p.currentToken(), fmt.Sprintf(
		"RECUPERA: pop %s (sincroniza com FOLLOW)", token.NonTerminalName(nonTerminal)))
}

func (p *Parser) currentTokenSynchronizes(follow map[int]bool) bool {
	current := p.currentToken()
	return current.Code == token.EOF || follow[current.Code]
}

func (p *Parser) recordError(line int, message string) {
	p.errors = append(p.errors, SyntaxError{Line: line, Message: message})
}

func (p *Parser) describeExpectedTokens(nonTerminal int, foundToken token.Token) string {
	var expectedNames []string
	for terminalCode := range parseTable[nonTerminal] {
		expectedNames = append(expectedNames, "'"+token.TerminalName(terminalCode)+"'")
	}
	sort.Strings(expectedNames)

	foundDescription := foundToken.Lexeme
	if foundToken.Code == token.EOF {
		foundDescription = "fim de programa"
	}

	if len(expectedNames) == 0 {
		return fmt.Sprintf("erro sintatico - token inesperado '%s'", foundDescription)
	}
	return fmt.Sprintf(
		"erro sintatico em %s - esperado %s, encontrou '%s'",
		token.NonTerminalName(nonTerminal),
		strings.Join(expectedNames, " ou "),
		foundDescription)
}

func lookupParseRule(nonTerminal, terminal int) (int, bool) {
	row, exists := parseTable[nonTerminal]
	if !exists {
		return 0, false
	}
	rule, found := row[terminal]
	return rule, found
}

func (p *Parser) printTraceHeader() {
	fmt.Printf("\n%-6s %-20s %-6s %-50s %s\n",
		"Passo", "Token", "Linha", "Acao", "Pilha")
	fmt.Println(strings.Repeat("-", 140))
}

func (p *Parser) logStep(t token.Token, action string) {
	p.stepNumber++
	fmt.Printf("%-6d %-20s %-6s %-50s %s\n",
		p.stepNumber,
		formatToken(t),
		formatLine(t),
		action,
		p.formatStack())
}

func formatToken(t token.Token) string {
	if t.Code == token.EOF {
		return "$ (27)"
	}
	return fmt.Sprintf("%s (%d)", t.Lexeme, t.Code)
}

func formatLine(t token.Token) string {
	if t.Code == token.EOF {
		return "-"
	}
	return fmt.Sprintf("%d", t.Line)
}

func (p *Parser) formatStack() string {
	if len(p.stack) == 0 {
		return "(vazia)"
	}
	symbolNames := make([]string, len(p.stack))
	for i, symbol := range p.stack {
		symbolNames[i] = token.SymbolName(symbol)
	}
	return strings.Join(symbolNames, " ")
}
