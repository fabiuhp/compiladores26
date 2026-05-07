package main

import (
	"compilador26/lexer"
	"compilador26/parser"
	"compilador26/token"
	"fmt"
	"os"
	"strings"
)

func main() {
	sourceFile := readSourceFileFromArguments()
	sourceCode := readFileOrExit(sourceFile)

	printRunHeader(sourceFile)

	tokens, lexicalErrors := runLexicalAnalysis(sourceCode)
	if len(lexicalErrors) > 0 {
		exitWithMessage("Analise sintatica abortada devido a erros lexicos.")
	}

	syntaxErrors := runSyntaxAnalysis(tokens)
	if len(syntaxErrors) > 0 {
		exitWithMessage("Programa rejeitado pela analise sintatica.")
	}

	fmt.Println("\nPrograma sintaticamente correto!")
}

func readSourceFileFromArguments() string {
	if len(os.Args) < 2 {
		fmt.Println("Uso: compilador26 <arquivo-fonte>")
		os.Exit(1)
	}
	return os.Args[1]
}

func readFileOrExit(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo: %v\n", err)
		os.Exit(1)
	}
	return string(content)
}

func printRunHeader(sourceFile string) {
	fmt.Printf("Arquivo: %s\n%s\n", sourceFile, strings.Repeat("=", 60))
}

func runLexicalAnalysis(sourceCode string) ([]token.Token, []lexer.LexicalError) {
	fmt.Println("\n=== ANALISE LEXICA ===\n")

	lex := lexer.New(sourceCode)
	tokens, errors := lex.Scan()

	for _, t := range tokens {
		fmt.Println(t)
	}
	fmt.Printf("\nTotal: %d tokens reconhecidos\n", len(tokens))

	if len(errors) > 0 {
		fmt.Println("\n=== ERROS LEXICOS ===\n")
		for _, e := range errors {
			fmt.Printf("  %s\n", e)
		}
		fmt.Printf("\nTotal: %d erro(s) lexico(s)\n", len(errors))
	}

	return tokens, errors
}

func runSyntaxAnalysis(tokens []token.Token) []parser.SyntaxError {
	fmt.Println("\n=== ANALISE SINTATICA ===")

	syntaxParser := parser.New(tokens)
	errors := syntaxParser.Parse()

	if len(errors) > 0 {
		fmt.Println("\n=== ERROS SINTATICOS ===\n")
		for _, e := range errors {
			fmt.Printf("  %s\n", e)
		}
		fmt.Printf("\nTotal: %d erro(s) sintatico(s)\n", len(errors))
	}

	return errors
}

func exitWithMessage(message string) {
	fmt.Println("\n" + message)
	os.Exit(1)
}
