# Script para gerar hashes SRI (Subresource Integrity) para arquivos JS
# Uso: .\scripts\generate-sri.ps1

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$JsDir = Join-Path $ScriptDir "..\assets\js"

Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "Gerando hashes SRI para arquivos JS..." -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""

if (-not (Test-Path $JsDir)) {
    Write-Error "Erro: Diretório $JsDir não encontrado!"
    exit 1
}

# Cria arquivo de saída
$OutputFile = Join-Path $ScriptDir "sri-hashes.txt"
"Hashes SRI gerados em: $(Get-Date)" | Out-File -FilePath $OutputFile
"" | Out-File -FilePath $OutputFile -Append

# Gera hashes para cada arquivo JS
Get-ChildItem -Path "$JsDir\*.js" | ForEach-Object {
    $file = $_.FullName
    $filename = $_.Name

    # Gera hash SHA384
    $hash = Get-FileHash -Path $file -Algorithm SHA384 | Select-Object -ExpandProperty Hash

    # Converte para base64 (necessário para SRI)
    $bytes = [byte[]] -split ($hash -replace '..', '0x$& ')
    $base64 = [Convert]::ToBase64String($bytes)

    Write-Host "Arquivo: " -NoNewline
    Write-Host $filename -ForegroundColor Yellow
    Write-Host "Hash: sha384-$base64"
    Write-Host ""

    # Salva no arquivo de saída
    "<!-- $filename -->" | Out-File -FilePath $OutputFile -Append
    "integrity=`"sha384-$base64`"" | Out-File -FilePath $OutputFile -Append
    "" | Out-File -FilePath $OutputFile -Append
}

Write-Host "==========================================" -ForegroundColor Green
Write-Host "Hashes salvos em: $OutputFile" -ForegroundColor Green
Write-Host "==========================================" -ForegroundColor Green
Write-Host ""
Write-Host "IMPORTANTE: Copie os hashes para o _layouts/default.html" -ForegroundColor Yellow
