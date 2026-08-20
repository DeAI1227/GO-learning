package main
import ("fmt")

func main() {
  var i,j string = "Hello","World"

  fmt.Println(i,j)
  fmt.Print(i,j",\n")
  fmt.Printf("i has value: %v and type: %T\n", i, i)
  fmt.Printf("%#v%%\n",i)
}