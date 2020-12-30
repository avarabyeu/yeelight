# yeelight
Golang API for [Yeelight](yeelight.com)

Yeelight is simple command line tool and Golang implementation API of Yeelight protocol
with notification listening support

## Installation
To install, run:
```sh
go get github.com/newbits/yeelight
```

## Usage
```go
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
```

## API Specification
Yeelight API Specification [can be found here] (https://www.yeelight.com/download/Yeelight_Inter-Operation_Spec.pdf)

## License
yeelight is distributed under the [MIT license](https://opensource.org/licenses/MIT)

## Legal
Yeelight® is a registered trademark of [Yeelight](https://www.yeelight.com/).

