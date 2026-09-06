package main
import ("fmt")

func main() {
  for i:=0; i <= 100; i+=10 {         //第一段是條件；第二段是每次判斷；第三段是每次迴圈結束後執行
    fmt.Println(i)
  }
  mn()
}
func mn() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      continue      //continue 會跳過這次剩下的程式，進入下一圈  //如果改成 break 會直接結束迴圈
    }
   fmt.Println(i)
  }
  in()
}

func in() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {                         //巢狀迴圈範例
      fmt.Println(adj[i],fruits[j])
    }
  }
  n()
}

func n() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)
  }
}//如果只要值不要 idx，可用底線 _ 丟掉第一個回傳值
