package main

import (
	"fmt"
	"github.com/dankozitza/gosb/sconf"
)

func main() {
	conf := sconf.Inst()
	//var conf sconf

	conf["hat"] = "butt"

	fmt.Print("main: conf (", conf, ")")

	return
}
