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

**Funcionalidade atual**: ✅ CONCLUÍDO - Ajustes de Tema e Idioma  
**Diretório**: `specs/002-theme-language-adjustments/`  
**Plano**: [plan.md](specs/002-theme-language-adjustments/plan.md)  
**Status**: **100% Implementado e Entregue**

### Artefatos Disponíveis
- [Especificação](specs/002-theme-language-adjustments/spec.md)
- [Plano de Implementação](specs/002-theme-language-adjustments/plan.md)
- [Tarefas](specs/002-theme-language-adjustments/tasks.md)

### Histórias de Usuário Implementadas
1. **US1 (P1)**: ✅ Ícones de tema - lua no modo claro, sol no modo escuro
2. **US2 (P1)**: ✅ Tema padrão escuro para novos visitantes
3. **US3 (P2)**: ✅ Validação da detecção automática de idioma via navegador
4. **US4 (P2)**: ✅ Animação fade-in no carregamento em todas as seções
5. **US5 (P1)**: ✅ Correções de segurança (XSS, permissões, polling)

### Implementações Realizadas
- CSS atualizado com controle de visibilidade para ícones sol/lua
- `DEFAULT_THEME` alterado para `'dark'`
- `FadeIn` adicionado aos módulos em app.js
- Efeito fade-in aplicado em todas as seções (header, about, experience, skills, education, contact, footer)
- `will-change: opacity` adicionado para performance
- Função `safeHTML` removida do gerador Go (prevenção XSS)
- Permissões de workflows restritas (menor privilégio)
- Polling com backoff implementado no workflow de release
- README.md atualizado com estrutura do projeto
- `.gitattributes` criado para destacar Go nas estatísticas

### Status da PR
- **PR #108**: [feat: ajustes de tema, idioma e correções de segurança](https://github.com/vagnerbarbosa/vagnerbarbosa.github.io/pull/108)
- Branch: `feat/theme-language-security-adjustments`
- Status: ✅ **Mergeado e Deployado**

### Próximos Passos
- Nenhum - funcionalidade completa e em produção
<!-- SPECKIT END -->
