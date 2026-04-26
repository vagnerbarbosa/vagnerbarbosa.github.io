# Modelo de Dados: LinkedIn Import CLI

**Feature**: LinkedIn Import CLI  
**Data**: 2026-04-26

---

## Diagrama de Entidades

```
┌─────────────────────────────────────────────────────────────────┐
│                          LinkedInCSV                           │
├─────────────────────────────────────────────────────────────────┤
│ Experiences.csv ──→ []Experience                                │
│ Education.csv ────→ []Education                                 │
│ Certifications.csv → []Certification                            │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                    LinkedInImporter (Service)                   │
├─────────────────────────────────────────────────────────────────┤
│ - ParseCSV(filepath, type) → []Entity                          │
│ - ConvertDates(dateStr) → dateStr (pt-BR)                      │
│ - SplitDescription(text) → []string (bullets)                │
│ - CompareWithConfig(entities, config) → Diff                 │
│ - ApplyChanges(diff, config) → Config                        │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                            Diff                                │
├─────────────────────────────────────────────────────────────────┤
│ Added: []Entity         - Novas entradas                        │
│ Modified: []Change      - Entradas alteradas                    │
│ Removed: []Entity       - Entradas removidas                    │
└─────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      ConfigPortfolio                           │
├─────────────────────────────────────────────────────────────────┤
│ experiences: []Experience                                      │
│ education: []Education                                         │
│ certifications: []Certification                                │
└─────────────────────────────────────────────────────────────────┘
```

---

## Entidades

### Experience (Experiência Profissional)

Representa uma experiência profissional importada do LinkedIn.

| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| `company` | string | Sim | Nome da empresa empregadora |
| `role` | string | Sim | Cargo ocupado |
| `start_date` | string | Sim | Data de início (formato: "Jan 2020" ou "Fev 2020" em pt) |
| `end_date` | string | Não | Data de término ou "Presente" |
| `description` | []string | Não | Lista de bullets com descrição da experiência |
| `location` | string | Não | Localização (cidade/país) |

**Regras de Validação:**
- `company` não pode ser vazio ou conter apenas espaços
- `role` não pode ser vazio
- `start_date` deve estar em formato válido (MMM YYYY)
- `end_date` pode estar vazio (indica posição atual)

---

### Education (Educação)

Representa uma formação acadêmica importada do LinkedIn.

| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| `institution` | string | Sim | Nome da instituição de ensino |
| `degree` | string | Sim | Grau acadêmico (ex: "Bachelor's degree") |
| `field` | string | Não | Área de estudo (ex: "Computer Science") |
| `start_date` | string | Não | Data de início (formato: "Jan 2020") |
| `end_date` | string | Não | Data de conclusão |
| `description` | []string | Não | Lista de bullets com detalhes adicionais |

**Regras de Validação:**
- `institution` não pode ser vazio
- `degree` não pode ser vazio
- Pelo menos uma das datas (start ou end) deve estar presente

---

### Certification (Certificação)

Representa uma certificação profissional importada do LinkedIn.

| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| `name` | string | Sim | Nome da certificação |
| `organization` | string | Sim | Organização emissora |
| `issue_date` | string | Sim | Data de emissão |
| `expiration_date` | string | Não | Data de expiração (se aplicável) |
| `credential_id` | string | Não | ID da credencial |
| `credential_url` | string | Não | URL de verificação |

**Regras de Validação:**
- `name` não pode ser vazio
- `organization` não pode ser vazio
- `issue_date` deve estar em formato válido

---

### Change (Mudança Detectada)

Representa uma diferença entre dados do LinkedIn e config.yaml atual.

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `entity_type` | string | Tipo da entidade (experience, education, certification) |
| `entity_id` | string | Identificador único (hash ou composite key) |
| `old_value` | interface{} | Valor atual no config.yaml |
| `new_value` | interface{} | Valor importado do LinkedIn |
| `change_type` | string | Tipo: "added", "modified", "removed" |
| `fields_changed` | []string | Lista de campos específicos modificados |

---

## Estruturas Auxiliares

### ConfigPortfolio

Estrutura completa do config.yaml do portfólio.

```yaml
title:
  en: string
  pt: string
name: string
about:
  en: string
  pt: string
social: Social
experiences: []Experience
certifications: []Certification
education: []Education
```

### Social

Links para redes sociais.

| Campo | Tipo |
|-------|------|
| `linkedin` | string (URL) |
| `github` | string (URL) |
| `youtube` | string (URL) |

---

## Identificação de Entidades

Para detectar duplicatas e alterações, cada entidade deve ter um identificador único:

### Experience
- Chave composta: `company + "#" + role`
- Exemplo: `"Anthropic#Senior Software Engineer"`

### Education
- Chave composta: `institution + "#" + degree + "#" + field`
- Exemplo: `"UFMG#Bacharel#Ciência da Computação"`

### Certification
- Chave composta: `name + "#" + organization`
- Exemplo: `"AWS Solutions Architect#Amazon Web Services"`

---

## Transições de Estado

### Fluxo de Importação

```
┌─────────┐    ┌──────────┐    ┌───────────┐    ┌──────────┐    ┌─────────┐
│  CSV    │───→│  Parse   │───→│ Convert   │───→│ Compare  │───→│ Confirm │
│ LinkedIn│    │  CSV     │    │ + Split   │    │  Diff    │    │ Apply   │
└─────────┘    └──────────┘    └───────────┘    └──────────┘    └─────────┘
                                                                     │
                                                                     ▼
                                                                ┌─────────┐
                                                                │ Config  │
                                                                │ Updated │
                                                                └─────────┘
```

### Estados da Mudança

- `pending`: Mudança detectada, aguardando confirmação
- `accepted`: Usuário confirmou, será aplicada
- `rejected`: Usuário rejeitou, será ignorada
- `applied`: Mudança já foi escrita no config.yaml

---

## Schema do CSV do LinkedIn

### Experiences.csv (colunas relevantes)

| Coluna | Mapeamento |
|--------|------------|
| `Company Name` | `company` |
| `Title` | `role` |
| `Started On` | `start_date` (requer conversão) |
| `Finished On` | `end_date` (requer conversão) |
| `Description` | `description` (requer split em bullets) |
| `Location` | `location` |

### Education.csv (colunas relevantes)

| Coluna | Mapeamento |
|--------|------------|
| `School Name` | `institution` |
| `Degree Name` | `degree` |
| `Field Of Study` | `field` |
| `Started On` | `start_date` |
| `Finished On` | `end_date` |
| `Description` | `description` |

### Certifications.csv (colunas relevantes)

| Coluna | Mapeamento |
|--------|------------|
| `Certification Name` | `name` |
| `Certification Authority` | `organization` |
| `Started On` | `issue_date` |
| `Finished On` | `expiration_date` |
