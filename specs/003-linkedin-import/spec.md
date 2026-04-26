# Especificação da Funcionalidade: LinkedIn Import CLI

**Branch da Funcionalidade**: `003-linkedin-import`  
**Criada**: 2026-04-26  
**Status**: Rascunho  
**Input**: Descrição do usuário: "Ferramenta CLI para importar experiências, educação e certificações do LinkedIn a partir do export CSV manual, converter datas para português, dividir descrições em bullets, comparar com config.yaml atual e aplicar mudanças com confirmação interativa"

## Cenários de Usuário & Testes *(obrigatório)*

### História de Usuário 1 - Importar Experiências Profissionais (Prioridade: P1)

Como usuário do site de portfólio, quero importar minhas experiências profissionais do LinkedIn para que eu não precise digitá-las manualmente no config.yaml.

**Por que esta prioridade**: Experiências profissionais são o conteúdo principal do portfólio e representam o maior volume de dados manuais. Automatizar essa importação economiza tempo significativo e reduz erros de digitação.

**Teste Independente**: Pode ser testado executando a ferramenta CLI com um arquivo CSV de experiências do LinkedIn e verificando se as entradas são parseadas corretamente com título da empresa, cargo, período e descrição.

**Cenários de Aceitação**:

1. **Dado que** o usuário possui um arquivo CSV de experiências exportado do LinkedIn, **Quando** executar a ferramenta com o comando de importação, **Então** o sistema deve ler e parsear todas as experiências do arquivo
2. **Dado que** uma experiência tem descrição em formato texto corrido, **Quando** o sistema processar a importação, **Então** a descrição deve ser dividida automaticamente em bullets (itens de lista)
3. **Dado que** as datas estão em formato inglês (ex: "Feb 2025"), **Quando** o sistema processar a importação, **Então** as datas devem ser convertidas para português (ex: "Fev 2025")

---

### História de Usuário 2 - Importar Educação (Prioridade: P2)

Como usuário do site de portfólio, quero importar minha formação acadêmica do LinkedIn para manter meu currículo atualizado automaticamente.

**Por que esta prioridade**: Educação é importante mas tem menor volume de dados que experiências. É a próxima prioridade após as experiências profissionais.

**Teste Independente**: Pode ser testado fornecendo um CSV de educação do LinkedIn e verificando se instituição, grau, área de estudo e datas são extraídos corretamente.

**Cenários de Aceitação**:

1. **Dado que** o usuário possui um arquivo CSV de educação exportado do LinkedIn, **Quando** executar a ferramenta, **Então** o sistema deve extrair instituição, grau, área de estudo e datas
2. **Dado que** a descrição da educação contém informações relevantes, **Quando** o sistema processar, **Então** essas informações devem ser preservadas e formatadas adequadamente

---

### História de Usuário 3 - Importar Certificações (Prioridade: P3)

Como usuário do site de portfólio, quero importar minhas certificações do LinkedIn para destacar minhas qualificações profissionais.

**Por que esta prioridade**: Certificações complementam o perfil mas são opcionais. Tem prioridade menor que experiências e educação.

**Teste Independente**: Pode ser testado fornecendo um CSV de certificações e verificando se nome da certificação, organização emissora e data de emissão são extraídos.

**Cenários de Aceitação**:

1. **Dado que** o usuário possui um arquivo CSV de certificações exportado do LinkedIn, **Quando** executar a ferramenta, **Então** o sistema deve extrair nome da certificação, organização emissora e datas
2. **Dado que** algumas certificações possuem data de expiração, **Quando** o sistema processar, **Então** essa informação deve ser preservada se presente

---

### História de Usuário 4 - Comparar e Aplicar Mudanças com Confirmação (Prioridade: P1)

Como usuário do site de portfólio, quero visualizar as diferenças entre os dados importados do LinkedIn e meu config.yaml atual, e decidir quais mudanças aplicar interativamente.

**Por que esta prioridade**: Esta é uma funcionalidade de segurança crítica que evita sobrescrição acidental de dados e dá controle total ao usuário sobre quais alterações serão feitas.

**Teste Independente**: Pode ser testado tendo um config.yaml existente e um CSV do LinkedIn, executando a ferramenta em modo dry-run e verificando se um diff claro é exibido.

**Cenários de Aceitação**:

1. **Dado que** existe um config.yaml atual e dados foram importados do LinkedIn, **Quando** o usuário executar a ferramenta, **Então** o sistema deve mostrar um diff visual comparando as diferenças
2. **Dado que** o usuário está visualizando o diff, **Quando** for solicitada confirmação, **Então** o usuário deve poder aceitar todas as mudanças, rejeitar todas, ou selecionar mudanças específicas
3. **Dado que** o usuário rejeitou uma mudança específica, **Quando** o sistema aplicar as alterações, **Então** essa mudança deve ser preservada conforme estava no config.yaml original
4. **Dado que** o usuário executou em modo dry-run, **Quando** a comparação for exibida, **Então** nenhuma alteração deve ser feita no arquivo config.yaml

---

### Casos de Borda

- O que acontece quando o arquivo CSV está vazio ou não pode ser lido?
- Como o sistema lida com datas mal formatadas ou inválidas no CSV do LinkedIn?
- O que acontece quando o config.yaml atual está mal formatado ou não existe?
- Como o sistema trata entradas duplicadas (mesma empresa/cargo ou mesma instituição/curso)?
- O que acontece quando a descrição está vazia ou contém apenas espaços?
- Como o sistema lida com caracteres especiais ou encoding incorreto no CSV?
- O que acontece se o usuário interromper o processo durante a confirmação interativa?

## Requisitos *(obrigatório)*

### Requisitos Funcionais

- **RF-001**: O sistema DEVE ler arquivos CSV exportados manualmente do LinkedIn (Experiences.csv, Education.csv, Certifications.csv)
- **RF-002**: O sistema DEVE converter datas do formato inglês para português (ex: "Jan 2020" → "Jan 2020", "Feb 2025" → "Fev 2025", "Present" → "Presente")
- **RF-003**: O sistema DEVE dividir descrições em formato texto corrido em bullets (itens de lista), identificando parágrafos ou frases separadas por quebras de linha/pontuação
- **RF-004**: O sistema DEVE carregar e parsear o arquivo config.yaml atual do projeto
- **RF-005**: O sistema DEVE comparar os dados importados do LinkedIn com o config.yaml existente e identificar:
  - Entradas novas (presentes no LinkedIn, ausentes no config.yaml)
  - Entradas modificadas (diferentes entre LinkedIn e config.yaml)
  - Entradas removidas (presentes no config.yaml, ausentes no LinkedIn)
- **RF-006**: O sistema DEVE exibir um diff visual colorido mostrando as diferenças entre os dados
- **RF-007**: O sistema DEVE permitir confirmação interativa das mudanças (yes/no para cada alteração, ou aceitar/rejeitar tudo)
- **RF-008**: O sistema DEVE fazer backup do config.yaml antes de aplicar qualquer alteração
- **RF-009**: O sistema DEVE aplicar apenas as mudanças confirmadas pelo usuário ao config.yaml
- **RF-010**: O sistema DEVE suportar modo dry-run onde mostra o diff mas não modifica o arquivo

### Entidades Chave

- **Experience**: Representa uma experiência profissional. Atributos: company (empresa), role (cargo), start_date (data início), end_date (data fim), description (descrição em formato de lista de bullets)
- **Education**: Representa uma formação acadêmica. Atributos: institution (instituição), degree (grau), field (área de estudo), start_date, end_date, description
- **Certification**: Representa uma certificação profissional. Atributos: name (nome), organization (organização emissora), issue_date (data emissão), expiration_date (data expiração, opcional)
- **Config**: Representa o arquivo config.yaml atual do portfólio. Contém as listas de experiences, education e certifications

## Critérios de Sucesso *(obrigatório)*

### Resultados Mensuráveis

- **CS-001**: Usuários conseguem importar experiências do LinkedIn em menos de 1 minuto (excluindo tempo de confirmação)
- **CS-002**: A ferramenta converte 100% das datas do formato LinkedIn para português corretamente
- **CS-003**: A ferramenta divide descrições em bullets com precisão superior a 90% (avaliação manual de amostra)
- **CS-004**: O diff mostra todas as diferenças de forma clara e compreensível para o usuário
- **CS-005**: 100% das alterações são confirmadas pelo usuário antes de serem aplicadas (nenhuma alteração automática)
- **CS-006**: Backup do config.yaml é criado com sucesso em 100% das execuções que modificam o arquivo
- **CS-007**: Usuários conseguem selecionar mudanças específicas (não apenas aceitar/rejeitar tudo) em menos de 30 segundos de interação

## Suposições

- Usuários possuem acesso ao export manual de dados do LinkedIn (Settings > Data privacy > Get a copy of your data)
- O formato CSV do LinkedIn segue o padrão atual de exportação (pode variar por idioma da conta)
- O config.yaml segue a estrutura atual do projeto de portfólio
- Usuários têm permissão de escrita no diretório do projeto
- A ferramenta será executada em ambiente local (não em CI/CD)
- Usuários preferem controle total sobre as mudanças via confirmação interativa vs. automação total

## Clarifications

### Session 2026-04-26

- **Q:** Como o sistema deve determinar que uma entrada no CSV corresponde a uma entrada existente no config.yaml?  
  **A:** Match por chave composta: nome da empresa/instituição normalizado (lowercase, trimmed) + cargo/grau + data de início. Esta abordagem lida com pequenas diferenças de formatação enquanto identifica corretamente a mesma posição lógica.
- **Q:** Qual deve ser o comportamento quando o usuário tenta importar um CSV que contém entradas duplicadas?  
  **A:** Alertar o usuário e permitir escolher: mesclar descrições, ignorar duplicata, ou manter ambas. Isso dá controle ao usuário sem perder dados acidentalmente.
- **Q:** Como devemos tratar descrições vazias ou com apenas espaços no CSV?  
  **A:** Ignorar campos vazios (não criar bullet) e manter descrição existente se houver.
- **Q:** O sistema deve persistir estado de importação parcialmente concluída?  
  **A:** Não persistir estado; usuário deve reiniciar do zero. Alinhado com princípio de simplicidade.
- **Q:** Qual deve ser o nome do comando CLI principal?  
  **A:** `linkedin-import` (kebab-case, descritivo). Segue convenções CLI Unix e descreve exatamente a função da ferramenta.
