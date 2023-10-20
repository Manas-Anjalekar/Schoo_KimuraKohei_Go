/* 演習1.4
 1から20まで数をカウントするプログラムをfor文を使って書いてください。
 但し、以下の条件の場合は数値ではなく指定された文字列を表示して下さい。
 ① 3の倍数のときはFizz
 ② 5の倍数のときはBuzz
 ③ 3の倍数かつ5の倍数のときはFizzBuzz
*/

// パッケージの宣言とインポート
package main
import "fmt"

// main関数の宣言
func main(){
	//
	for i:= 1; i <= 20; i++ {
		if i%3 == 0 {
			// fmt.Printlnの呼び出し
			fmt.Println("Fizz")

			// iを5で割ると余りが0だった場合、処理する
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			}
		} else if i%5 == 0{
			// fmt.Printlnの呼び出し
			fmt.Println("Buzz")
		} else {
			// iの値を表示
			fmt.Println(i)
		}
	}
}

/*
 結果 : 
 1
 2
 Fizz
 4
 Buzz
 Fizz
 7
 8
 Fizz
 Buzz
 11
 Fizz
 13
 14
 Fizz
 FizzBuzz
 16
 17
 Fizz
 19
 Buzz
*/