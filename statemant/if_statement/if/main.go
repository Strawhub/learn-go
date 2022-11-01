// 変数xを整数で割ったとき余りが0であれば出力
package main

import "fmt"

func main() {
    x := 27
    if x%3 == 0 {
        fmt.Println(x, "is even")
    }
}