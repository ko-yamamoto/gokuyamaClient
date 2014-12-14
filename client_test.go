package gokuyamaClient

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

func TestGetKeysByTag(t *testing.T) {

	gc.SetValueWithTag("ccc", "ddd", "testtag")

	ret, _ := gc.GetKeysByTag("testtag")
	if ret == nil {
		t.Errorf("GetKeysByTag result %s, want [aaa ccc]", ret)
	}

	expected := []string{"aaa", "ccc"}
	for i, tag := range ret {
		if tag != expected[i] {
			t.Errorf("GetKeysByTag result %s, want %s", tag, expected[i])
		}
	}

}

func TestRemoveValueByKey(t *testing.T) {

	ret, err := gc.RemoveValueByKey("aaa")
	if err != nil {
		t.Error(err)
	}
	if ret == false {
		t.Errorf("RemoveValueByKey result %t, want %t", ret, true)
	}

	getret, _ := gc.GetValue("aaa")
	if getret != "" {
		t.Errorf("getValue result %s, want \"\"", getret)
	}

}

func TestClose(t *testing.T) {

	err := gc.Close()
	if err != nil {
		t.Error(err)
	}
}
