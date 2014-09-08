package main

import (
	"fmt"
	"github.com/dankozitza/gosb/sconf"
)

func main() {
	conf := sconf.Inst()
	//var conf sconf

	conf["hat"] = "butt"

	fmt.Println(" main:  conf: ", conf)

	conf2 := sconf.Inst()

	fmt.Println(" main: conf2: ", conf2)

	return
}
