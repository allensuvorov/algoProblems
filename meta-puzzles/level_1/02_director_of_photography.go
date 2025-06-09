package main
import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"
// Write any import statements here

func getArtisticPhotographCount(N int32, C string, X int32, Y int32) int32 {
  // Write your code here
  prefSumArrayP := getPrefSumArray('P', C)
  prefSumArrayB := getPrefSumArray('B', C)
  
  var res int32
  for i, v := range C {
    if v == 'A' {
      leftPCount := getPrefSum(i - int(X), N, prefSumArrayP) - getPrefSum(i - int(Y) - 1, N, prefSumArrayP)
      
      var rightBCount int32
      if int32(i) + X < N {
        rightBCount = getPrefSum(i + int(Y), N, prefSumArrayB) - getPrefSum(i + int(X) - 1, N, prefSumArrayB)
      }
      fmt.Println(leftPCount, rightBCount)
      res += leftPCount * rightBCount

      leftBCount := getPrefSum(i - int(X), N, prefSumArrayB) - getPrefSum(i - int(Y) - 1, N, prefSumArrayB)
      var rightPCount int32
      if int32(i) + X < N {
        rightPCount = getPrefSum(i + int(Y), N, prefSumArrayP) - getPrefSum(i + int(X) - 1, N, prefSumArrayP)
      }
      fmt.Println(leftBCount, rightPCount)
      res += leftBCount * rightPCount
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

func getPrefSum(i int, N int32, prefSumArray []int32) int32 {
  // validate index i
  if i < 0 {
    return 0
  }
  if i >= int(N) {
    i = int(N) - 1
  }
  return prefSumArray[i]
}

/*
  APABA
     lr
^ 
P:01111
B 00011
*/
