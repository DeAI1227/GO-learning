// Day 03 複習：未初始化 / 部分 / 完整初始化
// 對照 archive/i.go（第二段）
package main

import "fmt"

func main() {
	arr1 := [5]int{}              // 零值
	arr2 := [5]int{1, 2}          // 部分初始化
	arr3 := [5]int{1, 2, 3, 4, 5} // 完整初始化

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
