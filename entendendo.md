## Tokenization

Tokenization é o processo de transformar uma sequencia de caracteres em um array de tokens que fazem sentido

"""
input: let nums = [A, 20, 30]

output : LET SYMBOL EQ BRACKET SYMBOL ...
"""

## Abstract Syntax Tree 
Vamos construir a AST processando os tokens que recebemos do lexer. O processo de construir um AST dos tokens é chamado de Parsing

## Objetivos

- construir um tokenizer

# Prat Parser

É uma analise que construi uma AST

## Binding Power BP

É como um token se liga aos tokens vizinhos. Um + vai ter um poder de ligação menor que um *

## Null Denotation (NUD)

type nud_handler func (p*parser) ast.Expr

Um token que tem um LUD handler, ou seja, nada é esperado a sua esquerda, como por exemplo, um prefixo e expressões unarias.

## Left Denotation (led):

type led_handler func(p*parser,left ast.Expr, bp binding_power) ast.Expr

Tokens que se espera estar entre ou depois de alguma expressão

## Lookup Tables

Usando lookup tables junto de funções NUD/LED, consiguimos criar um parser

## Exemplo de sintaxe

Declaração de variaveis em zig tem essa syntax:

Variavel Mútavel :
var nome: i32 = 42;  // Declara uma variável mutável do tipo `i32` com valor inicial 42


Variável Imutável (const):
const nome: i32 = 42;  // Declara uma constante (imutável) do tipo `i32` com valor inicial 42

Inferência de Tipo
Zig também pode inferir o tipo com base no valor inicial:

var idade = 25;       // Tipo inferido como `i32`
const pi = 3.1415;    // Tipo inferido como `f64`

### Ast Schema
type VarDeclarationStmt struct {
    identifier string
    Constant bool
    AssignedValue Expr
    ExplicitType Type
}