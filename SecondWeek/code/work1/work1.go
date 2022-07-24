package main

import (
	"fmt"
	"sync"
)

const (
	PRODUCEQUANTITY = 5  // 每个师傅生产5份羊肉串
	STORECAPACITY   = 10 // 放羊肉串容器容量为10
)

var (
	yrcStore       chan string    // 存放羊肉串
	prodCompletion chan string    // 记录是否是否烤完羊肉串
	judgeProdComp  bool           // 判断全部师傅是否烤完全部羊肉串
	wg             sync.WaitGroup // 等待组
)

func Producer(in chan<- string, producerId string, producerNum int) {
	// 师傅生产羊肉串并放进容器里，每次只能往容器里放一份
	defer wg.Done()
	for i := 1; i <= PRODUCEQUANTITY; i++ {
		if judgeProdComp == true {
			return
		}
		s := fmt.Sprintf("YRC-%v-%v", producerId, i)
		fmt.Printf("师傅%v 烤了羊肉串%v ...\n", producerId, i)
		in <- s
	}
	prodCompletion <- producerId
	if len(prodCompletion) >= producerNum {
		close(yrcStore)
		judgeProdComp = true
	}
}

func Consumer(out <-chan string) {
	// 顾客从容器里拿，每次只能从容器中取一份
	defer wg.Done()
	for {
		s, ok := <-out
		if !ok {
			fmt.Println("羊肉串卖完了，收摊，下次再来！")
			return
		}
		fmt.Printf("羊肉串%v已卖出...\n", s)
	}
}

func main() {
	pNum := 3 // 烤串师傅人数
	yrcStore = make(chan string, STORECAPACITY)
	prodCompletion = make(chan string, pNum)
	judgeProdComp = false

	for i := 1; i <= pNum; i++ {
		wg.Add(1)
		go Producer(yrcStore, fmt.Sprintf("P%v", i), pNum)
	}
	wg.Add(1)
	go Consumer(yrcStore)
	wg.Wait()
}

