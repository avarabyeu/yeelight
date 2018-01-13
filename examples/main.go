package main

import (
	"github.com/avarabyeu/yeelight"
	"github.com/prometheus/common/log"
	"fmt"
	"time"
	"context"
)

func main() {
	y, err := yeelight.Discover()
	checkError(err)

	on, err := y.GetProp("power")
	checkError(err)
	fmt.Printf("Power is %s", on[0].(string))

	notifications, done, err := y.Listen()
	checkError(err)
	go func() {
		<-time.After(time.Second)
		done <- struct{}{}
	}()
	for n := range notifications {
		fmt.Println(n)
	}

	context.Background().Done()

}

func checkError(err error) {
	if nil != err {
		log.Fatal(err)
	}
}
