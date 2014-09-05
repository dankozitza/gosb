package main

import (
	"./sconf"
	"fmt"
)

func main() {
	conf := sconf.New()

	conf["hat"] = "butt"

	fmt.Println(conf)
	return
}
