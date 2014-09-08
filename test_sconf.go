package main

import (
	"fmt"
	"github.com/dankozitza/gosb/sconf"
)

func main() {
	conf, err := sconf.Inst("first_init.ini")
	//var conf sconf

	fmt.Println(" main: err: [", err, "]")

	conf["hat"] = "butt"

	fmt.Println(" main:  conf: ", conf)

	conf2, err := sconf.Inst("seconf_init.ini")

	fmt.Println(" main: conf2: ", conf2, "\nerror: [", err, "]")

	return
}
