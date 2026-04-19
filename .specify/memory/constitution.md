
<!--
RELATÓRIO DE IMPACTO DA SINCRONIZAÇÃO
Versão: Nova (3.0.0)
Alterações: Criação inicial da constituição para a era pós-migração
- Adicionado: 5 princípios centrais refletindo lições aprendidas
- Adicionado: Seção de contexto histórico documentando jornada de 11 anos
- Adicionado: Seção de governança
Modelos a verificar: plan-template.md, spec-template.md, tasks-template.md
-->

# Constituição do vagnerbarbosa.github.io

## História do Projeto

### Capítulo I: O Início (2015-2026)

Este site nasceu em fevereiro de 2015 como um portfólio pessoal baseado em Jekyll. Por mais de uma década, serviu como uma representação estática da identidade profissional, evoluindo através de múltiplas iterações de design enquanto permanecia ancorado no ecossistema Ruby/Jekyll.

### Capítulo II: A Migração (Março de 2026)

A decisão de adotar a filosofia de design limpo e minimalista do [annamona.co](https://annamona.co) tornou-se o catalisador para uma migração completa de plataforma. Ao invés de adaptar a nova visão ao Jekyll, o site foi reconstruído a partir de primeiros princípios usando Go, alcançando:

- **Zero dependências de runtime** - apenas a toolchain de build requer Go
- **Builds em menos de um segundo** - de ~3s com Jekyll para ~0.5s com Go nativo
- **Preservação completa do conteúdo** - todo o histórico e conteúdo bilíngue mantidos
- **Ferramentas modernas** - minificação integrada, templating adequado (html/template) e CI/CD via GitHub Actions

A migração também serviu como campo de treinamento prático para desenvolvimento em Go, reforçando o princípio de que código em produção é o melhor ambiente de aprendizado.

### Capítulo III: Era Atual (v3.0.0+)

A iteração atual representa uma síntese de mais de 11 anos de aprendizados, destilados em cinco princípios não negociáveis.

---

## Princípios Centrais

### I. Minimalismo Intencional

Cada elemento deve merecer seu lugar. Sem frameworks, sem dependências de runtime, sem complexidade acidental. O design segue a filosofia do annamona.co: tipografia limpa, espaçamento generoso, conteúdo em primeiro lugar. O gerador em Go produz HTML/CSS/JS estáticos — nada mais é enviado para produção.

**Justificativa**: Performance, segurança e durabilidade. Sites estáticos sobrevivem à rotatividade de tecnologias. Este princípio é um resultado direto da observação da volatilidade de frontends de 2015-2026.

### II. Bilíngue por Padrão

Todo conteúdo existe em Português (PT) e Inglês (EN). i18n é infraestrutura primária, não um adicional. A troca de idioma é instantânea e preserva o estado. O sistema detecta preferências do navegador mas respeita escolhas explícitas do usuário.

**Justificativa**: Alcance global com acessibilidade local. O site serve tanto audiências brasileiras locais quanto recrutadores/conexões internacionais.

### III. Estabilidade Visual

O design inspirado no annamona.co é sacrossanto. Temas claro/escuro respeitam preferências do sistema com override manual. Transições respeitam `prefers-reduced-motion`. Tipografia e cores seguem uma escala rígida definida em variáveis CSS. Links de skip, contraste adequado e HTML semântico são obrigatórios.

**Justificativa**: A migração para a estética do annamona.co foi um esforço significativo; a consistência preserva esse investimento. Acessibilidade é não negociável.

### IV. Build Reprodutível

O gerador em Go produz output bit-idêntico dado input idêntico. Dependências são travadas (go.mod). GitHub Actions automatiza build e deploy para GitHub Pages. Sem deploy manual, sem "funciona na minha máquina".

**Justificativa**: O período 2015-2019 sofreu com builds locais inconsistentes. Reprodutibilidade é confiança.

### V. Código como Documentação

O código se explica através de nomenclatura e estrutura. Comentários explicam o "porquê", nunca o "o quê". O README permanece atualizado. Novas funcionalidades começam com especificações claras. Esta constituição deve ser atualizada quando princípios evoluírem.

**Justificativa**: A era Jekyll acumulou dívida técnica através de configurações obscuras e convenções não documentadas. Documentação explícita é mais barata que arqueologia.

---

## Restrições Técnicas

| Camada | Tecnologia |
|--------|------------|
| Gerador | Go 1.21+ |
| Templates | html/template (built-in) |
| Estilização | CSS3 com variáveis nativas |
| JavaScript | ES6+ vanilla, modular |
| Build | Gerador Go customizado + tdewolff/minify |
| Hospedagem | GitHub Pages com domínio customizado |
| CI/CD | GitHub Actions |

---

## Fluxo de Desenvolvimento

1. **Especificar antes de implementar** - Princípio YAGNI. Se não está na especificação, não é construído.
2. **Testes para lógica** - Código em Go carrega testes unitários (meta >70% de cobertura).
3. **Validação de acessibilidade** - Lighthouse e checagens manuais antes do merge.
4. **Commits semânticos** - formato de commit convencional.
5. **Revisão de PR obrigatória** - Sem pushes diretos para main.
6. **Guiado pela constituição** - Qualquer PR violando estes princípios deve atualizar a constituição ou será rejeitado.

---

## Governança

### Versionamento

Este projeto segue SemVer:

- **MAJOR**: Mudanças breaking na estrutura do site, esquemas de URL ou APIs públicas
- **MINOR**: Novas seções, funcionalidades significativas, refrescos de design
- **PATCH**: Atualizações de conteúdo, correções de bugs, ajustes menores de estilo

### Processo de Emenda

1. Propor mudança em um PR com justificativa
2. Atualizar este arquivo de constituição com bump de versão
3. Documentar no Relatório de Impacto da Sincronização (comentário HTML)
4. O PR deve demonstrar conformidade com princípios existentes ou justificar exceções principiadas

### Versões Históricas (Pré-Constituição)

| Era | Tecnologia | Período |
|-----|------------|---------|
| v1.x | Jekyll (inicial) | Fev 2015 – 2019 |
| v2.x | Jekyll (estabilizado) | 2019 – Mar 2026 |
| **v3.x** | **Gerador estático em Go** | **Mar 2026 – presente** |

---

**Versão**: 3.0.0 | **Ratificada**: 2026-04-19 | **Última Emenda**: 2026-04-19
