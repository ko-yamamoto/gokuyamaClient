package client

import (
	"testing"
)

const (
	hostname = "localhost"
	portNo   = 8888
)

var (
	gc GokuyamaClient
)

func TestConnect(t *testing.T) {

	err := gc.Connect(hostname, portNo)

	if err != nil {
		t.Error(err)
	}
	if gc.conn == nil {
		t.Error("gc.conn == nil")
	}

}

func TestSetValue(t *testing.T) {

	ret := gc.SetValue("aaa", "aaa")
	if ret != true {
		t.Errorf("setValue result %t, want true", ret)
	}
}

func TestGetValue(t *testing.T) {

	getret, _ := gc.GetValue("aaa")
	if getret != "aaa" {
		t.Errorf("getValue result %s, want aaa", getret)
	}
}

func TestSetValueWithTag(t *testing.T) {

	ret := gc.SetValueWithTag("aaa", "bbb", "testtag")
	if ret != true {
		t.Errorf("setValueWithTag result %t, want true", ret)
	}

	getret, _ := gc.GetValue("aaa")
	if getret != "bbb" {
		t.Errorf("getValue result %s, want bbb", getret)
	}

}

func TestClose(t *testing.T) {

	err := gc.Close()
	if err != nil {
		t.Error(err)
	}

}
