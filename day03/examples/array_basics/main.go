// Day 03 複習：陣列宣告的兩種寫法
// 對照 archive/H.go（第一段）
package main

import "fmt"

func main() {
	var arr1 = [...]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
