# Research: Implementação de Suíte de Testes de Integração e E2E

## Objetivo
Definir a melhor abordagem técnica para implementar testes de integração via Golden Files e testes E2E para o pipeline de geração de site estático.

## 1. Testes de Integração com Golden Files

### O que são?
A técnica de Golden Files consiste em salvar a saída correta de um processo em um arquivo. Em execuções futuras, a saída atual é comparada com este arquivo. Se divergirem, o teste falha.

### Abordagem em Go
Para validar YAMLs, a comparação bit-a-bit (`bytes.Equal`) é frágil devido a possíveis mudanças na ordem das chaves ou espaços em branco.

**Decisão de Design**:
- **Comparação Semântica**: Fazer o `Unmarshal` de ambos os arquivos (atual e golden) para `map[string]interface{}` e usar `reflect.DeepEqual`.
- **Atualização de Goldens**: Implementar uma flag `-update` nos testes para sobrescrever os arquivos Golden quando a mudança for intencional.

### Estrutura de Diretórios
```text
testdata/
├── input_csv/
│   └── experience.csv
└── golden/
    └── experience_expected.yaml
```

---

## 2. Testes End-to-End (E2E)

### Definição do Fluxo
O teste E2E deve simular a jornada completa:
`CSV Import $\rightarrow$ config.yaml $\rightarrow$ HTML Build $\rightarrow$ HTML Validation`.

### Estratégia de Isolamento
Para evitar que os testes alterem o `config.yaml` real do projeto, utilizaremos `t.TempDir()` do Go.
- O ambiente temporário conterá:
  - Uma cópia dos templates HTML.
  - Uma cópia dos assets básicos.
  - Um `config.yaml` inicial vazio.

### Validação do Output
Dado que o site é estático, não precisamos de um navegador (Headless Browser). 
**Técnica**:
- **Análise de Strings**: Verificar se tags específicas e conteúdos do CSV aparecem no HTML final.
- **Verificação de Existência**: Checar se `public/sitemap.xml`, `robots.txt` e `site.webmanifest` foram criados.

---

## 3. Conclusões e Trade-offs

| Abordagem | Pró | Contra |
| :--- | :--- | :--- |
| **Sementes de Dados Fixas** | Testes determinísticos e rápidos. | Não cobrem bugs aleatórios de runtime. |
| **Validação via Texto** | Extremamente rápido, sem dependências externas. | Não valida a renderização visual (CSS/JS). |
| ** la-linkedin/testdata** | Reaproveita dados existentes. | Requer manutenção manual dos CSVs de teste. |

**Veredito**: A combinação de Golden Files para dados + Validação de Texto para HTML oferece o melhor equilíbrio entre confiança e velocidade de build, alinhado ao princípio de Minimalismo Intencional.
