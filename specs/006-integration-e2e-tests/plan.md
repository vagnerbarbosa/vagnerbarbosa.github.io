# Plano de Implementação: Implementação de Suíte de Testes de Integração e E2E

**Branch**: `006-integration-e2e-tests` | **Data**: 2026-05-30 | **Spec**: [spec.md](spec.md)
**Input**: Especificação da funcionalidade de `/specs/006-integration-e2e-tests/spec.md`

## Resumo

Implementação de uma estratégia de testes abrangente para garantir a confiabilidade do pipeline de dados e do output final do site. A abordagem consiste em dois pilares:
1. **Testes de Integração (Golden Files)**: Validação bit-a-bit da transformação de CSV $\rightarrow$ YAML para evitar regressões na extração de dados do LinkedIn.
2. **Testes End-to-End (E2E)**: Validação do fluxo completo `Import $\rightarrow$ Generate $\rightarrow$ Public`, verificando a presença de dados importados no HTML final e a integridade dos arquivos de SEO/PWA.

## Contexto Técnico

**Linguagem/Versão**: Go 1.25+  
**Dependências Principais**: `html/template`, `tdewolff/minify`, `yaml.v3`  
**Armazenamento**: Arquivos YAML (`config.yaml`), CSVs de entrada e Golden Files  
**Testes**: `testing` (built-in Go) + `reflect.DeepEqual` para comparação de YAML  
**Plataforma Alvo**: Linux/macOS (GitHub Actions para CI)  
**Tipo de Projeto**: CLI Tool / Static Site Generator  
**Metas de Performance**: Execução total da suíte E2E em menos de 15 segundos.  
**Restrições**: Não utilizar navegadores Headless (Puppeteer/Playwright) para manter a simplicidade e velocidade do build.  
**Escala/Alcance**: Cobertura de 3 cenários principais (Sucesso, Parcial, Erro) e validação de 1 site completo.

## Verificação da Constituição

| Princípio | Alinhamento | Observação |
| :--- | :---: | :--- |
| **Minimalismo Intencional** | ✅ | Utiliza apenas a toolchain nativa do Go; sem frameworks de teste externos. |
| **Bilíngue por Padrão** | ✅ | Valida a renderização de conteúdo PT e EN no HTML final. |
| **Estabilidade Visual** | ✅ | Previne alterações acidentais na estrutura do HTML via testes E2E. |
| **Build Reprodutível** | ✅ | Golden Files garantem que o output de dados seja idêntico em qualquer ambiente. |
| **Código como Doc** | ✅ | A estratégia é documentada via Spec e Plan antes da implementação. |

## Estrutura do Projeto

### Documentação (esta funcionalidade)

```text
specs/006-integration-e2e-tests/
├── plan.md              # Este arquivo
├── research.md          # Estratégias de Golden Files e E2E em Go
├── data-model.md        # Mapeamento de fluxos de teste
├── quickstart.md        # Como executar e atualizar Golden Files
└── contracts/           # Definição de entradas/saídas esperadas
```

### Código Fonte (raiz do repositório)

```text
cmd/import-linkedin/
├── internal/
│   └── integration/       # Novos testes de integração (SUT: Pipeline)
│       ├── golden/        # Arquivos YAML de referência
│       └── pipeline_test.go
├── testdata/
│   └── input_csv/        # CSVs usados nos testes de integração
cmd/generator/
└── e2e_test.go           # Testes End-to-End (Fluxo completo)
public/                   # Validado via E2E
```

**Decisão de Estrutura**: Optamos por manter os testes de integração próximos ao módulo de importação (`cmd/import-linkedin/internal/integration`) e os testes E2E no nível do gerador ou em um pacote de testes global, já que eles orquestram múltiplos comandos.

## Rastreamento de Complexidade

Nenhuma violação da constituição detectada. A implementação segue rigorosamente os princípios de minimalismo e reprodutibilidade.

---

## Planejamento de Implementação

### Fase 1: Infraestrutura de Golden Files (Integração)
- [ ] Criar diretório `cmd/import-linkedin/testdata/golden/`.
- [ ] Implementar função helper para comparar YAMLs ignorando a ordem de chaves (usando `reflect.DeepEqual` após unmarshal).
- [ ] Implementar teste de integração que processa CSVs de teste e valida contra o Golden File.
- [ ] Criar comando para "atualizar" Golden Files quando mudanças no formato forem intencionais.

### Fase 2: Framework E2E (End-to-End)
- [ ] Implementar `E2ETestRunner` que:
    - Cria ambiente temporário.
    - Copia templates e assets necessários para o ambiente.
    - Executa a CLI de importação.
    - Executa o gerador de site.
- [ ] Implementar validações de output:
    - Verificar se `public/index.html` existe.
    - Verificar se strings chave do CSV de entrada aparecem no HTML.
    - Verificar se `sitemap.xml` e `robots.txt` foram gerados.

### Fase 3: Automação e CI
- [ ] Adicionar flag de teste para ignorar E2E em builds rápidos (opcional).
- [ ] Configurar GitHub Actions para rodar a suíte E2E em cada PR.

### Fase 4: Documentação de Qualidade
- [ ] Criar `docs/TESTING.md` detalhando a estratégia de testes (Unitários, Integração e E2E).
- [ ] Documentar o processo de atualização de Golden Files (comando `-update`).
- [ ] Atualizar o `README.md` na seção de "Desenvolvimento" para incluir a execução dos testes de integração.
