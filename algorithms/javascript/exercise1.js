const getR = () => {
  const n = 200

  let r = 0
  for (let i = 1; i <= n; i++) {
    for (let j = i + 1; j <= n; j++) {
      for (let k = i + j - 1; k <= n; k++) {
        r++
      }
    }
  }
  console.log(r)
}

getR()
