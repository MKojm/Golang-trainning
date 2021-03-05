package main

//importはパッケージごとに行う
import (
	"fmt"
	. "fmt"
	f "fmt"
	"section13/lesson1/foo"
)

//スコープ

func main() {
	fmt.Println(foo.Max)
	//fmt.Println(foo.min)

	f.Println(foo.ReturnMin())
	Println(foo.Max)
}
