# Plano de Implementação: Ajustes de Tema e Idioma

**Branch**: `002-theme-language-adjustments` | **Data**: 2026-04-19 | **Spec**: [spec.md](spec.md)
**Input**: Especificação da funcionalidade de `/specs/002-theme-language-adjustments/spec.md`

**Nota**: Este modelo é preenchido pelo comando `/speckit.plan`. Veja `.specify/templates/plan-template.md` para o fluxo de execução.

## Resumo

Implementar ajustes no sistema de tema e internacionalização do site: (1) Correção dos ícones de tema para mostrar lua no modo claro e sol no modo escuro; (2) Alteração do tema padrão para escuro; (3) Verificação da detecção automática de idioma baseada no navegador/SO; (4) Adição de animação fade-in no carregamento da página inspirada no annamona.co.

## Contexto Técnico

**Linguagem/Versão**: JavaScript ES6+ (vanilla), CSS3, Go 1.21+ (gerador estático)  
**Dependências Principais**: Nenhuma - projeto utiliza apenas APIs nativas do navegador  
**Armazenamento**: localStorage para persistência de preferências  
**Testes**: Testes manuais em navegadores (Chrome, Firefox, Safari)  
**Plataforma Alvo**: Navegadores modernos (Chrome 90+, Firefox 88+, Safari 14+, Edge 90+)  
**Tipo de Projeto**: Site estático gerador em Go  
**Metas de Performance**: Animações devem executar a 60fps, sem jank  
**Restrições**: Respeitar prefers-reduced-motion, funcionar sem JavaScript (degradado)  
**Escala/Alcance**: Site pessoal single-page, ~600 visitantes/mês

## Verificação da Constituição

| Princípio | Status | Justificativa |
|-----------|--------|---------------|
| I. Minimalismo Intencional | ✅ | Sem frameworks adicionais, apenas CSS/JS vanilla |
| II. Bilíngue por Padrão | ✅ | Detecção automática respeita princípio i18n |
| III. Estabilidade Visual | ✅ | Animação respeita prefers-reduced-motion |
| IV. Build Reprodutível | ✅ | Não afeta build, apenas assets estáticos |
| V. Código como Documentação | ✅ | Lógica simples, auto-explicativa |

## Estrutura do Projeto

### Documentação (esta funcionalidade)

```text
specs/002-theme-language-adjustments/
├── plan.md              # Este arquivo (output do comando /speckit.plan)
├── research.md          # Decisões técnicas investigadas
├── data-model.md        # N/A - sem mudanças de dados
├── quickstart.md        # Teste manual dos cenários
├── contracts/           # N/A - sem APIs externas
└── tasks.md             # Será criado pelo /speckit.tasks
```

### Código Fonte (raiz do repositório)

```text
assets/
├── css/
│   └── main.css          # Adicionar animações fade-in, classes de ícones
├── js/
│   ├── theme.js          # Alterar DEFAULT_THEME para 'dark', ícones sol/lua
│   ├── i18n.js           # Verificar detecção automática (já implementada)
│   └── fadein.js         # NOVO - módulo de animação fade-in
└── templates/
    └── partials/
        └── header.html   # Adicionar ícones sol e lua (ambos no DOM)
```

**Decisão de Estrutura**: Modificação em arquivos existentes sem alterar a arquitetura. Novo módulo JS dedicado para fade-in.

## Rastreamento de Complexidade

Nenhuma violação de princípios da constituição identificada.

---

## Decisões Técnicas (research.md)

### DT-001: Lógica dos Ícones de Tema

**Problema**: O botão de tema atual mostra apenas um ícone (sol), mas precisa mostrar lua no modo claro e sol no modo escuro.

**Opções Consideradas**:
- **A**: Substituir SVG via JavaScript (innerHTML) - Rejeitado: causaria flash e perda de estado
- **B**: CSS `display: none` para alternar entre dois SVGs no DOM - Escolhido: mais performático, sem FOUC
- **C**: SVG symbols + `<use>` com troca de referência - Rejeitado: complexidade desnecessária

**Decisão**: Implementar Opção B - dois ícones no HTML, controle de visibilidade via CSS/JS.

**Implementação**:
- HTML: Ambos ícones (sol e lua) presentes no botão
- CSS: Classes `.theme-icon-sun`, `.theme-icon-moon` controlam visibilidade
- JS: theme.js alterna classes baseado no tema atual
- Tema claro: mostra lua (indica que pode mudar para escuro)
- Tema escuro: mostra sol (indica que pode mudar para claro)

### DT-002: Tema Padrão Escuro

**Decisão**: Alterar `DEFAULT_THEME` de `'light'` para `'dark'` no arquivo theme.js.

**Impacto**: Novos visitantes verão tema escuro primeiro. Usuários com preferência salva não serão afetados.

**Validação**: Testar em navegação privada para confirmar comportamento.

### DT-003: Detecção Automática de Idioma

**Análise**: O código atual em `i18n.js` já implementa detecção via `navigator.language`:

```javascript
_detectBrowserLanguage: function() {
  const browserLang = navigator.language || navigator.userLanguage || 'pt';
  if (browserLang.toLowerCase().startsWith('pt')) {
    return 'pt';
  }
  return 'en';
}
```

**Decisão**: Verificar se funcionalidade está operacional. Possíveis ajustes:
- Expandir para mais variações de português (pt-PT, pt-AO, etc.)
- Garantir que fallback seja respeitado

**Status**: ✅ Já implementado, apenas validar.

### DT-004: Animação Fade-In (estilo annamona.co)

**Investigação no annamona.co**:
- Classe CSS: `fade-in`
- Keyframes: `fadeIn` com opacity 0 → 1
- Duração: 0.8s
- Timing: ease-in
- Respeita prefers-reduced-motion

**Implementação Proposta**:
```css
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.fade-in {
  opacity: 0;
  animation: fadeIn 0.8s ease-in forwards;
}

@media (prefers-reduced-motion: reduce) {
  .fade-in {
    opacity: 1;
    animation: none;
  }
}
```

**Aplicação**: Adicionar classe `fade-in` aos elementos principais (header, sections) via JavaScript no DOMContentLoaded.

### DT-005: Ordem de Carregamento

**Consideração**: Para evitar flash de conteúdo sem estilo:
1. CSS carrega primeiro (inline ou blocking)
2. Animação aplica-se após DOM pronto
3. Preferência de tema aplica-se imediatamente (antes de render)

**Solução**: Usar `document.documentElement.style.visibility = 'hidden'` no head, revelar após tema aplicado.

### DT-006: Correção de Segurança - XSS em safeHTML

**Problema**: A função `safeHTML` em `main.go` converte strings diretamente para `template.HTML` sem sanitização.

**Solução**: Remover a função `safeHTML` completamente. Não é necessária para este projeto pois todo conteúdo vem de config.yaml confiável, mas removê-la elimina o risco de uso indevido futuro.

### DT-007: Correção de Segurança - Permissões de Workflow

**Problema**: Permissão `contents: write` é excessiva para workflows de deploy.

**Solução**: Restringir permissões:
- `deploy.yml`: `contents: read`, `pages: write`, `id-token: write`
- `release.yml`: `contents: write` (necessário para criar tags), `actions: read`

### DT-008: Correção de Qualidade - Polling de Deploy

**Problema**: Sleep fixo de 30s no workflow de release é ineficiente.

**Solução**: Implementar polling com backoff:
```yaml
for i in {1..10}; do
  STATUS=$(curl -s ...)
  [ "$STATUS" == "built" ] && exit 0
  sleep 15
done
```

### DT-009: Otimização CSS - will-change

**Decisão**: Adicionar `will-change: opacity` às animações fade-in para melhor performance.

---

## Notas de Implementação

### Prioridade de Execução

1. **Tema Padrão Escuro** (maior impacto visual) - Alterar theme.js
2. **Ícones de Tema** (correção UX) - Modificar header.html + theme.js + CSS
3. **Fade-In** (polimento) - Criar fadein.js + CSS
4. **Verificação Idioma** (validação) - Revisar i18n.js

### Testes Manuais Requeridos

- [ ] Navegação privada: tema inicia escuro
- [ ] Modo claro: mostra ícone lua
- [ ] Modo escuro: mostra ícone sol
- [ ] Toggle: alterna ícones corretamente
- [ ] Navegador em EN: site aparece em inglês
- [ ] Navegador em PT: site aparece em português
- [ ] Animação fade-in executa suavemente
- [ ] prefers-reduced-motion desabilita animação
- [ ] Sem JavaScript: site funciona (degradado)
