package main
import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"
import "sort"

func getMaxAdditionalDinersCount(N int64, K int64, M int32, S []int64) int64 {
  sort.Slice(S, func(i, j int) bool {
    return S[i] < S[j]
  })
  
  var res int64
  
  res += getDinersCount(S[0] - 1 - K, K) // safe space before first diner

  for i := 1; i < len(S); i++ {
    safeSpace := S[i] - S[i-1] - 1 - 2*K //
    res += getDinersCount(safeSpace, K)
  }
  
  res += getDinersCount(N - S[len(S)-1] - K, K) // safe space after the last diner
  
  return res
}

func getDinersCount(safeSpace int64, K int64) int64 {
  var dinersCount int64
  if safeSpace > 0 {
      slot := 1 + K // one person plus distancing
      dinersCount += safeSpace / slot
      if safeSpace % slot != 0 {
        dinersCount++
      }
    }
  return dinersCount
}
