func getMinimumDeflatedDiscCount(N int32, R []int32) int32 {
  var count int32 = 0
  for i := len(R) - 2; i >= 0; i-- {    
    if R[i+1] < int32(i + 2) {
      return -1
    }
    if R[i] >= R[i+1] {
      R[i] = R[i+1] - 1
      count++
    }

  }
  return count
}
