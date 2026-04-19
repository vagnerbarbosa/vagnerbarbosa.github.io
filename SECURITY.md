# Security Policy

## Supported Versions

This is a static personal website, and we prioritize security in all deployments. The following versions are actively supported with security updates:

| Version | Supported          |
| ------- | ------------------ |
| 4.x.x   | :white_check_mark: |
| 3.x.x   | :x:                |
| 2.x.x   | :x:                |
| 1.x.x   | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability in this project, please report it responsibly:

1. **Do NOT** open a public issue
2. Send an email to: **security@vagnerbarbosa.com**
3. Include details about the vulnerability and steps to reproduce
4. Allow reasonable time for response before public disclosure

## Security Measures

This project implements the following security practices:

### Content Security
- HTML templates use Go's `html/template` package with auto-escaping to prevent XSS
- No `safeHTML` or raw HTML injection from user-provided content
- Input sanitization on all rendered content

### CI/CD Security
- GitHub Actions workflows run with minimum required permissions (principle of least privilege)
- No write permissions where only read is needed
- Secure polling mechanisms with backoff instead of fixed delays

### Dependencies
- Regular updates to dependencies via Dependabot
- Minimal dependency surface (Go standard library + minify)
- No JavaScript runtime dependencies in production

### Headers
The following security headers are configured:

- **Content-Security-Policy (CSP)**: Restricts content sources
- **X-Frame-Options: DENY**: Prevents clickjacking
- **X-Content-Type-Options: nosniff**: Prevents MIME type sniffing
- **Referrer-Policy**: Controls referrer information

## Acknowledgments

We appreciate responsible disclosure of security issues. Contributors who report valid vulnerabilities will be acknowledged (with permission).

---

Last updated: April 2026
