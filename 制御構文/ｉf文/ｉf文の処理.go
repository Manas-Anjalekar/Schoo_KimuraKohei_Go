// 制御構文 if文

// パッケージの宣言とインポート
package main
import "fmt"

// main関数の宣言
func main(){
	// 整数型変数の宣言及び初期化
	x := 10
	// 制御構文で判断させる
	// xを2で割ると余りが0だった場合、処理する
	if x%2 == 0{
		// fmt.Printlnの呼び出し
		fmt.Println(x, "は偶数です。")

		// xを5で割ると余りが0だった場合、処理する
		if x%5 == 0 {
			// fmt.Printlnの呼び出し
			fmt.Println(x, "は5の倍数です。")
		}
	} else {
		// fmt.Printlnの呼び出し
		fmt.Println(x, "は奇数です。")
	}
}

// 結果 : 10 は偶数です。10 は5の倍数です。