package main

import (
	"fmt"
	"github.com/dankozitza/gosb/sconf"
)

func main() {
	conf := sconf.Inst()
	//var conf sconf

	conf["hat"] = "butt"

	fmt.Println("main: conf (", conf, ")")

	return
}
