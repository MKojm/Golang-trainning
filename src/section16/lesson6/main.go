package main

import (
	"fmt"
	"os"
)

//fmt

func main() {
	/*
		fmt.Println("表示")

		//fmt標準
		fmt.Print("Hello")
		//改行
		fmt.Println("Hello!")
	*/
	/*
		//Println系　各々の文字列は半角スペースで区切られ、文字列の最後に改行が入る
		fmt.Println("Hello", "world!")
		fmt.Println("Hello", "world!")
	*/

	/*
		//Printf系　フォーマットを指定
		fmt.Printf("%s\n", "Hello")
		fmt.Printf("%#v\n", "Hello")
	*/

	/*
		//Sprint系　出力ではなくフォーマットした結果を文字列で返す。
		s := fmt.Sprint("Hello")
		s1 := fmt.Sprintf("%v\n", "Hello")
		s2 := fmt.Sprint("Hello")
		fmt.Println(s)
		fmt.Println(s1)
		fmt.Println(s2)
	*/

	//Fprint系　書き込み先指定
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("test1.txt")
	defer f.Close()

	fmt.Fprintln(f, "Fprint")

}
