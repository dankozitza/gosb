package main

import (
	"fmt"
	"github.com/dankozitza/gosb/sconf"
)

func main() {
	conf := sconf.Inst()
	//var conf make(sconf)

	conf["hat"] = "butt"

	fmt.Println(conf)
	return
}
