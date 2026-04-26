# Pesquisa Técnica: Extração Inteligente de Tech Stack

**Especificação**: [spec.md](spec.md) | **Plano**: [plan.md](plan.md)
**Data**: 2026-04-26

## Decisões Técnicas

### DT-001: Abordagem de Pattern Matching

**Decisão**: Usar expressões regulares (regex) para detecção de padrões

**Racional**:
- Padrões são bem definidos e previsíveis (lista fechada de prefixos)
- Regex oferece case-insensitive matching nativo em Go (`(?i)`)
- Performance adequada para o volume esperado (até 1000 experiências)
- Manutenibilidade simples - adicionar novos padrões é trivial

**Alternativas consideradas**:
- String matching simples: Rejeitado porque não lida bem com variações de espaçamento/case
- NLP/ML: Rejeitado - overkill para problema bem definido
- Parser customizado: Rejeitado - complexidade desnecessária

### DT-002: Estrutura de Dados para Resultado

**Decisão**: Criar struct `TechStackResult` com slices e metadata

**Racional**:
- Go não tem tuplas nativas - struct é idiomático
- Campos explícitos melhoram legibilidade vs retornar múltiplos valores
- Campo `Found` permite verificação booleana simples
- Campo `PatternMatched` auxilia em debugging

**Alternativas consideradas**:
- Retornar `([]string, string)`: Rejeitado - não é auto-documentado
- Retornar `([]string, string, bool)`: Aceitável mas verboso nos callers

### DT-003: Tratamento de Múltiplos Padrões

**Decisão**: Usar o ÚLTIMO bullet que contém padrão de tech stack

**Racional**:
- Conforme especificação do usuário, tech stack tipicamente é o último bullet
- Se múltiplos bullets contêm padrões, preservar o primeiro (descrição) e extrair do último
- Isso evita perda de dados acidental

**Exemplo**:
```
["Tech Lead", "Tools: OldTech", "Tecnologias: NewTech"]
→ Description: ["Tech Lead", "Tools: OldTech"]
→ TechStack: "NewTech"
```

### DT-004: Formato de Saída

**Decisão**: Usar separador `" • "` (bullet + espaços)

**Racional**:
- Consistente com formato atual em `config.yaml`
- Visualmente limpo no YAML e no site
- Unicode bullet (U+2022) é bem suportado em todos os sistemas

**Alternativas consideradas**:
- `", "`: Rejeitado - menos legível no YAML
- `" | "`: Considerado mas bullet é mais comum em portfólios
- `" · "`: Bullet médio - menos comum que bullet padrão

### DT-005: Localização da Lógica

**Decisão**: Criar novo arquivo `techstack.go` no pacote `transformer`

**Racional**:
- Separação de responsabilidades: `description.go` faz split, `techstack.go` faz extração
- Facilita testes unitários isolados
- Evita crescimento excessivo de um único arquivo
- Segue princípio "Código como Documentação" da constituição

### DT-006: Casos de Borda

**Decisão**: Implementar comportamentos específicos para casos de borda

| Caso | Comportamento |
|------|---------------|
| Tech stack no meio da descrição | Extrair e remover da lista, preservar outros bullets |
| Múltiplos tech stacks | Usar o último, preservar bullets anteriores |
| Tecnologia vazia após prefixo | Retornar tech_stack vazio, remover bullet |
| Separadores misturados | Suportar todos na mesma linha |
| Espaços extras | Trim em cada tecnologia individualmente |
| Tecnologia repetida | Preservar como está (não deduplicar) |

## Referências

- Go regex package: https://pkg.go.dev/regexp
- Go strings package: https://pkg.go.dev/strings
- Projeto similar em Python: https://github.com/johnmcconnell/jolly (para referência de patterns)
