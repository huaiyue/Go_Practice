/*
定时器
*/

package main

import (
	"fmt"
	"time"
)

func timer(duration time.Duration) chan bool{
	ch := make(chan bool)

	go func() {
		time.Sleep(duration)
		ch <- true
	}()

	return ch
}

func main()  {
	timeout := timer(time.Second)

	for {
		select {
		case <- timeout:
			fmt.Println("结束啦")
			return
		}
	}
}
