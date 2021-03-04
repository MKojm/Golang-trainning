package main 

import "fmt"
//スライス、copy

func main() {
	
	/*
	sl := []int{100, 200}
	sl2 := sl
	
	sl2[0] = 1000    //sl[0]もsl2[0]も指しているアドレスは同じ→どちらの値も変更される（参照型）
	fmt.Println(sl)
	
	var i int = 10
	i2 := i
	i2 = 100         //i2の値だけ変更される
	fmt.Println(i, i2)
	*/
	
	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)
	
	n := copy(sl2, sl)
	
	fmt.Println(n, sl2) //n=コピーに成功した数
}