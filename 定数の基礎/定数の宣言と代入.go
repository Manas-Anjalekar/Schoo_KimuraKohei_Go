// 定数の宣言と代入

// パッケージの宣言とインポート
package main
import "fmt"

// main関数の宣言
func main(){
	// 定数の宣言と代入
	const n = 1
	// fmt.Printlnの呼び出し
	fmt.Println(n)
	// コンパイルエラー
	n = 2
}

// 結果 : 定数の基礎/定数の宣言と代入.go:14:4: cannot assign to n