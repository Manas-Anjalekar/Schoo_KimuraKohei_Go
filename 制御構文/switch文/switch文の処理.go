// 制御構文 switch文

// パッケージの宣言とインポート
package main
import (
	"fmt"
	"math/rand"
	"time"
)

// main関数の宣言
func main(){
	// 整数型変数の宣言及び初期化
	n := time.Now().UnixNano()
	// ランダム化
	rand.Seed(n)
	// ランダム化された値を整数型変数に変換　
	r := rand.Intn(10)
	// fmt.Printlnの呼び出し
	fmt.Println(r)

	// switch文(制御構文)で判断させる
	switch r {
		case 9:
			fmt.Println("エクセレント!")
		case 6, 7, 8:
			fmt.Println("グレート。")
		case 3, 4, 5:
			fmt.Println("グッド。")
		default:
			fmt.Println("ノットバッド。")
	}
}

/*
結果 : 
一回目 : 
9
エクセレント!

二回目 : 
2
ノットバッド。
*/