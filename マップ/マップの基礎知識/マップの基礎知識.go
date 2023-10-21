// マップの基礎知識

// パッケージの宣言とインポート
package main
import "fmt"

// main関数の宣言
func main() {
	// マップの宣言と初期化
	scores := map[string]int{
		"佐藤" : 100,
		"鈴木" : 90,
		"田中" : 95,
	}

	// マップのすべての要素を表示
	fmt.Println("マップ : ", scores)
	// 田中さんだけの点数を表示
	fmt.Println("田中さんの点数は", scores["田中"], "点です。")

	// マップの要素の値を変換して削除
	scores["鈴木"] = 88
	delete(scores, "鈴木")
	// マップのすべての要素を表示
	fmt.Println("マップ : ", scores)

	// マップに要素があるかどうかを確認
	score, ok := scores["佐藤"]
	fmt.Println(ok, score)
}

/*
結果 : 
マップ :  map[佐藤:100 田中:95 鈴木:90]
田中さんの点数は 95 点です。
マップ :  map[佐藤:100 田中:95]
true 100
*/