package main;import("fmt")

func main (){
	a:=[]int{2,3}
	fmt.Println("lengh=",len(a))
	fmt.Print("容量=",cap(a))
	fmt.Print("切片練習")
	acd()//註解
}
func acd(){
	var b=[5]int{1:2}
	c:=b[1:4]//切片：取陣列的一部分
	fmt.Println(c,b)
	cd()
}
func cd() {
  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  // with omitted capacity
  myslice2 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))
}
