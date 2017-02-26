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
  const inArr = input.split('\n')
  inArr.pop()
  inArr.forEach((line) => {
    const lineArr = line.split(' ').map(Number)
    const i = lineArr[0]
    const j = lineArr[1]

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
    console.log(lineArr.join(' ') + ' ' + maxCount)
  })
}
