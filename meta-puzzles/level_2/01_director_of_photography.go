package main
import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"
// Write any import statements here

func getArtisticPhotographCount(N int32, C string, X int32, Y int32) int64 {
  prefSumArrayP := getPrefSumArray('P', C)
  prefSumArrayB := getPrefSumArray('B', C)
  
  var res int64
  for i := X; i < N - X; i++ {
    if C[i] == 'A' {
      leftPCount := getPrefSum(i - X, N, prefSumArrayP) - getPrefSum(i - Y - 1, N, prefSumArrayP)
      rightBCount := getPrefSum(i + Y, N, prefSumArrayB) - getPrefSum(i + X - 1, N, prefSumArrayB)
      res += int64(leftPCount * rightBCount)

      leftBCount := getPrefSum(i - X, N, prefSumArrayB) - getPrefSum(i - Y - 1, N, prefSumArrayB)
      rightPCount := getPrefSum(i + Y, N, prefSumArrayP) - getPrefSum(i + X - 1, N, prefSumArrayP)
      res += int64(leftBCount * rightPCount)
    }
  }
  return res
}

func getPrefSumArray(char rune, C string) []int32 {
  prefSumArray := make([]int32, len(C))
  var count int32
  for i, v := range C {
    if v == char {
      count++
    }
    prefSumArray[i] = count
  }
  return prefSumArray
}

func getPrefSum(i int32, N int32, prefSumArray []int32) int32 {
  // validate index i
  if i < 0 {
    return 0
  }
  if i >= N {
    i = N - 1
  }
  return prefSumArray[i]
}

