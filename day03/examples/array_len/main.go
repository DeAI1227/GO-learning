// Day 03 複習：len() 取陣列長度
// 對照 archive/i.go（第四段）
package main

import "fmt"

func main() {
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr2 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}
