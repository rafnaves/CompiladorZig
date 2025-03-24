package parser

import (
	"compilador/src/ast"
	"compilador/src/lexer"
)

func parse_stmt(p *parser) ast.Stmt{
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists{
		return stmt_fn(p)

	}

	expression := parse_expr(p, defalt_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}

func parse_var_decl_stmt(p *parser) ast.Stmt {
	IsConstant := p.advance().Kind == lexer.CONST
	varName := p.expectedError(lexer.IDENTIFIER, "Era esperado encontrar o nome da variavel dentro da declaração").Value
	p.expect(lexer.ASSIGNMENT)
	assignedValue := parse_expr(p, assignment)
	p.expect(lexer.SEMI_COLON)
	
	return ast.VarDeclStmt{
		IsConstant: IsConstant,
		VariableName: varName,
		AssignedValue: assignedValue,
	}

}