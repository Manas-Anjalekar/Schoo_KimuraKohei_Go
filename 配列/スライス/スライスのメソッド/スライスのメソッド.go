// スライスに値を追加と長さの測定

// パッケージの宣言とインポート
package main
import "fmt"

// main関数の宣言
func main() {
	// 配列の宣言と初期化
	s := []int{2, 3, 7, 11, 13, 15}
	// 配列を表示
	fmt.Println(s)
	// 配列の長さを表示
	fmt.Println("length:", len(s))

	// 配列に値を追加
	s = append(s, 17, 19)
	// 再び配列を表示
	fmt.Println(s)
	// 配列の長さを表示
	fmt.Println("length:", len(s))

	// for・rangeメソッドで配列の要素と数値にアクセス
	for i, v := range s {
		// 要素と数値を表示
		fmt.Println("要素【", i, "】 値 = ", v)
	}
}

/*
結果 : 
[2 3 7 11 13 15]
length: 6
[2 3 7 11 13 15 17 19]
length: 8
要素【 0 】 値 =  2
要素【 1 】 値 =  3
要素【 2 】 値 =  7
要素【 3 】 値 =  11
要素【 4 】 値 =  13
要素【 5 】 値 =  15
要素【 6 】 値 =  17
要素【 7 】 値 =  19
*/