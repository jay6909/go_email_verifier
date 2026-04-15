[![Build cross platform binaries and sync to repos](https://github.com/jay6909/go_email_verifier/actions/workflows/release.yml/badge.svg)](https://github.com/jay6909/go_email_verifier/actions/workflows/release.yml)

# Go Email Verifier (Core)

A high-performance email verification engine written in Go. This repo serves as the core logic for multiple language wrappers (Node.js, Python, etc.).

## Why Go?
Go's networking stack and concurrency allow us to perform DNS (MX record) lookups and SMTP handshakes much faster than interpreted languages.

## Features
- **MX Record Lookup**: Checks if the domain actually has mail servers.
- **Syntax Validation**: RFC-compliant email parsing.
- **Cross-Platform**: Compiled into shared libraries (.so, .dll, .dylib).
- **Zero-Dependency**: Built primarily using the Go standard library.

## Project Structure
- `/emailverifier`: Core Go logic.
- `main.go`: The C-Bridge for cross-language support.
- `.github/workflows`: Automated build and sync pipeline.
