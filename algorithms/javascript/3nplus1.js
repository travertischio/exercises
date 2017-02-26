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
  const i = inArr[0]
  const j = inArr[1]

  let maxCount = 0
  for (let k = i; k <= j; k++) {
    let count = 1
    for (let l = k; l > 1; count++) {
      if (l % 2 === 0) {
        l = l / 2
      } else {
        l = 3 * l + 1
      }
    }
    if (count > maxCount) maxCount = count
  }
  console.log(inArr.join(' ') + ' ' + maxCount)
}
