# Guia Rápido: Execução de Testes de Cobertura

## Como rodar os testes
Para executar todos os testes do projeto e gerar o relatório de cobertura:

```bash
go test ./... -cover
```

## Como analisar lacunas detalhadas
Para ver exatamente quais linhas não foram cobertas em um pacote específico:

1. Gere o arquivo de cobertura:
   ```bash
   go test -coverprofile=coverage.out ./cmd/generator
   ```
2. Visualize no navegador:
   ```bash
   go tool cover -html=coverage.out
   ```

## Estratégia de Adição de Testes
1. Identifique a linha não coberta no HTML.
2. Crie um caso de teste no arquivo `*_test.go` correspondente que force a execução daquele caminho.
3. Execute o teste e valide o incremento da porcentagem.
