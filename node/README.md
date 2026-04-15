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


## Batch Verification ⚡

For high-volume processing, use the `verifyEmailsBatch` method. This uses Go's native goroutines to verify thousands of emails in parallel, bypassing the overhead of the Node.js event loop.

### Usage

```javascript
const { verifyEmailsBatch } = require('@anclp/emailverifier');

const emails = [
  "valid@gmail.com",
  "typo@gmailll.com",
  "fake@nonexistent-domain.org"
];

const results = verifyEmailsBatch(emails);

// results is an array of booleans: [true, false, false]
results.forEach((isValid, index) => {
  console.log(`${emails[index]}: ${isValid ? '✅ Valid' : '❌ Invalid'}`);
});


Performance Metrics
On a standard consumer laptop (e.g., Asus TUF), this engine achieves:
Single Check: ~4.6ms
Batch Throughput: 2,800+ emails/second (verified via network MX lookups).
[!TIP]
Use Batch Verification when processing lists larger than 50 emails to maximize throughput.



