# Regras de Colaboração - Vagner Barbosa

## CRÍTICO: NUNCA fazer push direto na main

**QUANDO APLICAR:** SEMPRE. Sem exceções.

**AÇÃO OBRIGATÓRIA:**
1. **NUNCA** execute `git push origin main`
2. **NUNCA** execute `git push` estando na branch main
3. **SEMPRE** crie uma branch de feature: `git checkout -b feat/nome-da-feature`
4. **SEMPRE** faça push da branch de feature: `git push origin feat/nome-da-feature`
5. **SEMPRE** crie uma Pull Request via GitHub CLI: `gh pr create` ou pelo site
6. **SEMPRE** aguarde revisão e merge via PR

**PROTEÇÃO CONFIGURADA:**
- Hook pre-push bloqueia push na main localmente
- Branch protection deve ser configurada no GitHub (Settings > Branches > main)

**VIOLAÇÃO ANTERIOR:**
Commit 170ab22 foi feito direto na main em 19/04/2026. Foi revertido e refeito via PR #82.

**POR QUE:**
- Mantém histórico limpo e rastreável
- Permite revisão de código
- Previne erros em produção
- Alinha com princípio "Build Reprodutível" da constituição

---

## CRÍTICO: Sempre verificar status da PR antes de alterações

**Quando aplicar:** Antes de fazer qualquer commit, push ou alteração em branch existente.

**Ação obrigatória:**
1. Execute: `gh pr view <numero> --json state,mergedAt`
2. Se a PR estiver MERGED → PARE e crie uma nova branch a partir da main
3. Nunca continue trabalhando em branch de PR já mergeada

**Por que:** Evitar commits "perdidos" em branches que já foram integradas.

---

## Priorizar MCP quando custo for menor

**Quando aplicar:** Ao escolher entre usar MCP, CLI ou API direta.

**Regra:** Se o MCP (Model Context Protocol) oferecer a mesma funcionalidade com menor custo/complexidade, use-o em vez de CLI ou API direta.

---

## Verificar README.md após modificações

**Quando aplicar:** Após fazer alterações significativas no projeto.

**Ação obrigatória:**
1. Leia o README.md atual
2. Verifique se as alterações exigem atualização na documentação
3. Atualize o README.md se necessário

---

## Regras de Memória

- Sempre consulte o MEMORY.md no início da sessão
- Respeite todas as regras de feedback registradas
- Se uma regra for violada, pare e corrija imediatamente
- **Spec Kit**: Toda a documentação gerada via Spec Kit (specs, plans, tasks) deve ser escrita obrigatoriamente em Português.

---

<!-- SPECKIT START -->
## Histórico de Funcionalidades

### ✅ Concluídas

| Funcionalidade | PR | Status |
|---------------|-----|--------|
| Versionamento Automatizado | [#107](https://github.com/vagnerbarbosa/vagnerbarbosa.github.io/pull/107) | ✅ Mergeada |
| Ajustes de Tema, Idioma e Segurança | [#108](https://github.com/vagnerbarbosa/vagnerbarbosa.github.io/pull/108) | ✅ Mergeada |
| Fade-in Completo + SECURITY.md | [#110](https://github.com/vagnerbarbosa/vagnerbarbosa.github.io/pull/110) | ✅ Mergeada |
| LinkedIn Import CLI | [#124](https://github.com/vagnerbarbosa/vagnerbarbosa.github.io/pull/124) | ✅ Mergeada |
| Extração Inteligente de Tech Stack | [#126](https://github.com/vagnerbarbosa/vagnerbarbosa.github.io/pull/126) | 🔄 Em Revisão |

### 📁 Especificações Disponíveis
- [001-versionamento-automatizado](specs/001-versionamento-automatizado/) - ✅ Concluído
- [002-theme-language-adjustments](specs/002-theme-language-adjustments/) - ✅ Concluído
- [003-linkedin-import](specs/003-linkedin-import/) - ✅ Concluído
- [004-tech-stack-extraction](specs/004-tech-stack-extraction/) - 🔄 Em Desenvolvimento

### 🎯 Próxima Funcionalidade
- [005-test-coverage-total](specs/005-test-coverage-total/) - 🔄 Em Planejamento


---

## Melhorias Futuras

### Cobertura de Testes - Melhorias Realizadas

**Status atual (16/05/2026):**

| Pacote | Cobertura | Status |
|--------|-----------|--------|
| `cmd/generator` | 97.3% | ✅ Excelente |
| `cmd/import-linkedin/commands` | 79.2% | ✅ Bom |
| `cmd/import-linkedin/internal/comparator` | 98.3% | ✅ Excelente |
| `cmd/import-linkedin/internal/parser` | 96.3% | ✅ Excelente |
| `cmd/import-linkedin/internal/transformer` | 100% | ✅ Completo |
| `internal/config` | 100% | ✅ Completo |
| **Média Total** | **95.2%** | ✅ Excelente |

#### Melhorias Implementadas

1. **Bug Crítico Corrigido:** Loop infinito em `parser/certification.go` (linha 127) - trocado `continue` por `break` em caso de erro de leitura

2. **Novos Testes Adicionados:**
   - `transformer/description_test.go` - Cobertura 100%
   - `parser/education_test.go` - Testes completos para NewEducationParser, Close
   - `parser/experience_test.go` - Testes completos para NewExperienceParser, Close
   - `commands/import_test.go` - Testes para dry-run, backup, arquivos válidos
   - `commands/validate_test.go` - Testes para validação de education e certifications

#### Caminhos Não Cobertos (Restrições Técnicas)

- **`main()` em generator:** Chama `os.Exit()` - requer refatoração complexa
- **Caminos de erro de parsing em CSV:** Requer simulação de erros de I/O
- **Interação do usuário (UI):** Confirmação interativa - difícil de automatizar

**Conclusão:** 95.2% é cobertura excepcional para projeto Go. Os caminhos restantes são de erro extremos ou requerem mocks complexos.

---
<!-- SPECKIT END -->
