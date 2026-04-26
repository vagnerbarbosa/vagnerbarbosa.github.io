# Quickstart: Extração Inteligente de Tech Stack

**Especificação**: [spec.md](spec.md) | **Plano**: [plan.md](plan.md)

## Visão Rápida

Após implementação, o LinkedIn Import CLI detectará automaticamente tech stack em descrições de experiências.

## Como Funciona

### Antes (Comportamento Atual)

```yaml
experiences:
  - company: "Zup Innovation"
    role: "Engenheiro de Software"
    description:
      - "Referência técnica em FinOps"
      - "As principais tecnologias e ferramentas utilizadas: Java, Python, AWS"
```

### Depois (Novo Comportamento)

```yaml
experiences:
  - company: "Zup Innovation"
    role: "Engenheiro de Software"
    description:
      - "Referência técnica em FinOps"
    tech_stack: "Java • Python • AWS"
```

## Padrões Suportados

O sistema detecta automaticamente quando um bullet contém:

| Padrão (case-insensitive) | Exemplo detectado |
|---------------------------|-------------------|
| "As principais tecnologias e ferramentas utilizadas:" | `As principais tecnologias e ferramentas utilizadas: Java, Python` |
| "Tecnologias:" | `Tecnologias: Kubernetes, Docker` |
| "Tech Stack:" | `Tech Stack: Go \| Rust \| WASM` |
| "Technologies:" | `Technologies: Node.js, TypeScript` |
| "Stack:" | `Stack: AWS, GCP` |
| "Ferramentas:" | `Ferramentas: Terraform - Ansible` |
| "Tools:" | `Tools: React, Vue` |

## Formatos de Entrada

Todos estes formatos são automaticamente convertidos para `"Tecnologia • Tecnologia"`:

- `Java, Python, AWS` (vírgula)
- `Java; Python; AWS` (ponto-e-vírgula)
- `Java | Python | AWS` (pipe)
- `Java - Python - AWS` (hífen)
- `• Java • Python • AWS` (bullet)

## Uso

Nenhuma mudança necessária no uso da CLI:

```bash
go run cmd/import-linkedin/main.go import
```

O sistema processa automaticamente durante a importação do CSV.

## Validação

Para verificar se a extração funcionou corretamente:

1. Execute a importação com `--dry-run` para visualizar as mudanças
2. Verifique se `tech_stack` aparece no diff para experiências que contêm padrões
3. Confirme que bullets de descrição não contêm mais textos de tech stack
