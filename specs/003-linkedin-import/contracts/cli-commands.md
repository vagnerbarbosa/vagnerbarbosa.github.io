# Contrato: Comandos da CLI

**Feature**: LinkedIn Import CLI  
**Versão**: 1.0.0

---

## Visão Geral

A CLI expõe os seguintes comandos para importação de dados do LinkedIn.

---

## Comando: `import`

Importa dados do LinkedIn a partir de arquivos CSV.

### Uso

```bash
linkedin-import import [flags]
```

### Flags

| Flag | Curto | Tipo | Padrão | Descrição |
|------|-------|------|--------|-----------|
| `--experiences` | `-e` | string | `"Experiences.csv"` | Caminho para o CSV de experiências |
| `--education` | `-E` | string | `"Education.csv"` | Caminho para o CSV de educação |
| `--certifications` | `-c` | string | `"Certifications.csv"` | Caminho para o CSV de certificações |
| `--config` | `-C` | string | `"config.yaml"` | Caminho para o config.yaml do portfólio |
| `--dry-run` | `-d` | bool | `false` | Mostra diff sem aplicar mudanças |
| `--yes` | `-y` | bool | `false` | Aceita todas as mudanças sem confirmar |
| `--backup` | `-b` | bool | `true` | Cria backup do config.yaml antes de modificar |

### Exemplos

```bash
# Importação completa com confirmação interativa
linkedin-import import

# Importação apenas de experiências
linkedin-import import --experiences ./meu-experiences.csv

# Modo dry-run (visualizar sem aplicar)
linkedin-import import --dry-run

# Aceitar todas as mudanças automaticamente
linkedin-import import --yes

# Caminhos customizados
linkedin-import import -e ./data/Experiences.csv -E ./data/Education.csv -C ./config/portfolio.yaml
```

### Saída

**Sucesso (com mudanças):**
```
📊 Análise de Dados do LinkedIn
════════════════════════════════

✓ Experiências: 5 encontradas
✓ Educação: 2 encontradas
✓ Certificações: 3 encontradas

📋 Comparando com config.yaml atual...

┌─────────────────────────────────────┐
│ Resumo de Mudanças                  │
├─────────────────────────────────────┤
│ ➕ Novas:        2 entradas         │
│ ✏️  Modificadas: 1 entrada          │
│ ➖ Removidas:    0 entradas         │
└─────────────────────────────────────┘

🔄 Detalhes das Mudanças:

Experiências:
  ➕ Adicionar: Software Engineer @ Google
  ➕ Adicionar: Intern @ Microsoft
  ✏️  Modificar: Senior Dev @ Startup (descrição)

Deseja aplicar estas mudanças? (Y/n) >
```

**Sucesso (sem mudanças):**
```
✓ Nenhuma mudança detectada. Config.yaml já está atualizado.
```

**Erro:**
```
✗ Erro: arquivo Experiences.csv não encontrado
```

---

## Comando: `validate`

Valida os arquivos CSV do LinkedIn sem importar.

### Uso

```bash
linkedin-import validate [flags]
```

### Flags

| Flag | Curto | Tipo | Padrão | Descrição |
|------|-------|------|--------|-----------|
| `--experiences` | `-e` | string | `"Experiences.csv"` | Caminho para o CSV de experiências |
| `--education` | `-E` | string | `"Education.csv"` | Caminho para o CSV de educação |
| `--certifications` | `-c` | string | `"Certifications.csv"` | Caminho para o CSV de certificações |

### Exemplos

```bash
# Validar todos os arquivos
linkedin-import validate

# Validar apenas experiências
linkedin-import validate --experiences ./Experiences.csv
```

### Saída

**Sucesso:**
```
✓ Experiences.csv: 5 entradas válidas
✓ Education.csv: 2 entradas válidas
✓ Certifications.csv: 3 entradas válidas
```

**Erro:**
```
✗ Experiences.csv:
   Linha 12: data inválida "Feb 202" (esperado: MMM YYYY)
   Linha 15: campo "Company Name" vazio
```

---

## Comando: `version`

Exibe a versão da ferramenta.

### Uso

```bash
linkedin-import version
```

### Saída

```
linkedin-import version 1.0.0
Go version: go1.21.0
```

---

## Comando: `help`

Exibe ajuda para a CLI ou um comando específico.

### Uso

```bash
linkedin-import help
linkedin-import help [command]
```

### Exemplos

```bash
linkedin-import help
linkedin-import help import
```

---

## Códigos de Saída

| Código | Significado |
|--------|-------------|
| `0` | Sucesso |
| `1` | Erro genérico |
| `2` | Erro de parsing de CSV |
| `3` | Erro de parsing de YAML |
| `4` | Arquivo não encontrado |
| `5` | Permissão negada |
| `130` | Interrompido pelo usuário (Ctrl+C) |

---

## Fluxo de Interatividade

```
┌────────────────────────────────────────┐
│ 1. Executa importação                  │
│ 2. Parseia CSVs                        │
│ 3. Converte datas e descrições         │
│ 4. Compara com config.yaml             │
│ 5. Exibe diff                          │
└──────────────────┬─────────────────────┘
                   │
         ┌─────────▼─────────┐
         │ --dry-run?        │
         └─────────┬─────────┘
         ┌─────────▼─────────┐
         │ Sim → Sai com 0   │
         │ Não → Continua    │
         └─────────┬─────────┘
                   │
         ┌─────────▼─────────┐
         │ --yes?            │
         └─────────┬─────────┘
         ┌─────────▼─────────┐
         │ Sim → Aplica tudo │
         │ Não → Interativo  │
         └───────────────────┘
```

### Modo Interativo

Quando executado sem `--yes`, a CLI apresenta um menu interativo:

```
┌────────────────────────────────────────┐
│ Selecione as mudanças a aplicar:       │
├────────────────────────────────────────┤
│                                        │
│ ➕ Experiências:                       │
│   [✓] Software Engineer @ Google       │
│   [✓] Intern @ Microsoft               │
│                                        │
│ ✏️  Modificações:                      │
│   [✓] Senior Dev @ Startup             │
│                                        │
│                                        │
│   [A] Aplicar selecionadas             │
│   [S] Selecionar tudo                  │
│   [N] Não aplicar nenhuma              │
│   [Q] Cancelar                         │
└────────────────────────────────────────┘
```

Navegação:
- `↑/↓` ou `j/k`: Mover entre opções
- `Space` ou `Enter`: Selecionar/deselecionar
- `A`: Aplicar mudanças selecionadas
- `S`: Selecionar todas
- `N`: Desselecionar todas
- `Q`: Cancelar (sai com código 130)
