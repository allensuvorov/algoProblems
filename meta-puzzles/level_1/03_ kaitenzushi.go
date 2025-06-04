package main
import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

func getMaximumEatenDishCount(N int32, D []int32, K int32) int32 {
  justEaten := make(map[int32]bool)
  q := make([]int32, 0, K)
  var dishCount int32
  
  for _, dish := range D {
    if !justEaten[dish] {
      dishCount++
      justEaten[dish] = true
      q = append(q, dish) // enqueue
    }
    
    if len(q) > int(K) { // dequeue
      justEaten[q[0]] = false
      q = q[1:]
    }
    
  }
  return dishCount
}
