# Tarefas: Ajustes de Tema e Idioma

**Input**: Documentos de design de `/specs/002-theme-language-adjustments/`
**Pré-requisitos**: plan.md, spec.md

**Testes**: Testes manuais em navegador (não automatizados)

**Organização**: Tarefas agrupadas por história de usuário para permitir implementação e teste independentes.

## Formato: `[ID] [P?] [Story] Descrição`

- **[P]**: Pode rodar em paralelo
- **[Story]**: História de usuário (US1, US2, US3, US4)

---

## Fase 1: Setup (Infraestrutura Compartilhada)

**Propósito**: Preparação inicial e criação da branch

- [ ] T001 Criar branch `002-theme-language-adjustments` a partir da main
- [ ] T002 Verificar estado limpo do working tree antes de iniciar

**Checkpoint**: Branch criada e pronta para desenvolvimento

---

## Fase 2: História de Usuário 1 - Correção dos Ícones de Tema (Prioridade: P1) 🎯 MVP

**Objetivo**: Implementar dois ícones (sol e lua) no botão de tema, mostrando lua no modo claro e sol no modo escuro

**Teste Independente**: Abrir o site e verificar que o ícone muda ao alternar tema

### Implementação da US1

- [ ] T003 [US1] Adicionar ícone de lua (moon.svg) no header.html junto ao ícone de sol existente
- [ ] T004 [US1] Adicionar classes CSS para controle de visibilidade dos ícones em assets/css/main.css
- [ ] T005 [US1] Atualizar theme.js para alternar visibilidade dos ícones baseado no tema atual
- [ ] T006 [US1] Testar manualmente: modo claro mostra lua, modo escuro mostra sol

**Checkpoint**: Ícones alternam corretamente ao mudar tema

---

## Fase 3: História de Usuário 2 - Tema Padrão Escuro (Prioridade: P1) 🎯 MVP

**Objetivo**: Alterar o tema padrão de 'light' para 'dark' para novos visitantes

**Teste Independente**: Abrir site em navegação privada e verificar tema escuro

### Implementação da US2

- [ ] T007 [US2] Alterar `DEFAULT_THEME: 'light'` para `'dark'` em assets/js/theme.js
- [ ] T008 [US2] Verificar que preferências salvas continuam sendo respeitadas
- [ ] T009 [US2] Testar manualmente em navegação privada: tema inicia escuro

**Checkpoint**: Novos visitantes veem tema escuro; retornantes mantêm preferência

---

## Fase 4: História de Usuário 3 - Detecção Automática de Idioma (Prioridade: P2)

**Objetivo**: Validar que a detecção automática de idioma via navigator.language está funcionando

**Teste Independente**: Mudar idioma do navegador e verificar detecção em navegação privada

### Implementação da US3

- [ ] T010 [US3] Revisar função `_detectBrowserLanguage()` em assets/js/i18n.js
- [ ] T011 [US3] Verificar cobertura de variações de português (pt-BR, pt-PT, pt-AO)
- [ ] T012 [US3] Testar manualmente: navegador em EN → site em inglês
- [ ] T013 [US3] Testar manualmente: navegador em PT → site em português

**Checkpoint**: Idioma detectado automaticamente baseado no navegador

---

## Fase 5: História de Usuário 4 - Animação Fade-In (Prioridade: P2)

**Objetivo**: Adicionar animação fade-in suave no carregamento, estilo annamona.co

**Teste Independente**: Recarregar página e observar fade-in suave do conteúdo

### Implementação da US4

- [ ] T014 [P] [US4] Criar arquivo assets/js/fadein.js com módulo de animação
- [ ] T015 [P] [US4] Adicionar keyframes CSS fadeIn em assets/css/main.css
- [ ] T016 [US4] Adicionar classe .fade-in com duração 0.8s ease-in
- [ ] T017 [US4] Adicionar media query prefers-reduced-motion para desabilitar animação
- [ ] T018 [US4] Adicionar `will-change: opacity` para otimização de performance
- [ ] T019 [US4] Integrar módulo fadein.js em assets/js/app.js (inicialização)
- [ ] T020 [US4] Adicionar classe fade-in aos elementos principais (header, sections)
- [ ] T021 [US4] Testar manualmente: animação executa suavemente
- [ ] T022 [US4] Testar manualmente: prefers-reduced-motion desabilita animação

**Checkpoint**: Conteúdo aparece com fade-in suave ao carregar

---

## Fase 6: História de Usuário 5 - Correções de Segurança (Prioridade: P1) 🎯 CRÍTICO

**Objetivo**: Corrigir vulnerabilidades de segurança e problemas de qualidade identificados na revisão

**Teste Independente**: Verificar que correções foram aplicadas e workflows funcionam corretamente

### Implementação da US5

- [ ] T023 [P] [US5] Remover função `safeHTML` do cmd/generator/main.go (prevenir XSS)
- [ ] T024 [P] [US5] Restringir permissões em .github/workflows/deploy.yml (princípio do menor privilégio)
- [ ] T025 [P] [US5] Restringir permissões em .github/workflows/release.yml (princípio do menor privilégio)
- [ ] T026 [US5] Implementar polling com backoff em vez de sleep fixo no workflow release.yml
- [ ] T027 [US5] Verificar funcionamento do workflow de deploy após alterações
- [ ] T028 [US5] Testar manualmente: workflow de release funciona corretamente

**Checkpoint**: Correções de segurança aplicadas e validadas

---

## Fase 7: Polimento & Preocupações Transversais

**Propósito**: Ajustes finais e validação completa

- [ ] T029 [P] Validar todos os critérios de sucesso do spec.md
- [ ] T030 [P] Verificar acessibilidade (aria-labels, prefers-reduced-motion)
- [ ] T031 Verificar funcionamento sem JavaScript (degradado)
- [ ] T032 Atualizar quickstart.md com cenários de teste manuais
- [ ] T033 Atualizar README.md com estrutura de projeto atualizada
- [ ] T034 Executar build do site e verificar sem erros
- [ ] T035 Criar commit com mensagem seguindo Conventional Commits

**Checkpoint**: Funcionalidade completa e pronta para PR

---

## Dependências & Ordem de Execução

### Dependências de Fase

- **Fase 1 (Setup)**: Sem dependências
- **Fase 2 (US1 - Ícones)**: Depende do Setup
- **Fase 3 (US2 - Tema Padrão)**: Depende do Setup (pode rodar em paralelo com US1)
- **Fase 4 (US3 - Idioma)**: Depende do Setup (pode rodar em paralelo)
- **Fase 5 (US4 - Fade-In)**: Depende do Setup (pode rodar em paralelo)
- **Fase 6 (US5 - Segurança)**: Depende do Setup (pode rodar em paralelo, mas alta prioridade)
- **Fase 7 (Polimento)**: Depende de todas as histórias completas

### Oportunidades de Paralelismo

As histórias de usuário podem ser agrupadas para execução:
- **Grupo A (UX/UI)**: US1, US2, US4 - Alterações visuais
- **Grupo B (Lógica)**: US3 - Validação de i18n
- **Grupo C (Segurança)**: US5 - Correções críticas (prioridade máxima)

Tarefas marcadas com [P] dentro de cada história podem rodar em paralelo.

---

## Estratégia de Implementação

### Ordem de Prioridade (Segurança Primeiro)

1. **US5 (Segurança) - PRIORIDADE MÁXIMA**: Correções de vulnerabilidades
2. **US2 (Tema Padrão)**: Afeta primeira impressão dos visitantes
3. **US1 (Ícones)**: Complementa US2
4. **US4 (Fade-In)**: Polimento visual
5. **US3 (Idioma)**: Validação (já implementado)
6. **Polimento**: Documentação e finalização

### Fluxo Completo

1. Completar Fase 1: Setup
2. Completar Fase 2: US1 (Ícones)
3. Completar Fase 3: US2 (Tema Padrão)
4. Completar Fase 4: US3 (Idioma)
5. Completar Fase 5: US4 (Fade-In)
6. Completar Fase 6: US5 (Segurança) - PODE SER FEITO EM PARALELO
7. Completar Fase 7: Polimento
8. Criar PR para revisão

---

## Notas

- Todas as tarefas são de modificação em arquivos existentes ou criação de novos módulos
- Testes são manuais no navegador
- Foco em acessibilidade: prefers-reduced-motion, aria-labels
- Manter compatibilidade com navegadores modernos
- Build deve passar sem erros antes do commit
