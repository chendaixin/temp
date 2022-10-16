package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	// 2006-01-02 15:04:05为go的标准格式
	timeFormat1 = "2006-01-02 15:04:05"
	timeFormat2 = "2006-01-02-15-04-05"
	timeFormat3 = "2006-01-02 15:04:05.000" // 可以精确到毫秒ms
)

type subConfig struct {
	IPAddress string
	Port      int
	subInfo   string
}

func handleGO(p int, sw *sync.WaitGroup) {
	fmt.Printf("enter Go route[%d]\n", p)
	ticker := time.NewTicker(time.Second)
	for i := 0; i < p; i++ {
		<-ticker.C
	}
	ticker.Stop()
	fmt.Printf("leave Go route[%d]\n", p)

	sw.Done()
}

// 测试多个携程同步
func testMutilGO() {
	var wg sync.WaitGroup
	p := 3
	for ; p > 0; p-- {
		//每开辟一个协程就向等待组中+1
		wg.Add(1)

		go handleGO(p, &wg)
	}

	//阻塞等待wg中的协程数归零
	wg.Wait()
}

func handlenOnvifNotify(p int, sub subConfig, sw *sync.WaitGroup) {
	fmt.Printf("|%s| enter handlenOnvif Notify Go route[%d]\n", time.Now().Format(timeFormat3), p)

	ticker := time.(time.Second)
	for i := 0; i < 5; i++ {
		<-ticker.C
		fmt.Printf("|%s| Go route[%d],handing! [%d]\n", time.Now().Format(timeFormat3), p, sub.Port)
	}
	ticker.Stop()

	fmt.Printf("|%s| leave Go route[%d]\n", time.Now().Format(timeFormat3), p)
	sw.Done()
}

// 测试多个携程同步
func testMutilGOSub() {
	var wg sync.WaitGroup

	for p := 5; p > 0; p-- {
		var tmpConfig subConfig
		tmpConfig.IPAddress = "192.168.1.1"
		tmpConfig.Port = p
		tmpConfig.subInfo = "onvif"

		//每开辟一个协程就向等待组中+1
		wg.Add(1)

		go handlenOnvifNotify(p, tmpConfig, &wg)
	}

	//阻塞等待wg中的协程数归零
	wg.Wait()
}

func main() {
	fmt.Println("hello world")

	//testMutilGO()

	testMutilGOSub()

	fmt.Println("main over")
}
