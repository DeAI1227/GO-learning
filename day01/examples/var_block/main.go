// Day 01 複習：區塊式 var (...) 與多重賦值
// 對照 archive/B.go
package main

import "fmt"

func main() {
	var (
		A      int
		B      string = "as"
		C, d        = 1, 2
	)
	fmt.Println(A, B, C, d)
}
