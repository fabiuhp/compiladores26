package token

// Códigos dos terminais (conforme especificação)
const (
	INT    = 1  // int
	FLOAT  = 2  // float
	IF     = 3  // if
	ELSE   = 4  // else
	WHILE  = 5  // while
	RETURN = 6  // return
	PRINT  = 7  // print
	ID     = 8  // identificador
	NUM    = 9  // número (inteiro ou real)
	ASSIGN = 10 // =
	PLUS   = 11 // +
	MINUS  = 12 // -
	MULT   = 13 // *
	DIV    = 14 // /
	EQ     = 15 // ==
	NEQ    = 16 // !=
	LT     = 17 // <
	GT     = 18 // >
	LTE    = 19 // <=
	GTE    = 20 // >=
	LPAREN = 21 // (
	RPAREN = 22 // )
	LBRACE = 23 // {
	RBRACE = 24 // }
	COMMA  = 25 // ,
	SEMI   = 26 // ;
	EOF    = 27 // $ (fim de entrada - usado internamente pelo parser)
)

// Códigos dos não-terminais (conforme especificação)
const (
	NTProgram       = 28
	NTFunctionList  = 29
	NTFunctionListP = 30
	NTFunction      = 31
	NTParamListOpt  = 32
	NTParamList     = 33
	NTParamListP    = 34
	NTParam         = 35
	NTBlock         = 36
	NTDeclListOpt   = 37
	NTDeclList      = 38
	NTDeclListP     = 39
	NTVarDecl       = 40
	NTStmtListOpt   = 41
	NTStmtList      = 42
	NTStmtListP     = 43
	NTStmt          = 44
	NTAssignStmt    = 45
	NTReturnStmt    = 46
	NTPrintStmt     = 47
	NTIfStmt        = 48
	NTElsePart      = 49
	NTWhileStmt     = 50
	NTExpr          = 51
	NTRelExpr       = 52
	NTRelExprP      = 53
	NTRelOp         = 54
	NTAddExpr       = 55
	NTAddExprP      = 56
	NTMulExpr       = 57
	NTMulExprP      = 58
	NTFactor        = 59
	NTFactorTail    = 60
	NTArgListOpt    = 61
	NTArgList       = 62
	NTArgListP      = 63
	NTType          = 64
)

// Token representa um token reconhecido pelo analisador léxico
type Token struct {
	Code   int
	Lexeme string
	Line   int
}

// Keywords mapeia palavras reservadas para seus códigos
var Keywords = map[string]int{
	"int":    INT,
	"float":  FLOAT,
	"if":     IF,
	"else":   ELSE,
	"while":  WHILE,
	"return": RETURN,
	"print":  PRINT,
}

// terminalNames mapeia código do terminal para seu nome
var terminalNames = map[int]string{
	INT:    "int",
	FLOAT:  "float",
	IF:     "if",
	ELSE:   "else",
	WHILE:  "while",
	RETURN: "return",
	PRINT:  "print",
	ID:     "id",
	NUM:    "num",
	ASSIGN: "=",
	PLUS:   "+",
	MINUS:  "-",
	MULT:   "*",
	DIV:    "/",
	EQ:     "==",
	NEQ:    "!=",
	LT:     "<",
	GT:     ">",
	LTE:    "<=",
	GTE:    ">=",
	LPAREN: "(",
	RPAREN: ")",
	LBRACE: "{",
	RBRACE: "}",
	COMMA:  ",",
	SEMI:   ";",
	EOF:    "$",
}

// nonTerminalNames mapeia código do não-terminal para seu nome
var nonTerminalNames = map[int]string{
	NTProgram:       "Program",
	NTFunctionList:  "FunctionList",
	NTFunctionListP: "FunctionList'",
	NTFunction:      "Function",
	NTParamListOpt:  "ParamListOpt",
	NTParamList:     "ParamList",
	NTParamListP:    "ParamList'",
	NTParam:         "Param",
	NTBlock:         "Block",
	NTDeclListOpt:   "DeclListOpt",
	NTDeclList:      "DeclList",
	NTDeclListP:     "DeclList'",
	NTVarDecl:       "VarDecl",
	NTStmtListOpt:   "StmtListOpt",
	NTStmtList:      "StmtList",
	NTStmtListP:     "StmtList'",
	NTStmt:          "Stmt",
	NTAssignStmt:    "AssignStmt",
	NTReturnStmt:    "ReturnStmt",
	NTPrintStmt:     "PrintStmt",
	NTIfStmt:        "IfStmt",
	NTElsePart:      "ElsePart",
	NTWhileStmt:     "WhileStmt",
	NTExpr:          "Expr",
	NTRelExpr:       "RelExpr",
	NTRelExprP:      "RelExpr'",
	NTRelOp:         "RelOp",
	NTAddExpr:       "AddExpr",
	NTAddExprP:      "AddExpr'",
	NTMulExpr:       "MulExpr",
	NTMulExprP:      "MulExpr'",
	NTFactor:        "Factor",
	NTFactorTail:    "FactorTail",
	NTArgListOpt:    "ArgListOpt",
	NTArgList:       "ArgList",
	NTArgListP:      "ArgList'",
	NTType:          "Type",
}

func TerminalName(code int) string {
	if name, ok := terminalNames[code]; ok {
		return name
	}
	return "?"
}

func NonTerminalName(code int) string {
	if name, ok := nonTerminalNames[code]; ok {
		return name
	}
	return "?"
}

func SymbolName(code int) string {
	if code >= 1 && code <= 27 {
		return TerminalName(code)
	}
	if code >= 28 && code <= 64 {
		return NonTerminalName(code)
	}
	return "?"
}

func IsTerminal(code int) bool {
	return code >= 1 && code <= 27
}

func IsNonTerminal(code int) bool {
	return code >= 28 && code <= 64
}
