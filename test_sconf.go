package main

import (
	"fmt"
	"github.com/dankozitza/gosb/sconf"
)

func main() {
	conf1, err := sconf.New("new_test_sconf.ini")
	//var conf sconf

	fmt.Println(" main:   err: ", err)

	conf1["first_key"] = "first_value"

	fmt.Println(" main: conf1: ", conf1)

	conf2 := sconf.Inst()

	conf2["second_key"] = "second_value"

	fmt.Println("main: ran 'conf2[\"second_key\"] = \"second_value\"'")

	fmt.Println(" main: conf1: ", conf1)
	fmt.Println(" main: conf2: ", conf2)

	return
}
