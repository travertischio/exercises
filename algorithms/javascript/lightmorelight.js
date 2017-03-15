function lights(n) {
  return Math.sqrt(n) % 1 === 0
}

console.log(lights(10))
console.log(lights(81))
console.log(lights(3))
console.log(lights(6241))
console.log(lights(8191))
