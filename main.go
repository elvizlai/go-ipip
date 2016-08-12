package main

import (
	"github.com/17mon/go-ipip/ipip"
	"fmt"
	"time"
	"math/rand"
//	"sync"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(8)

	ipip.Load("C:/lovebizhi/tiantexin/17mon/mydata4vipday2.datx")
	var now time.Time
	/*
	now = time.Now()
	var wg sync.WaitGroup
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func(){
			ipip.Find(randomIp())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(now))

	now = time.Now()
	for i := 0; i < 1000000; i++ {
		ipip.Find2(randomIp())
	}
	fmt.Println(time.Now().Sub(now))
*/
	now = time.Now()
	for i := 0; i < 1000000; i++ {
		//ipip.FindLocation(randomIp())
	}
	fmt.Println(time.Now().Sub(now))

	fmt.Println(ipip.FindLocation("1.0.140.156"))
	fmt.Println(ipip.Find2("1.0.206.85"))
	fmt.Println(ipip.Find2("8.8.8.8"))
	fmt.Println(ipip.Find2("118.28.8.8"))
}

func randomIp() string {
	return fmt.Sprintf("%d.%d.%d.0", rand.Intn(255), rand.Intn(255), rand.Intn(255))
}