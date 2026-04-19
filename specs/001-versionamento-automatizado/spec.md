# Especificação da Funcionalidade: Versionamento Automatizado com Tags e Releases

**Branch da Funcionalidade**: `001-versionamento-automatizado`  
**Criada**: 2026-04-19  
**Status**: ✅ Concluído (Implementado)  
**Input**: Sistema de versionamento automatizado com geração de tags no GitHub e releases automáticas após deploy. Incluir criação de tags retroativas para versões históricas (v1.x, v2.x) no GitHub. O workflow deve detectar quando uma entrega é feita na main, criar uma nova tag semver, e após deploy bem-sucedido gerar release no GitHub com changelog automático.

## Esclarecimentos

### Sessão 2026-04-19

- **Q**: Como o sistema deve detectar e tratar breaking changes (major version)? → **A**: Usar `BREAKING CHANGE:` no corpo do commit ou `!` no tipo (ex: `feat!:`, `fix!:). Segue padrão Conventional Commits oficial.
- **Q**: Qual formato o changelog deve seguir? → **A**: Formato Keep a Changelog (Adicionado/Alterado/Obsoleto/Removido/Corrigido/Segurança) em português.
- **Q**: Qual evento deve disparar o workflow de versionamento? → **A**: Push na branch main (inclui merge via PR e push direto). Adicionar proteção de branch se necessário.
- **Q**: Como as tags retroativas devem ser criadas? → **A**: Workflow manual (workflow_dispatch) que aceita commit SHA e nome da tag como inputs. Mais transparente e reutilizável.
- **Q**: O que deve acontecer com a tag se o deploy falhar? → **A**: Tag permanece (não é deletada), mas release não é criada. Deploy pode ser reattemptado (idempotente).

## Cenários de Usuário & Testes *(obrigatório)*

### História de Usuário 1 - Deploy Automático com Tag (Prioridade: P1)

Como mantenedor do site, quero que cada merge na branch main gere automaticamente uma nova tag de versão seguindo semver, para que eu tenha um histórico versionado de todas as entregas sem esforço manual.

**Por que esta prioridade**: Elimina erro humano no versionamento, garante rastreabilidade completa de cada deploy e alinha com o princípio "Build Reprodutível" da constituição.

**Teste Independente**: Fazer merge de um PR na main e verificar que uma nova tag semver é criada automaticamente no GitHub (ex: v3.1.0 → v3.1.1 ou v3.2.0 dependendo do tipo de mudança).

**Cenários de Aceitação**:

1. **Dado que** um PR é mergeado na main com commit tipo `feat:`, **Quando** o workflow executa, **Então** uma tag minor é criada (ex: v3.1.0 → v3.2.0)
2. **Dado que** um PR é mergeado na main com commit tipo `fix:`, **Quando** o workflow executa, **Então** uma tag patch é criada (ex: v3.1.0 → v3.1.1)
3. **Dado que** não há mudanças desde a última tag, **Quando** o workflow executa, **Então** nenhuma nova tag é criada

---

### História de Usuário 2 - Release Automático no GitHub (Prioridade: P1)

Como mantenedor do site, quero que após um deploy bem-sucedido no GitHub Pages, uma release seja criada automaticamente no GitHub com changelog, para que usuários e stakeholders possam acompanhar o que foi entregue em cada versão.

**Por que esta prioridade**: Comunica mudanças de forma transparente, documenta o histórico do projeto e facilita rollback se necessário.

**Teste Independente**: Verificar que após deploy bem-sucedido, existe uma release no GitHub correspondente à tag criada, contendo changelog gerado automaticamente dos commits.

**Cenários de Aceitação**:

1. **Dado que** uma tag foi criada e o deploy foi bem-sucedido, **Quando** o workflow de release executa, **Então** uma release é publicada no GitHub
2. **Dado que** uma release é criada, **Quando** visualizo a página da release, **Então** vejo um changelog com os commits incluídos desde a última versão
3. **Dado que** o deploy falha, **Quando** o workflow detecta a falha, **Então** nenhuma release é criada (apenas tag)

---

### História de Usuário 3 - Tags Retroativas para Versões Históricas (Prioridade: P2)

Como mantenedor do site, quero criar tags retroativas para as eras v1.x (Jekyll inicial) e v2.x (Jekyll estável) no GitHub, para que o histórico completo do projeto seja preservado e documentado.

**Por que esta prioridade**: Completa o histórico versionado do projeto, respeitando os 11+ anos de evolução documentada na constituição.

**Teste Independente**: Verificar no GitHub que existem tags v1.0.0 (apontando para commit de 2015) e v2.0.0 (apontando para commit pré-migração Go).

**Cenários de Aceitação**:

1. **Dado que** identifico o commit inicial de 2015, **Quando** crio a tag v1.0.0, **Então** ela aparece no GitHub como uma release antiga
2. **Dado que** identifico o commit de estabilização em 2019, **Quando** crio a tag v2.0.0, **Então** ela aparece no GitHub como release intermediária
3. **Dado que** as tags retroativas existem, **Quando** visualizo o gráfico de releases, **Então** vejo a linha do tempo completa do projeto

---

## Requisitos *(obrigatório)*

### Requisitos Funcionais

- **RF-001**: O sistema DEVE detectar push na branch main automaticamente
- **RF-002**: O sistema DEVE analisar commits convencionais (conventional commits) para determinar o tipo de versão (major/minor/patch), incluindo detecção de breaking changes via `BREAKING CHANGE:` footer ou `!` no tipo
- **RF-003**: O sistema DEVE criar tags semver no formato v{MAJOR}.{MINOR}.{PATCH}
- **RF-004**: O sistema DEVE aguardar confirmação de deploy bem-sucedido antes de criar release
- **RF-005**: O sistema DEVE gerar changelog automaticamente a partir dos commits desde a última tag
- **RF-006**: O sistema DEVE criar releases no GitHub com título, descrição (changelog) e tag associada
- **RF-007**: O sistema DEVE permitir criação manual de tags retroativas via workflow_dispatch, aceitando commit SHA e nome da tag como inputs
- **RF-008**: O sistema NÃO DEVE criar releases para deploys falhos
- **RF-009**: O sistema DEVE gerar changelog no formato Keep a Changelog, traduzido para português (Adicionado/Alterado/Obsoleto/Removido/Corrigido/Segurança)

### Requisitos Não-Funcionais

- **RNF-001**: O workflow DEVE executar em menos de 2 minutos
- **RNF-002**: O changelog DEVE ser gerado em português (conforme constituição)
- **RNF-003**: O sistema DEVE ser compatível com GitHub Actions (já em uso no projeto)

## Critérios de Sucesso *(obrigatório)*

### Resultados Mensuráveis

- **CS-001**: 100% dos merges na main geram tags automaticamente dentro de 2 minutos
- **CS-002**: 100% dos deploys bem-sucedidos geram releases com changelog em português
- **CS-003**: Tags retroativas v1.0.0 e v2.0.0 existem no GitHub documentando o histórico
- **CS-004**: Zero releases criados para deploys falhos (nenhum falso positivo)

## Casos de Borda

- **CB-001**: O que acontece se dois merges ocorrem simultaneamente? → O workflow mais recente deve aguardar o anterior ou incrementar versão sequencialmente
- **CB-002**: Como lidar com commits que não seguem conventional commits? → Versão patch como padrão conservador
- **CB-003**: O que fazer se a criação da tag falha? → Workflow deve falhar e notificar, sem criar release parcial
- **CB-004**: O que acontece se deploy falha após tag criada? → Tag permanece, release não é criada, deploy pode ser reattemptado (idempotente)

## Suposições

- O projeto já usa GitHub Actions para deploy (confirmado na estrutura atual)
- O histórico de commits segue conventional commits (padrão adotado no projeto)
- O token GitHub tem permissões suficientes para criar tags e releases
- O commit de 2015 ainda existe no histórico (b81034a - Initial version)
- O commit de migração Go em 2026 é o marco para v3.0.0 (6aeecb4)

---

**Versão**: 1.0.0 | **Criada**: 2026-04-19
