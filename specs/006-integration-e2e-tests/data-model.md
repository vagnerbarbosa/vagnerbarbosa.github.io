# Modelo de Dados de Teste: Integração e E2E

Este documento define os fluxos de dados e as expectativas de saída para a suíte de testes.

## 1. Fluxo de Integração (Pipeline LinkedIn)

### Entrada (Input)
- **Fonte**: `cmd/import-linkedin/testdata/input_csv/*.csv`
- **Formato**: CSVs exportados do LinkedIn (Experiências, Educação, Certificações).

### Processamento
`Parser` $\rightarrow$ `Transformer` $\rightarrow$ `Comparator` $\rightarrow$ `YAML Generator`.

### Expectativa (Golden File)
- **Arquivo**: `cmd/import-linkedin/testdata/golden/*.yaml`
- **Validação**: O YAML final deve conter as entidades processadas com as seguintes regras:
  - `experiences`: Títulos limpos, datas normalizadas.
  - `education`: Instituições e períodos corretos.
  - `certifications`: Nomes de certificados e emissores validados.

---

## 2. Fluxo End-to-End (E2E)

### Componentes do Ambiente Temporário
Para cada teste E2E, é criado um ambiente com:
- `templates/` (Cópia do projeto)
- `assets/` (Cópia do projeto)
- `config.yaml` (Vazio ou com baseline mínima)

### Sequência de Ações
1. **Ação**: `import-linkedin import --input [CSV_TESTE]`
   - **Esperado**: `config.yaml` populado com dados do CSV.
2. **Ação**: `generator` (Execução do main.go do gerador)
   - **Esperado**: Criação da pasta `public/` com arquivos minificados.

### Validação de Output (HTML final)
| Alvo | Verificação | Critério de Sucesso |
| :--- | :--- | :--- |
| `public/index.html` | Busca de String | Presença de tags `<h2 class="education-title">` com conteúdo do CSV. |
| `public/sitemap.xml` | Existência e URL | Arquivo presente com a URL base configurada. |
| `public/robots.txt` | Conteúdo | Presença de `User-agent: *` e link para o sitemap. |
| `public/site.webmanifest` | Estrutura JSON | Validade do JSON e presença de cores de tema. |

## 3. Matriz de Cenários

| Cenário | Input | Expectativa de Integração | Expectativa E2E |
| :--- | :--- | :--- | :--- |
| **Cenário Feliz** | CSVs válidos e completos | YAML idêntico ao Golden | HTML com todas as seções preenchidas |
| **Dados Parciais** | CSV com colunas ausentes | YAML com campos nulos/defaults | HTML renderizando apenas campos disponíveis |
| **Erro de Parsing** | CSV malformado | Erro de importação reportado | Site gerado com dados anteriores ou vazio |
