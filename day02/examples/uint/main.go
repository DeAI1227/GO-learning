// Day 02 複習：uint 無符號整數與 %T 格式
// 對照 archive/G.go
package main

import "fmt"

func main() {
	var x uint = 500
	var y uint = 4500
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)
}
