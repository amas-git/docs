
class Input {
  constructor(message) {
    this.message  = message;
    this.optional = false;
    this.type = text;
  } 
  
  text() {
    this.type = 'text';
    return;
  }

  number() {
    this.type = 'number';
  }

  password() {
    this.type = 'passord'
  }

  optional() {
    return this;
  }
}

module.exports = {
  question, rl, questionPass, clearRL
};

if (!module.parent) {
  console.log("hello");
}
