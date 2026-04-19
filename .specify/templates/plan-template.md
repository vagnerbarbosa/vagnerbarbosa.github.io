# Plano de Implementação: [FUNCIONALIDADE]

**Branch**: `[###-nome-da-funcionalidade]` | **Data**: [DATA] | **Spec**: [link]
**Input**: Especificação da funcionalidade de `/specs/[###-nome-da-funcionalidade]/spec.md`

**Nota**: Este modelo é preenchido pelo comando `/speckit.plan`. Veja `.specify/templates/plan-template.md` para o fluxo de execução.

## Resumo

[Extraído da especificação da funcionalidade: requisito principal + abordagem técnica da pesquisa]

## Contexto Técnico

<!--
  AÇÃO NECESSÁRIA: Substitua o conteúdo desta seção com os detalhes técnicos
  do projeto. A estrutura aqui é apresentada de forma consultiva para guiar
  o processo de iteração.
-->

**Linguagem/Versão**: [ex: Python 3.11, Swift 5.9, Rust 1.75 ou PRECISA DE ESCLARECIMENTO]  
**Dependências Principais**: [ex: FastAPI, UIKit, LLVM ou PRECISA DE ESCLARECIMENTO]  
**Armazenamento**: [se aplicável, ex: PostgreSQL, CoreData, arquivos ou N/A]  
**Testes**: [ex: pytest, XCTest, cargo test ou PRECISA DE ESCLARECIMENTO]  
**Plataforma Alvo**: [ex: Linux server, iOS 15+, WASM ou PRECISA DE ESCLARECIMENTO]
**Tipo de Projeto**: [ex: library/cli/web-service/mobile-app/compiler/desktop-app ou PRECISA DE ESCLARECIMENTO]  
**Metas de Performance**: [específico do domínio, ex: 1000 req/s, 10k linhas/seg, 60 fps ou PRECISA DE ESCLARECIMENTO]  
**Restrições**: [específico do domínio, ex: <200ms p95, <100MB memória, offline-capable ou PRECISA DE ESCLARECIMENTO]  
**Escala/Alcance**: [específico do domínio, ex: 10k usuários, 1M LOC, 50 telas ou PRECISA DE ESCLARECIMENTO]

## Verificação da Constituição

*PORTÃO: Deve passar antes da pesquisa da Fase 0. Re-verificar após o design da Fase 1.*

[Portões determinados baseados no arquivo da constituição]

## Estrutura do Projeto

### Documentação (esta funcionalidade)

```text
specs/[###-funcionalidade]/
├── plan.md              # Este arquivo (output do comando /speckit.plan)
├── research.md          # Output da Fase 0 (comando /speckit.plan)
├── data-model.md        # Output da Fase 1 (comando /speckit.plan)
├── quickstart.md        # Output da Fase 1 (comando /speckit.plan)
├── contracts/           # Output da Fase 1 (comando /speckit.plan)
└── tasks.md             # Output da Fase 2 (comando /speckit.tasks - NÃO criado pelo /speckit.plan)
```

### Código Fonte (raiz do repositório)
<!--
  AÇÃO NECESSÁRIA: Substitua a árvore de placeholders abaixo com o layout concreto
  para esta funcionalidade. Delete opções não utilizadas e expanda a estrutura escolhida
  com caminhos reais (ex: apps/admin, packages/something). O plano entregue não deve
  incluir labels de Opção.
-->

```text
# [REMOVER SE NÃO USADO] Opção 1: Projeto único (PADRÃO)
src/
├── models/
├── services/
├── cli/
└── lib/

tests/
├── contract/
├── integration/
└── unit/

# [REMOVER SE NÃO USADO] Opção 2: Aplicação web (quando "frontend" + "backend" detectado)
backend/
├── src/
│   ├── models/
│   ├── services/
│   └── api/
└── tests/

frontend/
├── src/
│   ├── components/
│   ├── pages/
│   └── services/
└── tests/

# [REMOVER SE NÃO USADO] Opção 3: Mobile + API (quando "iOS/Android" detectado)
api/
└── [mesmo que backend acima]

ios/ ou android/
└── [estrutura específica da plataforma: módulos de feature, fluxos de UI, testes de plataforma]
```

**Decisão de Estrutura**: [Documente a estrutura selecionada e referencie os diretórios reais
capturados acima]

## Rastreamento de Complexidade

> **Preencha SOMENTE se a Verificação da Constituição tiver violações que devam ser justificadas**

| Violação | Por Que Necessário | Alternativa Mais Simples Rejeitada Porque |
|----------|-------------------|------------------------------------------|
| [ex: 4º projeto] | [necessidade atual] | [por que 3 projetos são insuficientes] |
| [ex: Padrão Repository] | [problema específico] | [por que acesso direto ao DB é insuficiente] |
