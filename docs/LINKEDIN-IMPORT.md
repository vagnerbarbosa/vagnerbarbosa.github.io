# LinkedIn Import CLI

Ferramenta CLI para importar dados do LinkedIn (experiências, educação e certificações) a partir do export CSV manual.

## Visão Geral

Esta ferramenta permite importar seu histórico profissional do LinkedIn diretamente para o arquivo `config.yaml` do seu portfólio, eliminando a necessidade de digitação manual.

## Funcionalidades

- ✅ Importação de experiências profissionais
- ✅ Importação de formação acadêmica
- ✅ Importação de certificações
- ✅ Conversão automática de datas (inglês → português)
- ✅ Divisão de descrições em bullets
- ✅ **Extração automática de tech stack** - Detecta tech stack em bullets e separa no campo `tech_stack`
- ✅ Diff visual colorido das mudanças
- ✅ Confirmação interativa das alterações
- ✅ Backup automático do config.yaml
- ✅ Modo dry-run para visualização sem aplicação

## Pré-requisitos

1. **Exportar dados do LinkedIn**:
   - Acesse: LinkedIn → Configurações → Conta → Gerenciar seus dados → Baixar seus dados
   - Selecione "Quer uma cópia adicional?" → Solicitar arquivo
   - Aguarde o email de confirmação (pode levar alguns minutos a 24h)
   - Baixe o arquivo ZIP e extraia os CSVs

2. **Localizar os arquivos CSV**:
   - `Positions.csv` (experiências profissionais)
   - `Education.csv` (formação acadêmica)
   - `Certifications.csv` (certificações)

## Instalação

A ferramenta já está incluída no projeto. Não requer instalação adicional.

## Uso

### Comando Básico

```bash
go run cmd/import-linkedin/main.go import
```

### Opções Disponíveis

```bash
# Importar com confirmação interativa (recomendado)
go run cmd/import-linkedin/main.go import

# Importar automaticamente sem confirmação (--yes)
go run cmd/import-linkedin/main.go import --yes

# Modo dry-run (visualizar sem aplicar)
go run cmd/import-linkedin/main.go import --dry-run

# Especificar caminho dos CSVs
./linkedin-import import --experiences /caminho/Positions.csv --education /caminho/Education.csv

# Validar CSVs antes de importar
./linkedin-import validate

# Mostrar versão
./linkedin-import version
```

## Fluxo de Uso

### 1. Preparar os Dados

Extraia o ZIP do LinkedIn e localize os arquivos CSV:

```
Downloads/
├── Positions.csv          # Experiências profissionais
├── Education.csv          # Formação acadêmica
└── Certifications.csv     # Certificações
```

### 2. Executar a Importação

```bash
cd /home/vagner-barbosa/Documentos/DevZone/vagnerbarbosa.github.io
go run cmd/import-linkedin/main.go import
```

### 3. Revisar o Diff

A ferramenta mostrará um diff colorido:

```diff
+ Nova experiência: Zup Innovation
+ Cargo: Desenvolvedor Backend Especialista
+ Período: Fev 2025 – Presente

~ Experiência modificada: Invillia
~ Cargo alterado: Analista → Tech Lead

- Experiência removida: Empresa X
```

### 4. Confirmar Mudanças

Para cada alteração, você pode:
- **Aceitar** (`y`): Aplica a mudança
- **Rejeitar** (`n`): Mantém o valor atual
- **Aceitar todas** (`a`): Aplica todas as mudanças restantes
- **Rejeitar todas** (`r`): Mantém todos os valores atuais

### 5. Verificar o Backup

Antes de modificar o `config.yaml`, a ferramenta cria automaticamente um backup:

```
config.yaml.backup.2026-04-26T14-30-00
```

## Tratamento de Dados

### Conversão de Datas

| Inglês (LinkedIn) | Português (config.yaml) |
|-------------------|-------------------------|
| Jan 2025          | Jan 2025                |
| Feb 2025          | Fev 2025                |
| Mar 2025          | Mar 2025                |
| Apr 2025          | Abr 2025                |
| May 2025          | Mai 2025                |
| Jun 2025          | Jun 2025                |
| Jul 2025          | Jul 2025                |
| Aug 2025          | Ago 2025                |
| Sep 2025          | Set 2025                |
| Oct 2025          | Out 2025                |
| Nov 2025          | Nov 2025                |
| Dec 2025          | Dez 2025                |
| Present           | Presente                |

### Divisão de Descrições

Descrições longas em texto corrido são automaticamente divididas em bullets:

**Entrada (CSV):**
```
"Como Especialista Backend, atuando em diversas frentes. Otimização de FinOps: Referência técnica na evolução do KPI. Liderança Técnica: Gestão direta de mais de 7 desenvolvedores."
```

**Saída (config.yaml):**
```yaml
details:
  - "Como Especialista Backend, atuando em diversas frentes"
  - "Otimização de FinOps: Referência técnica na evolução do KPI"
  - "Liderança Técnica: Gestão direta de mais de 7 desenvolvedores"
```

### Extração Automática de Tech Stack

Quando um bullet de descrição contém padrões de tech stack, a ferramenta detecta e extrai automaticamente:

**Entrada (CSV):**
```
"Referência técnica em FinOps para o Itaú\nGestão de 7+ desenvolvedores\nAs principais tecnologias e ferramentas utilizadas: Java, Python, AWS"
```

**Saída (config.yaml):**
```yaml
details:
  - "Referência técnica em FinOps para o Itaú"
  - "Gestão de 7+ desenvolvedores"
tech_stack: "Java • Python • AWS"
```

#### Padrões Detectados

A ferramenta reconhece os seguintes padrões (case-insensitive):
- `As principais tecnologias e ferramentas utilizadas:`
- `Tecnologias:` / `Technologies:`
- `Tech Stack:` / `Stack:`
- `Ferramentas:` / `Tools:`

#### Formatos Suportados

Tecnologias podem ser separadas por:
- Vírgula: `Java, Python, AWS`
- Pipe: `Java | Python | AWS`
- Hífen: `Java - Python - AWS`
- Bullet: `• Java • Python • AWS`
- Ponto-e-vírgula: `Java; Python; AWS`

O tech stack extraído é formatado com o separador ` • ` para consistência no portfólio.

### Identificação de Duplicatas

A ferramenta detecta entradas duplicadas (mesma empresa + cargo + data de início) e alerta o usuário com opções:
- **Mesclar**: Combina as descrições
- **Ignorar**: Mantém apenas uma
- **Manter ambas**: Importa como entradas separadas

## Casos de Uso

### Atualização Mensal

```bash
# Baixar novo export do LinkedIn
# Extrair CSVs
./linkedin-import import --yes
```

### Primeira Importação

```bash
# Recomendado: usar modo interativo
./linkedin-import import

# Revisar cuidadosamente cada mudança
# Confirmar apenas as alterações desejadas
```

### Teste Antes de Aplicar

```bash
# Visualizar mudanças sem aplicar
./linkedin-import import --dry-run

# Se satisfeito, executar de verdade
./linkedin-import import
```

## Solução de Problemas

### Erro: "Arquivo CSV não encontrado"

**Causa**: Os CSVs não estão no diretório esperado.

**Solução**: Especifique o caminho completo:
```bash
./linkedin-import import --experiences ~/Downloads/Positions.csv
```

### Erro: "config.yaml não encontrado"

**Causa**: Não está no diretório raiz do projeto.

**Solução**: Execute a partir da raiz do projeto:
```bash
cd /home/vagner-barbosa/Documentos/DevZone/vagnerbarbosa.github.io
go run cmd/import-linkedin/main.go import
```

### Descrições não foram divididas corretamente

**Causa**: Texto não contém pontuação ou quebras de linha claras.

**Solução**: Edite manualmente o `config.yaml` após importação.

### Datas não convertidas

**Causa**: Formato de data não reconhecido no CSV.

**Solução**: Verifique se o CSV está no formato padrão do LinkedIn (ex: "Feb 2025").

## Dicas

1. **Sempre faça backup**: Embora a ferramenta crie backup automaticamente, mantenha seu próprio controle de versão.

2. **Use dry-run primeiro**: Especialmente na primeira vez, execute com `--dry-run` para entender as mudanças.

3. **Revisão humana**: A ferramenta facilita a importação, mas revise o resultado final no `config.yaml`.

4. **Mantenha o export do LinkedIn atualizado**: Faça novos exports periodicamente para refletir mudanças recentes.

## Referência de Comandos

| Comando | Descrição |
|---------|-----------|
| `import` | Importa dados dos CSVs para o config.yaml |
| `validate` | Valida os arquivos CSV sem importar |
| `version` | Mostra versão da ferramenta |
| `--yes` | Aceita todas as mudanças automaticamente |
| `--dry-run` | Simula importação sem modificar arquivos |
| `--help` | Mostra ajuda detalhada |

## Veja Também

- [Especificação](../specs/003-linkedin-import/spec.md) - Detalhes técnicos da funcionalidade
- [README principal](../README.md) - Informações gerais do projeto
