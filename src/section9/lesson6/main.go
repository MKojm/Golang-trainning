package main

import "fmt"

//map [key]value の型

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}  //暗黙的宣言

	fmt.Println(m2)

	m3 := map[int]string{        //行を変えても行末に,をつければOK
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)   //からのmap
	fmt.Println(m4)

	m4[1] = "JAPAN"              //値の追加
	m4[2] = "USA"
	fmt.Println(m4)

	fmt.Println(m["A"])           //値の取り出し
	fmt.Println(m4[2])
	fmt.Println(m4[3])

	_, ok:= m4[3]
	if !ok{
		fmt.Println("error")
	}
	//fmt.Println(s,ok)

	m4[2] = "US"
	fmt.Println(m4)

	m4[3] = "CHINA"
	fmt.Println(m4)

	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))


}