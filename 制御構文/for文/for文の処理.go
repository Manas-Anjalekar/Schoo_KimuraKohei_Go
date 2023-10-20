// 制御構文 for文

// パッケージの宣言とインポート
package main
import "fmt"

// main関数の宣言
func main(){
	// 整数型変数の宣言及び初期化
	sum := 0
	// for 初期化式; 継続条件式; 反復式 { 処理 }
	for i := 0; i < 10; i++ {
		// sum整数型変数とiの足し算(省略)
		sum += i
	}
	// fmt.Printlnの呼び出し
	fmt.Println("iが9になるまでの足し算 : ", sum)

	for sum < 1000 {
		// 自分での足し算(省略)
		sum += sum
		// fmt.Printlnの呼び出し
		fmt.Println("1000に超えるまでのsumの値 : ", sum)
	}
}

// 結果 : 
// iが9になるまでの足し算 :  45
// 1000に超えるまでのsumの値 :  90
// 1000に超えるまでのsumの値 :  180
// 1000に超えるまでのsumの値 :  360
// 1000に超えるまでのsumの値 :  720
// 1000に超えるまでのsumの値 :  1440