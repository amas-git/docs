const vm = require('vm');


function seval(code, context, opts) {
  const sandbox = {};
  const resultKey = 'SAFE_EVAL_' + Math.floor(Math.random() * 1000000);
  sandbox[resultKey] = {};

  const clearContext = `
    (function() {
      Function = undefined;
      const keys = Object.getOwnPropertyNames(this).concat(['constructor']);
      keys.forEach((key) => {
        const item = this[key];
        if (!item || typeof item.constructor !== 'function') return;
        this[key].constructor = undefined;
      });
    })();
  `
  code = clearContext + resultKey + '=' + code;
  if (context) {
    Object.keys(context).forEach(function (key) {
      sandbox[key] = context[key];
    })
  }
  vm.runInNewContext(code, sandbox, opts);
  return sandbox[resultKey];
}
const x=1;
console.log('hello');
const r = seval('2000+a === 100 || (1 === 1)', {a:400, x:999});
//console.log(seval('x; var fs = require("fs")', {x:999}));

//expr().var('TESTCASE', sheet1, a, b);

//expr('@sheet[a,b] > 100');

(function() {
  console.log(Function);
})();

<<<<<<< HEAD
console.log(vm.runInNewContext('`${a}+1`', {a:1}));
=======

class Rule {
  constructor(id) {
    this.id = id;
  }
  test(expr) {
    let c = eval(expr);
    let r = {
      DO(action) {
        if (c && action) {
          eval(action)
        }
        return r;
      },
      orElse(action) {
        if (!c && action) {
          eval(action);
        }
        return r;
      }
    }
    return r;
  }
}

let xx = 1;
new Rule('is_even').test('xx % 2 ===  0').DO(`console.log('YES')`).orElse(`console.log('NO')`);

>>>>>>> ecdd2757dd0a73c45579416d7033e1a68d1d286c
