# Analisador Léxico e Sintático para Zig

Este projeto implementa um analisador léxico e sintático para a linguagem de programação Zig. Ele é capaz de ler código-fonte, realizar a tokenização, construir a Árvore de Sintaxe Abstrata (AST) e exibir os resultados.

## O que ele faz?

- **Leitura do código Zig**: O arquivo de código-fonte Zig é lido e armazenado como uma string.
- **Tokenização**: O código é dividido em partes utilizando expressões regulares para identificar números, operadores, palavras-chave e símbolos.
- **Geração de Tokens**: Cada elemento do código é representado por um token, estruturado para representar seus valores e tipos.
- **Análise Sintática**: A partir dos tokens gerados, a Árvore de Sintaxe Abstrata (AST) é construída, representando a estrutura do código de forma hierárquica.
- **Exibição dos Tokens e AST**: Os tokens gerados e a AST são exibidos no console, proporcionando uma visualização clara da análise léxica e sintática.

## ⚙️ Como ele faz?

1. **Leitura do arquivo**: O código-fonte Zig é lido como uma string no `main.go`.
2. **Criação do analisador léxico**: O lexer utiliza expressões regulares para identificar padrões no código e gerar os tokens.
3. **Construção da AST (Parsing)**: A partir dos tokens gerados pelo lexer, a função de parsing constrói a Árvore de Sintaxe Abstrata, processando as expressões com base nos tokens lidos.
4. **Avanço caractere por caractere**: O lexer percorre a string do código, aplicando as regras de reconhecimento de tokens.
5. **Armazenamento de tokens**: Cada token identificado é adicionado a uma lista para posterior análise sintática.
6. **Exibição dos resultados**: Os tokens e a AST gerada são impressos no console.

## 🛠️ Como usar?

1. Certifique-se de que o código Zig a ser analisado esteja no diretório `./examples/`.
2. Compile e execute o código Go com o seguinte comando:

   ```bash
   go run src/main.go


## 🚀 Como usar?

1. Clone o repositório:
   ```bash
   git clone <url-do-repositório>

2. Navegue até o diretório do projeto:
   ```bash
   cd <diretório-do-projeto>

3. Compile e execute
   ```bash
   go run main.go <caminho-do-arquivo-zig>



