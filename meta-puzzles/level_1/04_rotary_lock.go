package main
import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"
// Write any import statements here

func getMinCodeEntryTime(N int32, M int32, C []int32) int64 {
  // Write your code here
  var res int64
  var prev int32 = 1
  for _, curr := range C {
    diff := abs(curr, prev)
    if diff > N/2 {
      diff = N - diff
    }
    res += int64(diff)
    prev = curr
  } 
  return res
}

func abs (a, b int32) int32 {
  if a < b {
    return b - a
  }
  return a - b
}
