// Day 02 複習：fmt.Println / Print / Printf 與格式字串
// 對照 archive/D.go（原版 fmt.Print(i,j",\n") 語法錯誤，這裡已修正）
package main

import "fmt"

func main() {
	var i, j string = "Hello", "World"

	fmt.Println(i, j)
	fmt.Print(i, j, "\n")
	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("%#v%%\n", i)
}
