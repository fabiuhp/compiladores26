package parser

import "compilador26/token"

var productionRightHandSide = map[int][]int{
	1:  {token.NTFunctionList},
	2:  {token.NTFunction, token.NTFunctionListP},
	3:  {token.NTFunction, token.NTFunctionListP},
	4:  {},
	5:  {token.NTType, token.ID, token.LPAREN, token.NTParamListOpt, token.RPAREN, token.NTBlock},
	6:  {token.NTParamList},
	7:  {},
	8:  {token.NTParam, token.NTParamListP},
	9:  {token.COMMA, token.NTParam, token.NTParamListP},
	10: {},
	11: {token.NTType, token.ID},
	12: {token.LBRACE, token.NTDeclListOpt, token.NTStmtListOpt, token.RBRACE},
	13: {token.NTDeclList},
	14: {},
	15: {token.NTVarDecl, token.NTDeclListP},
	16: {token.NTVarDecl, token.NTDeclListP},
	17: {},
	18: {token.NTType, token.ID, token.SEMI},
	19: {token.NTStmtList},
	20: {},
	21: {token.NTStmt, token.NTStmtListP},
	22: {token.NTStmt, token.NTStmtListP},
	23: {},
	24: {token.NTAssignStmt},
	25: {token.NTIfStmt},
	26: {token.NTWhileStmt},
	27: {token.NTPrintStmt},
	28: {token.NTReturnStmt},
	29: {token.NTBlock},
	30: {token.ID, token.ASSIGN, token.NTExpr, token.SEMI},
	31: {token.RETURN, token.NTExpr, token.SEMI},
	32: {token.PRINT, token.LPAREN, token.NTExpr, token.RPAREN, token.SEMI},
	33: {token.IF, token.LPAREN, token.NTExpr, token.RPAREN, token.NTStmt, token.NTElsePart},
	34: {token.ELSE, token.NTStmt},
	35: {},
	36: {token.WHILE, token.LPAREN, token.NTExpr, token.RPAREN, token.NTStmt},
	37: {token.NTRelExpr},
	38: {token.NTAddExpr, token.NTRelExprP},
	39: {token.NTRelOp, token.NTAddExpr},
	40: {},
	41: {token.EQ},
	42: {token.NEQ},
	43: {token.LT},
	44: {token.GT},
	45: {token.LTE},
	46: {token.GTE},
	47: {token.NTMulExpr, token.NTAddExprP},
	48: {token.PLUS, token.NTMulExpr, token.NTAddExprP},
	49: {token.MINUS, token.NTMulExpr, token.NTAddExprP},
	50: {},
	51: {token.NTFactor, token.NTMulExprP},
	52: {token.MULT, token.NTFactor, token.NTMulExprP},
	53: {token.DIV, token.NTFactor, token.NTMulExprP},
	54: {},
	55: {token.LPAREN, token.NTExpr, token.RPAREN},
	56: {token.ID, token.NTFactorTail},
	57: {token.NUM},
	58: {token.LPAREN, token.NTArgListOpt, token.RPAREN},
	59: {},
	60: {token.NTArgList},
	61: {},
	62: {token.NTExpr, token.NTArgListP},
	63: {token.COMMA, token.NTExpr, token.NTArgListP},
	64: {},
	65: {token.INT},
	66: {token.FLOAT},
}

var ruleDescription = map[int]string{
	1:  "Program -> FunctionList",
	2:  "FunctionList -> Function FunctionList'",
	3:  "FunctionList' -> Function FunctionList'",
	4:  "FunctionList' -> e",
	5:  "Function -> Type id ( ParamListOpt ) Block",
	6:  "ParamListOpt -> ParamList",
	7:  "ParamListOpt -> e",
	8:  "ParamList -> Param ParamList'",
	9:  "ParamList' -> , Param ParamList'",
	10: "ParamList' -> e",
	11: "Param -> Type id",
	12: "Block -> { DeclListOpt StmtListOpt }",
	13: "DeclListOpt -> DeclList",
	14: "DeclListOpt -> e",
	15: "DeclList -> VarDecl DeclList'",
	16: "DeclList' -> VarDecl DeclList'",
	17: "DeclList' -> e",
	18: "VarDecl -> Type id ;",
	19: "StmtListOpt -> StmtList",
	20: "StmtListOpt -> e",
	21: "StmtList -> Stmt StmtList'",
	22: "StmtList' -> Stmt StmtList'",
	23: "StmtList' -> e",
	24: "Stmt -> AssignStmt",
	25: "Stmt -> IfStmt",
	26: "Stmt -> WhileStmt",
	27: "Stmt -> PrintStmt",
	28: "Stmt -> ReturnStmt",
	29: "Stmt -> Block",
	30: "AssignStmt -> id = Expr ;",
	31: "ReturnStmt -> return Expr ;",
	32: "PrintStmt -> print ( Expr ) ;",
	33: "IfStmt -> if ( Expr ) Stmt ElsePart",
	34: "ElsePart -> else Stmt",
	35: "ElsePart -> e",
	36: "WhileStmt -> while ( Expr ) Stmt",
	37: "Expr -> RelExpr",
	38: "RelExpr -> AddExpr RelExpr'",
	39: "RelExpr' -> RelOp AddExpr",
	40: "RelExpr' -> e",
	41: "RelOp -> ==",
	42: "RelOp -> !=",
	43: "RelOp -> <",
	44: "RelOp -> >",
	45: "RelOp -> <=",
	46: "RelOp -> >=",
	47: "AddExpr -> MulExpr AddExpr'",
	48: "AddExpr' -> + MulExpr AddExpr'",
	49: "AddExpr' -> - MulExpr AddExpr'",
	50: "AddExpr' -> e",
	51: "MulExpr -> Factor MulExpr'",
	52: "MulExpr' -> * Factor MulExpr'",
	53: "MulExpr' -> / Factor MulExpr'",
	54: "MulExpr' -> e",
	55: "Factor -> ( Expr )",
	56: "Factor -> id FactorTail",
	57: "Factor -> num",
	58: "FactorTail -> ( ArgListOpt )",
	59: "FactorTail -> e",
	60: "ArgListOpt -> ArgList",
	61: "ArgListOpt -> e",
	62: "ArgList -> Expr ArgList'",
	63: "ArgList' -> , Expr ArgList'",
	64: "ArgList' -> e",
	65: "Type -> int",
	66: "Type -> float",
}

var parseTable = map[int]map[int]int{
	token.NTProgram: {
		token.INT: 1, token.FLOAT: 1,
	},
	token.NTFunctionList: {
		token.INT: 2, token.FLOAT: 2,
	},
	token.NTFunctionListP: {
		token.INT: 3, token.FLOAT: 3,
		token.EOF: 4,
	},
	token.NTFunction: {
		token.INT: 5, token.FLOAT: 5,
	},
	token.NTParamListOpt: {
		token.INT: 6, token.FLOAT: 6,
		token.RPAREN: 7,
	},
	token.NTParamList: {
		token.INT: 8, token.FLOAT: 8,
	},
	token.NTParamListP: {
		token.COMMA:  9,
		token.RPAREN: 10,
	},
	token.NTParam: {
		token.INT: 11, token.FLOAT: 11,
	},
	token.NTBlock: {
		token.LBRACE: 12,
	},
	token.NTDeclListOpt: {
		token.INT: 13, token.FLOAT: 13,
		token.ID: 14, token.IF: 14, token.WHILE: 14,
		token.PRINT: 14, token.RETURN: 14,
		token.LBRACE: 14, token.RBRACE: 14,
	},
	token.NTDeclList: {
		token.INT: 15, token.FLOAT: 15,
	},
	token.NTDeclListP: {
		token.INT: 16, token.FLOAT: 16,
		token.ID: 17, token.IF: 17, token.WHILE: 17,
		token.PRINT: 17, token.RETURN: 17,
		token.LBRACE: 17, token.RBRACE: 17,
	},
	token.NTVarDecl: {
		token.INT: 18, token.FLOAT: 18,
	},
	token.NTStmtListOpt: {
		token.ID: 19, token.IF: 19, token.WHILE: 19,
		token.PRINT: 19, token.RETURN: 19, token.LBRACE: 19,
		token.RBRACE: 20,
	},
	token.NTStmtList: {
		token.ID: 21, token.IF: 21, token.WHILE: 21,
		token.PRINT: 21, token.RETURN: 21, token.LBRACE: 21,
	},
	token.NTStmtListP: {
		token.ID: 22, token.IF: 22, token.WHILE: 22,
		token.PRINT: 22, token.RETURN: 22, token.LBRACE: 22,
		token.RBRACE: 23,
	},
	token.NTStmt: {
		token.ID: 24, token.IF: 25, token.WHILE: 26,
		token.PRINT: 27, token.RETURN: 28, token.LBRACE: 29,
	},
	token.NTAssignStmt: {token.ID: 30},
	token.NTReturnStmt: {token.RETURN: 31},
	token.NTPrintStmt:  {token.PRINT: 32},
	token.NTIfStmt:     {token.IF: 33},
	token.NTElsePart: {
		token.ELSE: 34,
		token.ID: 35, token.IF: 35, token.WHILE: 35,
		token.PRINT: 35, token.RETURN: 35,
		token.LBRACE: 35, token.RBRACE: 35,
	},
	token.NTWhileStmt: {
		token.WHILE: 36,
	},
	token.NTExpr: {
		token.LPAREN: 37, token.ID: 37, token.NUM: 37,
	},
	token.NTRelExpr: {
		token.LPAREN: 38, token.ID: 38, token.NUM: 38,
	},
	token.NTRelExprP: {
		token.EQ: 39, token.NEQ: 39, token.LT: 39,
		token.GT: 39, token.LTE: 39, token.GTE: 39,
		token.SEMI: 40, token.RPAREN: 40, token.COMMA: 40,
	},
	token.NTRelOp: {
		token.EQ: 41, token.NEQ: 42, token.LT: 43,
		token.GT: 44, token.LTE: 45, token.GTE: 46,
	},
	token.NTAddExpr: {
		token.LPAREN: 47, token.ID: 47, token.NUM: 47,
	},
	token.NTAddExprP: {
		token.PLUS: 48, token.MINUS: 49,
		token.EQ: 50, token.NEQ: 50, token.LT: 50, token.GT: 50,
		token.LTE: 50, token.GTE: 50,
		token.SEMI: 50, token.RPAREN: 50, token.COMMA: 50,
	},
	token.NTMulExpr: {
		token.LPAREN: 51, token.ID: 51, token.NUM: 51,
	},
	token.NTMulExprP: {
		token.MULT: 52, token.DIV: 53,
		token.PLUS: 54, token.MINUS: 54,
		token.EQ: 54, token.NEQ: 54, token.LT: 54, token.GT: 54,
		token.LTE: 54, token.GTE: 54,
		token.SEMI: 54, token.RPAREN: 54, token.COMMA: 54,
	},
	token.NTFactor: {
		token.LPAREN: 55, token.ID: 56, token.NUM: 57,
	},
	token.NTFactorTail: {
		token.LPAREN: 58,
		token.MULT: 59, token.DIV: 59,
		token.PLUS: 59, token.MINUS: 59,
		token.EQ: 59, token.NEQ: 59, token.LT: 59, token.GT: 59,
		token.LTE: 59, token.GTE: 59,
		token.SEMI: 59, token.RPAREN: 59, token.COMMA: 59,
	},
	token.NTArgListOpt: {
		token.LPAREN: 60, token.ID: 60, token.NUM: 60,
		token.RPAREN: 61,
	},
	token.NTArgList: {
		token.LPAREN: 62, token.ID: 62, token.NUM: 62,
	},
	token.NTArgListP: {
		token.COMMA:  63,
		token.RPAREN: 64,
	},
	token.NTType: {
		token.INT: 65, token.FLOAT: 66,
	},
}

var followSet = map[int]map[int]bool{
	token.NTProgram:       {token.EOF: true},
	token.NTFunctionList:  {token.EOF: true},
	token.NTFunctionListP: {token.EOF: true},
	token.NTFunction: {
		token.INT: true, token.FLOAT: true, token.EOF: true,
	},
	token.NTType:         {token.ID: true},
	token.NTParamListOpt: {token.RPAREN: true},
	token.NTParamList:    {token.RPAREN: true},
	token.NTParamListP:   {token.RPAREN: true},
	token.NTParam: {
		token.COMMA: true, token.RPAREN: true,
	},
	token.NTBlock: {
		token.INT: true, token.FLOAT: true, token.EOF: true,
		token.ID: true, token.IF: true, token.WHILE: true,
		token.PRINT: true, token.RETURN: true,
		token.LBRACE: true, token.RBRACE: true, token.ELSE: true,
	},
	token.NTDeclListOpt: declarationFollow(),
	token.NTDeclList:    declarationFollow(),
	token.NTDeclListP:   declarationFollow(),
	token.NTVarDecl: {
		token.INT: true, token.FLOAT: true,
		token.ID: true, token.IF: true, token.WHILE: true,
		token.PRINT: true, token.RETURN: true,
		token.LBRACE: true, token.RBRACE: true,
	},
	token.NTStmtListOpt: {token.RBRACE: true},
	token.NTStmtList:    {token.RBRACE: true},
	token.NTStmtListP:   {token.RBRACE: true},
	token.NTStmt:        statementFollow(),
	token.NTAssignStmt:  statementFollow(),
	token.NTReturnStmt:  statementFollow(),
	token.NTPrintStmt:   statementFollow(),
	token.NTIfStmt:      statementFollow(),
	token.NTWhileStmt:   statementFollow(),
	token.NTElsePart:    statementFollow(),
	token.NTExpr: {
		token.SEMI: true, token.RPAREN: true, token.COMMA: true,
	},
	token.NTRelExpr: {
		token.SEMI: true, token.RPAREN: true, token.COMMA: true,
	},
	token.NTRelExprP: {
		token.SEMI: true, token.RPAREN: true, token.COMMA: true,
	},
	token.NTRelOp: {
		token.LPAREN: true, token.ID: true, token.NUM: true,
	},
	token.NTAddExpr:     relationalFollow(),
	token.NTAddExprP:    relationalFollow(),
	token.NTMulExpr:     additiveFollow(),
	token.NTMulExprP:    additiveFollow(),
	token.NTFactor:      multiplicativeFollow(),
	token.NTFactorTail:  multiplicativeFollow(),
	token.NTArgListOpt:  {token.RPAREN: true},
	token.NTArgList:     {token.RPAREN: true},
	token.NTArgListP:    {token.RPAREN: true},
}

func declarationFollow() map[int]bool {
	return map[int]bool{
		token.ID: true, token.IF: true, token.WHILE: true,
		token.PRINT: true, token.RETURN: true,
		token.LBRACE: true, token.RBRACE: true,
	}
}

func statementFollow() map[int]bool {
	follow := declarationFollow()
	follow[token.ELSE] = true
	return follow
}

func relationalFollow() map[int]bool {
	return map[int]bool{
		token.EQ: true, token.NEQ: true, token.LT: true, token.GT: true,
		token.LTE: true, token.GTE: true,
		token.SEMI: true, token.RPAREN: true, token.COMMA: true,
	}
}

func additiveFollow() map[int]bool {
	follow := relationalFollow()
	follow[token.PLUS] = true
	follow[token.MINUS] = true
	return follow
}

func multiplicativeFollow() map[int]bool {
	follow := additiveFollow()
	follow[token.MULT] = true
	follow[token.DIV] = true
	return follow
}
