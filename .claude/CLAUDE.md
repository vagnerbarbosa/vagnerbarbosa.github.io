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

---

<!-- SPECKIT START -->
## Plano Ativo do Speckit

**Funcionalidade atual**: ✅ CONCLUÍDO - Versionamento Automatizado com Tags e Releases  
**Diretório**: `specs/001-versionamento-automatizado/`  
**Plano**: [plan.md](specs/001-versionamento-automatizado/plan.md)  
**Status**: **100% Completo** - Todas as 34 tarefas finalizadas e testadas

### Artefatos Disponíveis
- [Especificação](specs/001-versionamento-automatizado/spec.md)
- [Pesquisa](specs/001-versionamento-automatizado/research.md)
- [Modelo de Dados](specs/001-versionamento-automatizado/data-model.md)
- [Plano de Implementação](specs/001-versionamento-automatizado/plan.md)
- [Tarefas](specs/001-versionamento-automatizado/tasks.md)

### Resultado da Implementação
- **10 tags criadas**: v1.0.0, v2.0.0 (retroativas), v2.1.0-v2.2.2, v3.0.0, v4.0.0-v4.2.0
- **6 releases automáticas** publicadas no GitHub
- **Versionamento**: `feat:` → minor, `fix:` → patch, `feat!:` → major
- **Changelog**: Em português (Adicionado/Alterado/Corrigido/Quebra de Compatibilidade)
- **Concorrência**: Workflows executam sequencialmente sem conflitos
<!-- SPECKIT END -->
