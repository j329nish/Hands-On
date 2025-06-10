// 直接実行
// go run sample.go

// ビルドして実行
// go build sample.go
// ./sample

package main

import "fmt"

func main() {
	fmt.Println("hello, world")
	Print()
	Var()
	Const()
}

// Print・Println・Printf
func Print() {
	num := 123
	str := "ABC"
	fmt.Print("num=", num, " str=", str, "\n")
	fmt.Println("num =", num, "str =", str)
	fmt.Printf("num=%d str=%s\n", num, str)
}

// %v(デフォルト形式)、%#v(Go言語表記)、%t(真偽値)、%d(整数)、
// %s(文字列)、%c(文字)、%f(小数)、%F(小数)、%e(浮動小数点数e)、
// %E(浮動小数点数E)、%g(%f/%e自動選択)、%b(2進数)、%o(8進数)、
// %O(0o付き8進数)、%x(16進数)、%X(16進数大文字)、%U(Unicode)、
// %p(ポインタ)、%q("..."付き文字列)、%T(型表示)、%%(パーセント)

// 変数 var
func Var() {
	var a1 int = 1
	var a2 = 2
	a3 := 3
	var (
		a4 int = 4
		a5 int = 5
	)
	a1 = 11
	name, age := "Yamada", 26
	fmt.Println(a1, a2, a3, a4, a5, name, age)
}

// 定数 const
func Const() {
	const foo1 = 100
	const (
		foo2 = 100
		baa  = 200
	)
	fmt.Println(foo1, foo2, baa)
}

// コメント
/* コメント */

// 一行で書ける
// num = 123; str = "ABC";

// 型 type
// bool							真偽値(true or false)
// int8/int16/int32/int64		nビット整数
// uint8/uint16/uint32/uint64	nビット非負整数
// float32/float64				nビット浮動小数点数
// complex64/complex128			nビット虚数
// byte							1バイトデータ(uint8と同義)
// rune							1文字(int32と同義)
// uint							uint32 または uint64
// int							int32 または int64
// uintptr						ポインタを表現するのに十分な非負整数
// string						文字列

// func Type() {
// 	type
// }
