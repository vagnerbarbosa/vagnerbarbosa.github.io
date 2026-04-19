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
## Histórico de Funcionalidades

### ✅ Concluídas

| Funcionalidade | PR | Status |
|---------------|-----|--------|
| Versionamento Automatizado | [#107](https://github.com/vagnerbarbosa/vagnerbarbosa.github.io/pull/107) | ✅ Mergeada |
| Ajustes de Tema, Idioma e Segurança | [#108](https://github.com/vagnerbarbosa/vagnerbarbosa.github.io/pull/108) | ✅ Mergeada |
| Fade-in Completo + SECURITY.md | [#110](https://github.com/vagnerbarbosa/vagnerbarbosa.github.io/pull/110) | ✅ Mergeada |

### 📁 Especificações Disponíveis
- [001-versionamento-automatizado](specs/001-versionamento-automatizado/) - ✅ Concluído
- [002-theme-language-adjustments](specs/002-theme-language-adjustments/) - ✅ Concluído

### 🎯 Próxima Funcionalidade
Nenhuma funcionalidade ativa. Use `/speckit-specify` para iniciar uma nova.

---

## Melhorias Futuras

### Cobertura de Testes - 100%

**Status atual:** 93.0% (PR #121)  
**Gap para 100%:** ~7% (principalmente em `cmd/generator/main.go`)

#### Bloqueadores Identificados

1. **`main()` - 0% cobertura**
   - Problema: Chama `os.Exit(1)` que encerra o processo de teste
   - Solução potencial: Refatorar para retornar código de erro em vez de chamar `os.Exit()`
   ```go
   // Atual
   func main() {
       if err := run(); err != nil {
           fmt.Fprintf(os.Stderr, "Error: %v\n", err)
           os.Exit(1)
       }
   }
   
   // Refatorado para testabilidade
   func main() {
       os.Exit(runWithExitCode())
   }
   
   func runWithExitCode() int {
       if err := run(); err != nil {
           fmt.Fprintf(os.Stderr, "Error: %v\n", err)
           return 1
       }
       return 0
   }
   ```

2. **Casos de erro em `copyAndMinifyDir()` e `copyAndMinifyFile()`**
   - Problema: Alguns caminhos de erro requerem condições de corrida ou permissões de arquivo específicas
   - Solução potencial: Usar interface `fs.FS` para permitir mocking do filesystem

3. **Testes com permissões de arquivo**
   - Problema: Windows e Unix tratam permissões diferentemente
   - Solução potencial: Usar build tags ou abstração de filesystem

#### Estratégia para Alcançar 100%

1. **Refatoração leve:** Extrair lógica de `main()` para função testável
2. **Filesystem abstraction:** Interface para operações de arquivo permitindo mocks
3. **Testes de integração:** Testar o binário gerado em vez de apenas pacotes

**Prioridade:** Baixa - 93% é excelente cobertura para um projeto Go

---
<!-- SPECKIT END -->
