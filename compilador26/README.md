# compilador26

Compilador da disciplina de Compiladores 2026/1 — UNESC.

## Integrantes

- Érick Lúcio de Oliveira
- Fábio Henrique Pereira
- Giovana Olivo Cittadin
- Gustavo de Almeida Kammer

## Etapas implementadas

- **E3**: Analisador léxico
- **E4**: Analisador sintático descendente preditivo tabular (LL(1) sem backtracking)

## Estrutura do projeto

```
compilador26/
├── main.go                 orquestracao do fluxo: lexer -> parser
├── token/
│   └── token.go            tipos, constantes dos terminais e nao-terminais
├── lexer/
│   └── lexer.go            algoritmo do analisador lexico
├── parser/
│   ├── grammar.go          dados estaticos: producoes, tabela M(X,a), conjuntos FOLLOW
│   └── parser.go           algoritmo do analisador sintatico (LL(1) tabular)
├── exemplos/               arquivos-fonte de exemplo
├── RELATORIO_E4.md         conjuntos First, Follow e tabela de Parsing
└── README.md
```

A separação `grammar.go` / `parser.go` no parser é proposital: um arquivo guarda os **dados** da gramática (produções, tabela LL(1), FOLLOW), o outro guarda o **algoritmo**. Quem lê o parser entende o algoritmo sem se distrair com a tabela gigantesca.

## Como rodar

### 1. Instalar o Go

Pra baixar o Go, você pode entrar aqui, é bem simples e tem instalador pra todo OS:

> https://go.dev/dl/

Se você tiver em distro Linux Ubuntu-based (Mint, Pop!_OS, elementaryOS, etc.), roda esse comando aqui que é mais rápido:

```bash
sudo apt update && sudo apt install golang-go
```

Pra confirmar que instalou certinho, abre o terminal e roda:

```bash
go version
```

Tem que aparecer algo tipo `go version go1.25.x ...`. Precisa ser **Go 1.25 ou superior**.

### 2. (Opcional, mas recomendado) VSCode com extensão Go

Pra ficar chique demais da conta, você baixa a extensão **Go** do VSCode (a oficial, do "Go Team at Google"). Depois roda `Ctrl+Shift+P` (no Windows/Linux) ou `Cmd+Shift+P` (no Mac), digita **Install Tools** e instala todas. Daí o teu VSCode vira um foguete pra programar em Golang: autocomplete, formatação automática, ir pra definição, encontrar referências, tudo funcionando.

### 3. Compilar e rodar

Abre o terminal, navega até a pasta `compilador26` (`cd caminho/para/compilador26`) e segue o fluxo do teu sistema operacional:

#### Mac

```bash
go build -o compilador26 .
./compilador26 exemplos/exemplo1.txt
```

#### Linux (Mint, Ubuntu e similares)

```bash
go build -o compilador26 .
./compilador26 exemplos/exemplo1.txt
```

#### Windows (PowerShell ou CMD)

```powershell
go build -o compilador26.exe .
.\compilador26.exe exemplos\exemplo1.txt
```

> Observação: no Windows o executável precisa terminar em `.exe`, e os caminhos usam barra invertida `\` em vez de `/`.

### 4. O que aparece na tela

A saída mostra, em ordem:

1. Lista de tokens reconhecidos (Token, Lexema, Linha)
2. Erros léxicos, se houver
3. Trace completo da análise sintática (Passo, Token, Linha, Ação, Pilha)
4. Erros sintáticos detectados, se houver

Pra testar com outros exemplos, basta trocar o nome do arquivo no comando: `exemplo2.txt`, `exemplo3.txt`, etc.

## Arquivos de exemplo

| Arquivo | Conteúdo |
|---|---|
| `exemplo1.txt` | Fatorial — válido. Cobre funções, `while`, `if/else`, comentário de linha **e** de bloco |
| `exemplo2.txt` | Média — válido. Múltiplas funções, expressões aritméticas e relacionais |
| `exemplo3.txt` | Programa com 5 erros léxicos diferentes |
| `exemplo4.txt` | Programa sem erro léxico, com 1 erro sintático (falta `;`) |
| `exemplo5.txt` | Programa com múltiplos erros sintáticos para demonstrar a recuperação em modo pânico |

---

## Diferenciais da implementação

Esta seção descreve os pontos onde fomos além do mínimo solicitado.

### 1. Tratamento de erros sintáticos com recuperação em modo pânico

A maioria dos parsers para no primeiro erro. A Aula 4 (Erros Sintáticos) é enfática: o compilador **deve** continuar a análise para detectar mais erros e ajudar o programador. Implementamos a recuperação seguindo as técnicas exatas dos slides do professor.

#### O algoritmo do parser detecta erro em duas situações

1. O símbolo corrente da entrada **não corresponde** ao terminal no topo da pilha.
2. O símbolo corrente da entrada **não possui produção** correspondente a partir do não-terminal no topo da pilha (`M(X, a) = ∅`).

#### Técnica 5 — terminal no topo da pilha não casa

Quando isso acontece, o parser:

1. Reporta o erro com a linha
2. Faz `pop` do terminal da pilha (descarta o símbolo da pilha, não da entrada)
3. Continua a análise

A pilha **encurta** após a correção, garantindo que não haja laço infinito.

#### Técnica 1 — não-terminal sem produção para o token atual

Esta é a técnica clássica do modo pânico:

1. Reporta o erro com a linha
2. Descarta tokens da entrada **até encontrar um símbolo do conjunto FOLLOW** do não-terminal
3. Quando encontra, faz `pop` do não-terminal (sincroniza)
4. Continua a análise

#### Demonstração na prática (exemplo5.txt)

Programa com erro: `x = 10` sem `;`.

```
Passo  Token         Linha   Acao
32     y (8)         8       ERRO em MulExpr': token 'y' inesperado
33     y (8)         8       RECUPERA: descarta 'y' (modo panico)
34     = (10)        8       RECUPERA: descarta '=' (modo panico)
35     20 (9)        8       RECUPERA: descarta '20' (modo panico)
36     ; (26)        8       RECUPERA: pop MulExpr' (sincroniza com FOLLOW)
```

O parser identificou o erro na expansão de `MulExpr'`, descartou os tokens `y`, `=`, `20` (que não estão em `FOLLOW(MulExpr')`) e parou em `;` (que está no FOLLOW), sincronizando a pilha. **Após a recuperação a análise continua e detecta os próximos erros.**

#### Proteções contra laço infinito

- **Limite de 50 erros**: depois disso a análise é abortada (mensagens em cascata costumam ser falsos positivos gerados pela recuperação)
- **Cada operação encurta a pilha ou avança a entrada**: garante progresso a cada iteração
- **Sincronização com `$`**: se ao descartar tokens chegar ao fim de entrada, ainda assim sincroniza

### 2. Relatório com FIRST, FOLLOW e Tabela de Parsing

O edital E4 cita explicitamente: *"As tabelas de First, Follow e Parser"*. O arquivo `RELATORIO_E4.md` traz tudo isso documentado:

- 37 conjuntos FIRST calculados
- 37 conjuntos FOLLOW calculados (com as três regras dos slides)
- Tabela M(X, a) completa: 37 não-terminais × 27 terminais
- Justificativa da resolução do **dangling else** (regra 34 vence regra 35)
- Pseudocódigo do algoritmo conforme slides 23–24 da Aula 4

### 3. Lexer baseado em expressões regulares

Em vez de implementar o autômato finito caractere por caractere, implementamos cada estado terminal como um padrão regex. O resultado é equivalente ao autômato (regex compila para autômato internamente), mas a tabela de regras se lê como uma especificação:

```go
emit(`^==`, token.EQ),
emit(`^!=`, token.NEQ),
malformedNumber(`^[0-9]+\.`),
skip(`^//[^\n]*`),
startBlockComment(`^/\*`),
```

A primeira regra que casar vence, então a ordem de prioridade fica explícita (operadores compostos antes dos simples, números mal formados antes dos válidos).

### 4. Detecção continuada de erros léxicos

O lexer **não para** ao encontrar um erro léxico — ele continua e reporta todos os erros encontrados no arquivo de uma vez. Isso evita o ciclo cansativo de "rodar, ver erro, corrigir, rodar de novo".

Detecta 5 categorias de erro léxico:

- Caractere inválido
- Identificador mal formado
- Identificador acima de 30 caracteres
- Número mal formado (formas: `12.`, `12.abc`, `123abc`, `12.5abc`)
- Comentário de bloco não fechado

### 5. Saída no estilo do professor

A saída do léxico segue exatamente o formato do exemplo de código fornecido pelo professor (`LexicoProgram.py`):

```
Token: 1 - Lexema: int - Linha: 5
Token: 8 - Lexema: fatorial - Linha: 5
Token: 21 - Lexema: ( - Linha: 5
```

A saída do sintático mostra a cada modificação da pilha as quatro informações que o edital pede: **código do token**, **token (lexema)**, **linha**, e **pilha**.

### 6. Cinco exemplos em vez de três

O edital pede no mínimo 3 arquivos de exemplo. Entregamos 5, cobrindo:

- 2 programas válidos cobrindo todos os recursos da linguagem
- 1 programa com erros léxicos (5 erros distintos)
- 1 programa com erro sintático simples
- 1 programa com múltiplos erros sintáticos para demonstrar a recuperação em modo pânico
