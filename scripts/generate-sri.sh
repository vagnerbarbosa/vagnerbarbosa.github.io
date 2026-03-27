#!/bin/bash

# Script para gerar hashes SRI (Subresource Integrity) para arquivos JS
# Uso: ./scripts/generate-sri.sh

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
JS_DIR="$SCRIPT_DIR/../assets/js"

echo "=========================================="
echo "Gerando hashes SRI para arquivos JS..."
echo "=========================================="
echo ""

if [ ! -d "$JS_DIR" ]; then
    echo "Erro: Diretório $JS_DIR não encontrado!"
    exit 1
fi

# Cria arquivo de saída
OUTPUT_FILE="$SCRIPT_DIR/sri-hashes.txt"
echo "Hashes SRI gerados em: $(date)" > "$OUTPUT_FILE"
echo "" >> "$OUTPUT_FILE"

# Gera hashes para cada arquivo JS
for file in "$JS_DIR"/*.js; do
    if [ -f "$file" ]; then
        filename=$(basename "$file")
        hash=$(openssl dgst -sha384 -binary "$file" | openssl base64 -A)

        echo "Arquivo: $filename"
        echo "Hash: sha384-$hash"
        echo ""

        # Salva no arquivo de saída
        echo "<!-- $filename -->" >> "$OUTPUT_FILE"
        echo "integrity=\"sha384-$hash\"" >> "$OUTPUT_FILE"
        echo "" >> "$OUTPUT_FILE"
    fi
done

echo "=========================================="
echo "Hashes salvos em: $OUTPUT_FILE"
echo "=========================================="
echo ""
echo "IMPORTANTE: Copie os hashes para o _layouts/default.html"
echo ""
echo "Exemplo de uso:"
echo '<script src="/assets/js/utils.js"'
echo '        integrity="sha384-..."'
echo '        crossorigin="anonymous"'
echo '        defer></script>'
