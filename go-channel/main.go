package main

import (
	"fmt"
	"time"
)

func main() {
	// 只读 channel
	// var readOnlyChan <-chan int

	// // 只写channel
	// var writeOnlyChan chan<- int

	// // 读写channel
	// var ch chan int

	// // 直接初始化 make
	// readOnlyChan1 := make(<-chan int, 2)  // 只读且带缓存区的channel
	// readOnlyChan2 := make(<-chan int)     // 只读且不带缓存区的channel

	// writeOnlyChan1 := make(chan<- int, 4)  // 只写且带缓存区的channel
	// witeOnlyChan2 := make(chan<- int) //只写不带缓存区的channel

	// ch := make(chan int, 10)

	// ch := make(chan int)
	// go foo1(ch)
	// fmt.Println(<-ch)

	// ch2 := make(chan int)
	// go unbufferChan(ch2)
	// for i:=0; i<10; i++ {
    //     fmt.Println("received", <-ch2)
    // }

	ch3 := make(chan string, 3)
	ch3 <- "tom"
	ch3 <- "jerry"
	ch3 <- "bill"
	close(ch3) // close
	bufferChan(ch3)

	ch := make(chan int)
	quitChan := make(chan int)
	go func() {
		for i:=0; i<10; i++ {
			fmt.Println(<-ch)
		}
		quitChan <- 0
	}()
	fibonacci(ch, quitChan)

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sums(s[:len(s)/2], c)
	go sums(s[len(s)/2:], c)
	fmt.Println("Sum:", <-c + <-c)

	ch4 := make(chan bool, 1)
	go worker(ch4)
	<-ch4

	// channel 超时处理
	ch5 := make(chan int)
	quitChan2 := make(chan bool)
	go func() {
        for {
			select {
			case v := <- ch5:
				fmt.Println(v)
			case <-time.After(time.Second * time.Duration(3)):
				quitChan2 <- true
				fmt.Println("timeout, send notice")
				return
			}
		}
    }()
	for i:=0; i<4; i++ {
		ch5 <- i
	}

	<-quitChan2  // 输出值，相当于收到通知，解除主程序阻塞
	fmt.Println("main quit out")


	//死锁：未初始化数据，写入，读出，关闭 都会造成死锁
	/*
	1. channel 要用make 进行初始化操作
	2.  读取和写入都要配对出现，并且不能在同一个goroutine中
	3. 一定先用go起一个协程执行读取或者写入操作
	4. 多次写入数据，for读取数据时，写入这注意关闭channel 
	*/

}

func foo1(ch chan<- int) {
	ch <- 10
}

// 不带缓冲区的channel
func unbufferChan(ch chan int) {
	for i:=0; i<10; i++ {
		fmt.Println("send", i)
		ch <- i
	}
}

// 带缓冲区的channel
func bufferChan(ch chan string) {
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println("received", v)
		}
	}
}

func fibonacci(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("fibonacci quit")
            return
		}
	}
}

func sums(s []int, c chan int) {
	sum := 0
	for _, v := range s {
        sum += v
    }
	c <- sum
}

func worker(done chan bool) {
	fmt.Println("Worker started")
    defer fmt.Println("Worker finished")
    done <- true
}