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
  for (; numerator < denominator; numerator = numerator + numerator, divisor++) {}
  console.log(divisor)
}
