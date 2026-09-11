package main

import "fmt"

// 本檔練習：函式 func
// 對應課：參數、回傳值、多回傳值、命名回傳

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

	fmt.Println("=== 6) 小練習：slice 總和 ===")
	nums := []int{10, 20, 30, 40}
	fmt.Println("總和=", sumSlice(nums))
	fmt.Println("最大=", maxSlice(nums))
}

func sayHello() {
	fmt.Println("Hello, Go function!")
}

func greet(name string) {
	fmt.Println("你好,", name)
}

func add(a int, b int) int {
	return a + b
}

// 同型別參數可縮寫：func add(a, b int) int

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func rectArea(width, height int) (area int) {
	area = width * height
	return // 命名回傳可裸 return
}

func sumSlice(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

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
