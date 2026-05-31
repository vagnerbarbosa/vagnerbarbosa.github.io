# Especificação da Funcionalidade: Implementação de Suíte de Testes de Integração e E2E

**Branch da Funcionalidade**: `006-integration-e2e-tests`  
**Criada**: 2026-05-30  
**Status**: Rascunho  
**Input**: Descrição do usuário: "Implementação de Suíte de Testes de Integração e E2E: Implementar uma estratégia de testes abrangente além dos testes unitários. Isso inclui testes de integração para o pipeline de dados do LinkedIn (entrada CSV $\rightarrow$ saída YAML usando Golden Files) e testes E2E (End-to-End) que validem o fluxo completo desde a importação de dados até a geração e validação do HTML final na pasta `public/`. O objetivo é garantir a confiabilidade do comportamento do sistema e a corretude do output final do site."

## Clarificações

### Session 2026-05-30
- Q: Profundidade da Validação do HTML $\rightarrow$ A: Conjunto Mínimo: Nome do Usuário, 1 Título de Experiência, 1 Instituição de Ensino

## Cenários de Usuário & Testes *(obrigatório)*

### História de Usuário 1 - Validação do Pipeline de Dados via Golden Files (Prioridade: P1)

Como desenvolvedor, desejo que todo o pipeline de processamento de dados do LinkedIn (Parsing $\rightarrow$ Transforming $\rightarrow$ Comparing) seja validado contra arquivos de saída esperados (Golden Files) para garantir que mudanças no código não alterem a extração de dados de forma não intencional.

**Por que esta prioridade**: É a base de confiança do sistema. Se a extração de dados falhar ou mudar silenciosamente, todas as informações do site serão comprometidas.

**Teste Independente**: Pode ser testado executando o pipeline de importação com arquivos CSV de teste e comparando o YAML resultante com o arquivo Golden.

**Cenários de Aceitação**:

1. **Dado que** existe um arquivo CSV de entrada e um arquivo Golden (YAML esperado), **Quando** o pipeline de importação é executado, **Então** o resultado deve ser idêntico ao arquivo Golden.
2. **Dado que** o pipeline processa dados malformados, **Quando** a importação é executada, **Então** o sistema deve gerar o output esperado para casos de erro, mantendo a consistência.

---

### História de Usuário 2 - Validação do Fluxo End-to-End (Prioridade: P1)

Como mantenedor do projeto, desejo validar que o ciclo completo (Importação de Dados $\rightarrow$ Geração do Site $\rightarrow$ Publicação em `public/`) funciona corretamente, garantindo que as alterações em qualquer módulo não quebrem a entrega final.

**Por que esta prioridade**: É a única forma de garantir que o site realmente "funciona" como um todo, validando a integração entre a CLI de importação e o gerador estático.

**Teste Independente**: Pode ser testado executando um script que dispara a importação, roda o gerador e verifica a existência e integridade dos arquivos em `public/`.

**Cenários de Aceitação**:

1. **Dado que** um conjunto de dados de teste é importado, **Quando** o gerador é executado, **Então** a pasta `public/` deve conter um `index.html` válido.
2. **Dado que** o site foi gerado, **Quando** o conteúdo de `public/index.html` é analisado, **Então** as informações importadas (ex: títulos de experiências) devem estar presentes no HTML final.

---

### História de Usuário 3 - Prevenção de Regressão em Assets e SEO (Prioridade: P2)

Como desenvolvedor, desejo garantir que as otimizações de SEO e Assets (Sitemap, Robots, Manifest) sejam geradas corretamente após cada build, evitando a perda de pontuações Lighthouse.

**Por que esta prioridade**: Garante que as melhorias de performance e SEO implementadas recentemente não sejam removidas acidentalmente por mudanças no gerador.

**Teste Independente**: Pode ser testado verificando a presença e o formato básico dos arquivos de SEO (`sitemap.xml`, `robots.txt`, `site.webmanifest`) após a execução do gerador.

**Cenários de Aceitação**:

1. **Dado que** o gerador foi executado, **Quando** os arquivos de SEO são verificados, **Então** `sitemap.xml` e `robots.txt` devem existir e conter a URL base correta.
2. **Dado que** o site foi gerado, **Quando** o `site.webmanifest` é analisado, **Então** ele deve conter as cores de tema e ícones configurados.

### Casos de Borda

- O que acontece quando o arquivo `config.yaml` está vazio durante o teste E2E? (O sistema deve gerar um site básico ou erro controlado).
- Como o sistema lida com a ausência de arquivos CSV durante o teste de integração? (Deve reportar erro de arquivo não encontrado sem crashar).
- O que acontece se o gerador falhar mas a importação tiver tido sucesso? (O teste E2E deve falhar explicitamente no passo de geração).

## Requisitos *(obrigatório)*

### Requisitos Funcionais

- **RF-001**: O sistema DEVE implementar testes de integração para o pipeline de importação do LinkedIn utilizando a técnica de Golden Files.
- **RF-002**: O sistema DEVE validar a integridade do YAML gerado comparando-o bit-a-bit com a versão de referência.
- **RF-003**: O sistema DEVE executar um teste E2E que integre `cmd/import-linkedin` e `cmd/generator`.
- **RF-004**: O sistema DEVE validar que o HTML gerado em `public/index.html` contém, no mínimo, o nome do usuário, um título de experiência e uma instituição de ensino importados.
- **RF-005**: O sistema DEVE verificar a existência e a validade básica dos arquivos de SEO (`sitemap.xml`, `robots.txt`) e PWA (`site.webmanifest`) após o build.
- **RF-006**: A suíte de testes E2E DEVE ser capaz de rodar em um ambiente isolado (ex: pasta temporária) para não afetar o `config.yaml` de produção.

### Entidades Chave

- **Golden File**: Arquivo de referência que contém a saída "perfeita" e esperada de um processamento de dados.
- **E2E Pipeline**: Sequência de comandos `import` $\rightarrow$ `generate` $\rightarrow$ `verify`.
- **Test Data**: Conjunto de CSVs e configurações simplificadas usadas exclusivamente para disparar os testes.

## Critérios de Sucesso *(obrigatório)*

### Resultados Mensuráveis

- **CS-001**: 100% de concordância entre a saída do pipeline de importação e os arquivos Golden definidos.
- **CS-002**: O teste E2E deve completar o ciclo completo (Import $\rightarrow$ Build $\rightarrow$ HTML Check) em menos de 15 segundos.
- **CS-003**: Detecção automática de regressão: qualquer alteração não intencional no HTML final deve causar a falha do teste E2E.
- **CS-004**: Cobertura de cenários: ao menos 3 cenários de importação (sucesso, parcial, erro) devem ser validados via integração.

## Suposições

- Os testes E2E serão executados em ambiente local e/ou CI (GitHub Actions).
- O diretório `public/` será limpo antes de cada execução de teste E2E para evitar falsos positivos.
- A validação do HTML será feita via busca de strings simples ou análise de tags, sem a necessidade de um navegador completo (Headless browser), priorizando a velocidade.
- As tags de versão do projeto serão mantidas consistentes durante a execução dos testes.
