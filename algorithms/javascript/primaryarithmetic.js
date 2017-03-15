function carry(twoNumbers) {
  const firstNumer = twoNumbers[0].split('').reverse()
  const secondNumer = twoNumbers[1].split('').reverse()

  let carrys = 0
  let previousCarry = 0
  const longerLength = firstNumer.length > secondNumer.length ? firstNumer.length : secondNumer.length
  for(let i = 0; i < longerLength; i++) {
    const first = firstNumer[i]
    const second = secondNumer[i]
    if (first && second) {
      const a = parseInt(first, 10)
      const b = parseInt(second, 10)
      if (a + b + previousCarry > 9) {
        previousCarry = 1
        carrys++
      } else {
        previousCarry = 0
      }
    } else if (previousCarry === 1 && (first === '9' || second === '9')) {
      previousCarry = 1
      carrys++
    }
  }
  return(carrys)
}


console.log(carry(['123', '456']))
console.log(carry(['555', '555']))
console.log(carry(['123', '594']))
