// 演習1.7
/*
for~range文を使ってマップscoresの要素をすべて表示するコードを追加してください。
*/

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

	// for制御構文でマップにアクセス
	for k, v:= range scores {
		// マップの要素を表示
		fmt.Println(k, v)
	}
}

/*
結果 : 
田中 95
佐藤 100
鈴木 90
*/