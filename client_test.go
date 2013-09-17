package client

import (
	"testing"
)

func TestSetValue(t *testing.T) {

	var gc GokuyamaClient
	err := gc.Connect("localhost", 8888)
	if err != nil {
		t.Error(err)
	}
	if gc.conn == nil {
		t.Error("gc.conn == nil")
	}

	ret := gc.SetValue("aaa", "aaa")
	if ret != true {
		t.Errorf("setValue result %t, want true", ret)
	}

	getret, _ := gc.GetValue("aaa")
	if getret != "aaa" {
		t.Errorf("getValue result %s, want aaa", getret)
	}

	err = gc.Close()
	if err != nil {
		t.Error(err)
	}

}
