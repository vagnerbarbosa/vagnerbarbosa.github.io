# Especificação da Funcionalidade: Extração Inteligente de Tech Stack

**Branch da Funcionalidade**: `004-tech-stack-extraction`
**Criada**: 2026-04-26
**Status**: Rascunho
**Input**: Descrição do usuário: Implementar inteligência no LinkedIn Import CLI para detectar e extrair tech stack de bullets de descrição, armazenando no campo TechStack separadamente.

## Cenários de Usuário & Testes *(obrigatório)*

### História de Usuário 1 - Extração Automática de Tech Stack (Prioridade: P1)

Como usuário do LinkedIn Import CLI, quero que o sistema detecte automaticamente quando um bullet point na descrição da experiência contém informações de tech stack, extraindo-as para o campo específico `TechStack` em vez de mantê-lo como bullet comum.

**Por que esta prioridade**: Esta é a funcionalidade core que resolve o problema de perda de dados estruturados durante a importação do LinkedIn.

**Teste Independente**: Pode ser testado importando um CSV com descrições contendo padrões de tech stack e verificando se o campo `TechStack` é preenchido corretamente e o bullet é removido da descrição.

**Cenários de Aceitação**:

1. **Dado que** uma experiência no CSV tem descrição "Referência técnica em FinOps\nAs principais tecnologias e ferramentas utilizadas: Java, Python, AWS", **Quando** o parser processar o CSV, **Então** o campo `Description` deve conter apenas ["Referência técnica em FinOps"] e `TechStack` deve ser "Java • Python • AWS"

2. **Dado que** uma experiência tem descrição "Gestão de equipe\nTecnologias: Kubernetes, Docker, Terraform", **Quando** o parser processar, **Então** `TechStack` deve ser "Kubernetes • Docker • Terraform"

3. **Dado que** uma experiência tem descrição sem padrões de tech stack, **Quando** o parser processar, **Então** `TechStack` deve estar vazio e todos os bullets devem permanecer em `Description`

---

### História de Usuário 2 - Suporte a Múltiplos Padrões e Formatos (Prioridade: P2)

Como usuário, quero que o sistema reconheça diferentes padrões de escrita para tech stack (tanto em português quanto em inglês) e diferentes separadores entre tecnologias.

**Por que esta prioridade**: Diferentes usuários podem usar diferentes convenções no LinkedIn, e o sistema deve ser flexível o suficiente para lidar com isso.

**Teste Independente**: Pode ser testado com CSVs contendo variações de padrões ("Tech Stack:", "Technologies:", "Ferramentas:", etc.) e diferentes separadores (vírgula, pipe, bullet, etc.).

**Cenários de Aceitação**:

1. **Dado que** uma descrição contém "Tech Stack: Go | Python | AWS", **Quando** processado, **Então** `TechStack` deve ser "Go • Python • AWS"

2. **Dado que** uma descrição contém "Tools: Terraform - Ansible - Puppet", **Quando** processado, **Então** `TechStack` deve ser "Terraform • Ansible • Puppet"

3. **Dado que** uma descrição contém "Tecnologias utilizadas: React, Node.js, TypeScript", **Quando** processado, **Então** `TechStack` deve ser "React • Node.js • TypeScript"

---

### História de Usuário 3 - Preservação de Descrições sem Tech Stack (Prioridade: P3)

Como usuário, quero que descrições que não contêm tech stack sejam processadas normalmente, sem alterações ou perda de conteúdo.

**Por que esta prioridade**: Garante que a nova funcionalidade não quebre comportamentos existentes para experiências sem tech stack explícito.

**Teste Independente**: Pode ser testado importando experiências sem padrões de tech stack e verificando que a descrição permanece intacta.

**Cenários de Aceitação**:

1. **Dado que** uma experiência tem descrição "Liderança técnica e mentoria", **Quando** processada, **Então** `Description` deve conter ["Liderança técnica e mentoria"] e `TechStack` deve estar vazio

2. **Dado que** uma experiência tem múltiplos bullets normais sem padrões de tech stack, **Quando** processada, **Então** todos os bullets devem permanecer em `Description`

---

### Casos de Borda

- O que acontece quando o padrão de tech stack aparece no meio da descrição (não no final)?
- Como o sistema lida com múltiplas ocorrências de padrões de tech stack na mesma descrição?
- O que acontece quando as tecnologias estão em formato inválido ou vazio após o prefixo?
- Como o sistema lida com bullets que contêm tanto descrição quanto tech stack na mesma linha?

## Requisitos *(obrigatório)*

### Requisitos Funcionais

- **RF-001**: O sistema DEVE detectar padrões de tech stack em bullets de descrição (case-insensitive)
- **RF-002**: O sistema DEVE suportar os seguintes padrões: "As principais tecnologias e ferramentas utilizadas:", "Tecnologias:", "Tech Stack:", "Technologies:", "Stack:", "Ferramentas:", "Tools:"
- **RF-003**: O sistema DEVE extrair as tecnologias após o padrão detectado e armazená-las no campo `TechStack`
- **RF-004**: O sistema DEVE converter diferentes separadores (vírgula, ponto, pipe, hífen, bullet) para o formato padronizado com " • "
- **RF-005**: O sistema DEVE remover o bullet contendo tech stack da lista de descrições
- **RF-006**: O sistema DEVE preservar bullets sem padrões de tech stack inalterados
- **RF-007**: O sistema DEVE suportar extração de tech stack tanto para conteúdo em português quanto em inglês
- **RF-008**: Se múltiplos bullets contiverem padrões de tech stack, o sistema DEVE usar apenas o último bullet e ignorar os anteriores

### Entidades Chave *(inclua se a funcionalidade envolve dados)*

- **Experience**: Representa uma experiência profissional importada do LinkedIn. Atributos relevantes:
  - `Description`: Lista de bullets descrevendo atividades (array de strings)
  - `TechStack`: String formatada com tecnologias separadas por " • "

## Critérios de Sucesso *(obrigatório)*

### Resultados Mensuráveis

- **CS-001**: Tech stack é extraído corretamente em 100% dos casos onde o padrão é claramente identificável
- **CS-002**: Zero falsos positivos - bullets que não contêm tech stack nunca são removidos indevidamente
- **CS-003**: O formato de saída do tech stack é consistente (separado por " • ") independente do formato de entrada
- **CS-004**: A importação de CSVs sem tech stack continua funcionando exatamente como antes (backwards compatibility)

## Suposições

- O usuário tipicamente coloca tech stack como o último bullet da descrição
- Os padrões de tech stack são seguidos por dois pontos (:) ou hífen (-)
- As tecnologias são separadas por caracteres comuns: vírgula, ponto-e-vírgula, pipe, hífen, ou bullet
- O LinkedIn exporta descrições com quebras de linha representando bullets
