package main

import (
	"fmt"
	"strings"
)

//strings

func main() {
	/*
		//文字列を結合する
		s1 := strings.Join([]string{"A","B","C"},",")
		s2 := strings.Join([]string{"A","B","C"},"")
		fmt.Println(s1,s2)
	*/
	/*
		//文字列に含まれる部分文字列を検索する(返り値が-1の場合は部分文字がないことを表す)
		i1 := strings.Index("ABCDE", "A")
		i2 := strings.Index("ABCDE", "D")
		fmt.Println(i1, i2)

		i3 := strings.LastIndex("ABCDABCD", "BC")
		fmt.Println(i3)

		i4 := strings.IndexAny("ABC", "ABC")

		i5 := strings.LastIndexAny("ABC", "ABC")
		fmt.Println(i4, i5)

		b1 := strings.HasPrefix("ABC", "A") //Aで始まるか
		b2 := strings.HasSuffix("ABC", "C") //Cで終わるか
		fmt.Println(b1, b2)

		b3 := strings.Contains("ABC", "B")
		b4 := strings.ContainsAny("ABCDE", "BD")
		fmt.Println(b3, b4)

		i6 := strings.Count("ABC", "B")
		i7 := strings.Count("ABC", "")  //文字列の長さ+1
		fmt.Println(i6, i7)
	*/
	/*
		//文字列を繰り返して結合する(第二引数に負の整数を指定するとエラー)
		s3 := strings.Repeat("ABC", 4)
		s4 := strings.Repeat("ABC", 0)
		fmt.Println(s3, s4)
	*/
	/*
		//	文字列の置換
		s5 := strings.Replace("AAAAAA", "A", "B", 1)
		s6 := strings.Replace("AAAAAA", "A", "B", -1) //該当するすべての文字を置換
		fmt.Println(s5, s6)
	*/
	/*
		//文字列を分割する
		s7 := strings.Split("A,B,C,D,E", ",")
		fmt.Println(s7)
		s8 := strings.SplitAfter("A,B,C,D,E", ",")
		fmt.Println(s8)
		s9 := strings.SplitN("A,B,C,D,E", ",", 2)
		fmt.Println(s9)
		s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
		fmt.Println(s10)
	*/
	/*
		//大文字、小文字の変換
		sl1 := strings.ToLower("ABC")
		sl2 := strings.ToLower("E")

		sl3 := strings.ToUpper("abc")
		sl4 := strings.ToUpper("e")
		fmt.Println(sl1, sl2, sl3, sl4)
	*/
	/*
		//文字列から空白を取り除く
		sl5 := strings.TrimSpace("   -   ABC    -    ")
		//全角
		sl6 := strings.TrimSpace("　　　　ABC　　　　")
		fmt.Println(sl5, sl6)
	*/

	//文字列からスペースで区切られたフィールドを取り出す
	s18 := strings.Fields("a b c")
	fmt.Println(s18)
	fmt.Println(s18[1])

}
