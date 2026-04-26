# Pesquisa Técnica: LinkedIn Import CLI

**Data**: 2026-04-26  
**Feature**: LinkedIn Import CLI  
**Contexto**: Go 1.21+, CSV parsing, YAML manipulation, CLI interativa

---

## 1. CSV Parsing em Go

### Decision
Usar `encoding/csv` da standard library do Go.

### Rationale
- A standard library é robusta, bem testada e suficiente para o parsing de CSVs do LinkedIn
- Suporta RFC 4180 (formato padrão de CSV)
- Não adiciona dependências externas (alinha com Princípio I: Minimalismo Intencional)
- Performance adequada para arquivos de tamanho moderado (exportações LinkedIn típicas)

### Alternatives Considered
| Biblioteca | Pros | Cons |
|------------|------|------|
| `encoding/csv` (stdlib) | Zero deps, bem documentada, RFC 4180 compliant | Menos flexível para edge cases |
| `github.com/josephcopenhaver/csv-go` | Zero-allocation, mais performante | Dependência externa, overkill para este caso |

### Implementation Notes
- Usar streaming (`Read()` em loop) ao invés de `ReadAll()` para economia de memória
- Configurar `FieldsPerRecord` = 0 para auto-detecção
- Habilitar `LazyQuotes` para lidar com possíveis inconsistências no CSV do LinkedIn
- Sempre verificar `io.EOF` e `ParseError` detalhados

---

## 2. YAML Manipulation em Go

### Decision
Usar `go.yaml.in/yaml/v4` (sucessor oficial do gopkg.in/yaml.v3).

### Rationale
- `gopkg.in/yaml.v3` foi marcado como unmaintained em Abril 2025
- A YAML Organization assumiu manutenção oficial em `go.yaml.in/yaml`
- v4 está em desenvolvimento ativo com novas features
- API compatível com v3 para migração suave
- Preserva comentários e formatação original do YAML (importante para config.yaml)

### Alternatives Considered
| Biblioteca | Status | Recommendation |
|------------|--------|----------------|
| `gopkg.in/yaml.v3` | 🔒 Unmaintained | Evitar - apenas security fixes |
| `go.yaml.in/yaml/v4` | ✅ Active development | **Usar para novos projetos** |
| `github.com/goccy/go-yaml` | ✅ Active | Boa performance, API diferente |
| `sigs.k8s.io/yaml` | ✅ Active | Usada por Kubernetes, feature set limitado |

### Implementation Notes
- Usar `yaml.Node` para preservar estrutura e comentários ao modificar o arquivo
- Para leitura simples, pode usar unmarshaling direto em structs
- Para escrita, usar `yaml.Node` ou `Encoder` com `Indent` configurado

---

## 3. CLI Interativa em Go

### Decision
Usar `github.com/charmbracelet/huh` v2.0.3 para prompts interativos.

### Rationale
- Biblioteca oficial do ecossistema Charmbracelet
- Suporte built-in para confirmações yes/no (campo `Confirm`)
- Acessibilidade integrada (screen readers)
- Temas customizáveis (5 temas built-in incluindo Dracula e Catppuccin)
- API simples e declarativa
- Compatível com `bubbletea` para extensão futura

### Alternatives Considered
| Biblioteca | Use Case | Status |
|------------|----------|--------|
| `charmbracelet/huh` | Forms e prompts | **Recomendado** |
| `charmbracelet/bubbletea` | TUIs complexas | Usar huh que é built on top |
| `erikgeiser/promptkit` | Standalone prompts | Boa alternativa |
| `indaco/prompti` | Confirmações customizadas | Dialog/inline modes |

### Implementation Notes
- `huh.NewConfirm()` para yes/no simples
- `huh.NewForm()` para múltiplas perguntas em sequência
- Usar `huh.WithTheme()` para consistência visual
- Suporta `Affirmative()` e `Negative()` para labels customizadas

---

## 4. Diff Visual Colorido

### Decision
Usar `github.com/sergi/go-diff` combinado com customização para saída colorida no terminal.

### Rationale
- `go-diff` é a implementação Go do algoritmo diff
- Biblioteca leve e focada
- Permite construir visualização colorida customizada usando `lipgloss` (Charmbracelet)

### Alternatives Considered
| Biblioteca | Pros | Cons |
|------------|------|------|
| `github.com/sergi/go-diff` | Implementação padrão, leve | Sem saída colorida nativa |
| `github.com/pmezard/go-difflib` | Simples, Python-like | Menos ativa |
| Custom + `lipgloss` | Controle total | Mais código para manter |

### Implementation Notes
- Usar `diffmatchpatch` do `go-diff` para gerar patches
- Aplicar cores via `github.com/charmbracelet/lipgloss`:
  - Verde para adições (+)
  - Vermelho para remoções (-)
  - Amarelo para modificações (~)
- Criar diff por seção (experiences, education, certifications)

---

## 5. Conversão de Datas (Inglês → Português)

### Decision
Implementar mapa de conversão customizado para meses e termos.

### Rationale
- O formato do LinkedIn usa abreviações de meses em inglês (Jan, Feb, Mar...)
- Necessário converter para português (Jan → Jan, Feb → Fev, Mar → Mar...)
- Termo "Present" deve converter para "Presente"
- Solução simples não requer bibliotecas externas

### Implementation Notes
```go
var monthMap = map[string]string{
    "Jan": "Jan", "Feb": "Fev", "Mar": "Mar",
    "Apr": "Abr", "May": "Mai", "Jun": "Jun",
    "Jul": "Jul", "Aug": "Ago", "Sep": "Set",
    "Oct": "Out", "Nov": "Nov", "Dec": "Dez",
}
// "Present" → "Presente"
```

---

## 6. Divisão de Descrição em Bullets

### Decision
Usar heurística simples: dividir por quebras de linha (`\n\n`) ou frases separadas por ponto final.

### Rationale
- Descrições do LinkedIn costumam ser parágrafos ou listas
- Quebras de linha duplas geralmente indicam separação de tópicos
- Frases terminadas em `. ` seguidas de maiúscula indicam novos bullets
- Solução simples atinge precisão >90% (requisito CS-003)

### Implementation Notes
- Tentar primeiro separar por `\n\n` (parágrafos)
- Se resultado for apenas 1 item, tentar separar por ponto final seguido de maiúscula
- Limpar espaços em branco extras
- Ignorar bullets vazios

---

## Referências

- [encoding/csv - Go Packages](https://pkg.go.dev/encoding/csv@go1.25.6)
- [yaml/go-yaml - GitHub](https://github.com/yaml/go-yaml)
- [go.yaml.in/yaml/v4 - pkg.go.dev](https://pkg.go.dev/go.yaml.in/yaml/v4)
- [charmbracelet/huh - GitHub](https://github.com/charmbracelet/huh)
- [charmbracelet/lipgloss - GitHub](https://github.com/charmbracelet/lipgloss)
- [sergi/go-diff - GitHub](https://github.com/sergi/go-diff)
