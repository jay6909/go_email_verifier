const koffi = require('koffi');
const path = require('path');


const extenstion=process.platform === 'win32' ? 'dll' : process.platform === 'darwin' ? 'dylib' : 'so';
const libPath = path.join(__dirname, `libverifier.${extenstion}`);

const lib = koffi.load(libPath);


const Verify = lib.func('int Verify(const char *)');
const VerifyBatch = lib.func('const char * VerifyBatch(const char *)');

function verifyEmail(email) {
  return Verify(email);
}
// 2. Batch verification
function verifyEmailsBatch(emails) {
  // Convert JS array to JSON string for Go
  const jsonInput = JSON.stringify(emails);
  
  // Call the Go native function
  const jsonResult = VerifyBatch(jsonInput);
  
  // Parse the result back to a JS array
  return JSON.parse(jsonResult);
}

module.exports = {
  verifyEmail,verifyEmailsBatch
};