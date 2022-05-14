package main
import (
	"fmt"
	"os/user"
	"time"
)


// 一番初めに実行される特別な関数
// パッケージの初期化に用いられる
// main 関数よりも先に呼ばれる
func init() {
	fmt.Println("Init!")
}

// 関数の書き方
func bazz() {
	fmt.Println("Bazz")
}

func main(){
	bazz()
	fmt.Println("Hello","World",time.Now())
	fmt.Println(user.Current())
}