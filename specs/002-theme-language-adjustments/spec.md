# Especificação da Funcionalidade: Ajustes de Tema e Idioma

**Branch da Funcionalidade**: `002-theme-language-adjustments`  
**Criada**: 2026-04-19  
**Status**: ✅ Concluído (Implementado via PR #108)  
**Input**: Descrição do usuário: "Ajuste nos ícones de tema para mostrar lua no modo claro e sol no modo escuro, alteração do tema padrão para escuro (dark), e verificação da detecção automática de idioma baseada nas configurações do navegador/SO do usuário"

## Cenários de Usuário & Testes *(obrigatório)*

### História de Usuário 1 - Correção dos Ícones de Tema (Prioridade: P1)

Como visitante do site, quero ver o ícone de lua quando o site está em modo claro e o ícone de sol quando está em modo escuro, para que o botão de alternância indique claramente qual tema será ativado ao clicar.

**Por que esta prioridade**: É um bug de UX - o ícone atual não reflete o estado correto do tema, confundindo o usuário sobre qual ação será tomada.

**Teste Independente**: Pode ser testado visualmente ao acessar o site - o ícone do botão de tema deve ser uma lua no modo claro e um sol no modo escuro.

**Cenários de Aceitação**:

1. **Dado que** o site está no modo claro, **Quando** eu olhar o botão de tema, **Então** devo ver o ícone de lua
2. **Dado que** o site está no modo escuro, **Quando** eu olhar o botão de tema, **Então** devo ver o ícone de sol
3. **Dado que** o site está no modo claro, **Quando** eu clicar no botão de tema, **Então** o tema muda para escuro e o ícone muda para sol

---

### História de Usuário 2 - Tema Padrão Escuro (Prioridade: P1)

Como visitante do site pela primeira vez, quero que o tema escuro seja o padrão, para que a experiência inicial seja consistente com preferências modernas de design.

**Por que esta prioridade**: Define a primeira impressão do site para novos visitantes e alinha com tendências de design minimalista escuro.

**Teste Independente**: Pode ser testado abrindo o site em navegação privada (sem cache) - o tema deve iniciar escuro automaticamente.

**Cenários de Aceitação**:

1. **Dado que** nunca visitei o site antes (sem dados salvos), **Quando** eu acessar o site, **Então** o tema deve iniciar escuro
2. **Dado que** já visitei o site e escolhi o tema claro, **Quando** eu acessar novamente, **Então** minha preferência salva deve ser respeitada
3. **Dado que** o tema padrão é escuro, **Quando** o site carregar, **Então** a variável CSS deve refletir o tema escuro

---

### História de Usuário 3 - Detecção Automática de Idioma (Prioridade: P2)

Como visitante internacional do site, quero que o idioma seja detectado automaticamente baseado nas configurações do meu navegador/SO, para que eu veja o conteúdo no idioma correto sem precisar manualmente alternar.

**Por que esta prioridade**: Melhora a experiência para usuários internacionais, mas não bloqueia o uso já que o toggle manual está disponível.

**Teste Independente**: Pode ser testado alterando o idioma do navegador para inglês e abrindo o site em navegação privada - o site deve aparecer em inglês automaticamente.

**Cenários de Aceitação**:

1. **Dado que** meu navegador está configurado em português, **Quando** eu acessar o site pela primeira vez, **Então** o conteúdo deve aparecer em português
2. **Dado que** meu navegador está configurado em inglês, **Quando** eu acessar o site pela primeira vez, **Então** o conteúdo deve aparecer em inglês
3. **Dado que** eu já escolhi um idioma manualmente, **Quando** eu acessar novamente, **Então** minha preferência salva deve ser respeitada

---

### História de Usuário 4 - Animação Fade-In no Carregamento (Prioridade: P2)

Como visitante do site, quero ver o conteúdo aparecer suavemente com uma animação fade-in ao carregar a página, para uma experiência visual mais elegante e refinada, similar ao site annamona.co.

**Por que esta prioridade**: Adiciona um toque de sofisticação visual que alinha o site com o design de referência do annamona.co, melhorando a percepção de qualidade sem afetar a performance.

**Teste Independente**: Pode ser testado ao recarregar a página - o conteúdo deve aparecer gradualmente com uma animação suave de fade-in em vez de aparecer instantaneamente.

**Cenários de Aceitação**:

1. **Dado que** eu acesso o site, **Quando** a página carregar, **Então** o conteúdo deve aparecer com animação fade-in suave
2. **Dado que** a animação fade-in está ativa, **Quando** eu observar o carregamento, **Então** a duração deve ser aproximadamente 0.8 segundos com curva ease-in
3. **Dado que** eu tenho preferência por movimento reduzido, **Quando** eu acessar o site, **Então** a animação deve ser respeitada (reduzida ou desabilitada)

---

### História de Usuário 5 - Correções de Segurança e Qualidade (Prioridade: P1) 🎯 CRÍTICO

Como administrador do site, quero que vulnerabilidades de segurança e problemas de qualidade de código sejam corrigidos, para garantir a integridade e segurança do site.

**Por que esta prioridade**: Problemas de segurança como XSS e permissões excessivas em workflows podem comprometer o site e seus visitantes.

**Teste Independente**: Verificar que as correções foram aplicadas através de análise estática de código e testes manuais.

**Cenários de Aceitação**:

1. **Dado que** o gerador de site processa conteúdo do config.yaml, **Quando** houver tentativa de injeção XSS, **Então** o conteúdo deve ser sanitizado antes de renderização
2. **Dado que** os workflows de CI/CD executam, **Quando** verificar permissões, **Então** devem ter apenas os scopes necessários (princípio do menor privilégio)
3. **Dado que** o workflow de release aguarda deploy, **Quando** o deploy estiver em andamento, **Então** deve usar polling em vez de sleep fixo
4. **Dado que** animações CSS são aplicadas, **Quando** renderizar elementos, **Então** devem usar `will-change` para otimização de performance

---

### Casos de Borda

- O que acontece quando o navegador está configurado para um idioma não suportado (ex: espanhol)?
- Como o sistema lida quando o usuário tem preferência de tema salva mas o navegador muda para modo claro/escuro do sistema?
- O que acontece se o localStorage estiver desabilitado no navegador?
- Como o sistema lida com usuários que têm preferência por movimento reduzido configurada no sistema operacional?
- O que acontece se o JavaScript estiver desabilitado - a animação fade-in ainda funciona de forma degradada?
- Como o sistema lida com tentativas de injeção de scripts via config.yaml?
- O que acontece se o workflow de release não conseguir verificar o status do deploy?

## Requisitos *(obrigatório)*

### Requisitos Funcionais

- **RF-001**: O sistema DEVE exibir o ícone de lua quando o tema atual é claro
- **RF-002**: O sistema DEVE exibir o ícone de sol quando o tema atual é escuro
- **RF-003**: O sistema DEVE usar o tema escuro como padrão para novos visitantes
- **RF-004**: O sistema DEVE detectar o idioma preferido do navegador do usuário
- **RF-005**: O sistema DEVE priorizar português quando o navegador indica pt-*
- **RF-006**: O sistema DEVE usar inglês como fallback para idiomas não suportados
- **RF-007**: O sistema DEVE respeitar preferências salvas do usuário sobre as detecções automáticas
- **RF-008**: O sistema DEVE alternar entre os dois ícones de forma suave e acessível
- **RF-009**: O sistema DEVE aplicar animação fade-in nos elementos principais ao carregar a página
- **RF-010**: O sistema DEVE respeitar a preferência do usuário por movimento reduzido (prefers-reduced-motion)
- **RF-011**: O sistema DEVE sanitizar conteúdo do config.yaml antes de renderização para prevenir XSS
- **RF-012**: O sistema DEVE configurar permissões mínimas necessárias nos workflows de CI/CD
- **RF-013**: O sistema DEVE implementar polling com backoff para aguardar deploy em workflows
- **RF-014**: O sistema DEVE otimizar animações CSS com `will-change` para melhor performance

### Entidades Chave

- **Preferência de Tema**: Representa a escolha do usuário (claro/escuro), persistida entre sessões
- **Preferência de Idioma**: Representa a escolha do usuário (PT/EN), persistida entre sessões

## Critérios de Sucesso *(obrigatório)*

### Resultados Mensuráveis

- **CS-001**: 100% dos novos visitantes devem ver o tema escuro na primeira visita
- **CS-002**: O ícone correto deve ser exibido em 100% dos casos para cada estado de tema
- **CS-003**: Usuários com navegador em inglês devem ver o site em inglês sem intervenção manual
- **CS-004**: A transição de ícones deve ocorrer sem flashes ou atrasos perceptíveis (< 100ms)
- **CS-005**: Preferências salvas do usuário devem persistir com 100% de confiabilidade
- **CS-006**: A animação fade-in deve ter duração de 0.8s com timing ease-in
- **CS-007**: A animação deve ser desabilitada automaticamente quando prefers-reduced-motion está ativo
- **CS-008**: O conteúdo do config.yaml deve ser sanitizado 100% das vezes antes da renderização
- **CS-009**: Workflows devem ter apenas permissões necessárias (nenhuma permissão excessiva)
- **CS-010**: O polling de deploy deve verificar status a cada 15s até 10 tentativas

## Suposições

- Usuários têm navegadores modernos que suportam a API navigator.language
- O localStorage está disponível para persistência de preferências
- Ícones SVG de sol e lua serão utilizados para representação visual
- A detecção de idioma considera apenas os dois idiomas suportados (PT/EN)
- O navegador suporta CSS animations e media query prefers-reduced-motion
- O efeito fade-in é aplicado aos elementos principais da página (header, content sections)
