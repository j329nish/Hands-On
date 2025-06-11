// 直接実行
// go run sample.go

// ビルドして実行
// go build sample.go
// ./sample

package main

import (
	"errors" // goto
	"fmt"
	"os" // defer
	"time"
)

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

// 型を定義
func Type() {
	type UtcTime1 string
	type JstTime1 string
	type (
		UtcTime2 string
		JstTime2 string
	)
	var t1 UtcTime1 = "00:00:00"
	var t2 JstTime1 = "09:00:00"
	fmt.Println(t1, t2)
	// t1 = t2 とするとerror
}

// 型変換
func TypeT() {
	var a1 uint16 = 1234
	var a2 uint32 = uint32(a1)
	fmt.Println(a1, a2)
}

// リテラル・値
// nil		無しを示す特別な値
// true		真偽値の真
// false	真偽値の偽
// 1234		整数
// 1_234	整数(カンマ区切りの代わりに_を使用。_は無視される)
// 0777		8進数
// 0o755	8進数(0Oも可)
// 0x89ab	16進数(0Xも可)
// 0b1111	2進数(0Bも可)
// 123.4	小数
// 1.23e4	浮動小数点数(1.23E4も可)
// 1.23i	虚数
// "ABC"	文字列
// 'A'		文字(rune)

// エスケープシーケンス
// \a			ベル(U+0007)
// \b			バックスペース(U+0008)
// \t			タブ(U+0009)
// \n			改行(U+000A)
// \v			垂直タブ(U+000B)
// \f			フォームフィード(U+000C)
// \r			キャリッジリターン(U+000D)
// \"			ダブルクォート(U+0022)
// \'			シングルクォート(U+0027)
// \\			バックスラッシュ(U+005C)
// \x42			ASCII文字(U+0000～U+00FF)
// \u30A2		Unicode(U+0000～U+FFFF)
// \U0001F604	Unicode(U+0000～U+10FFFF)

// 配列 array
func Array() {
	a1 := [3]string{}
	a1[0] = "Red"
	a1[1] = "Green"
	a1[2] = "Blue"
	fmt.Println(a1[0], a1[1], a1[2])
}

// スライス slice
func Slice() {
	a1 := []string{}
	a1 = append(a1, "Red")
	a1 = append(a1, "Green")
	a1 = append(a1, "Blue")
	fmt.Println(a1[0], a1[1], a1[2])

	a := []int{}
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a))
	}

	bufa := make([]byte, 0, 1024)
	fmt.Println(len(bufa), cap(bufa))
}

// マップ map
func Map() {
	a1 := map[string]int{
		"x": 100,
		"y": 200,
	}
	fmt.Println(a1["x"])
	a1["z"] = 300
	delete(a1, "z")
	fmt.Println(len(a1))
	_, ok := a1["z"]
	if ok {
		fmt.Println("Exist")
	} else {
		fmt.Println("Not exist")
	}

	for key, value := range a1 {
		fmt.Printf("%s, %d\n", key, value)
	}
}

// 制御構文 if
func If() {
	x, y := 1, 2
	if x > y {
		fmt.Println("x is greater than y")
	} else if x < y {
		fmt.Println("y is greater than x")
	} else {
		fmt.Println("x equal y")
	}
}

// 制御構文 switch
func Switch() {
	var mode string = "running"
	switch mode {
	case "running":
		fmt.Println("実行中")
		/* fallthroughで次の条件も処理 */
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("不明")
	}
}

// 制御構文 for
func For() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// whileの代わり
	x, y := 1, 2
	for x < y {
		x++
	}

	// while (True)
	n := 0
	for {
		n++
		if n > 10 {
			break
		} else if n%2 == 1 {
			continue
		} else {
			fmt.Println(n)
		}
	}

	colors := [...]string{"Red", "Green", "Blue"}
	for i, color := range colors {
		fmt.Printf("%d: %s\n", i, color)
	}
}

// 制御構文 goto
func Goto() (string, error) {
	var err error
	filename := ""
	data := ""

	filename, err = GetFileName()
	if err != nil {
		fmt.Println(err)
		goto Done
	}

	data, err = ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		goto Done
	}

	fmt.Println(data)

Done:
	return data, err
}

func GetFileName() (string, error) {
	return "sample.txt", nil
}

func ReadFile(filename string) (string, error) {
	return "Hello world!", errors.New("Can't read file")
}

// 関数 func
func FuncA() {
	fmt.Println(add(5, 3))
}

func add(x int, y int) int {
	return x + y
}

func FuncB(a int, b ...int) {
	fmt.Printf("a=%d\n", a)
	for i, num := range b {
		fmt.Printf("b[%d]=%d\n", i, num)
	}
}

// 構造体 struct
type Person struct {
	name string
	age  int
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name, p.age
}

func Struct() {
	var p1 Person
	p1.SetPerson("Yamada", 26)
	name, age := p1.GetPerson()
	fmt.Printf("%s(%d)\n", name, age)
}

// インタフェース interface
func (p Person) ToString() string {
	return p.name
}

func (p Person) PrintOut() {
	fmt.Println(p.ToString())
}

type Book struct {
	title string
}

func (b Book) ToString() string {
	return b.title
}
func (b Book) PrintOut() {
	fmt.Println(b.ToString())
}

func Interface_func1() {
	a1 := Person{name: "山田太郎", age: 26}
	a2 := Book{title: "吾輩は猫である"}
	a1.PrintOut()
	a2.PrintOut()
}

type Printable interface {
	ToString() string
}

func PrintOut1(p Printable) {
	fmt.Println(p.ToString())
}

func Interface_func2() {
	a1 := Person{name: "山田太郎"}
	a2 := Book{title: "吾輩は猫である"}
	PrintOut1(a1)
	PrintOut1(a2)
}

func Interface1() {
	Interface_func1()
	Interface_func2()
}

// interface {}型
func InterfaceAny(a interface{}) {
	fmt.Printf("%d\n", a.(int))
}

func PrintOut2(a interface{}) {
	q, ok := a.(Printable)
	if ok {
		fmt.Println(q.ToString())
	} else {
		fmt.Println("Not printable.")
	}
}

func Interface_func3(a interface{}) {
	switch a.(type) {
	case bool:
		fmt.Printf("%t\n", a)
	case int:
		fmt.Printf("%d\n", a)
	case string:
		fmt.Printf("%s\n", a)
	}
}

func Interface_func4() {
	p1 := map[string]interface{}{
		"name": "Yamada",
		"age":  26,
	}
	fmt.Println(p1)

	type any interface{}
	type dict map[string]any

	p2 := dict{
		"name": "Yamada",
		"age":  26,
		"address": dict{
			"zip": "123-4567",
			"tel": "012-3456-7890",
		},
	}

	name := p2["name"]
	tel := p2["address"].(dict)["tel"]
	fmt.Println(p2)
	fmt.Println(name, tel)
}

func Interface2() {
	a1 := Book{title: "吾輩は猫である"}
	InterfaceAny(1)
	PrintOut2(a1)
	Interface_func3(true)
	Interface_func4()
}

// ポインタ pointer
func Pointer() {
	var a1 int
	var p1 *int
	p1 = &a1
	*p1 = 123
	fmt.Println(a1)

	var a2 int = 123
	var a3 int = 123
	fn(a2, &a3)
	fmt.Println(a1, a2)

	a4 := Person{"Tanaka", 26}
	p2 := &a4
	fmt.Println(a4.name, a4.age)
	fmt.Println((*p2).name)
	fmt.Println(p2.name)
}

func fn(b1 int, b2 *int) {
	b1 = 456
	*b2 = 456
}

// 領域確保 new
func New() {
	bookList := []*Book{}

	for i := 0; i < 10; i++ {
		book := new(Book)
		book.title = fmt.Sprintf("Title#%d", i)
		bookList = append(bookList, book)
	}
	for _, book := range bookList {
		fmt.Println(book.title)
	}
}

// 遅延実行 defer
func Defer() {
	fp, err := os.Open("sample.txt")
	if err != nil {
		return
	}
	defer fp.Close()
}

// ゴルーチン goroutine
func GofuncA() {
	for i := 0; i < 10; i++ {
		fmt.Print("A")
		time.Sleep(10 * time.Millisecond)
	}
}

func Goroutine_funcA() {
	go GofuncA()
	for i := 0; i < 10; i++ {
		fmt.Print("M")
		time.Sleep(10 * time.Millisecond)
	}
}

func GofuncB(chA chan<- string) {
	time.Sleep(1 * time.Second)
	chA <- "Finished"
}

func Goroutine_funcB() {
	chA := make(chan string)
	defer close(chA)
	go GofuncB(chA)
	msg := <-chA
	fmt.Println(msg)
}

func GofuncC(chA chan<- string) {
	time.Sleep(1 * time.Second)
	chA <- "GofuncC Finished"
}

func GofuncD(chB chan<- string) {
	time.Sleep(2 * time.Second)
	chB <- "GofuncD Finished"
}

func Goroutine_funcC() {
	chA := make(chan string)
	chB := make(chan string)
	defer close(chA)
	defer close(chB)
	finflagA := false
	finflagB := false
	go GofuncC(chA)
	go GofuncD(chB)
	for {
		select {
		case msg := <-chA:
			finflagA = true
			fmt.Println(msg)
		case msg := <-chB:
			finflagB = true
			fmt.Println(msg)
		}
		if finflagA && finflagB {
			break
		}
	}
}

func Goroutine() {
	Goroutine_funcA()
	Goroutine_funcB()
	Goroutine_funcC()
}

func main() {
	fmt.Println("hello, world")
	Print()
	Var()
	Const()
	Type()
	TypeT()
	Array()
	Slice()
	Map()
	If()
	Switch()
	For()
	Goto()
	FuncA()
	FuncB(1, 2, 3, 4, 5)
	Struct()
	Interface1()
	Interface2()
	Pointer()
	New()
	Defer()
	Goroutine()
}
