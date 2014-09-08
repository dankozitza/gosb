package main

import (
	"fmt"
	"github.com/dankozitza/gosb/sconf"
)

func main() {
	//conf := sconf.New()
	var conf sconf

	conf["hat"] = "butt"

	fmt.Println(conf)
	return
}
