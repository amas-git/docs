const readline = require('readline');
const _ = require('lodash');

function dup(ch, length) {
  const xs = [];
  for (let i=0; i<length; i++) {
    xs.push(ch);
  }
  return xs.join('');
}

class Input {
  constructor(key, text, parser) {
    this._key     = key;
    this._message = text;
    this._parser  = parser || ((text) => text);
    this._opt     = false;
    this._mask    = undefined;
    this._filter  = undefined;
    this._validator = null;
  } 

  parse(text, ...args) {
    return this._parser(text, ...args);
  }

  // TODO: implement optional
  opt() {
    this.opt = true;
    return this;
  }

  mask(ch) {
    this._mask  = ch;
    return this;
  }

  filter(regexp=/./) {
    this._filter = regexp;
  }

  validator(validator) {
    this._validator = validator;
    return this;
  }
}

class ConsoleForm {
  static REGEX_PASSWORD = /[a-zA-Z0-9!"#$%&'()*+,.\/:;<=>?@\[\]^_`{|}~-]/;
  constructor() {
    this.inputs = [];
    this.value = {};
  }

  async _question(text, mask, filter) {
    let rl = readline.createInterface({
      input: process.stdin,
      output: process.stdout,
      terminal: true
    });

    return new Promise((resolve, reject) => {
      rl.question(text, (a) => {
        rl.history = rl.history.slice(1);
        rl.close();
        resolve(_.trim(a));
      });
      
      if (mask) {
        rl._writeToOutput = function (_) {
          let out = '';
          if (_.startsWith(text)) {
            out = text + dup(mask, _.length - text.length);
          } else {
            if (filter && filter.test(_)) {
              out = mask;
            } else {
              // done
              rl.output.write('\n');
              return;
            }
          }
          rl.output.write(out);
        };
      }
    });
  }

  text(key, prompt, parser) {
    let input = new Input(key, prompt, parser);
    this.inputs.push(input);
    return input;
  }

  int53(key, prompt) {
    return this.text(key, prompt, (text) => {
      let n = Number.parseInt(text);
      return Number.isNaN(n) ? undefined : n;
    });
  }

  password(key, prompt, dict = ConsoleForm.REGEX_PASSWORD) {
    return this.text(key, prompt).mask('*').filter(dict);
  }
  
  async commit() {
    for (let input of this.inputs) {
      let text = null;
      while (true) {
        text = await this._question(input._message, input._mask, input._filter);
        if (_.isEmpty(text)) {
          continue;
        }

        let value = input.parse(text);
        if (value === null || value === undefined) {
          this.error(`Type Error`);
          continue;
        }

        if (input._validator) {
          const accepted = !!input._validator(value)
          if (!accepted) {
            continue;
          }
        }
        this.value[input._key] = value;
        break;
      }
    }
    return this.value;
  }

  error(msg) {
    // TODO: implement error message
  }
}

module.exports = { ConsoleForm }

if (!module.parent) {
  (async () => {
    let form = new ConsoleForm();
    form.text('name', 'Enter Name:');
    form.int53('age', 'Enter Age:').validator(n => n > 0);
    form.password('password1', 'Enter Password 1:');
    form.password('password2', 'Enter Password 2:');
    form.text('email', 'Enter Email:');
    form.text('email', 'Enter Email:');
    let r = await form.commit();
    console.log(r);
  })();
}