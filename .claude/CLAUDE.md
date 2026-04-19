# Regras de Colaboração - Vagner Barbosa

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

**Funcionalidade atual**: Versionamento Automatizado com Tags e Releases  
**Diretório**: `specs/001-versionamento-automatizado/`  
**Plano**: [plan.md](specs/001-versionamento-automatizado/plan.md)  
**Status**: Fase 1 completa, pronto para tasks

### Artefatos Disponíveis
- [Especificação](specs/001-versionamento-automatizado/spec.md)
- [Pesquisa](specs/001-versionamento-automatizado/research.md)
- [Modelo de Dados](specs/001-versionamento-automatizado/data-model.md)
- [Plano de Implementação](specs/001-versionamento-automatizado/plan.md)
<!-- SPECKIT END -->
