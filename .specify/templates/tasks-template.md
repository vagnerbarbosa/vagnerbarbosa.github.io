---

description: "Modelo de lista de tarefas para implementação de funcionalidades"
---

# Tarefas: [NOME DA FUNCIONALIDADE]

**Input**: Documentos de design de `/specs/[###-nome-da-funcionalidade]/`
**Pré-requisitos**: plan.md (obrigatório), spec.md (obrigatório para histórias de usuário), research.md, data-model.md, contracts/

**Testes**: Os exemplos abaixo incluem tarefas de teste. Testes são OPCIONAIS - inclua-os apenas se explicitamente solicitados na especificação da funcionalidade.

**Organização**: Tarefas são agrupadas por história de usuário para permitir implementação e teste independentes de cada história.

## Formato: `[ID] [P?] [Story] Descrição`

- **[P]**: Pode rodar em paralelo (arquivos diferentes, sem dependências)
- **[Story]**: A qual história de usuário esta tarefa pertence (ex: HU1, HU2, HU3)
- Inclua caminhos de arquivo exatos nas descrições

## Convenções de Caminho

- **Projeto único**: `src/`, `tests/` na raiz do repositório
- **App web**: `backend/src/`, `frontend/src/`
- **Mobile**: `api/src/`, `ios/src/` ou `android/src/`
- Caminhos mostrados abaixo assumem projeto único - ajuste baseado na estrutura do plan.md

<!-- 
  ============================================================================
  IMPORTANTE: As tarefas abaixo são TAREFAS DE EXEMPLO apenas para fins ilustrativos.
  
  O comando /speckit.tasks DEVE substituir estas com tarefas reais baseadas em:
  - Histórias de usuário do spec.md (com suas prioridades P1, P2, P3...)
  - Requisitos da funcionalidade do plan.md
  - Entidades do data-model.md
  - Endpoints de contracts/
  
  Tarefas DEVEM ser organizadas por história de usuário para que cada história possa ser:
  - Implementada independentemente
  - Testada independentemente
  - Entregue como um incremento MVP
  
  NÃO mantenha estas tarefas de exemplo no arquivo tasks.md gerado.
  ============================================================================
-->

## Fase 1: Setup (Infraestrutura Compartilhada)

**Propósito**: Inicialização do projeto e estrutura básica

- [ ] T001 Criar estrutura de projeto conforme plano de implementação
- [ ] T002 Inicializar projeto [linguagem] com dependências [framework]
- [ ] T003 [P] Configurar ferramentas de lint e formatação

---

## Fase 2: Fundacional (Pré-requisitos Bloqueantes)

**Propósito**: Infraestrutura central que DEVE estar completa antes de QUALQUER história de usuário ser implementada

**⚠️ CRÍTICO**: Nenhum trabalho de história de usuário pode começar até esta fase estar completa

Exemplos de tarefas fundacionais (ajuste baseado no seu projeto):

- [ ] T004 Configurar esquema de banco de dados e framework de migrações
- [ ] T005 [P] Implementar framework de autenticação/autorização
- [ ] T006 [P] Configurar estrutura de rotas e middleware de API
- [ ] T007 Criar modelos/entidades base de que todas as histórias dependem
- [ ] T008 Configurar tratamento de erros e infraestrutura de logging
- [ ] T009 Configurar gerenciamento de configuração de ambiente

**Checkpoint**: Fundação pronta - implementação de histórias de usuário pode agora começar em paralelo

---

## Fase 3: História de Usuário 1 - [Título] (Prioridade: P1) 🎯 MVP

**Objetivo**: [Breve descrição do que esta história entrega]

**Teste Independente**: [Como verificar que esta história funciona por conta própria]

### Testes para História de Usuário 1 (OPCIONAL - apenas se testes solicitados) ⚠️

> **NOTA: Escreva estes testes PRIMEIRO, garanta que FALHEM antes da implementação**

- [ ] T010 [P] [HU1] Teste de contrato para [endpoint] em tests/contract/test_[nome].py
- [ ] T011 [P] [HU1] Teste de integração para [jornada do usuário] em tests/integration/test_[nome].py

### Implementação da História de Usuário 1

- [ ] T012 [P] [HU1] Criar modelo [Entidade1] em src/models/[entidade1].py
- [ ] T013 [P] [HU1] Criar modelo [Entidade2] em src/models/[entidade2].py
- [ ] T014 [HU1] Implementar [Serviço] em src/services/[servico].py (depende de T012, T013)
- [ ] T015 [HU1] Implementar [endpoint/funcionalidade] em src/[localizacao]/[arquivo].py
- [ ] T016 [HU1] Adicionar validação e tratamento de erros
- [ ] T017 [HU1] Adicionar logging para operações da história de usuário 1

**Checkpoint**: Neste ponto, a História de Usuário 1 deve estar totalmente funcional e testável independentemente

---

## Fase 4: História de Usuário 2 - [Título] (Prioridade: P2)

**Objetivo**: [Breve descrição do que esta história entrega]

**Teste Independente**: [Como verificar que esta história funciona por conta própria]

### Testes para História de Usuário 2 (OPCIONAL - apenas se testes solicitados) ⚠️

- [ ] T018 [P] [HU2] Teste de contrato para [endpoint] em tests/contract/test_[nome].py
- [ ] T019 [P] [HU2] Teste de integração para [jornada do usuário] em tests/integration/test_[nome].py

### Implementação da História de Usuário 2

- [ ] T020 [P] [HU2] Criar modelo [Entidade] em src/models/[entidade].py
- [ ] T021 [HU2] Implementar [Serviço] em src/services/[servico].py
- [ ] T022 [HU2] Implementar [endpoint/funcionalidade] em src/[localizacao]/[arquivo].py
- [ ] T023 [HU2] Integrar com componentes da História de Usuário 1 (se necessário)

**Checkpoint**: Neste ponto, as Histórias de Usuário 1 E 2 devem ambas funcionar independentemente

---

## Fase 5: História de Usuário 3 - [Título] (Prioridade: P3)

**Objetivo**: [Breve descrição do que esta história entrega]

**Teste Independente**: [Como verificar que esta história funciona por conta própria]

### Testes para História de Usuário 3 (OPCIONAL - apenas se testes solicitados) ⚠️

- [ ] T024 [P] [HU3] Teste de contrato para [endpoint] em tests/contract/test_[nome].py
- [ ] T025 [P] [HU3] Teste de integração para [jornada do usuário] em tests/integration/test_[nome].py

### Implementação da História de Usuário 3

- [ ] T026 [P] [HU3] Criar modelo [Entidade] em src/models/[entidade].py
- [ ] T027 [HU3] Implementar [Serviço] em src/services/[servico].py
- [ ] T028 [HU3] Implementar [endpoint/funcionalidade] em src/[localizacao]/[arquivo].py

**Checkpoint**: Todas as histórias de usuário devem agora ser independentemente funcionais

---

[Adicione mais fases de histórias de usuário conforme necessário, seguindo o mesmo padrão]

---

## Fase N: Polimento & Preocupações Transversais

**Propósito**: Melhorias que afetam múltiplas histórias de usuário

- [ ] TXXX [P] Atualizações de documentação em docs/
- [ ] TXXX Limpeza e refatoração de código
- [ ] TXXX Otimização de performance entre todas as histórias
- [ ] TXXX [P] Testes unitários adicionais (se solicitados) em tests/unit/
- [ ] TXXX Hardening de segurança
- [ ] TXXX Executar validação do quickstart.md

---

## Dependências & Ordem de Execução

### Dependências de Fase

- **Setup (Fase 1)**: Sem dependências - pode começar imediatamente
- **Fundacional (Fase 2)**: Depende da conclusão do Setup - BLOQUEIA todas as histórias de usuário
- **Histórias de Usuário (Fase 3+)**: Todas dependem da conclusão da fase Fundacional
  - Histórias de usuário podem então prosseguir em paralelo (se houver pessoal)
  - Ou sequencialmente em ordem de prioridade (P1 → P2 → P3)
- **Polimento (Fase Final)**: Depende que todas as histórias de usuário desejadas estejam completas

### Dependências entre Histórias de Usuário

- **História de Usuário 1 (P1)**: Pode começar após Fundacional (Fase 2) - Sem dependências em outras histórias
- **História de Usuário 2 (P2)**: Pode começar após Fundacional (Fase 2) - Pode integrar com HU1 mas deve ser independentemente testável
- **História de Usuário 3 (P3)**: Pode começar após Fundacional (Fase 2) - Pode integrar com HU1/HU2 mas deve ser independentemente testável

### Dentro de Cada História de Usuário

- Testes (se incluídos) DEVEM ser escritos e FALHAR antes da implementação
- Modelos antes de serviços
- Serviços antes de endpoints
- Implementação core antes de integração
- História completa antes de prosseguir para próxima prioridade

### Oportunidades de Paralelismo

- Todas as tarefas de Setup marcadas [P] podem rodar em paralelo
- Todas as tarefas Fundacionais marcadas [P] podem rodar em paralelo (dentro da Fase 2)
- Uma vez que a fase Fundacional completa, todas as histórias de usuário podem começar em paralelo (se a capacidade do time permitir)
- Todos os testes para uma história de usuário marcados [P] podem rodar em paralelo
- Modelos dentro de uma história marcados [P] podem rodar em paralelo
- Diferentes histórias de usuário podem ser trabalhadas em paralelo por diferentes membros do time

---

## Exemplo de Paralelismo: História de Usuário 1

```bash
# Lançar todos os testes para História de Usuário 1 juntos (se testes solicitados):
Tarefa: "Teste de contrato para [endpoint] em tests/contract/test_[nome].py"
Tarefa: "Teste de integração para [jornada do usuário] em tests/integration/test_[nome].py"

# Lançar todos os modelos para História de Usuário 1 juntos:
Tarefa: "Criar [Entidade1] modelo em src/models/[entidade1].py"
Tarefa: "Criar [Entidade2] modelo em src/models/[entidade2].py"
```

---

## Estratégia de Implementação

### MVP Primeiro (Apenas História de Usuário 1)

1. Completar Fase 1: Setup
2. Completar Fase 2: Fundacional (CRÍTICO - bloqueia todas as histórias)
3. Completar Fase 3: História de Usuário 1
4. **PARE e VALIDE**: Teste a História de Usuário 1 independentemente
5. Deploy/demo se estiver pronto

### Entrega Incremental

1. Completar Setup + Fundacional → Fundação pronta
2. Adicionar História de Usuário 1 → Testar independentemente → Deploy/Demo (MVP!)
3. Adicionar História de Usuário 2 → Testar independentemente → Deploy/Demo
4. Adicionar História de Usuário 3 → Testar independentemente → Deploy/Demo
5. Cada história adiciona valor sem quebrar histórias anteriores

### Estratégia de Time Paralelo

Com múltiplos desenvolvedores:

1. Time completa Setup + Fundacional juntos
2. Uma vez Fundacional pronto:
   - Desenvolvedor A: História de Usuário 1
   - Desenvolvedor B: História de Usuário 2
   - Desenvolvedor C: História de Usuário 3
3. Histórias completam e integram independentemente

---

## Notas

- Tarefas [P] = arquivos diferentes, sem dependências
- Label [Story] mapeia tarefa para história de usuário específica para rastreabilidade
- Cada história de usuário deve ser completável e testável independentemente
- Verifique que testes falham antes de implementar
- Commit após cada tarefa ou grupo lógico
- Pare em qualquer checkpoint para validar história independentemente
- Evite: tarefas vagas, conflitos de mesmo arquivo, dependências entre histórias que quebram independência
