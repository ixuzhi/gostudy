package main

import "fmt"
/*
信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：

ch := make(chan int, 100)
仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞
*/
func main(){
	ch := make(chan int ,2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	ch1 := make(chan int ,1)
	ch1 <- 1
	fmt.Println(<-ch1)
	ch1 <- 2
	fmt.Println(<-ch1)
	
	fmt.Println(<-ch1)

	ch2 := make(chan int ,1)
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

}