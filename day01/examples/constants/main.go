// Day 01 複習：const 常數
// 對照 archive/C.go（原版 B 重複定義會編譯失敗，這裡已修正）
package main

import "fmt"

const (
	A int    = 1
	B        = 3.14
	C        = "Hi!"
	MaxRetry = 26 // 原本第二個 B 改名，避免重複
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(MaxRetry)
}
