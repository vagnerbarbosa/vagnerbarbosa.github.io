# Modelo de Dados: Extração Inteligente de Tech Stack

**Especificação**: [spec.md](spec.md) | **Plano**: [plan.md](plan.md)
**Data**: 2026-04-26

## Entidades

### Experience (Atualizado)

Representa uma experiência profissional importada do LinkedIn.

```go
type Experience struct {
    Company     string   `yaml:"company" json:"company"`
    Role        string   `yaml:"role" json:"role"`
    StartDate   string   `yaml:"start_date" json:"start_date"`
    EndDate     string   `yaml:"end_date" json:"end_date"`
    Description []string `yaml:"description" json:"description"`
    TechStack   string   `yaml:"tech_stack,omitempty" json:"tech_stack,omitempty"`  // NOVO
    Location    string   `yaml:"location,omitempty" json:"location,omitempty"`
}
```

#### Campos

| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| Company | string | Sim | Nome da empresa |
| Role | string | Sim | Cargo/título |
| StartDate | string | Sim | Data de início |
| EndDate | string | Não | Data de término (ou "Presente") |
| Description | []string | Não | Lista de bullets descrevendo atividades |
| TechStack | string | Não | Tecnologias formatadas separadas por " • " |
| Location | string | Não | Localização da vaga |

#### Regras de Validação

- `Company`, `Role`, `StartDate` são obrigatórios (validação existente)
- `TechStack` é opcional (omitempty)
- Quando presente, `TechStack` deve conter tecnologias separadas por " • "

### TechStackResult (Novo - Internal)

Resultado da extração de tech stack.

```go
type TechStackResult struct {
    CleanedBullets []string // Bullets sem o tech stack
    TechStack      string     // Tech stack formatado ou vazio
    Found          bool       // Indica se tech stack foi encontrado
    PatternMatched string     // Qual padrão foi detectado (para debug)
}
```

## Fluxo de Transformação

### Entrada (CSV do LinkedIn)

```
Description: "Referência técnica em FinOps para o Itaú\nGestão de 7+ desenvolvedores\nAs principais tecnologias e ferramentas utilizadas: Java, Python, AWS"
```

### Processo

1. **SplitDescription** (existente)
   - Input: Texto bruto do CSV
   - Output: `[]string{"Referência técnica em FinOps para o Itaú", "Gestão de 7+ desenvolvedores", "As principais tecnologias e ferramentas utilizadas: Java, Python, AWS"}`

2. **ExtractTechStack** (novo)
   - Input: `[]string` do passo 1
   - Detecta padrão no último elemento
   - Extrai: "Java, Python, AWS"
   - Converte para: "Java • Python • AWS"
   - Output: `TechStackResult`

### Saída (YAML)

```yaml
experiences:
  - company: "Zup Innovation (Itaú)"
    role: "Engenheiro de Software Especialista"
    start_date: "Fev 2025"
    end_date: "Presente"
    description:
      - "Referência técnica em FinOps para o Itaú"
      - "Gestão de 7+ desenvolvedores"
    tech_stack: "Java • Python • AWS"
```

## Patterns de Extração

### Prefixos (case-insensitive)

| Padrão | Regex |
|--------|-------|
| As principais tecnologias e ferramentas utilizadas: | `(?i)as principais tecnologias e ferramentas utilizadas[:\s]*` |
| Tecnologias: | `(?i)tecnologias[:\s]*` |
| Tech Stack: | `(?i)tech stack[:\s]*` |
| Technologies: | `(?i)technologies[:\s]*` |
| Stack: | `(?i)stack[:\s]*` |
| Ferramentas: | `(?i)ferramentas[:\s]*` |
| Tools: | `(?i)tools[:\s]*` |

### Separadores

| Caractere | Descrição |
|-----------|-----------|
| `,` | Vírgula (mais comum) |
| `;` | Ponto-e-vírgula |
| `\|` | Pipe |
| `-` | Hífen (quando seguido de espaço) |
| `•` | Bullet |
| `◦` | Círculo vazio |
| `○` | Círculo |
| `●` | Círculo cheio |
| `·` | Ponto médio |
| ` ` | Espaço (fallback) |

## Exemplos de Casos

### Caso 1: Sucesso - Padrão PT com vírgula

**Entrada**: `"Tecnologias: Java, Python, AWS"`
**Saída**: 
- Description: `[]` (removido)
- TechStack: `"Java • Python • AWS"`

### Caso 2: Sucesso - Padrão EN com bullet

**Entrada**: `"Tech Stack: • Go • Rust • WASM"`
**Saída**:
- Description: `[]` (removido)
- TechStack: `"Go • Rust • WASM"`

### Caso 3: Sucesso - Padrão no meio

**Entrada**: `["Liderança técnica", "Tools: Terraform, Ansible", "Mentoria"]`
**Saída**:
- Description: `["Liderança técnica", "Mentoria"]`
- TechStack: `"Terraform • Ansible"`

### Caso 4: Sem padrão

**Entrada**: `["Liderança técnica", "Mentoria de devs"]`
**Saída**:
- Description: `["Liderança técnica", "Mentoria de devs"]`
- TechStack: `""` (vazio)

### Caso 5: Múltiplos padrões

**Entrada**: `["Tech: Java", "Tech Stack: Python"]`
**Saída**:
- Description: `["Tech: Java"]` (primeiro preservado)
- TechStack: `"Python"` (último usado)
