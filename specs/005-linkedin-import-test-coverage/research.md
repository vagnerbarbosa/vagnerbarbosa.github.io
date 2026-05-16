# Pesquisa de Lacunas de Cobertura: Cobertura de Testes Total

## Objetivo
Identificar as linhas e caminhos exatos que impedem o projeto de atingir 100% de cobertura de instruções.

## Análise por Módulo

### 1. Gerador de Site (`cmd/generator`)
- **Lacuna Principal**: A função `main()` chama `os.Exit(1)` em caso de erro. Como `os.Exit` encerra o processo do teste, as linhas dentro do bloco `if err != nil` do `main()` nunca são marcadas como cobertas.
- **Decisão**: Refatorar `main()` para delegar a execução a uma função `run()` que retorna erro, e usar `os.Exit(runWithExitCode())`.
- **Outras Lacunas**: Erros de permissão de arquivo em `copyAndMinifyDir` e `copyAndMinifyFile` são difíceis de simular sem mocks de filesystem.

### 2. Importador LinkedIn - Parser (`internal/parser`)
- **Lacunas**:
    - `ConvertDate`: Formatos de data que não batem com os regexes de `MMM YYYY` ou `YYYY` retornam a string original. Esse caminho de fallback precisa de testes.
    - `CSVReader.Next()`: Linhas com menos colunas do que o cabeçalho (preenchidas com strings vazias) precisam de validação explícita.
    - `ValidateColumns`: Casos onde múltiplas colunas estão ausentes.

### 3. Importador LinkedIn - Transformer (`internal/transformer`)
- **Lacunas**:
    - `ExtractTechStack`: Casos onde existem múltiplos padrões de tech stack no mesmo texto (deve validar que o último é o escolhido).
    - `parseTechStack`: Normalização de separadores complexos ou strings de tech stack vazias após a limpeza.
    - `SplitDescription`: Casos onde o texto não contém nem parágrafos, nem bullets, nem frases claras, forçando o fallback para um único bullet.

### 4. Importador LinkedIn - Comparator (`internal/comparator`)
- **Lacunas**:
    - `ApplyChanges`: A lógica de atualização do mapa de experiências/educação/certificações precisa de testes que validem a remoção total de itens.
    - `FormatID`: Casos com ponteiros nulos para as entidades.

### 5. Comandos CLI (`cmd/import-linkedin/commands`)
- **Lacunas**:
    - `root.go`: Comandos desconhecidos e a flag de ajuda.
    - `import.go`: Falha na leitura de qualquer um dos três arquivos CSV obrigatórios.
    - `validate.go`: Falha na validação de colunas de um CSV específico.

## Conclusão
Atingir 100% é viável através de:
1. Refatoração do `main()` do gerador.
2. Expansão do `testdata` com arquivos CSV "quebrados".
3. Implementação de testes de tabela (table-driven tests) para todas as funções de parsing e transformação.
