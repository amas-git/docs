import * as cheerio from 'cheerio';

const print = (m) => console.log(m);

function log(target: any, key: string, descriptor: any) {
    const original = descriptor.value;
    descriptor.value = function (...args: any[]) {
        // Call the original method
        const result = original.apply(this, args);
        // Log the call, and the result
        console.log(`${key} with args ${JSON.stringify(args)} returned
${JSON.stringify(result)}`);
        // Return the result
        return result;
    }
    return descriptor;
}

function test(a) {
    return a;
}
const NAME: string = `amas`;
//---------------------------------------------------------[ Enum ]
enum Color {
    RED = 0xF00,
    GREEN = 0x0F0,
    BLUE = 0x00F
};

// const将影响ts的编译结果
const enum Color2 {
    RED = 0xF00,
    GREEN = 0x0F0,
    BLUE = 0x00F
};

enum Flag {
    A = 0,
    B = 1,
    C = 2,
    D = 4,
    E = 8
};

const names: string[] = [];

const color = Color.RED;  // TSC: var color = Color.RED;
const color2 = Color2.RED; // TSC: var color2 = 3840 /* RED */;

//---------------------------------------------------------[ Function Types ]
function add(a: number, b: number): number {
    return a + b;
}

let add2: (a: number, b: number) => number = (a, b) => a + b;
//---------------------------------------------------------[ Object Type ]
let person: { name: string, age: number };
person = {
    name: "amas",
    age: 19,
    //ext: 1 // ERRORS
}

type Person = {
    name: string,
    age: number
}

const amas: Person = {
    name: "amas",
    age: 18
}
//---------------------------------------------------------[ Primitive Type ]
let s: string;
let n: never;
let u: undefined;
let N: number;
let x: null;
let v: void;
//---------------------------------------------------------[ Literal Types ]
type Even = 2 | 4 | 6 | 8;
// let x: Even = 1; // ERROR
//---------------------------------------------------------[ Union Types ]
type StringOrError = string | Error;
//---------------------------------------------------------[ Intersection Types ]
interface A {
    foo(): void;
}

interface B {
    bar(): void;
}

type C = A & B;
let c: C;
//---------------------------------------------------------[ Array Types ]
let xsC: C[] = [];

//---------------------------------------------------------[ Tuple Types ]
type Point = [number, number];
let xPoint: Point = [1, 2]

//---------------------------------------------------------[ Index Type]
interface Book {
    readonly name: string,
    isbn: string,
    author?: string, // Optional
    seria: string | null // 可空
}

interface BookMap {
    [index: string]: Book;
}

let book1: Book = { name: "book1", isbn: "98767", seria: null };

// book1.name = "xxx" // ERROR

let bookDict: BookMap = {};
// bookDict[1] = 1; // ERROR
bookDict[1] = book1;
bookDict["book1"] = book1;
bookDict["book2"] = { name: "book2", isbn: "9999", seria: null };
bookDict["book3"] = { name: "book3", isbn: "9999", seria: "11" };
// bookDict["book4"] = {name: "book3", isbn:"9999"}; // ERROR, no seria
//---------------------------------------------------------[ Mapped Types]


type Package = {
    readonly name: string;

}

// ? type和interface的区别
// 1. type名是唯一的，重复定义会报错
// 2. interface同名会merge
// 3. type里可以遍历type的key,有点meta编程的意思
type ReadOnly<T> = { readonly [k in keyof T]: T[k]; }
type Optional<T> = { [k in keyof T]?: T[k]; }
type Nullable<T> = { [k in keyof T]: T[k] | null; }

type Keys = "firstname" | "surname"
type DudeType = {
    [key in Keys]: string
}

const test1: DudeType = {
    firstname: "Pawel",
    surname: "Grzybek"
}
// 4. type不可以循环引用(circularly references)
type Tx = string | Ty;
// type Ty = Tx // ERROR
type Ty = never;

type ReadOnlyBook = ReadOnly<Book>;

let rbook: ReadOnlyBook = { name: "1", isbn: "2", seria: "99" };
// rbook.isbn = null; ERROR

// 5. interface可以多继承
interface IA {
    ia();
}

interface IB {
    ib();
}

interface IAB extends IA,IB {

}


//---------------------------------------------------------[ Type Assertion ]
const name = "amas";
const obj = { a: 'a', b: 'b' }
//let num1:number = <number> name; // ERROR
//let sobj:string = <string> obj;  // ERROR
let sobj: string = <string>obj.toString();
//---------------------------------------------------------[ Type Guard ]
print((typeof 1));     // js
print((typeof rbook)); // js

//---------------------------------------------------------[ Discriminated Unions ]

//---------------------------------------------------------[ Override ]
class A {
    get(x: 'hello'): string;
    get(x: string): string {
        return x === "hello" ? "hello" : x;
    }
}

let aa: A = new A();
print(aa.get("hello"));
//---------------------------------------------------------[ Generics Function ]
function reverse<T>(list: T[]): T[] {
    const reversedList: T[] = [];
    for (let i = (list.length - 1); i >= 0; i--) {
        reversedList.push(list[i]);
    }
    return reversedList;
}

print(reverse([1, 2, 3]));
//print(reverse(null)); ? 为什么可以是null
//---------------------------------------------------------[ Naming Space ]
export class AAA {

}

namespace Car {
    export function a() {
        return "Car.a";
    }
}

namespace Goo {
    export function a() {
        return "Goo.b";
    }
}

print(Car.a());
print(Goo.a());
//---------------------------------------------------------[ export ]
export class CCC {
    
    public test(a) {
        return a;
    }
}

export function hello() {
}

export default class DDD {
}

export namespace Ns {
    export class Car {

    }
}

// ERROR: default只能有一个
// export default class DD2 {
//
// }
//----------------------------------------------------------[ NOT NULL ]
// 必须打开--strictNullChecks
function f(x: {}): void {

}

f(0);           // ok
f("");          // ok
f(true);        // ok
f({});          // ok
f(null);        // error
f(undefined);   // error
print(+0 === -0);
print(NaN === NaN);