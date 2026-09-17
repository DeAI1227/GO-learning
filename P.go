package main

import "fmt"

// P.go｜map 課（30 分鐘）
// 重點：鍵→值、value/ok 安全讀取、函式包查詢
// 跑法：go run P.go

func main() {
	fmt.Println("===== 1) 建立 map（字面量）=====")
	ages := map[string]int{
		"惟理": 20,
		"小明": 18,
	}
	fmt.Println(ages)

	fmt.Println("\n===== 2) make 建立空 map 再新增 =====")
	scores := make(map[string]int)
	scores["國文"] = 90
	scores["數學"] = 85
	fmt.Println(scores)

	fmt.Println("\n===== 3) 讀取：簡單版 vs 安全版 =====")
	fmt.Println("惟理的年齡 =", ages["惟理"]) // 確定存在可用
	value, ok := ages["小華"]              // 不確定時用 ok
	if ok {
		fmt.Println("找到小華:", value)
	} else {
		fmt.Println("沒有小華這個鍵（ok=false）")
	}

	fmt.Println("\n===== 4) 新增／修改／刪除 =====")
	ages["小美"] = 19  // 新增
	ages["惟理"] = 21  // 修改
	delete(ages, "小明") // 刪除
	fmt.Println(ages)

	fmt.Println("\n===== 5) range 遍歷 =====")
	for name, age := range ages {
		fmt.Println(name, age)
	}

	fmt.Println("\n===== 6) 包成函式查詢（接上堂課）=====")
	age, found := getAge(ages, "惟理")
	if found {
		fmt.Println("getAge 找到惟理 =", age)
	}
	_, found = getAge(ages, "小明")
	fmt.Println("getAge 找小明 found =", found)

	fmt.Println("\n===== 7) 小練習：名字→城市 =====")
	cities := map[string]string{
		"惟理": "台北",
		"小美": "高雄",
	}
	fmt.Println("hasCity(惟理) =", hasCity(cities, "惟理"))
	fmt.Println("hasCity(小明) =", hasCity(cities, "小明"))
}

// 查年齡：回傳 (值, 是否找到)
func getAge(ages map[string]int, name string) (int, bool) {
	value, ok := ages[name]
	return value, ok
}

// 是否有這個名字的城市資料
func hasCity(m map[string]string, name string) bool {
	_, ok := m[name]
	return ok
}
