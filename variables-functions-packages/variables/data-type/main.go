package main

import (
	"fmt"
	"strconv"
)

//stconvパッケージでstring を int に変換する、またその逆を行う
func main() {
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
