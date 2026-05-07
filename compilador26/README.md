# Compilador da disciplina de Compiladores 2026/1 — UNESC.

## Integrantes

- Érick Lúcio de Oliveira
- Fábio Henrique Pereira
- Giovana Olivo Cittadin
- Gustavo de Almeida Kammer

## Etapas implementadas

- **E3**: Analisador léxico
- **E4**: Analisador sintático descendente preditivo tabular (LL(1) sem backtracking)

## Estrutura

```
compilador26/
├── main.go              ponto de entrada
├── token/token.go       códigos dos terminais e não-terminais, struct Token
├── lexer/lexer.go       analisador léxico (regras como expressões regulares)
├── parser/parser.go     analisador sintático tabular com modo pânico
├── exemplos/            arquivos-fonte de exemplo
├── RELATORIO_E4.md      conjuntos First, Follow e tabela de Parsing
└── README.md
```

## Como compilar

```bash
go build -o compilador26 .
```

Requer Go 1.25 ou superior.

## Como executar

```bash
./compilador26 exemplos/exemplo1.txt
```

A saída mostra:

1. Lista de tokens reconhecidos (Token, Lexema, Linha)
2. Erros léxicos, se houver
3. Trace completo da análise sintática (Passo, Token, Linha, Ação, Pilha)
4. Erros sintáticos detectados, se houver

## Exemplos

- `exemplo1.txt`: fatorial — válido (funções, `while`, `if/else`, comentário de linha e de bloco)
- `exemplo2.txt`: média — válido (múltiplas funções, expressões aritméticas e relacionais)
- `exemplo3.txt`: programa com 5 erros léxicos
- `exemplo4.txt`: programa sem erro léxico, com 1 erro sintático
- `exemplo5.txt`: programa com múltiplos erros sintáticos (demonstra recuperação em modo pânico)
