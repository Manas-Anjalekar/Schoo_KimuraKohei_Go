// 演習1.2
// xとyに値を代入する前にfmt.Printlnでxとyの値を表示するとどうなるでしょうか。

// パッケージの宣言及びインポート
package main
import "fmt"

// main関数の宣言
func main(){
	// 複数の整数変数を宣言
	var x int
	var y int
	// fmt.Printlnの呼び出し
	fmt.Println(x, y)
	// 複数の整数変数を代入
	x = 1
	y = 3
	// 再びfmt.Printlnを処理
	fmt.Println(x, y)
}

// 結果 : 0 0 1 3
