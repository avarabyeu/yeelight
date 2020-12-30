package yeelight

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/newbits/yeelight/color"
)

type Device struct {
	Address string
	Random  *rand.Rand
}

func (d Device) New() *Device {
	return &d
}

//Power state toggle for light
func (d Device) Power() error {
	_, err := d.executeCommand("toggle")

	return err
}

func (d Device) Color(value string) error {
	c, _ := color.HexStringToRgbInt(value)

	_, err := d.executeCommand("set_rgb", c)

	return err
}

func (d Device) Brightness(brightness int) error {

	_, err := d.executeCommand("set_bright", brightness)
	return err
}

func (d Device) Timer(minutes int) error {
	_, err := d.executeCommand("cron_add", 0, minutes)
	return err
}

func (d Device) StopTimer() error {
	_, err := d.executeCommand("cron_del", 0)
	return err
}

func (d *Device) randID() int {
	i := d.Random.Intn(100)

	return i
}

func (d *Device) newCommand(name string, params []interface{}) *Command {
	return &Command{
		Method: name,
		ID:     d.randID(),
		Params: params,
	}
}

//executeCommand executes command with provided parameters
func (d *Device) executeCommand(name string, params ...interface{}) (*CommandResult, error) {
	return d.execute(d.newCommand(name, params))
}

//executeCommand executes command
func (d *Device) execute(cmd *Command) (*CommandResult, error) {

	conn, err := net.Dial("tcp", d.Address)
	if nil != err {
		return nil, fmt.Errorf("cannot open connection to %s. %s", d.Address, err)
	}

	time.Sleep(time.Second)
	conn.SetReadDeadline(time.Now().Add(timeout))

	//write request/command
	b, _ := json.Marshal(cmd)
	fmt.Println(fmt.Sprintf("%v", string(b)))

	fmt.Fprint(conn, string(b)+crlf)

	//wait and read for response
	res, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("cannot read command result %s", err)
	}

	var rs CommandResult
	err = json.Unmarshal([]byte(res), &rs)

	fmt.Println(string([]byte(res)))

	if nil != err {
		return nil, fmt.Errorf("cannot parse command result %s", err)
	}

	if nil != rs.Error {
		return nil, fmt.Errorf("command execution error. Code: %d, Message: %s", rs.Error.Code, rs.Error.Message)
	}

	return &rs, nil
}

//parseAddr parses address from ssdp response
func parseAddr(msg string) string {
	if strings.HasSuffix(msg, crlf) {
		msg = msg + crlf
	}

	resp, err := http.ReadResponse(bufio.NewReader(strings.NewReader(msg)), nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()

	return strings.TrimPrefix(resp.Header.Get("LOCATION"), "yeelight://")
}

//closeConnection closes network connection
func closeConnection(c net.Conn) {
	if nil != c {
		c.Close()
	}
}
