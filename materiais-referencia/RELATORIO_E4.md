# Relatório da Etapa E4 — Analisador Sintático

## Trabalho Semestral de Compiladores — 2026/1

- Érick Lúcio de Oliveira
- Fábio Henrique Pereira
- Giovana Olivo Cittadin
- Gustavo de Almeida Kammer

---

## 1. Visão Geral

Este relatório apresenta a construção do analisador sintático **descendente preditivo tabular sem backtracking** (LL(1)) para a gramática definida no manual da linguagem (E1).

A construção segue o algoritmo apresentado na Aula 4 (slides 23–24), cumprindo as quatro condições para análise LL(1):

1. A gramática é livre de contexto.
2. A gramática está fatorada.
3. A gramática não possui recursão à esquerda.
4. Para todo A ∈ N tal que A ::=* ε, FIRST(A) ∩ FOLLOW(A) = ∅.

O parser usa três estruturas:

- **Entrada**: sequência de tokens vinda do analisador léxico (E3)
- **Pilha**: inicializada com `$` no fundo e o símbolo inicial `Program` no topo
- **Tabela de Parsing M(X, a)**: matriz onde X é não-terminal e a é terminal

---

## 2. Gramática

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

## 3. Conjuntos FIRST

| Não-terminal      | FIRST                                        |
|-------------------|----------------------------------------------|
| Program           | { int, float }                               |
| FunctionList      | { int, float }                               |
| FunctionList'     | { int, float, ε }                            |
| Function          | { int, float }                               |
| ParamListOpt      | { int, float, ε }                            |
| ParamList         | { int, float }                               |
| ParamList'        | { , , ε }                                    |
| Param             | { int, float }                               |
| Block             | { { }                                        |
| DeclListOpt       | { int, float, ε }                            |
| DeclList          | { int, float }                               |
| DeclList'         | { int, float, ε }                            |
| VarDecl           | { int, float }                               |
| StmtListOpt       | { id, if, while, print, return, {, ε }       |
| StmtList          | { id, if, while, print, return, { }          |
| StmtList'         | { id, if, while, print, return, {, ε }       |
| Stmt              | { id, if, while, print, return, { }          |
| AssignStmt        | { id }                                       |
| ReturnStmt        | { return }                                   |
| PrintStmt         | { print }                                    |
| IfStmt            | { if }                                       |
| ElsePart          | { else, ε }                                  |
| WhileStmt         | { while }                                    |
| Expr              | { (, id, num }                               |
| RelExpr           | { (, id, num }                               |
| RelExpr'          | { ==, !=, <, >, <=, >=, ε }                  |
| RelOp             | { ==, !=, <, >, <=, >= }                     |
| AddExpr           | { (, id, num }                               |
| AddExpr'          | { +, -, ε }                                  |
| MulExpr           | { (, id, num }                               |
| MulExpr'          | { *, /, ε }                                  |
| Factor            | { (, id, num }                               |
| FactorTail        | { (, ε }                                     |
| ArgListOpt        | { (, id, num, ε }                            |
| ArgList           | { (, id, num }                               |
| ArgList'          | { , , ε }                                    |
| Type              | { int, float }                               |

---

## 4. Conjuntos FOLLOW

Aplicando as três regras do FOLLOW (Aula 4, slide 26):

1. FOLLOW do símbolo inicial contém `$`.
2. Para `A → αBβ`, FIRST(β) (exceto ε) está em FOLLOW(B).
3. Para `A → αB` (ou `A → αBβ` com ε ∈ FIRST(β)), FOLLOW(A) está em FOLLOW(B).

| Não-terminal      | FOLLOW                                                                  |
|-------------------|-------------------------------------------------------------------------|
| Program           | { $ }                                                                   |
| FunctionList      | { $ }                                                                   |
| FunctionList'     | { $ }                                                                   |
| Function          | { int, float, $ }                                                       |
| ParamListOpt      | { ) }                                                                   |
| ParamList         | { ) }                                                                   |
| ParamList'        | { ) }                                                                   |
| Param             | { , , ) }                                                               |
| Block             | { int, float, $, id, if, while, print, return, {, }, else }             |
| DeclListOpt       | { id, if, while, print, return, {, } }                                  |
| DeclList          | { id, if, while, print, return, {, } }                                  |
| DeclList'         | { id, if, while, print, return, {, } }                                  |
| VarDecl           | { int, float, id, if, while, print, return, {, } }                      |
| StmtListOpt       | { } }                                                                   |
| StmtList          | { } }                                                                   |
| StmtList'         | { } }                                                                   |
| Stmt              | { id, if, while, print, return, {, }, else }                            |
| AssignStmt        | { id, if, while, print, return, {, }, else }                            |
| ReturnStmt        | { id, if, while, print, return, {, }, else }                            |
| PrintStmt         | { id, if, while, print, return, {, }, else }                            |
| IfStmt            | { id, if, while, print, return, {, }, else }                            |
| WhileStmt         | { id, if, while, print, return, {, }, else }                            |
| ElsePart          | { id, if, while, print, return, {, }, else }                            |
| Type              | { id }                                                                  |
| Expr              | { ; , ) , , }                                                           |
| RelExpr           | { ; , ) , , }                                                           |
| RelExpr'          | { ; , ) , , }                                                           |
| RelOp             | { ( , id, num }                                                         |
| AddExpr           | { ==, !=, <, >, <=, >=, ; , ) , , }                                     |
| AddExpr'          | { ==, !=, <, >, <=, >=, ; , ) , , }                                     |
| MulExpr           | { +, -, ==, !=, <, >, <=, >=, ; , ) , , }                               |
| MulExpr'          | { +, -, ==, !=, <, >, <=, >=, ; , ) , , }                               |
| Factor            | { *, /, +, -, ==, !=, <, >, <=, >=, ; , ) , , }                         |
| FactorTail        | { *, /, +, -, ==, !=, <, >, <=, >=, ; , ) , , }                         |
| ArgListOpt        | { ) }                                                                   |
| ArgList           | { ) }                                                                   |
| ArgList'          | { ) }                                                                   |

---

## 5. Tabela de Parsing M(X, a)

A tabela é construída a partir dos conjuntos FIRST e FOLLOW:

- Para cada produção `A → α`: para cada terminal `t ∈ FIRST(α)`, M[A, t] = número da regra.
- Se `ε ∈ FIRST(α)`: para cada terminal `t ∈ FOLLOW(A)`, M[A, t] = número da regra ε.

A entrada vazia indica erro sintático.

| NT \\ T            | int | float | if  | else | while | return | print | id  | num | =   | +   | -   | *   | /   | ==  | !=  | <   | >   | <=  | >=  | (   | )   | {   | }   | ,   | ;   | $   |
|--------------------|-----|-------|-----|------|-------|--------|-------|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|-----|
| Program            | 1   | 1     |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| FunctionList       | 2   | 2     |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| FunctionList'      | 3   | 3     |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 4   |
| Function           | 5   | 5     |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| ParamListOpt       | 6   | 6     |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 7   |     |     |     |     |     |
| ParamList          | 8   | 8     |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| ParamList'         |     |       |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 10  |     |     | 9   |     |     |
| Param              | 11  | 11    |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| Block              |     |       |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 12  |     |     |     |     |
| DeclListOpt        | 13  | 13    | 14  |      | 14    | 14     | 14    | 14  |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 14  | 14  |     |     |     |
| DeclList           | 15  | 15    |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| DeclList'          | 16  | 16    | 17  |      | 17    | 17     | 17    | 17  |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 17  | 17  |     |     |     |
| VarDecl            | 18  | 18    |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| StmtListOpt        |     |       | 19  |      | 19    | 19     | 19    | 19  |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 19  | 20  |     |     |     |
| StmtList           |     |       | 21  |      | 21    | 21     | 21    | 21  |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 21  |     |     |     |     |
| StmtList'          |     |       | 22  |      | 22    | 22     | 22    | 22  |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 22  | 23  |     |     |     |
| Stmt               |     |       | 25  |      | 26    | 28     | 27    | 24  |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 29  |     |     |     |     |
| AssignStmt         |     |       |     |      |       |        |       | 30  |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| ReturnStmt         |     |       |     |      |       | 31     |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| PrintStmt          |     |       |     |      |       |        | 32    |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| IfStmt             |     |       | 33  |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| ElsePart           |     |       | 35  | 34   | 35    | 35     | 35    | 35  |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 35  | 35  |     |     |     |
| WhileStmt          |     |       |     |      | 36    |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |
| Expr               |     |       |     |      |       |        |       | 37  | 37  |     |     |     |     |     |     |     |     |     |     |     | 37  |     |     |     |     |     |     |
| RelExpr            |     |       |     |      |       |        |       | 38  | 38  |     |     |     |     |     |     |     |     |     |     |     | 38  |     |     |     |     |     |     |
| RelExpr'           |     |       |     |      |       |        |       |     |     |     |     |     |     |     | 39  | 39  | 39  | 39  | 39  | 39  |     | 40  |     |     | 40  | 40  |     |
| RelOp              |     |       |     |      |       |        |       |     |     |     |     |     |     |     | 41  | 42  | 43  | 44  | 45  | 46  |     |     |     |     |     |     |     |
| AddExpr            |     |       |     |      |       |        |       | 47  | 47  |     |     |     |     |     |     |     |     |     |     |     | 47  |     |     |     |     |     |     |
| AddExpr'           |     |       |     |      |       |        |       |     |     |     | 48  | 49  |     |     | 50  | 50  | 50  | 50  | 50  | 50  |     | 50  |     |     | 50  | 50  |     |
| MulExpr            |     |       |     |      |       |        |       | 51  | 51  |     |     |     |     |     |     |     |     |     |     |     | 51  |     |     |     |     |     |     |
| MulExpr'           |     |       |     |      |       |        |       |     |     |     | 54  | 54  | 52  | 53  | 54  | 54  | 54  | 54  | 54  | 54  |     | 54  |     |     | 54  | 54  |     |
| Factor             |     |       |     |      |       |        |       | 56  | 57  |     |     |     |     |     |     |     |     |     |     |     | 55  |     |     |     |     |     |     |
| FactorTail         |     |       |     |      |       |        |       |     |     |     | 59  | 59  | 59  | 59  | 59  | 59  | 59  | 59  | 59  | 59  | 58  | 59  |     |     | 59  | 59  |     |
| ArgListOpt         |     |       |     |      |       |        |       | 60  | 60  |     |     |     |     |     |     |     |     |     |     |     | 60  | 61  |     |     |     |     |     |
| ArgList            |     |       |     |      |       |        |       | 62  | 62  |     |     |     |     |     |     |     |     |     |     |     | 62  |     |     |     |     |     |     |
| ArgList'           |     |       |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     | 64  |     |     | 63  |     |     |
| Type               | 65  | 66    |     |      |       |        |       |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |     |

### 5.1 Resolução do conflito do dangling else

A entrada M[ElsePart, else] poderia ter regra 34 (de FIRST(else Stmt)) e regra 35 (porque else ∈ FOLLOW(ElsePart)). Adotamos a convenção padrão: o `else` casa com o `if` mais próximo, então a regra 34 tem prioridade.

---

## 6. Algoritmo do Analisador

Conforme apresentado nos slides 23–24 da Aula 4 (descendente preditivo tabular sem backtracking):

```
Inicializa pilha com [$, Program]

Repita
    X = topo da pilha
    a = token corrente da entrada

    Se X = $ então
        Se a = $: ACEITA
        Senão: erro

    Senão Se X é terminal então
        Se X = a: pop X, avança entrada
        Senão: erro (pop X — modo pânico Técnica 5)

    Senão (X é não-terminal):
        Se M[X, a] tem regra:
            pop X
            empilha conteúdo da regra (em ordem reversa)
        Senão:
            erro (modo pânico Técnica 1: descarta tokens
                  até encontrar um símbolo de FOLLOW(X),
                  então pop X)
```

---

## 7. Recuperação de Erros (Modo Pânico)

Conforme ensinado na Aula 4 (Erros Sintáticos), o parser não para no primeiro erro. Usamos as técnicas:

- **Técnica 1**: quando M[X, a] não existe (X não-terminal), descarta tokens da entrada até encontrar um símbolo em FOLLOW(X), depois faz pop de X.
- **Técnica 5**: quando o terminal no topo da pilha não casa com a entrada, faz pop do terminal e segue.

Há um limite de 50 erros para evitar avalanche e laço infinito. Cada operação de recuperação encurta a pilha ou avança a entrada, garantindo progresso.

---

## 8. Saída do Parser

A cada modificação da pilha, o parser exibe:

- **Passo**: número sequencial
- **Token**: lexema e código numérico
- **Linha**: linha onde o token foi encontrado
- **Ação**: regra aplicada, consumo de terminal, erro detectado, ou ação de recuperação
- **Pilha**: conteúdo da pilha do fundo (`$`) até o topo

Ao final, exibe `ACEITA` (programa aceito) ou a lista completa de erros sintáticos detectados.

---

## 9. Arquivos de Exemplo

- **exemplo1.txt**: programa válido (fatorial com função, while, if/else, comentário de bloco e linha)
- **exemplo2.txt**: programa válido (média com múltiplas funções e expressões)
- **exemplo3.txt**: programa com 5 erros léxicos
- **exemplo4.txt**: programa sem erro léxico, com erro sintático (falta `;`)
- **exemplo5.txt**: programa com múltiplos erros sintáticos para demonstrar recuperação em modo pânico
