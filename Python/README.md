# Go言語

## 例題（A-Conflict）[[URL](https://atcoder.jp/contests/abc409/tasks/abc409_a)]

```go
package main
import "fmt"

func main() {
  var N int
  var T, A string
  Answer := "No"
  fmt.Scan(&N)
  fmt.Scan(&T)
  fmt.Scan(&A)
  
  for i:=0; i<N; i++ {
    if T[i] == 'o' && T[i] == A[i] {
      Answer = "Yes"
      break
    }
  }
  fmt.Println(Answer)
}
```

## 実行方法
```go
// 直接実行
go run sample.go

// コンパイルして実行
go build sample.go
./sample
```
