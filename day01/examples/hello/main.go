// Day 01 複習：第一個程式、package、import、外部套件
// 對照 archive/hello.go
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
}
