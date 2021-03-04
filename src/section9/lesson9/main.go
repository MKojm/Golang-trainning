package main
import (
	"fmt"
	"time"
)
//channel
//チャネルとゴルーチン

func reciever(c chan int){
	for {
		i := <-c           //channelにデータが入ったらデータを受信して出力
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	//fmt.Println(<-ch1)
	go reciever(ch1)    //channelにデータが入るのを待機
	go reciever(ch2)
	
	i := 0
	for i < 100{
		ch1 <- i    //channelにデータを送信
		ch2 <- i
		time.Sleep(50*time.Millisecond)
		i++
	}

}