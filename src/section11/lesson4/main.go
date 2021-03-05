package main

import "fmt"

//構造体
//埋め込み

type Point struct {
	A int
	B string
	C float64
}

type BigPoint struct {
	Point //データ型と名前が一緒の場合、データ型は省略可
	A     int
}

func (p *Point) Set(i int) {
	p.A = i
}

func main() {
	bp := BigPoint{}
	fmt.Println(bp)
	bp.Point.A = 100
	bp.Point.B = "Apple"
	bp.Point.C = 2.8

	fmt.Println(bp.A)
	bp.A = 1000
	fmt.Println(bp.A)

	bp2 := BigPoint{
		Point: Point{
			A: 100,
			B: "Banana",
			C: 3.5,
		},
	}
	fmt.Println(bp2)

	bp2.Point.Set(2000)
	fmt.Println(bp2)
}
