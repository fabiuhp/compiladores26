# TRABALHO SEMESTRAL – ETAPA E1

## Manual da Linguagem

- Érick Lúcio de Oliveira
- Fábio Henrique Pereira
- Giovana Olivo Cittadin
- Gustavo de Almeida Kammer

---

## 1. Introdução

Neste documento é apresentada a primeira etapa do trabalho semestral da disciplina de Compiladores. O objetivo desta etapa é definir o manual de uma linguagem de programação simplificada, descrevendo seus comandos, a sintaxe de cada um, exemplos de uso, as regras léxicas e a definição dos erros léxicos que serão tratados ao longo do projeto.

A linguagem foi pensada para ser simples de entender e viável de implementar nas próximas etapas, principalmente no analisador léxico e no analisador sintático.

---

## 2. Visão geral da linguagem

A linguagem definida para este trabalho é uma linguagem imperativa baseada em funções. Ela possui suporte para:

- definição de funções com parâmetros e retorno;
- declaração de variáveis locais;
- comando de atribuição;
- comando de saída (`print`);
- comando de retorno (`return`);
- estrutura de decisão (`if` / `else`);
- estrutura de repetição (`while`);
- expressões aritméticas e relacionais;
- chamada de funções com passagem de argumentos;
- comentários de linha e de bloco;
- agrupamento de comandos em blocos.

---

## 3. Estrutura geral do programa

Um programa nesta linguagem é composto por uma ou mais definições de funções. Não existe uma palavra reservada `program`. O programa é diretamente uma lista de funções.

### 3.1 Sintaxe geral

```
<tipo> nomeFuncao(<parâmetros>) {
    <declarações de variáveis>
    <comandos>
}
```

### 3.2 Explicação

Cada função possui um tipo de retorno (`int` ou `float`), um nome (identificador), uma lista de parâmetros (que pode ser vazia) entre parênteses, e um bloco de código delimitado por chaves.

Um detalhe importante: dentro de um bloco, todas as declarações de variáveis devem vir **antes** dos comandos. Não é permitido declarar variáveis no meio dos comandos.

### 3.3 Exemplo

```
int main() {
    int x;
    x = 10;
    print(x);
    return 0;
}
```

---

## 4. Tipos de dados

A linguagem possui dois tipos de dados:

- `int`: números inteiros (ex.: `0`, `5`, `42`)
- `float`: números reais com ponto decimal (ex.: `3.14`, `0.5`, `10.0`)

Os tipos são utilizados na declaração de variáveis, nos parâmetros de funções e no retorno das funções.

---

## 5. Declaração de variáveis

### Sintaxe

```
<tipo> <identificador> ;
```

### Explicação

As variáveis são declaradas informando o tipo seguido do nome e um ponto e vírgula. Não é permitido atribuir valor no momento da declaração. A atribuição deve ser feita em um comando separado.

Todas as declarações devem aparecer no início do bloco, antes de qualquer comando.

### Exemplos

```
int idade;
float salario;
int x;
float media;
```

---

## 6. Funções

### 6.1 Definição de função

#### Sintaxe

```
<tipo> <identificador> ( <parâmetros> ) {
    <declarações>
    <comandos>
}
```

#### Explicação

Toda função precisa ter um tipo de retorno (`int` ou `float`), um nome, uma lista de parâmetros (pode ser vazia) e um bloco com o corpo da função. Como toda função tem tipo de retorno, espera-se que ela contenha ao menos um comando `return`.

#### Exemplo

```
int soma(int a, int b) {
    int resultado;
    resultado = a + b;
    return resultado;
}
```

### 6.2 Parâmetros

Os parâmetros são declarados entre parênteses, separados por vírgula. Cada parâmetro possui um tipo e um nome.

```
float media(float x, float y) {
    float m;
    m = (x + y) / 2;
    return m;
}
```

Quando a função não recebe parâmetros, os parênteses ficam vazios:

```
int principal() {
    return 0;
}
```

### 6.3 Chamada de função

#### Sintaxe

```
<identificador> ( <argumentos> )
```

#### Explicação

Uma função é chamada pelo seu nome, passando os argumentos entre parênteses. A chamada pode ser usada dentro de expressões, atribuições ou como argumento de outros comandos.

#### Exemplos

```
resultado = soma(3, 5);
print(fatorial(n));
x = calcula(a, b + 1);
```

---

## 7. Comandos da linguagem

### 7.1 Comando de atribuição

#### Sintaxe

```
<identificador> = <expressão> ;
```

#### Explicação

O comando de atribuição armazena o resultado de uma expressão em uma variável já declarada.

#### Exemplos

```
x = 5;
y = x + 10;
media = (nota1 + nota2) / 2;
```

### 7.2 Comando de saída (`print`)

#### Sintaxe

```
print ( <expressão> ) ;
```

#### Explicação

O comando `print` exibe na saída o valor resultante de uma expressão.

#### Exemplos

```
print(x);
print(x + y);
print(media(a, b));
```

### 7.3 Comando de retorno (`return`)

#### Sintaxe

```
return <expressão> ;
```

#### Explicação

O comando `return` encerra a execução da função e devolve o valor da expressão para quem a chamou.

#### Exemplos

```
return 0;
return resultado;
return a + b;
```

### 7.4 Comando de decisão (`if` / `else`)

#### Sintaxe

```
if ( <expressão> ) <comando>
```

ou com `else`:

```
if ( <expressão> ) <comando> else <comando>
```

#### Explicação

O `if` avalia uma expressão. Se o resultado for verdadeiro (diferente de zero), executa o comando logo após. Se houver `else` e a condição for falsa, executa o comando do `else`.

**Importante:** nesta linguagem **não existe** a palavra `then`. O comando vem direto após o fechamento do parêntese.

Quando se deseja executar mais de um comando, basta usar um bloco `{ }`.

A parte `else` é opcional, ou seja, pode haver um `if` sem `else`.

#### Exemplos

```
if (x > 0)
    print(x);

if (media >= 6) {
    print(media);
    return 1;
} else {
    print(0);
    return 0;
}
```

### 7.5 Comando de repetição (`while`)

#### Sintaxe

```
while ( <expressão> ) <comando>
```

#### Explicação

O `while` executa um comando (ou bloco) repetidamente enquanto a expressão for verdadeira (diferente de zero).

**Importante:** nesta linguagem **não existe** a palavra `do`. O comando vem direto após o fechamento do parêntese.

#### Exemplos

```
while (x < 10) {
    print(x);
    x = x + 1;
}

while (n > 1)
    n = n - 1;
```

### 7.6 Bloco de comandos

#### Sintaxe

```
{
    <declarações>
    <comandos>
}
```

#### Explicação

O bloco agrupa declarações e comandos entre chaves. Pode ser usado como corpo de uma função, dentro de um `if`, de um `while`, ou como um comando isolado. As declarações sempre devem vir antes dos comandos no bloco.

#### Exemplo

```
{
    int temp;
    temp = x;
    x = y;
    y = temp;
}
```

---

## 8. Expressões

### 8.1 Operadores aritméticos

| Operador | Operação       |
|----------|----------------|
| `+`      | soma           |
| `-`      | subtração      |
| `*`      | multiplicação  |
| `/`      | divisão        |

Parênteses podem ser usados para alterar a ordem de avaliação.

```
x = 2 + 3;
y = a * b;
z = (x + y) / 2;
```

### 8.2 Operadores relacionais

| Operador | Significado    |
|----------|----------------|
| `==`     | igual          |
| `!=`     | diferente      |
| `<`      | menor que      |
| `>`      | maior que      |
| `<=`     | menor ou igual |
| `>=`     | maior ou igual |

São usados principalmente nas condições do `if` e do `while`.

```
if (x == 0)
    print(0);

while (i < 10) {
    print(i);
    i = i + 1;
}
```

### 8.3 Precedência de operadores

Da mais alta para a mais baixa:

1. `*` , `/` (multiplicação e divisão)
2. `+` , `-` (soma e subtração)
3. `==` , `!=` , `<` , `>` , `<=` , `>=` (relacionais)

Operadores de mesma precedência são avaliados da esquerda para a direita. Parênteses alteram a ordem quando necessário.

---

## 9. Comentários

### 9.1 Comentário de linha

Inicia com `//` e vai até o final da linha. Tudo após `//` é ignorado.

```
// isto é um comentário
x = 10;  // comentário no fim da linha
```

### 9.2 Comentário de bloco

Inicia com `/*` e termina com `*/`. Pode ocupar várias linhas.

```
/* este é um
   comentário de bloco */
x = 20;
```

Comentários são completamente ignorados pelo analisador léxico. Um comentário de bloco não fechado é tratado como erro léxico.

---

## 10. Palavras reservadas

As palavras abaixo têm significado especial e **não podem** ser usadas como nomes de variáveis ou funções:

```
int    float    if    else    while    return    print
```

---

## 11. Delimitadores e operadores

### Delimitadores

| Símbolo | Descrição           |
|---------|---------------------|
| `(`     | parêntese esquerdo  |
| `)`     | parêntese direito   |
| `{`     | chave esquerda      |
| `}`     | chave direita       |
| `;`     | ponto e vírgula     |
| `,`     | vírgula             |

### Operadores

| Símbolo | Descrição      |
|---------|----------------|
| `=`     | atribuição     |
| `+`     | soma           |
| `-`     | subtração      |
| `*`     | multiplicação  |
| `/`     | divisão        |
| `==`    | igual          |
| `!=`    | diferente      |
| `<`     | menor que      |
| `>`     | maior que      |
| `<=`    | menor ou igual |
| `>=`    | maior ou igual |

---

## 12. Regras léxicas

### 12.1 Identificadores

Os identificadores são usados para nomear variáveis e funções.

**Regras de formação:**

- Devem começar com letra (`a-z`, `A-Z`) ou sublinhado (`_`)
- Após o primeiro caractere, podem conter letras, dígitos (`0-9`) ou sublinhado
- Tamanho máximo de 30 caracteres

**Expressão regular:**

```
[A-Za-z_][A-Za-z0-9_]{0,29}
```

**Exemplos válidos:** `x`, `idade`, `_media`, `valor1`, `nota_final`

**Exemplos inválidos:** `1x` (começa com número), `@valor` (caractere inválido)

### 12.2 Números inteiros

Sequência de um ou mais dígitos.

**Expressão regular:**

```
[0-9]+
```

**Exemplos:** `0`, `7`, `123`, `99999`

### 12.3 Números reais

Sequência de dígitos, seguida de ponto `.`, seguida de pelo menos um dígito.

**Expressão regular:**

```
[0-9]+\.[0-9]+
```

**Exemplos válidos:** `1.0`, `3.14`, `0.5`, `25.75`

**Exemplos inválidos:** `10.` (falta dígito após o ponto), `.5` (falta dígito antes do ponto), `1..2` (dois pontos)

### 12.4 Espaços em branco

Espaços, tabulações e quebras de linha são ignorados pelo analisador léxico, servindo apenas como separadores entre tokens.

### 12.5 Operadores compostos

Os operadores `>=`, `<=`, `==` e `!=` devem ser reconhecidos como um único token. O analisador léxico deve considerar sempre o maior lexema possível (ex.: `==` é um operador de igualdade, não dois sinais de `=` separados).

---

## 13. Tabela de tokens

### Terminais

| Código | Token   |
|--------|---------|
| 1      | int     |
| 2      | float   |
| 3      | if      |
| 4      | else    |
| 5      | while   |
| 6      | return  |
| 7      | print   |
| 8      | id      |
| 9      | num     |
| 10     | =       |
| 11     | +       |
| 12     | -       |
| 13     | *       |
| 14     | /       |
| 15     | ==      |
| 16     | !=      |
| 17     | <       |
| 18     | >       |
| 19     | <=      |
| 20     | >=      |
| 21     | (       |
| 22     | )       |
| 23     | {       |
| 24     | }       |
| 25     | ,       |
| 26     | ;       |
| 27     | $ (fim de entrada) |

**Observação:** tanto números inteiros quanto reais são reconhecidos pelo token `num` (código 9). A distinção entre inteiro e real pode ser feita pelo valor armazenado como atributo do token.

### Não-terminais

| Código | Símbolo          |
|--------|------------------|
| 28     | \<Program>       |
| 29     | \<FunctionList>  |
| 30     | \<FunctionList'> |
| 31     | \<Function>      |
| 32     | \<ParamListOpt>  |
| 33     | \<ParamList>     |
| 34     | \<ParamList'>    |
| 35     | \<Param>         |
| 36     | \<Block>         |
| 37     | \<DeclListOpt>   |
| 38     | \<DeclList>      |
| 39     | \<DeclList'>     |
| 40     | \<VarDecl>       |
| 41     | \<StmtListOpt>   |
| 42     | \<StmtList>      |
| 43     | \<StmtList'>     |
| 44     | \<Stmt>          |
| 45     | \<AssignStmt>    |
| 46     | \<ReturnStmt>    |
| 47     | \<PrintStmt>     |
| 48     | \<IfStmt>        |
| 49     | \<ElsePart>      |
| 50     | \<WhileStmt>     |
| 51     | \<Expr>          |
| 52     | \<RelExpr>       |
| 53     | \<RelExpr'>      |
| 54     | \<RelOp>         |
| 55     | \<AddExpr>       |
| 56     | \<AddExpr'>      |
| 57     | \<MulExpr>       |
| 58     | \<MulExpr'>      |
| 59     | \<Factor>        |
| 60     | \<FactorTail>    |
| 61     | \<ArgListOpt>    |
| 62     | \<ArgList>       |
| 63     | \<ArgList'>      |
| 64     | \<Type>          |

---

## 14. Erros léxicos

Erro léxico é qualquer sequência de caracteres que não pode ser reconhecida como um token válido.

### 14.1 Caractere inválido

Caractere que não faz parte do alfabeto da linguagem.

```
x = 10 @ 2;    // '@' é inválido
valor = 5 # 1; // '#' é inválido
```

### 14.2 Identificador mal formado

Identificador que não respeita as regras de formação (ex.: começa com número).

```
1abc            // começa com número
nome-completo   // contém '-'
```

### 14.3 Identificador acima do tamanho máximo

Identificadores com mais de 30 caracteres são considerados inválidos.

```
variavel_com_nome_extremamente_grande_demais
```

### 14.4 Número mal formado

Sequência numérica que não forma nem inteiro nem real válido.

```
12.     // falta dígito após o ponto
.5      // falta dígito antes do ponto
1..2    // dois pontos decimais
123abc  // mistura número e letra
```

### 14.5 Comentário de bloco não fechado

Comentário aberto com `/*` sem o respectivo `*/`.

```
/* comentário sem fechar
x = 10;
```

---

## 15. Gramática da linguagem

Abaixo está a gramática completa da linguagem, que servirá de base para a construção do autômato finito e dos analisadores nas próximas etapas:

```
 1. <Program>       ::= <FunctionList>
 2. <FunctionList>  ::= <Function> <FunctionList'>
 3. <FunctionList'> ::= <Function> <FunctionList'>
 4. <FunctionList'> ::= ε
 5. <Function>      ::= <Type> id ( <ParamListOpt> ) <Block>
 6. <ParamListOpt>  ::= <ParamList>
 7. <ParamListOpt>  ::= ε
 8. <ParamList>     ::= <Param> <ParamList'>
 9. <ParamList'>    ::= , <Param> <ParamList'>
10. <ParamList'>    ::= ε
11. <Param>         ::= <Type> id
12. <Block>         ::= { <DeclListOpt> <StmtListOpt> }
13. <DeclListOpt>   ::= <DeclList>
14. <DeclListOpt>   ::= ε
15. <DeclList>      ::= <VarDecl> <DeclList'>
16. <DeclList'>     ::= <VarDecl> <DeclList'>
17. <DeclList'>     ::= ε
18. <VarDecl>       ::= <Type> id ;
19. <StmtListOpt>   ::= <StmtList>
20. <StmtListOpt>   ::= ε
21. <StmtList>      ::= <Stmt> <StmtList'>
22. <StmtList'>     ::= <Stmt> <StmtList'>
23. <StmtList'>     ::= ε
24. <Stmt>          ::= <AssignStmt>
25. <Stmt>          ::= <IfStmt>
26. <Stmt>          ::= <WhileStmt>
27. <Stmt>          ::= <PrintStmt>
28. <Stmt>          ::= <ReturnStmt>
29. <Stmt>          ::= <Block>
30. <AssignStmt>    ::= id = <Expr> ;
31. <ReturnStmt>    ::= return <Expr> ;
32. <PrintStmt>     ::= print ( <Expr> ) ;
33. <IfStmt>        ::= if ( <Expr> ) <Stmt> <ElsePart>
34. <ElsePart>      ::= else <Stmt>
35. <ElsePart>      ::= ε
36. <WhileStmt>     ::= while ( <Expr> ) <Stmt>
37. <Expr>          ::= <RelExpr>
38. <RelExpr>       ::= <AddExpr> <RelExpr'>
39. <RelExpr'>      ::= <RelOp> <AddExpr>
40. <RelExpr'>      ::= ε
41. <RelOp>         ::= ==
42. <RelOp>         ::= !=
43. <RelOp>         ::= <
44. <RelOp>         ::= >
45. <RelOp>         ::= <=
46. <RelOp>         ::= >=
47. <AddExpr>       ::= <MulExpr> <AddExpr'>
48. <AddExpr'>      ::= + <MulExpr> <AddExpr'>
49. <AddExpr'>      ::= - <MulExpr> <AddExpr'>
50. <AddExpr'>      ::= ε
51. <MulExpr>       ::= <Factor> <MulExpr'>
52. <MulExpr'>      ::= * <Factor> <MulExpr'>
53. <MulExpr'>      ::= / <Factor> <MulExpr'>
54. <MulExpr'>      ::= ε
55. <Factor>        ::= ( <Expr> )
56. <Factor>        ::= id <FactorTail>
57. <Factor>        ::= num
58. <FactorTail>    ::= ( <ArgListOpt> )
59. <FactorTail>    ::= ε
60. <ArgListOpt>    ::= <ArgList>
61. <ArgListOpt>    ::= ε
62. <ArgList>       ::= <Expr> <ArgList'>
63. <ArgList'>      ::= , <Expr> <ArgList'>
64. <ArgList'>      ::= ε
65. <Type>          ::= int
66. <Type>          ::= float
```

---

## 16. Exemplo de programa válido

```
/* Programa que calcula o fatorial de um numero
   e imprime o resultado */

int fatorial(int n) {
    int resultado;
    resultado = 1;
    while (n > 1) {
        resultado = resultado * n;
        n = n - 1;
    }
    return resultado;
}

float media(float a, float b) {
    float m;
    m = (a + b) / 2;
    return m;
}

int main() {
    int x;
    int fat;
    float m;

    x = 5;
    fat = fatorial(x);
    print(fat);

    // calcula e imprime a media
    m = media(8.0, 6.5);
    if (m >= 7.0)
        print(m);
    else
        print(0);

    return 0;
}
```

---

## 17. Exemplo de programa com erros léxicos

```
int main() {
    int 1x;
    float valor@;
    1x = 12.;
    /* comentario nao fechado
    print(1x);
}
```

### Erros identificados

- `1x`: identificador mal formado (começa com número)
- `@`: caractere inválido
- `12.`: número real mal formado (falta dígito após o ponto)
- `/* comentario nao fechado`: comentário de bloco sem fechamento
