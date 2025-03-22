package ast

import "compilador/src/lexer"

// expressões literais
type NumberExpr struct{
	Value float64
}

func(n NumberExpr) expr(){}

type StringExpr struct{
	Value string
}

func(n StringExpr) expr(){}

type SymbolExpr struct{
	Value string
}

func(n SymbolExpr) expr(){}


// Expressões complexas

// 10 + 5 * 2

type BinaryExpr struct{
	Left 	Expr
	Operator lexer.Token
	Right 	Expr

}

func(n BinaryExpr) expr() {}