func getUniformIntegerCountInInterval(A int64, B int64) int32 {
  // generate uniform numbers 1 to 10^12
  var count int32 = 0
  var base int64 = 0
  num := A
  for num <= B {
    base = base * 10 + 1
    for f := 1; f <= 9; f++ {
      num = base * int64(f)
      if A <= num && num <= B {
        count++
      }
    }
  }
  
  return count
}
