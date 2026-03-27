package main

import (
	"compilador26/lexer"
	"compilador26/parser"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: compilador26 <arquivo-fonte>")
		fmt.Println("Exemplo: compilador26 exemplos/exemplo1.txt")
		os.Exit(1)
	}

	filename := os.Args[1]
	source, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Erro ao ler arquivo '%s': %v\n", filename, err)
		os.Exit(1)
	}

	fmt.Printf("Arquivo: %s\n", filename)
	fmt.Println(fmt.Sprintf("%s", string(repeatChar('=', 60))))

	// === ANÁLISE LÉXICA ===
	fmt.Println("\n=== ANALISE LEXICA ===\n")
	l := lexer.New(string(source))
	tokens, lexErrors := l.Scan()

	fmt.Printf("%-8s %-8s %s\n", "Linha", "Codigo", "Token")
	fmt.Println(repeatChar('-', 40))
	for _, tok := range tokens {
		fmt.Printf("%-8d %-8d %s\n", tok.Line, tok.Code, tok.Lexeme)
	}
	fmt.Printf("\nTotal: %d tokens reconhecidos\n", len(tokens))

	// Exibe erros léxicos, se houver
	if len(lexErrors) > 0 {
		fmt.Println("\n=== ERROS LEXICOS ===\n")
		for _, e := range lexErrors {
			fmt.Printf("  %s\n", e.String())
		}
		fmt.Printf("\nTotal: %d erro(s) lexico(s)\n", len(lexErrors))
		fmt.Println("\nAnalise sintatica abortada devido a erros lexicos.")
		os.Exit(1)
	}

	// === ANÁLISE SINTÁTICA ===
	fmt.Println("\n=== ANALISE SINTATICA ===")
	p := parser.New(tokens)
	parseErr := p.Parse()

	if parseErr != nil {
		fmt.Printf("\nERRO SINTATICO: %v\n", parseErr)
		os.Exit(1)
	}

	fmt.Println("\nPrograma sintaticamente correto!")
}

func repeatChar(ch byte, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = ch
	}
	return string(b)
}
