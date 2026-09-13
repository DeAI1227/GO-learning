package main

import "fmt"

// 本檔練習：函式 func
// 教學格式：看學習手冊第 17 章——每個程式碼概念都有「加了什麼／為什麼這樣寫」

func main() {
	fmt.Println("=== 1) 無參數、無回傳 ===")
	sayHello()

	fmt.Println("=== 2) 有參數 ===")
	greet("惟理")

	fmt.Println("=== 3) 有回傳值 ===")
	sum := add(3, 5)
	fmt.Println("3 + 5 =", sum)

	fmt.Println("=== 4) 多回傳值 ===")
	q, r := divmod(17, 5)
	fmt.Println("17 / 5 → 商=", q, "餘=", r)

	fmt.Println("=== 5) 命名回傳 ===")
	fmt.Println("面積=", rectArea(4, 6))

	fmt.Println("=== 6) 小練習：slice 總和／最大 ===")
	nums := []int{10, 20, 30, 40}
	fmt.Println("總和=", sumSlice(nums))
	fmt.Println("最大=", maxSlice(nums))
}

// 加了什麼：無參數、無回傳的函式
// 為什麼：只做事（印字），不需要算出值回傳
func sayHello() {
	fmt.Println("Hello, Go function!")
}

// 加了什麼：參數 name string
// 為什麼：把「會變的名字」做成參數，同一個函式可服務不同人名
func greet(name string) {
	fmt.Println("你好,", name)
}

// 加了什麼：兩個 int 參數 + 回傳一個 int
// 為什麼：把加法邏輯命名成 add，之後可重用；return 把結果送回呼叫端
func add(a int, b int) int {
	return a + b
}

// 同型別參數可縮寫：func add(a, b int) int
// 為什麼：少重複寫型別，讀起來更乾淨

// 加了什麼：一次回傳兩個 int（商、餘）
// 為什麼：成對結果用多回傳最自然，不必硬塞進怪結構
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

// 加了什麼：命名回傳 area，最後可裸 return
// 為什麼：回傳值先取名，函式內賦值後自動送出；初學先會看即可
func rectArea(width, height int) (area int) {
	area = width * height
	return
}

// 加了什麼：接收 []int，用 range 累加後 return
// 為什麼：你已會的迴圈邏輯，包成函式後任何地方都能重用
func sumSlice(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// 加了什麼：先處理空 slice，再找最大值
// 為什麼：空 slice 不能直接取 [0]；先假設第一個最大再比較較安全
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
