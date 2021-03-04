package main

import "fmt"

// appned make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)
	
	sl = append(sl, 300) //append(変数, 追加したい値)
	fmt.Println(sl)
	
	sl = append(sl, 400, 500, 600)  //複数の値を追加できる
	fmt.Println(sl) 
	
	sl2 := make([]int, 5)
	fmt.Println(sl2)
	
	fmt.Println(len(sl2))
	
	fmt.Println(cap(sl2))

	sl3 := make([]int, 5,10)  //(型, 要素数, キャパシティ)
	
	fmt.Println(len(sl3))
	
	fmt.Println(cap(sl3))
	
	sl3 = append(sl3, 1,2,3,4,5,6,7)
	
	fmt.Println(len(sl3))
	
	fmt.Println(cap(sl3))
}
