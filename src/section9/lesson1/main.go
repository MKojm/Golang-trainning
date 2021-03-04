package main

import "fmt"
//スライス、宣言、操作
func main() {
	
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2) //値を格納
	
	sl3 := []string{"A", "B"}  //配列との違いは、[]の中に要素数を指定しない
	fmt.Println(sl3)
	
	sl4 := make([]int, 5)  //make([]型, 要素数)　int型の初期値は0
	fmt.Println(sl4)

	sl2[0] = 1000        //値の代入
	fmt.Println(sl2)
	
	sl5 := []int{1,2,3,4,5}  
	fmt.Println(sl5)
	
	fmt.Println(sl5[0])  //値の取り出し
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:2])
	fmt.Println(sl5[2:])
	fmt.Println(sl5[:])
	fmt.Println(sl5[len(sl5)-1])
	fmt.Println(sl5[1 : len(sl5)-1])
}