// Day 02 複習：bool 的四種宣告方式
// 對照 archive/F.go
package main

import "fmt"

func main() {
	var b1 bool = true // 有型別 + 初始值
	var b2 = true        // 型別推斷 + 初始值
	var b3 bool          // 有型別，零值為 false
	b4 := true           // 短宣告

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)
}
