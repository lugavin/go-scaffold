package test

import (
	"fmt"
	"testing"
	"time"
)

/*
 * 当使用无缓冲的channel时，发送和接收操作都是阻塞的。也就是说，发送操作必须等待接收方准备好接收数据，而接收操作也必须等待发送方准备好发送数据。
 */
func TestChannel(t *testing.T) {
	// 创建一个无缓冲的channel
	ch := make(chan int)
	// 创建一个容量为3的有缓冲的channel
	//ch := make(chan int, 3)

	// 启动一个goroutine，向channel中发送数据
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Send before:", i)
			ch <- i
			fmt.Println("Send after:", i)
		}
		// 关闭channel
		close(ch)
	}()

	// 从channel中接收数据，并打印出来
	for {
		// 通过ok来判断channel是否已经关闭
		if val, ok := <-ch; ok {
			fmt.Println("Received:", val)
		} else {
			fmt.Println("Channel closed!")
			break
		}
	}
}

func TestSelect(t *testing.T) {
	// 创建两个channel
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 启动一个goroutine，向ch1中发送数据
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(500 * time.Millisecond)
		}
		close(ch1)
	}()

	// 启动一个goroutine，向ch2中发送数据
	go func() {
		for i := 5; i < 10; i++ {
			ch2 <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(ch2)
	}()

	// 在两个channel中等待数据
	for {
		select {
		case val, ok := <-ch1:
			if ok {
				fmt.Println("Received from ch1:", val)
			} else {
				fmt.Println("ch1 closed")
			}
		case val, ok := <-ch2:
			if ok {
				fmt.Println("Received from ch2:", val)
			} else {
				fmt.Println("ch2 closed")
			}
		default:
			fmt.Println("No data received")
		}
	}
}
