// Package client provides GokuyamaClient for connecting to okuyama servers,
// setting key-value and getting value.
package main

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"strings"
)

// GokuyamaClient has a connection with okuyama
type GokuyamaClient struct {
	conn net.Conn
}

// connect to a okuyama master node
func (gc *GokuyamaClient) Connect(hostname string, portNo int) error {
	var err error
	gc.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", hostname, portNo))
	return err
}

// close connection to a okuyama master node
func (gc *GokuyamaClient) Close() error {
	var err error
	err = gc.conn.Close()
	return err
}

// set a key-value to okuyama
func (gc *GokuyamaClient) SetValue(key string, value string) bool {

	fmt.Fprintf(gc.conn, fmt.Sprintf("1,%s,(B),0,%s\n",
		base64.StdEncoding.EncodeToString([]byte(key)),
		base64.StdEncoding.EncodeToString([]byte(value))))

	status, _ := bufio.NewReader(gc.conn).ReadString('\n')

	if status != "1,true,OK" {
		return true
	} else {
		return false
	}

}

// set a key-value with a tag to okuyama
func (gc *GokuyamaClient) SetValueWithTag(key string, value string, tag string) bool {

	fmt.Fprintf(gc.conn, fmt.Sprintf("1,%s,%s,0,%s\n",
		base64.StdEncoding.EncodeToString([]byte(key)),
		base64.StdEncoding.EncodeToString([]byte(tag)),
		base64.StdEncoding.EncodeToString([]byte(value))))

	status, _ := bufio.NewReader(gc.conn).ReadString('\n')

	if status != "1,true,OK" {
		return true
	} else {
		return false
	}

}

// get value by a key from okuyama
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

// get keys by a tag from okuyama
func (gc *GokuyamaClient) GetKeysByTag(tag string) ([]string, error) {

	fmt.Fprintf(gc.conn, fmt.Sprintf("3,%s,False\n", base64.StdEncoding.EncodeToString([]byte(tag))))

	ret, err := bufio.NewReader(gc.conn).ReadString('\n')
	if ret == "" {
		fmt.Println(err)
	}

	rets := strings.Split(ret, ",")

	if rets[1] == "true" {

		tags := strings.Split(rets[2], ":")

		for i, tag := range tags {
			data, err := base64.StdEncoding.DecodeString(tag)
			if err != nil {
				fmt.Println(err)
			}
			tags[i] = string(data)
		}

		return tags, err
	} else {
		return nil, nil
	}
}

// remove key-value by a key from okuyama
func (gc *GokuyamaClient) RemoveValueByKey(key string) (bool, error) {

	fmt.Fprintf(gc.conn, fmt.Sprintf("5,%s,0\n", base64.StdEncoding.EncodeToString([]byte(key))))

	ret, err := bufio.NewReader(gc.conn).ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return false, err
	}
	if ret == "" {
		return false, errors.New("unknown error")
	}

	rets := strings.Split(ret, ",")

	if rets[1] == "true" {
		return true, err
	} else {
		return false, nil
	}
}
