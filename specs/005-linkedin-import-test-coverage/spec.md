# Especificação da Funcionalidade: Cobertura de Testes Total do Projeto

## Status
- **Versão**: 1.1.0
- **Status**: Rascunho
- **Dono**: Vagner Barbosa

## 1. Visão Geral
Esta funcionalidade visa atingir 100% de cobertura de instruções (statement coverage) para todo o projeto, incluindo o Gerador de Site Estático (`cmd/generator`) e a ferramenta CLI LinkedIn Import (`cmd/import-linkedin`). O objetivo é eliminar "pontos cegos" em toda a base de código, garantindo que cada caminho lógico, incluindo manipuladores de erro e casos de borda, seja validado por testes automatizados.

## 2. Cenários de Usuário e Testes
### Cenário 1: Execução da Suíte Completa
- **Dado**: Um desenvolvedor deseja verificar a integridade geral de todo o projeto.
- **Quando**: Executando a suíte de testes completa (`go test ./...`).
- **Então**: Todos os testes passam e o relatório de cobertura agregado mostra 100% de instruções executadas para todos os pacotes.

### Cenário 2: Validação de Casos de Borda (Gerador e Importador)
- **Dado**: Insumos malformados, como arquivos de configuração YAML inválidos para o gerador ou arquivos CSV corrompidos para o importador.
- **Quando**: O sistema processa esses arquivos.
- **Então**: O sistema lida com o erro graciosamente sem causar pânico, e os caminhos de código de manipulação de erro são marcados como cobertos.

### Cenário 3: Extração de Tech Stack e Minificação
- **Dado**: Descrições complexas para a extração de tech stack ou assets com caracteres inesperados para a minificação.
- **Quando**: As respectivas lógicas de transformação e minificação são executadas.
- **Então**: Os resultados correspondem à saída esperada e a lógica de processamento é integralmente validada.

## 3. Requisitos Funcionais

### 3.1 Cobertura do Gerador (`cmd/generator`)
- [ ] A lógica de `main()` deve ser testável e coberta (ex: refatorando para retornar erro em vez de `os.Exit`).
- [ ] Todas as ramificações de `copyAndMinifyDir` e `copyAndMinifyFile` devem ser executadas, incluindo falhas de permissão e arquivos inexistentes.
- [ ] A renderização de templates e a minificação do HTML final devem ser integralmente validadas.

### 3.2 Cobertura do Parser do Importador
- [ ] Toda a lógica de leitura de CSV em `internal/parser/csv.go` deve ser coberta, incluindo `NewCSVReaderFromIO` e `ValidateColumns`.
- [ ] Cada ramificação em `ConvertDate` e `ValidateDate` deve ser executada, incluindo formatos inválidos e manipulação de "Presente".
- [ ] Cada entidade do parser (`Certification`, `Education`, `Experience`) deve ter testes para parsing bem-sucedido e falhas de validação.

### 3.3 Cobertura do Transformer do Importador
- [ ] Cada padrão regex em `techStackPatterns` deve ser acionado por pelo menos um caso de teste.
- [ ] Todos os tipos de separadores em `techSeparators` devem ser testados durante a normalização.
- [ ] `SplitDescription` deve ser testado para as três estratégias de divisão: parágrafos, bullets e frases.

### 3.4 Cobertura do Comparator do Importador
- [ ] Toda a lógica de comparação de entidades (`experiencesEqual`, `educationEqual`, `certificationsEqual`) deve ser testada com objetos idênticos e diferentes.
- [ ] `CompareAll` deve ser testado para cenários onde algumas entidades são adicionadas, algumas modificadas e algumas removidas.
- [ ] `ApplyChanges` deve ser verificado para garantir que o modelo `config.yaml` seja atualizado corretamente.

### 3.5 Cobertura de Comandos CLI
- [ ] Todos os comandos CLI do importador (`import`, `validate`, `version`) devem ser testados.
- [ ] Caminhos de erro para arquivos ausentes ou configuração inválida devem ser cobertos.
- [ ] A lógica da flag `Dry Run` deve ser verificada para garantir que nenhum arquivo seja escrito quando ativada.

## 4. Critérios de Sucesso
- **Quantitativo**: A cobertura de instruções para TODOS os pacotes do projeto é de exatamente 100.0%.
- **Qualitativo**: Cada "ponto cego" identificado (especialmente em `cmd/generator/main.go` e nos parsers do LinkedIn Import) possui um caso de teste correspondente.
- **Verificável**: `go test ./... -cover` não produz linhas não cobertas para nenhum pacote do projeto.

## 5. Suposições
- Os testes continuarão a usar o diretório `testdata` para arquivos de entrada.
- A versão do toolchain do Go é consistente entre os ambientes de desenvolvimento e CI.
- "100% de cobertura" refere-se à cobertura de instruções (statement coverage).

## 6. Entidades Chave
- **Relatório de Cobertura**: A saída da ferramenta de cobertura do Go usada para rastrear o progresso.
- **Caso de Borda (Edge Case)**: Um conjunto de entradas que testa os limites do sistema (ex: strings vazias, comprimento máximo, caracteres inválidos).
- **Dados de Mock**: Conteúdo CSV ou arquivos de configuração sintéticos usados para acionar caminhos de erro específicos.
