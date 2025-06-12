// 値
// 12345		整数
// 12_345_678	カンマの代わりに_を使用して読みやすく
// 12345u32	    u32型の12345
// 0xfff		16進数
// 0o777		8進数
// 0b11001100	2進数
// 'あ'		    文字(char)
// "..."		文字列(&str)
// r"..."		raw文字列
// r#"..."#	    ダブルクォートをそのまま使用できる文字列
// b'a'		    1バイト文字(u8)
// b"abc"	    バイト配列(&[u8])
// br"..."		rawバイト配列(&[u8])

// 変数・定数 let, mut, const
fn vari_con() {
    let n1 = 0;
    let mut n2 = 0;
    const MAX_POINTS: u32 = 100_000;
    println!("{} {} {}", n1, n2, MAX_POINTS);
    n2 = 3;
    println!("n2 = {}", n2);
}

// 型
// bool 			    真偽値(true/false)
// i8	    	    	符号付き8ビット整数
// u8		        	符号無し8ビット整数
// i16	    		    符号付き16ビット整数
// u16    			    符号無し16ビット整数
// i32	    		    符号付き32ビット整数
// u32		        	符号無し32ビット整数
// i64		    	    符号付き64ビット整数
// u64       			符号無し64ビット整数
// i128	    		    符号付き128ビット整数
// u128		        	符号無し128ビット整数
// isize	    	    ポインタサイズと同じ符号付き整数
// usize        		ポインタサイズと同じ符号無し整数
// f32  		    	32ビット浮動小数点数
// f64			        64ビット浮動小数点数
// char 		    	文字(U+0000～U+D7FF, U+E000～U+10FFFF)
// str	        		文字列(&strとして使用することが多い)
// (type, type, ...)	タプル
// [type; len]		    配列
// Vec<type>		    ベクタ
// &type			    typeへの参照
// &mut type		    typeへのミュータブルな参照
// &[type]			    type型要素を持つスライス

// 型変換 as
fn as_con() {
    let x: i32 = 123;
    let y: i64 = x as i64;
    println!("{} {}", x, y);
}

// 構造体 struct
struct Point {
    x: i32,
    y: i32,
}

fn _struct() {
    let p = Point { x: 100, y: 200 };
    println!("{} {}", p.x, p.y);
}

// 共用体 union
union MyUnion {
    f1: u32,
    f2: u32,
}

fn _union() {
    let u = MyUnion { f1: 123 };
    unsafe {
        println!("{}", u.f1);
        println!("{}", u.f2);
    }
}

// 列挙型 enum
#[derive(Debug)]
enum Color {
    Red,
    Green,
    Blue,
}

fn _enum() {
    let (color1, color2, color3) = (Color::Red, Color::Green, Color::Blue);
    println!("{:?} {:?} {:?}", color1, color2, color3);
}

// タプル tup
fn _tup() {
    let tup = (10, "20", 30);
    println!("{} {} {}", tup.0, tup.1, tup.2);
}

// 配列 array
fn _array() {
    let arr = [10, 20, 30];
    println!("{} {} {}", arr[0], arr[1], arr[2]);

    for v in &arr {
        println!("{}", v);
    }
}

// ベクタ vec
fn _vec() {
    let mut vect = vec![10, 20, 30];
    vect.push(40);
    println!("{} {} {} {}", vect[0], vect[1], vect[2], vect[3]);

    for v in &vect {
        println!("{}", v);
    }
}

// ハッシュマップ HashMap
fn _hashmap() {
    use std::collections::HashMap;
    let mut map = HashMap::new();
    map.insert("x", 10);
    map.insert("y", 20);
    map.insert("z", 30);
    println!("{} {} {}", map["x"], map["y"], map["z"]);

    for (k, v) in &map {
        println!("{} {}", k, v);
    }
}

// 文字列 &str = String
fn _str() {
    let mut name1: &str = "Yamada";
    println!("{}", name1);
    name1 = "Tanaka";

    let mut name2 = String::from("Yamada");
    println!("{}", name2);
    name2 = "Tanaka".to_string();
    name2.push_str(" Taro");
    println!("{} {}", name1, name2)
}

// 演算子
// fn(...) -> type	関数の型定義
// expr;			行の終わり
// 'label			ラベル
// expr..expr		範囲
// macro!(...)		マクロ呼び出し
// macro![...]		マクロ呼び出し
// macro!{...}		マクロ呼び出し
// [type; len]		配列

// ヒープ領域 Box
fn _box() {
    let p: Box<Point> = Box::new(Point { x: 100, y: 200 });
    println!("{} {}", p.x, p.y);
}

impl Drop for Point {
    fn drop(&mut self) {
        println!("Bye!");
    }
}

// スライス &var[n..m]
fn _var() {
    let s = String::from("ABCDEFGH");
    let s1 = &s[0..3];
    let s2 = &s[3..6];
    println!("{} {}", s1, s2);

    let a = [10, 20, 30, 40, 50, 60, 70, 80];
    let a1 = &a[0..3];
    let a2 = &a[3..6];
    println!("{:?} {:?}", a1, a2);
}

// 関数 fn
fn add1(x: i32, y: i32) -> i32 {
    return x + y;
}

fn add2(x: i32, y: i32) -> i32 {
    x + y
}

fn _add() {
    println!("{}", add1(2, 3));
    println!("{}", add2(3, 4));
}

// クロージャー
fn _closure() {
    let square = |x: i32| x * x;
    println!("{}", square(9));

    let msg = String::from("Hello");
    let func = move || {
        println!("{}", msg);
    };
    func();
}

// マクロ macro_rules!
macro_rules! log {
    ($x:expr) => {
        println!("{}", $x);
    };
}

// 条件分岐 if
fn _if() {
    let n: i32 = 2;
    if n == 1 {
        println!("One");
    } else if n == 2 {
        println!("Two");
    } else {
        println!("Others");
    }
    let s = if n == 1 { "OK!" } else { "NG!" };
    println!("{}", s);
}

// 繰り返し while
fn _while() {
    let mut n = 0;
    while n < 10 {
        println!("{}", n);
        n += 1;
    }
}

// 繰り返し for
fn _for() {
    for i in 0..10 {
        println!("{}", i);
    }
}

// ループ loop, ループ制御 break, continue
fn _loop() {
    let mut n = 0;
    loop {
        n += 1;
        if n == 2 {
            continue;
        }
        if n == 8 {
            break;
        }
        println!("{}", n);
    }
}

// マッチ match
fn _match() {
    let x = 2;
    match x {
        1 => println!("One"),
        2 => println!("Two"),
        3 => println!("Three"),
        _ => println!("More"),
    }
}

// インプリメンテーション impl
struct Rect1 {
    width: u32,
    height: u32,
}

impl Rect1 {
    fn area(&self) -> u32 {
        self.width * self.height
    }
}

fn _impl() {
    let r = Rect1 {
        width: 200,
        height: 300,
    };
    println!("{}", r.area());
}

// トレイト trait
trait Printable {
    fn print(&self);
}

impl Printable for Rect1 {
    fn print(&self) {
        println!("width:{}, height:{}", self.width, self.height)
    }
}

struct Rect2<T> {
    width: T,
    height: T,
}

impl<T> Printable for Rect2<T>
where
    T: std::fmt::Display,
{
    fn print(self: &Rect2<T>) {
        println!("{}x{}", self.width, self.height);
    }
}

fn _trait() {
    let r1 = Rect1 {
        width: 200,
        height: 300,
    };
    r1.print();

    let r2: Rect2<i32> = Rect2 {
        width: 100,
        height: 200,
    };
    let r3: Rect2<i64> = Rect2 {
        width: 100,
        height: 200,
    };
    r2.print();
    r3.print();
}

// イテレータ iterator

// マルチスレッド thread

// 非同期関数 async, await

// クレート crate

// モジュール mod, pub, use, as

// 参照型

// 所有権・移動・参照・借用

// 型エイリアス type

// 型を調べる typeof

// 外部関数の呼び出し extern

// 静的変数 static

fn main() {
    println!("Hello, world!");
    vari_con();
    as_con();
    _struct();
    _union();
    _enum();
    _tup();
    _array();
    _vec();
    _hashmap();
    _str();
    _box();
    _var();
    _add();
    _closure();
    log!("ABC...");
    _if();
    _while();
    _for();
    _loop();
    _match();
    _impl();
    _trait();
}
