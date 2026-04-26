# Quickstart: LinkedIn Import CLI

**Feature**: LinkedIn Import CLI  
**Data**: 2026-04-26

---

## Instalação

### Requisitos

- Go 1.21 ou superior
- Acesso aos arquivos CSV exportados do LinkedIn

### Build

```bash
# Clone o repositório
git clone https://github.com/vagnerbarbosa/vagnerbarbosa.github.io.git
cd vagnerbarbosa.github.io

# Instale a ferramenta
go install ./cmd/linkedin-import

# Ou compile localmente
go build -o linkedin-import ./cmd/linkedin-import
```

---

## Configuração Inicial

### 1. Exportar Dados do LinkedIn

1. Acesse [linkedin.com](https://linkedin.com) e faça login
2. Vá em **Settings & Privacy** > **Data privacy** > **Get a copy of your data**
3. Selecione **"Want something in particular?"**
4. Escolha:
   - Connections (opcional)
   - Profile
   - **Salve os arquivos CSV em um diretório de trabalho**

### 2. Preparar Ambiente

```bash
# Crie um diretório de trabalho
mkdir ~/linkedin-export
cd ~/linkedin-export

# Coloque os arquivos CSV exportados aqui
# Experiences.csv, Education.csv, Certifications.csv
```

---

## Uso Básico

### Primeira Execução (Modo Dry-Run)

Sempre execute primeiro em modo dry-run para visualizar as mudanças:

```bash
linkedin-import import --dry-run
```

### Importação Completa

```bash
# Com confirmação interativa (recomendado)
linkedin-import import

# Aceitar todas as mudanças automaticamente
linkedin-import import --yes

# Com caminhos customizados
linkedin-import import \
  --experiences ./data/Experiences.csv \
  --education ./data/Education.csv \
  --config ./config/portfolio.yaml
```

### Validação de Arquivos

Antes de importar, valide os arquivos CSV:

```bash
linkedin-import validate
```

---

## Exemplos

### Exemplo 1: Importar apenas experiências

```bash
linkedin-import import --education "" --certifications ""
```

### Exemplo 2: Importação automática em script

```bash
#!/bin/bash
# update-portfolio.sh

set -e

# Download do export mais recente do LinkedIn
# (adapte conforme seu método de download)

# Importa automaticamente
linkedin-import import --yes --backup

echo "Portfolio atualizado com sucesso!"
```

### Exemplo 3: Verificar antes de aplicar

```bash
# Salva o diff em um arquivo
linkedin-import import --dry-run > diff.txt

# Revisa manualmente
cat diff.txt

# Se aprovado, aplica
linkedin-import import --yes
```

---

## Desenvolvimento

### Executar Testes

```bash
# Todos os testes
go test ./...

# Com cobertura
go test -cover ./...

# Testes específicos do módulo de importação
go test ./cmd/linkedin-import/...
```

### Estrutura do Projeto

```
cmd/linkedin-import/
├── main.go              # Entry point
├── commands/
│   ├── import.go        # Comando de importação
│   ├── validate.go        # Comando de validação
│   └── root.go           # Comandos root e help
├── internal/
│   ├── parser/
│   │   ├── csv.go        # Parsing de CSV
│   │   └── date.go       # Conversão de datas
│   ├── transformer/
│   │   └── description.go # Split em bullets
│   ├── comparator/
│   │   └── diff.go       # Lógica de diff
│   ├── config/
│   │   └── yaml.go       # Manipulação de YAML
│   └── ui/
│       ├── confirm.go    # Prompts interativos
│       └── diff.go       # Visualização de diff
```

### Adicionar Novo Comando

Para adicionar um novo comando à CLI:

1. Crie um arquivo em `cmd/linkedin-import/commands/`
2. Implemente a interface `Command`
3. Registre no `root.go`

Exemplo:

```go
// cmd/linkedin-import/commands/export.go
package commands

import "github.com/spf13/cobra"

var exportCmd = &cobra.Command{
    Use:   "export",
    Short: "Exporta dados do config.yaml para CSV",
    RunE: func(cmd *cobra.Command, args []string) error {
        // Implementação
        return nil
    },
}

func init() {
    rootCmd.AddCommand(exportCmd)
}
```

---

## Troubleshooting

### Erro: "Experiences.csv não encontrado"

**Solução**: Especifique o caminho correto:

```bash
linkedin-import import --experiences /caminho/completo/Experiences.csv
```

### Erro: "encoding incorreto no CSV"

**Solução**: O LinkedIn exporta em UTF-8. Se houver problemas, verifique a codificação:

```bash
file Experiences.csv
# Deve mostrar: UTF-8 Unicode text
```

### Erro: "config.yaml já existe"

Isso é esperado! A ferramenta faz backup automaticamente. Se quiser sobrescrever:

```bash
# Remove o backup anterior se necessário
rm config.yaml.backup.*

# Ou use --backup=false (não recomendado)
linkedin-import import --backup=false
```

### Datas não estão sendo convertidas

Verifique o formato das datas no CSV. Deve estar como: "Feb 2025" ou "Jan 2020 - Present"

---

## Dicas

1. **Sempre use --dry-run primeiro**: Visualize antes de aplicar
2. **Mantenha backups**: A ferramenta faz backup automaticamente, mas mantenha seu próprio versionamento
3. **Valide antes**: Use `validate` para verificar erros no CSV
4. **Edite manualmente se necessário**: Nem toda descrição será perfeitamente dividida em bullets - revise e ajuste manualmente

---

## Referências

- [Especificação completa](./spec.md)
- [Modelo de dados](./data-model.md)
- [Contratos da CLI](./contracts/cli-commands.md)
- [Pesquisa técnica](./research.md)
