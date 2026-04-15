# @anclp/emailverifier

[![NPM Version](https://shields.io)](https://npmjs.com)
[![Build Status](https://github.com)](https://github.com)
[![NPM Downloads](https://shields.io)](https://npmjs.com)

[![Publish to npm](https://github.com/jay6909/node_email_verifier/actions/workflows/release.yml/badge.svg)](https://github.com/jay6909/node_email_verifier/actions/workflows/release.yml)
# @anclp/emailverifier

Ultra-fast email verification for Node.js, powered by a Go-compiled native engine.

## Installation

```bash
npm install @anclp/emailverifier

Why this package?

Most Node.js email validators only check syntax with RegEx. This package uses a native Go binary to perform actual DNS lookups, making it significantly more reliable and faster for bulk verification.
Usage
javascript
const { verifyEmail } = require('@anclp/emailverifier');

// Perform a full check (Syntax + MX Records)
const isValid = verifyEmail("test@example.com");

if (isValid) {
    console.log("Email is valid and has active mail servers!");
} else {
    console.log("Invalid email or non-existent domain.");
}






