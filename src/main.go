package main

import (
	"os"

	"compilador/src/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/01.lang")
	source := string(bytes)

	tokens := lexer.Tokenize(source)

	for _, token := range tokens {
		token.Debug()
	}
}
