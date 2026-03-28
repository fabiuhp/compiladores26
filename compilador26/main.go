package main

import (
	"compilador26/lexer"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: compilador26 <arquivo-fonte>")
		os.Exit(1)
	}

	source, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Erro ao ler arquivo: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Arquivo: %s\n%s\n", os.Args[1], strings.Repeat("=", 60))

	fmt.Println("\n=== ANALISE LEXICA ===\n")
	l := lexer.New(string(source))
	tokens, lexErrors := l.Scan()

	fmt.Printf("%-8s %-8s %s\n%s\n", "Linha", "Codigo", "Token", strings.Repeat("-", 40))
	for _, tok := range tokens {
		fmt.Printf("%-8d %-8d %s\n", tok.Line, tok.Code, tok.Lexeme)
	}
	fmt.Printf("\nTotal: %d tokens reconhecidos\n", len(tokens))

	if len(lexErrors) > 0 {
		fmt.Println("\n=== ERROS LEXICOS ===\n")
		for _, e := range lexErrors {
			fmt.Printf("  %s\n", e)
		}
		fmt.Printf("\nTotal: %d erro(s) lexico(s)\n", len(lexErrors))
	}
}
