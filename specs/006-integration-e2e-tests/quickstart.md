# Quickstart: Executando Testes de Integração e E2E

Este guia descreve como executar, validar e atualizar a suíte de testes de alta ordem do projeto.

## 1. Executando os Testes

Toda a suíte de testes é integrada ao comando padrão do Go.

### Executar todos os testes (incluindo unitários)
```bash
go test ./...
```

### Executar apenas testes de integração e E2E
Como esses testes podem ser mais lentos ou exigir arquivos externos, eles são marcados com a build tag `integration`.
```bash
go test -tags=integration ./...
```

---

## 2. Trabalhando com Golden Files

Se você alterar a lógica de transformação de dados (ex: mudar como as datas são formatadas), os testes de integração falharão porque o output não baterá com o Golden File.

### Como atualizar os arquivos de referência
Se a mudança for intencional e correta, você pode atualizar os Golden Files automaticamente:
```bash
go test -tags=integration -update ./cmd/import-linkedin/...
```
*O sistema escreverá o novo output no arquivo `.yaml` de referência e o teste passará a usá-lo como novo padrão.*

---

## 3. Debugging de Testes E2E

Os testes E2E utilizam diretórios temporários (`t.TempDir()`). Se um teste falhar e você quiser inspecionar o HTML gerado:

1. No código do teste, altere o caminho do diretório temporário para um caminho fixo (ex: `/tmp/e2e-debug`).
2. Execute o teste.
3. Abra a pasta no navegador ou editor de texto para analisar o `index.html` gerado.

## 4. CI/CD Integration
Os testes são executados automaticamente no GitHub Actions durante o merge para `main`. Se qualquer Golden File divergir ou o HTML final não conter as strings esperadas, o build falhará.
