const koffi = require('koffi');
const path = require('path');


const extenstion=process.platform === 'win32' ? 'dll' : process.platform === 'darwin' ? 'dylib' : 'so';
const libPath = path.join(__dirname, `libverifier.${extenstion}`);

const lib = koffi.load(libPath);


const Verify = lib.func('int Verify(const char *)');

function verifyEmail(email) {
  return Verify(email);
}

module.exports = {
  verifyEmail
};