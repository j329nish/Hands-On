# C++

## 例題（A-Conflict）[[URL](https://atcoder.jp/contests/abc409/tasks/abc409_a)]

```cpp
#include <iostream>
#include <string>
using namespace std;

int main() {
  int N;
  string T, A, Answer = "No";
  cin >> N;
  cin >> T;
  cin >> A;
  
  for (int i=0; i<N; i++) {
    if (T[i] == 'o' and T[i] == A[i]) {
      Answer = "Yes";
      break;
    }
  }
  cout << Answer << endl;
  
  return 0;
}
```

## 実行方法
```cpp
// コンパイルして実行
g++ sample.cpp -o sample
./sample
```
