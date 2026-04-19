#!/bin/bash

# Script para gerar changelog em português no formato Keep a Changelog
# Usage: ./generate-changelog.sh <versao> <sha-anterior>

set -e

VERSION="${1:-v3.0.0}"
PREVIOUS_SHA="${2:-}"

# Se não houver SHA anterior, pegar a última tag
if [ -z "$PREVIOUS_SHA" ]; then
  PREVIOUS_SHA=$(git describe --tags --abbrev=0 2>/dev/null || echo "")
fi

# Configurar range de commits
if [ -n "$PREVIOUS_SHA" ]; then
  COMMIT_RANGE="${PREVIOUS_SHA}..HEAD"
else
  COMMIT_RANGE="HEAD"
fi

# Gerar changelog temporário
CHANGELOG=$(mktemp)

# Cabeçalho
echo "## [${VERSION}] - $(date +%Y-%m-%d)" > "$CHANGELOG"
echo "" >> "$CHANGELOG"

# Mapeamento de tipos para seções em português
declare -A SECTIONS=(
  ["feat"]="### Adicionado"
  ["fix"]="### Corrigido"
  ["docs"]="### Documentação"
  ["style"]="### Estilo"
  ["refactor"]="### Alterado"
  ["perf"]="### Performance"
  ["test"]="### Testes"
  ["chore"]="### Manutenção"
  ["build"]="### Build"
  ["ci"]="### CI/CD"
  ["revert"]="### Revertido"
)

# Extrair commits por tipo
for type in feat fix docs style refactor perf test chore build ci revert; do
  section="${SECTIONS[$type]}"
  commits=$(git log "$COMMIT_RANGE" --grep="^$type" --pretty=format:"- %s" 2>/dev/null || true)

  # Limpar prefixo do commit (feat: , fix: , etc)
  cleaned_commits=$(echo "$commits" | sed "s/^- $type\(!\)\?:\s*/- /" | sed "s/^- $type(\([^)]*\))\(!\)\?:\s*/- /" || true)

  if [ -n "$cleaned_commits" ]; then
    echo "$section" >> "$CHANGELOG"
    echo "" >> "$CHANGELOG"
    echo "$cleaned_commits" >> "$CHANGELOG"
    echo "" >> "$CHANGELOG"
  fi
done

# Verificar breaking changes
breaking=$(git log "$COMMIT_RANGE" --grep="BREAKING CHANGE" --pretty=format:"- %s" 2>/dev/null || true)
if [ -n "$breaking" ]; then
  echo "### Quebra de Compatibilidade" >> "$CHANGELOG"
  echo "" >> "$CHANGELOG"
  echo "$breaking" >> "$CHANGELOG"
  echo "" >> "$CHANGELOG"
fi

# Adicionar link de comparação se houver tag anterior
if [ -n "$PREVIOUS_SHA" ]; then
  echo "[${VERSION}]: https://github.com/vagnerbarbosa/vagnerbarbosa.github.io/compare/${PREVIOUS_SHA}...${VERSION}" >> "$CHANGELOG"
fi

# Output final
cat "$CHANGELOG"

# Limpar arquivo temporário
rm -f "$CHANGELOG"
