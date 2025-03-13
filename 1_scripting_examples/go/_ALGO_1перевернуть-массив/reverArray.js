const array1 = [1,2,3,4,5,6,7,8,9]
const array2 = [10,9,8,7,6,5,4,3,2,1]
const array3 = [10,9,8,7,6,5,4,3,2,1,0]

function revertArray(arr) {
  const result = []

  const len = arr.length
  const mid = len / 2

  for (let idx = 0; idx <= mid; idx++) {
    result.push(arr[len - idx -1])
  }

  let past = arr.slice(0, Math.ceil(mid) - 1)
  let pastLen = past.length - 1
  for (; pastLen >= 0; pastLen--) {
    result.push(past[pastLen])
  }

  return result
}

const revertArray2 = (arr) => {
  lnt = arr.length
  const result = []

  for (let idx = 0; idx <= lnt; idx++) {
    result.push(0)
  }

  for (let idx = 0; idx <= lnt - 1; idx++) {
    result[idx] = arr[lnt-idx-1]
    result[lnt-idx-1] = arr[idx]
  }

  return result
}

console.log(revertArray2(array1))
console.log(revertArray2(array2))
console.log(revertArray2(array3))