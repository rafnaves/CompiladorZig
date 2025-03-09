package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	STRING
	IDENTIFIER

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	ASSIGNMENT // =
	EQUALS // ==
	NOT // !
	NOT_EQUALS

	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS

	AND

	DOT 
	DOT_dOT 
	SEMI_COLON 
	COLON 
	QUESTION 
	COMMA 

	PLUS_PLUS 
	MINUS_MINUS 
	PLUS_EQUALS 
	MINUS_EQUALS 

	PLUS 
	DASH 
	SLASH 
	STAR 
	PERCENT

	DOT_DOT
	NULLISH_ASSIGNMENT


	//PALAVRAS RESERVADAS

	ADDRSPACE
	ALIGN
	ALLOWZERO


	ANYFRAME
	ANYTYPE

	ASM
	ASYNC
	AWAIT

	BREAK

	CATCH
	COMPTIME
	CONST
	CONTINUE

	DEFER
	ELSE
	ENUM
	ERRDEFER
	ERROR
	EXPORT
	EXTERN

	FALSE
	FN
	FOR

	IF
	INLINE

	NOALIAS
	NOSUSPEND
	NULL

	OPAQUE
	OR
	ORELSE

	PACKED
	PUB

	RESUME
	RETURN

	STRUCT
	SUSPEND
	SWITCH

	TEST
	THREADLOCAL
	TRUE
	TRY

	UNDEFINED
	UNION
	UNREACHABLE
	USINGNAMESPACE

	VAR
	VOLATILE

	WHILE


)

type Token struct {
	Kind  TokenKind // Antes era `kind`, precisa ser `Kind`
	Value string    // Antes era `value`, precisa ser `Value`
}


func (token Token) isOneOfMany (expectedTokens ...TokenKind) bool{
	for _, expexpected := range expectedTokens{
		if expexpected == token.Kind{
			return true
		}
	}
	return false
}

func(token Token) Debug () {
	if token.isOneOfMany(IDENTIFIER, NUMBER, STRING) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)

	} else{
		fmt.Printf("%s ()\n", TokenKindString(token.Kind))
	}

}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}



func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case IDENTIFIER:
		return "identifier"
	case OPEN_BRACKET:
		return "open_bracket"
	case CLOSE_BRACKET:
		return "close_bracket"
	case OPEN_CURLY:
		return "open_curly"
	case CLOSE_CURLY:
		return "close_curly"
	case OPEN_PAREN:
		return "open_paren"
	case CLOSE_PAREN:
		return "close_paren"
	case ASSIGNMENT:
		return "assignment"
	case EQUALS:
		return "equals"
	case NOT:
		return "not"
	case NOT_EQUALS:
		return "not_equals"
	case LESS:
		return "less"
	case LESS_EQUALS:
		return "less_equals"
	case GREATER:
		return "greater"
	case GREATER_EQUALS:
		return "greater_equals"
	case OR:
		return "or"
	case AND:
		return "and"
	case DOT:
		return "dot"
	case DOT_DOT:
		return "dot_dot"
	case SEMI_COLON:
		return "semi_colon"
	case COLON:
		return "colon"
	case QUESTION:
		return "question"
	case COMMA:
		return "comma"
	case PLUS_PLUS:
		return "plus_plus"
	case MINUS_MINUS:
		return "minus_minus"
	case PLUS_EQUALS:
		return "plus_equals"
	case MINUS_EQUALS:
		return "minus_equals"
	case PLUS:
		return "plus"
	case DASH:
		return "dash"
	case SLASH:
		return "slash"
	case STAR:
		return "star"
	case PERCENT:
		return "percent"
	case ADDRSPACE:
		return "addrspace"
	case ALIGN:
		return "align"
	case ALLOWZERO:
		return "allowzero"
	case ANYFRAME:
		return "anyframe"
	case ANYTYPE:
		return "anytype"
	case ASM:
		return "asm"
	case ASYNC:
		return "async"
	case AWAIT:
		return "await"
	case BREAK:
		return "break"
	case CATCH:
		return "catch"
	case COMPTIME:
		return "comptime"
	case CONST:
		return "const"
	case CONTINUE:
		return "continue"
	case DEFER:
		return "defer"
	case ELSE:
		return "else"
	case ENUM:
		return "enum"
	case ERRDEFER:
		return "errdefer"
	case ERROR:
		return "error"
	case EXPORT:
		return "export"
	case EXTERN:
		return "extern"
	case FALSE:
		return "false"
	case FN:
		return "fn"
	case FOR:
		return "for"
	case IF:
		return "if"
	case INLINE:
		return "inline"
	case NOALIAS:
		return "noalias"
	case NOSUSPEND:
		return "nosuspend"
	case NULL:
		return "null"
	case OPAQUE:
		return "opaque"
	case ORELSE:
		return "orelse"
	case PACKED:
		return "packed"
	case PUB:
		return "pub"
	case RESUME:
		return "resume"
	case RETURN:
		return "return"
	case STRUCT:
		return "struct"
	case SUSPEND:
		return "suspend"
	case SWITCH:
		return "switch"
	case TEST:
		return "test"
	case THREADLOCAL:
		return "threadlocal"
	case TRUE:
		return "true"
	case TRY:
		return "try"
	case UNDEFINED:
		return "undefined"
	case UNION:
		return "union"
	case UNREACHABLE:
		return "unreachable"
	case USINGNAMESPACE:
		return "usingnamespace"
	case VAR:
		return "var"
	case VOLATILE:
		return "volatile"
	case WHILE:
		return "while"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}
