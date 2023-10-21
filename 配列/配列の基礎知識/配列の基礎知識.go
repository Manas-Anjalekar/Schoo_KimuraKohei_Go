// 配列の基礎知識

// パッケージの宣言とインポート
package main
import "fmt" 

// main関数の宣言
func main() {
	// 配列の宣言(省略)
	arr := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Printlnの呼び出し
	fmt.Println(arr)
	// 配列の0,1と5ポジションにある値を表示
	fmt.Println(arr[0], arr[1], arr[5])
}

/*
結果 : 
[2 3 5 7 11 13]
2 3 13
*/