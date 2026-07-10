
const { performance } = require("perf_hooks");
const { verifyEmailsBatch } = require("@anclp/node_email_verifier")

const count = 10000;
const emails = Array.from({ length: count }, (_, i) => `test${i}@gmail.com`);

console.log(`--- Starting Batch Benchmark: ${count} emails ---`);

const start = performance.now();
const results = verifyEmailsBatch(emails);
const end = performance.now();

const durationInSeconds = (end - start) / 1000;
const throughput = Math.round(count / durationInSeconds);

console.log(`Batch Duration: ${(end - start).toFixed(3)}ms`);
console.log(`Throughput: ${throughput} emails/sec`);