package main 
import "fmt"
//channel
//複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造
//宣言、捜査

func main() {
	var ch1 chan int  //双方向のchannelとして機能
	
	//受信専用
	//var ch2 <-chan int

	//送信専用
	//var ch3 chan<- int
	
	ch1 = make(chan int)
	
	ch2 := make(chan int)
	
	fmt.Println(cap(ch1)) //容量（バッファサイズ）
	fmt.Println(cap(ch2))
	
	ch3 := make(chan int, 5)   //容量（バッファサイズ）を指定
	fmt.Println(cap(ch3))
	
	ch3 <- 1     //ch3に1というデータを送信
	fmt.Println(len(ch3))  
	
	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))  
	
	i := <-ch3      //ch3からデータを受信
	fmt.Println(i)
	fmt.Println("len", len(ch3)) 
	
	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3)) 
	
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3)) 
	
	ch3 <- 1
	fmt.Println("len", len(ch3)) 
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3)) 
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6
}