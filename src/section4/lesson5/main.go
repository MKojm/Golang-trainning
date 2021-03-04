package main

import "fmt"

func main() {
	var s string = "Hello GOlang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "300"
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	test
	    test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	fmt.Println(string(s[0]))
}
