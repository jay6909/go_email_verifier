# @anclp/node_email_verifier

[![NPM Version](https://img.shields.io/npm/v/@anclp/node_email_verifier)](https://www.npmjs.com/package/@anclp/node_email_verifier)
[![Build Status](https://github.com/jay6909/node_email_verifier/actions/workflows/release.yml/badge.svg)](https://github.com/jay6909/node_email_verifier/actions/workflows/release.yml)
[![NPM Downloads](https://img.shields.io/npm/dm/@anclp/node_email_verifier)](https://www.npmjs.com/package/@anclp/node_email_verifier)
[![Publish to npm](https://github.com/jay6909/node_email_verifier/actions/workflows/release.yml/badge.svg)](https://github.com/jay6909/node_email_verifier/actions/workflows/release.yml)


Ultra-fast email verification for Node.js, powered by a Go-compiled native engine.

## Installation

```bash
npm install @anclp/node_email_verifier

Why this package?

Most Node.js email validators only check syntax with RegEx. This package uses a native Go binary to perform actual DNS lookups, making it significantly more reliable and faster for bulk verification.
Usage
javascript
const { verifyEmail } = require('@anclp/node_email_verifier');

// Perform a full check (Syntax + MX Records)
const isValid = verifyEmail("test@example.com");

if (isValid) {
    console.log("Email is valid and has active mail servers!");
} else {
    console.log("Invalid email or non-existent domain.");
}






