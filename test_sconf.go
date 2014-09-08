package main

import (
	"github.com/dankozitza/gosb/sconf"
	"fmt"
)

func main() {
	conf := sconf.New()

	conf["hat"] = "butt"

	fmt.Println(conf)
	return
}
