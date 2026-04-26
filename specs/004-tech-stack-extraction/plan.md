# Plano de Implementação: Extração Inteligente de Tech Stack

**Branch**: `004-tech-stack-extraction` | **Data**: 2026-04-26 | **Spec**: [spec.md](spec.md)
**Input**: Especificação da funcionalidade de `/specs/004-tech-stack-extraction/spec.md`

**Nota**: Este modelo é preenchido pelo comando `/speckit.plan`. Veja `.specify/templates/plan-template.md` para o fluxo de execução.

## Resumo

Implementar inteligência no LinkedIn Import CLI para detectar automaticamente tech stack em bullets de descrição de experiências. O sistema deve reconhecer padrões multilíngues (PT/EN), extrair as tecnologias em diferentes formatos de entrada e convertê-las para um formato padronizado separado por " • ", armazenando no campo `TechStack` do model `Experience`.

## Contexto Técnico

**Linguagem/Versão**: Go 1.25
**Dependências Principais**: charmbracelet/huh (UI), tdewolff/minify (minificação)
**Armazenamento**: YAML files (config.yaml)
**Testes**: go test (built-in)
**Plataforma Alvo**: Linux/macOS/Windows CLI
**Tipo de Projeto**: CLI tool (LinkedIn Import)
**Metas de Performance**: Parse de CSV em < 1s para 100+ experiências
**Restrições**: Manter compatibilidade com formato atual de config.yaml
**Escala/Alcance**: Importação de CSVs do LinkedIn com até 1000 linhas

## Verificação da Constituição

*PORTÃO: Deve passar antes da pesquisa da Fase 0. Re-verificar após o design da Fase 1.*

### Princípios da Constituição do Projeto

| Princípio | Status | Justificativa |
|-----------|--------|---------------|
| I. Minimalismo Intencional | ✅ Passa | A solução é simples: pattern matching e extração de texto |
| II. Bilíngue por Padrão | ✅ Passa | Suporte a padrões em PT e EN conforme especificado |
| III. Estabilidade Visual | ✅ Passa | Não afeta UI/frontend |
| IV. Build Reprodutível | ✅ Passa | Apenas transformação de dados, comportamento deterministico |
| V. Código como Documentação | ✅ Passa | Nomes de funções devem ser claros (ExtractTechStack, etc) |

## Estrutura do Projeto

### Documentação (esta funcionalidade)

```text
specs/004-tech-stack-extraction/
├── plan.md              # Este arquivo
├── research.md          # Decisões técnicas (Fase 0)
├── data-model.md        # Modelos de dados atualizados (Fase 1)
├── quickstart.md        # Guia rápido de uso
├── contracts/           # Interfaces da CLI
└── tasks.md             # Tarefas de implementação (Fase 2)
```

### Código Fonte (raiz do repositório)

```text
cmd/import-linkedin/
├── commands/
│   └── import.go        # Entry point do comando import
├── internal/
│   ├── models/
│   │   └── experience.go # Model Experience (adicionar TechStack)
│   ├── parser/
│   │   └── experience.go # Parser CSV (integrar extração)
│   └── transformer/
│       ├── description.go     # SplitDescription existente
│       └── techstack.go       # NOVO: Extração de tech stack
└── tests/
    └── unit/              # Testes unitários
```

**Decisão de Estrutura**: Criar novo arquivo `techstack.go` no pacote `transformer` para manter separação de responsabilidades. O campo `TechStack` será adicionado ao struct `Experience` em `models/experience.go`.

## Rastreamento de Complexidade

> Nenhuma violação de constituição identificada. A implementação é minimalista e focada.

## Design Técnico

### Componentes Principais

1. **TechStackExtractor** (`transformer/techstack.go`)
   - Função principal: `ExtractTechStack(bullets []string) (cleanedBullets []string, techStack string)`
   - Pattern matching com regex para detectar prefixos
   - Parsing de tecnologias com múltiplos separadores
   - Formatação de saída padronizada

2. **Patterns Suportados (case-insensitive)**
   ```go
   var techStackPatterns = []string{
       `as principais tecnologias e ferramentas utilizadas[:\s]`,
       `tecnologias[:\s]`,
       `tech stack[:\s]`,
       `technologies[:\s]`,
       `stack[:\s]`,
       `ferramentas[:\s]`,
       `tools[:\s]`,
   }
   ```

3. **Separadores Suportados**
   ```go
   var separators = []string{
       `,`, `;`, `|`, `-`, `•`, `◦`, `○`, `●`, `·`,
   }
   ```

4. **Flow de Processamento**
   ```
   CSV Description → SplitDescription → ExtractTechStack → 
   [cleanedBullets (sem tech stack), techStack (formatado)]
   ```

### Alterações em Arquivos Existentes

1. **models/experience.go**: Adicionar campo `TechStack string`
2. **parser/experience.go**: Chamar `ExtractTechStack` após `SplitDescription`
3. **internal/config/yaml.go**: Mapear `TechStack` no YAML de saída
