package main

import (
	"github.com/17mon/go-ipip/ipip"
	"fmt"
	"time"
	"math/rand"
)

func main() {

	ipip.Load("C:/lovebizhi/tiantexin/17mon/mydata4vipday2.datx")
	var now time.Time
	now = time.Now()
	for i := 0; i < 1000000; i++ {
		ipip.Find(randomIp())
	}
	fmt.Println(time.Now().Sub(now))

	now = time.Now()
	for i := 0; i < 1000000; i++ {
		ipip.Find2(randomIp())
	}
	fmt.Println(time.Now().Sub(now))

	now = time.Now()
	for i := 0; i < 1000000; i++ {
		ipip.FindLocation(randomIp())
	}
	fmt.Println(time.Now().Sub(now))

	fmt.Println(ipip.FindLocation("118.228.228.8"))
	fmt.Println(ipip.Find2("118.228.228.8"))

	fmt.Println(randomIp())
	fmt.Println(randomIp())
}

func randomIp() string {
	return fmt.Sprintf("%d.%d.%d.0", rand.Intn(255), rand.Intn(255), rand.Intn(255))
}