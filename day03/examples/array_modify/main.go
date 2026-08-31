// Day 03 複習：修改陣列元素
// 對照 archive/i.go（第一段）
package main

import "fmt"

func main() {
	prices := [3]int{10, 20, 30}

	prices[2] = 50
	fmt.Println(prices)
}
