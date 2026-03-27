# Scripts de Segurança

Este diretório contém scripts para manutenção de segurança do site.

---

## 🔐 SRI (Subresource Integrity)

Os hashes SRI garantem que os arquivos JavaScript não foram modificados no servidor.

### Quando regenerar os hashes?

- ✅ Após modificar qualquer arquivo em `assets/js/`
- ✅ Após atualizar dependências
- ✅ Antes de deploy em produção

### Como regenerar

**Windows (PowerShell):**
```powershell
cd scripts
.\generate-sri.ps1
```

**Linux/Mac/Git Bash:**
```bash
cd scripts
chmod +x generate-sri.sh
./generate-sri.sh
```

O script irá gerar um arquivo `sri-hashes.txt` com os hashes atualizados.

### Como atualizar o HTML

Copie os hashes gerados e atualize o arquivo `_layouts/default.html`:

```html
<script src="/assets/js/utils.js"
        integrity="sha384-NOVO_HASH_AQUI"
        crossorigin="anonymous"
        defer></script>
```

**IMPORTANTE:** Se o hash não corresponder ao arquivo real, o navegador **BLOQUEARÁ** o script!

---

## 📋 Checklist de Segurança

Antes de cada deploy:

- [ ] Executar script de geração de hashes SRI
- [ ] Atualizar `_layouts/default.html` com novos hashes
- [ ] Testar se o site funciona corretamente
- [ ] Verificar console do navegador por erros de SRI

---

## 🛡️ Notas de Segurança

### CSP (Content Security Policy)

A CSP está configurada no `_layouts/default.html` via meta tag. Limitações do GitHub Pages:

- ✅ Funciona: CSP via meta tag
- ❌ Não funciona: Headers HTTP customizados

Para headers HTTP completos, use **Cloudflare** como proxy (veja `SECURITY_AUDIT.md` na raiz).

### Validação de Hashes

Você pode verificar um hash manualmente:

```bash
openssl dgst -sha384 -binary assets/js/utils.js | openssl base64 -A
```

O resultado deve ser idêntico ao hash no HTML.
