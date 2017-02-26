let totalInput = ''
process.stdin.resume()
process.stdin.setEncoding('ascii')
process.stdin.on('data', (input) => {
  totalInput += input
})

process.stdin.on('end', () => {
  processData(totalInput)
})

const processData = (input) => {
  const inArr = input.split(' ').map(Number)
  let numerator = inArr[0]
  const denominator = inArr[1]
  let divisor = 1
  let n = numerator
  for (; n < denominator; n = n + numerator, divisor++) {}
  let remainder = 0
  if (n !== denominator) {
    divisor--
    remainder = denominator - (n - numerator)
  }
  if (remainder) {
    console.log(divisor + ' Remainder ' + remainder)
  } else {
    console.log(divisor)
  }
}
