function* simpleGenerator() {
  yield 1;
  yield 2;
  yield 3;
}

// Создаем экземпляр генератора
const generator = simpleGenerator();

// Используем метод next() для получения значений
console.log(generator.next().value); // 1
console.log(generator.next().value); // 2
console.log(generator.next().value); // 3
console.log(generator.next().value); // undefined
console.log(generator.next().done);  // true
