package main
import ("fmt")

func main() {
  var x = 10
  x=x+5//=x+=5
  fmt.Println(x)
  min()
}

func min() {
  x:= 20
  y:= 18
  if x > y {
    fmt.Println("x is greater than y")
  }
  man()
}


func man() {
  temperature := 14
  if temperature > 15 {
    fmt.Println("It is warm out there")
  } else {                             //注意：else 必須跟在同一行的 if 結束大括號後面，不能換到下一行
    fmt.Println("It is cold out there")
  }
  mn()
}

func mn() {
  time := 22
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
  in()
}

func in() {
  num := 20
  if num >= 10 {
    fmt.Println("Num is more than 10.")
    if num > 15 {
      fmt.Println("Num is also more than 15.")//If裡面再If
     }
  } else {
    fmt.Println("Num is less than 10.")
  }
}
