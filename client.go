package client

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
	"strings"
)

type GokuyamaClient struct {
	conn net.Conn
}

func (gc *GokuyamaClient) Connect(hostname string, portNo int) error {
	var err error
	gc.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", hostname, portNo))
	return err
}

func (gc *GokuyamaClient) Close() error {
	var err error
	err = gc.conn.Close()
	return err
}

func (gc *GokuyamaClient) SetValue(key string, value string) bool {

	fmt.Fprintf(gc.conn, fmt.Sprintf("1,%s,(B),0,%s\n", base64.StdEncoding.EncodeToString([]byte(key)), base64.StdEncoding.EncodeToString([]byte(value))))
	status, _ := bufio.NewReader(gc.conn).ReadString('\n')

	if status != "1,true,OK" {
		return true
	} else {
		return false
	}

}

func (gc *GokuyamaClient) GetValue(key string) (string, error) {

	fmt.Fprintf(gc.conn, fmt.Sprintf("2,%s\n", base64.StdEncoding.EncodeToString([]byte(key))))
	ret, err := bufio.NewReader(gc.conn).ReadString('\n')

	if ret == "" {
		fmt.Println(err)
	}

	rets := strings.Split(ret, ",")

	if rets[1] == "true" {
		data, err := base64.StdEncoding.DecodeString(rets[2])
		if err != nil {
			fmt.Println(err)
		}
		return string(data), err
	} else {
		return "", nil
	}
}
