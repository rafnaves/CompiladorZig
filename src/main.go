package main

import (
	"compilador/src/lexer"
	"compilador/src/parser"
	"os"

	"github.com/sanity-io/litter"
)

func main() {
	bytes, _ := os.ReadFile("./examples/03.lang")
	tokens := lexer.Tokenize(string(bytes))

	ast := parser.Parse(tokens)
	litter.Dump(ast)
}


// Para analisar apenas lexamente
/*
func main() {
	bytes, _ := os.ReadFile("./examples/01.lang")
	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}
}
*/