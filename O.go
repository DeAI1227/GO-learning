package main

import "fmt"

// O.go｜函式課完整版
// 核心：定義 vs 呼叫；avg 會呼叫 sumSlice
// 跑法：go run O.go

func main() {
	fmt.Println("===== 1) 定義 vs 呼叫 =====")
	// 下面這些都是「呼叫」：函式名後面有 ()
	sayHello()
	greet("惟理")

	sum := add(3, 5) // 呼叫 add，並用 sum 接住回傳值
	fmt.Println("add(3,5) =", sum)
	fmt.Println("sub(10,4) =", sub(10, 4))

	fmt.Println("\n===== 2) 多回傳 =====")
	q, r := divmod(17, 5)
	fmt.Println("divmod(17,5) → 商=", q, "餘=", r)
	qOnly, _ := divmod(17, 5) // _ 丟掉不需要的回傳值
	fmt.Println("只要商 =", qOnly)

	fmt.Println("\n===== 3) 命名回傳 =====")
	fmt.Println("rectArea(4,6) =", rectArea(4, 6))

	fmt.Println("\n===== 4) 函式呼叫函式：avg → sumSlice =====")
	nums := []int{10, 20, 30, 40}
	fmt.Println("nums =", nums)
	fmt.Println("sumSlice(nums) =", sumSlice(nums)) // 先單獨呼叫加總
	fmt.Println("maxSlice(nums) =", maxSlice(nums))
	fmt.Println("avg(nums) =", avg(nums)) // avg 內部會再呼叫 sumSlice
}

// 定義：無參數、無回傳。呼叫寫法：sayHello()
func sayHello() {
	fmt.Println("Hello, Go function!")
}

// 定義：有參數。呼叫寫法：greet("惟理")
func greet(name string) {
	fmt.Println("你好,", name)
}

// 定義：單一回傳。呼叫寫法：sum := add(3, 5)
func add(a, b int) int {
	return a + b
}

// 同型別參數縮寫練習
func sub(a, b int) int {
	return a - b
}

// 多回傳：一次送出商與餘
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

// 命名回傳：回傳值先取名 area
func rectArea(width, height int) (area int) {
	area = width * height
	return
}

// 專責加總。誰需要總和，就呼叫它。
func sumSlice(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 找最大值（先擋空 slice）
func maxSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	for _, n := range nums[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

// 平均 = 總和 / 個數
// 這裡會呼叫 sumSlice：先請它算總和，再拿來回傳值去除
func avg(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	// 執行順序：
	// 1) 先呼叫 sumSlice(nums) 得到總和
	// 2) 轉 float64 後再除以長度
	return float64(sumSlice(nums)) / float64(len(nums))
}
