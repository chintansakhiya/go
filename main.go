package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//var b bytes.Buffer
	wg := &sync.WaitGroup{}
	ch := make(chan int, 3)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch) //1
		fmt.Println(<-ch) //2
		fmt.Println(<-ch) //3
		fmt.Println(<-ch) //4
		fmt.Println(<-ch) //5
		time.Sleep(1 * time.Second)
		fmt.Println("from buffer")
		for k := range ch {

			fmt.Println(k)
		}

		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {

		fmt.Println("Starting now")
		ch <- 50
		ch <- 52
		ch <- 54
		ch <- 56
		ch <- 58
		ch <- 60
		// ch <- 62
		// ch <- 64

		close(ch)
		fmt.Println("Closing now")
		wg.Done()
	}(ch, wg)
	wg.Wait()

}
