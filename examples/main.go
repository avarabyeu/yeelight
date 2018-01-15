package main

import (
	"context"
	"fmt"
	"github.com/avarabyeu/yeelight"
	"log"
	"time"
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
		<-time.After(time.Minute * 30)
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
