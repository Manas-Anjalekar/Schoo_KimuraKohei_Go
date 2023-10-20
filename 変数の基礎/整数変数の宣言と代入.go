// 変数の宣言及び代入方法

// パッケージの宣言及びインポート
package main
import "fmt"

// main関数の宣言
func main() {
	// 複数の整数変数の宣言
	var x int
	var y int
	// 複数の整数変数を代入
	x = 1
	y = 2
	// fmt.Printlnの呼び出し
	fmt.Println(x, y)
	// 自分での足し算の省略
	x += y
	// 再びfmt.Printlnを処理
	fmt.Println(x, y)
}

// 結果 : 1 2　3 2
