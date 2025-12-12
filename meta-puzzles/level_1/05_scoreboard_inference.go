func max(a, b int32) int32 {
  if a > b {
    return a
  }
  return b
}

func getMinProblemCount(N int32, S []int32) int32 {
  var maxNum int32
  var ones int32
  
  for _, v := range S {
    maxNum = max(maxNum, v)
    if v % 2 == 1 {
      ones = 1
    }
  }
  
  return maxNum / 2 + ones
}
