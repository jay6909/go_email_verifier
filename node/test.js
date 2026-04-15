const { verifyEmail } = require(".");


//check good email
console.log(verifyEmail("test@example.com"));

//check bad email
console.log(verifyEmail("test@exampasase.com"));