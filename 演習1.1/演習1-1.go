// 演習1.1 
// timeパッケージのNow関数を使って現在時刻を表示するプログラムを作ってください。

// パッケージの宣言及びインポート
package main
import (
	"fmt"
	"time"
)

// main関数の宣言
func main(){
	// fmt.Printlnの呼び出し
	fmt.Println(time.Now())
}

// 結果 : 2023-10-20 21:02:51.207596244 +0530 IST m=+0.000049922