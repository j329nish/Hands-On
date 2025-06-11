# Python

## 例題（A-Conflict）[[URL](https://atcoder.jp/contests/abc409/tasks/abc409_a)]

```python
N = int(input())
T = input()
A = input()
Answer = 'No'
for i in range(N):
  if (T[i] == 'o' and A[i] == 'o'):
    Answer = 'Yes'
    break
print(Answer)
```

## 実行方法
```python
## 直接実行
python sample.py
```
