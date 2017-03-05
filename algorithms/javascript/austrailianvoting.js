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
  const cases = parseInt(inArr.shift(), 10)
  let inputIndex = 1 // skip first blank line

  for (let i = 0; i < cases; i++, inputIndex++) {
    const candidateCount = parseInt(inArr[inputIndex++], 10)
    const candidates = inArr.slice(inputIndex, inputIndex + candidateCount)
    inputIndex += candidateCount

    let activeCandidates = {}
    let votes = {}
    for (let num = 1; num <= candidates.length; num++) {
      activeCandidates[num] = true
    }

    const startOfBallots = inputIndex
    let totalBallots = 0
    let winner = ''
    while (winner === '') {
      for (let num = 1; num <= candidates.length; num++) {
        votes[num] = 0
      }
      for (let ballotCount = 0; inArr[startOfBallots + ballotCount] !== ''; ballotCount++) {
        if (ballotCount > totalBallots) totalBallots = ballotCount
        const ballot = inArr[startOfBallots + ballotCount].split(' ')
        for (let voteIndex = 0; voteIndex < ballot.length; voteIndex++) {
          const vote = parseInt(ballot[voteIndex], 10)
          if (activeCandidates[vote]) {
            votes[vote]++
            break // Successful vote
          }
        }
      }
      let worstTotal = votes['1']
      for (candidate in votes) {
        if (votes[candidate] < worstTotal) worstTotal = votes[candidate]
        if (votes[candidate] / totalBallots > 0.5) winner = candidates[candidate - 1]
      }
      for (candidate in votes) {
        if (votes[candidate] === worstTotal) activeCandidates[candidate] = false
      }
    }
    console.log(winner)
    inputIndex += totalBallots
  }
}
