package main
import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

func getMaximumEatenDishCount(N int32, D []int32, K int32) int32 {
  set := make(map[int32]bool) // set of previous K eaten
  q := make([]int32, 0, K) // queue of previous K eaten
  var res int32
  
  for _, dish := range D {
    if !set[dish] {
      res++
      set[dish] = true
      q = append(q, dish) // enqueue
    }
    
    if len(q) > int(K) { // dequeue
      set[q[0]] = false
      q = q[1:]
    }
    
  }
  return res
}
