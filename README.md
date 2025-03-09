Até agora, o projeto **realiza a análise léxica** do código-fonte Zig. Ele funciona da seguinte maneira:  

### 🔹 **O que ele faz?**  
1. **Lê um arquivo de código Zig** e o armazena como uma string.  
2. **Tokeniza (divide em partes) o código** usando expressões regulares para identificar elementos como números, operadores, palavras-chave e símbolos.  
3. **Gera uma lista de tokens**, que representam cada elemento do código de maneira estruturada.  
4. **Imprime os tokens no console**, ajudando a visualizar a análise léxica.  

### 🔹 **Como ele faz?**  
✅ **Leitura do arquivo:** O código-fonte Zig é lido como uma string no `main.go`.  
✅ **Criação do analisador léxico:** O `lexer` usa expressões regulares para identificar padrões no código.  
✅ **Avanço caractere por caractere:** O lexer percorre a string e aplica as regras de reconhecimento de tokens.  
✅ **Armazenamento de tokens:** Cada token identificado é adicionado a uma lista.  
✅ **Exibição dos tokens:** Cada token é impresso com seu tipo e valor.  
