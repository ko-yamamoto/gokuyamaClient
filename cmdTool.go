package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	cmdHostname = flag.String("h", "localhost", "hostname of okuyama master node")
	cmdPortNo   = flag.Int("p", 8888, "port number of okuyama master node")
	key         = flag.String("k", "", "key for setting or getting value")
	value       = flag.String("v", "", "value for setting to okuyama")
)

func main() {

	flag.Parse()

	if *key == "" && *value == "" {
		fmt.Println("Option -k is required")
		os.Exit(1)
	}

	var gc GokuyamaClient
	gc.Connect(*cmdHostname, *cmdPortNo)

	if *value == "" {
		// get value
		ret, err := gc.GetValue(*key)

		if err != nil {
			fmt.Errorf("error: %s\n", err)
		} else {
			fmt.Printf("result: %s\n", ret)
		}

	} else {
		// set key-value
		ret := gc.SetValue(*key, *value)
		if ret == true {
			fmt.Println("registerd")
		} else {
			fmt.Errorf("error: %s\n", ret)
		}
	}
}
