# Estratégia de Testes

Este documento descreve a estratégia de testes para o site de portfólio e seu pipeline de importação.

## Visão Geral

O projeto utiliza uma abordagem de testes em múltiplas camadas para garantir a integridade dos dados, desde a importação do LinkedIn até a renderização final do HTML.

## Camadas de Teste

### 1. Testes Unitários
- **Foco**: Componentes pequenos e isolados (parsers, transformadores, gerenciadores de configuração).
- **Objetivo**: Validar que funções individuais se comportem corretamente sob várias entradas.
- **Execução**: `go test ./...`

### 2. Testes de Integração (Pipeline de Importação)
- **Foco**: O fluxo de arquivos CSV $\rightarrow$ configuração YAML.
- **Método**: **Golden Files**.
  - CSVs de referência são processados.
  - O YAML resultante é comparado com um arquivo de referência "golden".
  - É utilizada a comparação semântica (unmarshaling para structs) para ignorar diferenças insignificantes de formatação.
- **Objetivo**: Garantir que o pipeline de importação seja determinístico e correto.

### 3. Testes End-to-End (E2E)
- **Foco**: Fluxo completo de CSV $\rightarrow$ YAML $\rightarrow$ HTML.
- **Método**:
  - Um ambiente temporário é criado com um subconjunto de templates e assets.
  - O pipeline de importação é disparado usando dados de teste em CSV.
  - O gerador do site é executado.
  - O `index.html` resultante é escaneado em busca de "marcadores" específicos (strings únicas) que devem estar presentes para cada seção (Experiência, Educação, Certificações).
- **Objetivo**: Garantir que os dados importados realmente sejam renderizados no site final.

## Dados de Teste

Os dados de teste para integração e E2E estão localizados em:
`cmd/import-linkedin/testdata/e2e/`

## Executando os Testes

### Todos os Testes
```bash
go test -v ./...
```

### Apenas Testes E2E
```bash
go test -v ./cmd/generator/...
```

## Portões de Qualidade (Quality Gates)

- **Cobertura de Código**: O projeto busca alta cobertura de testes, particularmente no pipeline de importação (meta >95%).
- **CI/CD**: Todos os testes são executados automaticamente em cada Pull Request via GitHub Actions.
