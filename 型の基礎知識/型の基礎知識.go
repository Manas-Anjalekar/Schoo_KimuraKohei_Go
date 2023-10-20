// 型の基礎知識

// パッケージの宣言とインポート
package main
import "fmt"

// main関数の宣言
func main(){
	// 複数の変数の宣言と代入
	var x int = 10 // 整数型
	var s string = "abc" // 文字列型
	var f float32 = 3.99 // 実数型
	var b bool = true // ブーリアン型
	// fmt.Printlnの呼び出し
	fmt.Println(x, s, f, b)
	// コンパイルエラー
	// x = f
	// キャスト
	x = int(f)
	// x実数型変数を表示
	fmt.Println("x : ", x)
}

// 結果 : 10 abc 3.99 true x :  3