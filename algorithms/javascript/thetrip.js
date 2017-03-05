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
  for (let i = 0; i < inArr.length; ) {
    const people = parseInt(inArr[i], 10)
		if (people > 0) {
			const nextTrip = i + people + 1
			const thisTrip = inArr.slice(i + 1, nextTrip)
			i = nextTrip
			let sum = 0
			thisTrip.forEach((spent) => {
				sum += parseFloat(spent)
			})
			const thisTripAvg = sum / people

			let changeDown = 0
			let changeUp = 0
			thisTrip.forEach((value) => {
				const val = parseFloat(value)
				if (val - thisTripAvg > 0.01) {
					changeDown += val - thisTripAvg
				} else {
					changeUp += thisTripAvg - val
				}
			})
			let smallerChange = changeUp < changeDown ? changeUp : changeDown
      if ((smallerChange * 100.0) % 10 > 0) smallerChange -= 0.005 // Always round down
			console.log('$' + smallerChange.toFixed(2))
		} else {
			i++
		}
  }
}
