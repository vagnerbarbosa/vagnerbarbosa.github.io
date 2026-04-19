# Especificação da Funcionalidade: [NOME DA FUNCIONALIDADE]

**Branch da Funcionalidade**: `[###-nome-da-funcionalidade]`  
**Criada**: [DATA]  
**Status**: Rascunho  
**Input**: Descrição do usuário: "$ARGUMENTS"

## Cenários de Usuário & Testes *(obrigatório)*

<!--
  IMPORTANTE: Histórias de usuário devem ser PRIORIZADAS como jornadas ordenadas por importância.
  Cada história/jornada do usuário deve ser INDEPENDENTEMENTE TESTÁVEL - ou seja, se você implementar apenas UMA delas,
  você ainda deve ter um MVP (Produto Mínimo Viável) que entrega valor.
  
  Atribua prioridades (P1, P2, P3, etc.) para cada história, onde P1 é a mais crítica.
  Pense em cada história como uma fatia independente de funcionalidade que pode ser:
  - Desenvolvida independentemente
  - Testada independentemente
  - Deployada independentemente
  - Demonstrada para usuários independentemente
-->

### História de Usuário 1 - [Título Breve] (Prioridade: P1)

[Descreva esta jornada do usuário em linguagem simples]

**Por que esta prioridade**: [Explique o valor e por que ela tem este nível de prioridade]

**Teste Independente**: [Descreva como isto pode ser testado independentemente - ex: "Pode ser totalmente testado por [ação específica] e entrega [valor específico]"]

**Cenários de Aceitação**:

1. **Dado que** [estado inicial], **Quando** [ação], **Então** [resultado esperado]
2. **Dado que** [estado inicial], **Quando** [ação], **Então** [resultado esperado]

---

### História de Usuário 2 - [Título Breve] (Prioridade: P2)

[Descreva esta jornada do usuário em linguagem simples]

**Por que esta prioridade**: [Explique o valor e por que ela tem este nível de prioridade]

**Teste Independente**: [Descreva como isto pode ser testado independentemente]

**Cenários de Aceitação**:

1. **Dado que** [estado inicial], **Quando** [ação], **Então** [resultado esperado]

---

### História de Usuário 3 - [Título Breve] (Prioridade: P3)

[Descreva esta jornada do usuário em linguagem simples]

**Por que esta prioridade**: [Explique o valor e por que ela tem este nível de prioridade]

**Teste Independente**: [Descreva como isto pode ser testado independentemente]

**Cenários de Aceitação**:

1. **Dado que** [estado inicial], **Quando** [ação], **Então** [resultado esperado]

---

[Adicione mais histórias de usuário conforme necessário, cada uma com uma prioridade atribuída]

### Casos de Borda

<!--
  AÇÃO NECESSÁRIA: O conteúdo desta seção representa placeholders.
  Preencha-os com os casos de borda corretos.
-->

- O que acontece quando [condição de limite]?
- Como o sistema lida com [cenário de erro]?

## Requisitos *(obrigatório)*

<!--
  AÇÃO NECESSÁRIA: O conteúdo desta seção representa placeholders.
  Preencha-os com os requisitos funcionais corretos.
-->

### Requisitos Funcionais

- **RF-001**: O sistema DEVE [capacidade específica, ex: "permitir que usuários criem contas"]
- **RF-002**: O sistema DEVE [capacidade específica, ex: "validar endereços de email"]  
- **RF-003**: Usuários DEVEM ser capazes de [interação chave, ex: "redefinir sua senha"]
- **RF-004**: O sistema DEVE [requisito de dados, ex: "persistir preferências do usuário"]
- **RF-005**: O sistema DEVE [comportamento, ex: "registrar todos os eventos de segurança"]

*Exemplo de marcação de requisitos pouco claros:*

- **RF-006**: O sistema DEVE autenticar usuários via [PRECISA DE ESCLARECIMENTO: método de auth não especificado - email/senha, SSO, OAuth?]
- **RF-007**: O sistema DEVE reter dados do usuário por [PRECISA DE ESCLARECIMENTO: período de retenção não especificado]

### Entidades Chave *(inclua se a funcionalidade envolve dados)*

- **[Entidade 1]**: [O que representa, atributos principais sem implementação]
- **[Entidade 2]**: [O que representa, relacionamentos com outras entidades]

## Critérios de Sucesso *(obrigatório)*

<!--
  AÇÃO NECESSÁRIA: Defina critérios de sucesso mensuráveis.
  Estes devem ser agnósticos de tecnologia e mensuráveis.
-->

### Resultados Mensuráveis

- **CS-001**: [Métrica mensurável, ex: "Usuários podem completar criação de conta em menos de 2 minutos"]
- **CS-002**: [Métrica mensurável, ex: "Sistema lida com 1000 usuários simultâneos sem degradação"]
- **CS-003**: [Métrica de satisfação do usuário, ex: "90% dos usuários completam tarefa primária na primeira tentativa"]
- **CS-004**: [Métrica de negócio, ex: "Reduzir tickets de suporte relacionados a [X] em 50%"]

## Suposições

<!--
  AÇÃO NECESSÁRIA: O conteúdo desta seção representa placeholders.
  Preencha-os com as suposições corretas baseadas em padrões razoáveis
  escolhidos quando a descrição da funcionalidade não especificou certos detalhes.
-->

- [Suposição sobre usuários alvo, ex: "Usuários têm conectividade de internet estável"]
- [Suposição sobre limites de escopo, ex: "Suporte mobile está fora do escopo para v1"]
- [Suposição sobre dados/ambiente, ex: "Sistema de autenticação existente será reutilizado"]
- [Dependência em sistema/serviço existente, ex: "Requer acesso à API de perfil de usuário existente"]
