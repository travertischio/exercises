function jumpers(arr) {
  const len = arr.length
  let previous = 0
  let status = 'Jolly'

  arr.forEach((num) => {
    if (num < 1) status = 'Not Jolly'
    if (Math.abs(num - previous) > len) status = 'Not Jolly'

    previous = num
  })
  return status
}

console.log(jumpers([4,5,6,2,3,1]))
console.log(jumpers([4,1,4,2,3]))
console.log(jumpers([4,5,6,0,3,1]))
console.log(jumpers([5,1,4,2,-1,6]))
