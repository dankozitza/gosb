package sconf

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"syscall"
	"testing"
)

var file_path string = "sconf_test.json"
var pconf Sconf = Init(file_path, Sconf{"pkey": "pval"})

func TestInst(t *testing.T) {

	conf := Inst()

	fmt.Println("Init conf:", conf)
}

func TestMap(t *testing.T) {

	conf := Inst()

	d := dummy{"structval", 3}

	conf["key1"] = "val1"
	conf["key2"] = 2
	conf["key3"] = d

	fmt.Println("TestMap: conf:", conf)
}

func TestNew(t *testing.T) {

	conf := New("blah", nil)
	fmt.Println("New conf:", conf)

	if conf["key1"] != nil {
		t.Fail()
	}
}

func TestContents(t *testing.T) {

	conf := Inst()

	good := true

	d2 := dummy{"structval", 3}

	var dconf map[string]interface{} = make(map[string]interface{})
	dconf["pkey"] = "pval"
	dconf["key1"] = "val1"
	dconf["key2"] = 2
	dconf["key3"] = d2

	if conf["key1"] != dconf["key1"] {
		good = false
	}
	if conf["key2"] != dconf["key2"] {
		good = false
	}
	if conf["key3"] != dconf["key3"] {
		good = false
	}

	if !good {
		fmt.Println("TestContents: sconf map does not match dummy map")
		fmt.Println("sconf:", conf)
		fmt.Println("dummy:", dconf)
		t.Fail()
	}
}

func TestSave(t *testing.T) {
	conf := Inst()
	conf.Save(file_path)

	dummyfile := "{\n   \"key1\": \"val1\",\n   \"key2\": 2,\n   \"key3\": {},\n   \"pkey\": \"pval\"\n}"

	fi, err := os.Open(file_path)
	if err != nil {
		fmt.Println("TestSave: could not open saved config file:", file_path)
		t.Fail()
		return
	}

	buff := make([]byte, 1024)
	n, err := fi.Read(buff)
	if err != nil && err != io.EOF {
		fmt.Println("TestSave: could not read from config file:", file_path,
			"err:", err)
		t.Fail()
		return
	}
	if string(buff[:n]) != dummyfile {
		fmt.Println("TestSave: config file does not match dummy file:")
		fmt.Println("saved file:", string(buff[:n]))
		fmt.Println("dummy file:", dummyfile)
		t.Fail()
	}
}

func TestHTTPHandler(t *testing.T) {
	var jsm HTTPHandler
	http.Handle("/sconf", jsm)
	fmt.Println("TestHTTPHandler: added handler to http")
	//http.ListenAndServe("localhost:9000", nil)
}

func TestClean(t *testing.T) {
	fmt.Println("TestClean: removing", file_path)
	syscall.Exec("/usr/bin/rm", []string{"rm", file_path}, nil)
}

type dummy struct {
	s string
	i int
}
