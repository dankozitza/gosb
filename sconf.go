package sconf

import (
	"fmt"
	"encoding/json"
	"io"
	"os"
)

type ErrNilSettingsMap int
func (e ErrNilSettingsMap) Error() string {
	return "sconf: settings map is nil!"
}
type ErrUpdateSettings string
func (e ErrUpdateSettings) Error() string {
	return "sconf: could not update: " + string(e)
}
type ErrSecondCallToNew int
func (e ErrSecondCallToNew) Error() string {
	return "sconf: sconf.New() can not be called a second time!"
}

type sconf map[string]string

var settings sconf = make(sconf)
var config_file_path string
var New_called bool = false

func New(cfp string) (sconf, error) {

	if (New_called) {
		return nil, ErrSecondCallToNew(1)
	}
	New_called = true

	if (settings == nil) {
		return nil, ErrNilSettingsMap(1)
	}

	config_file_path = cfp

	err := settings.Update()
	if (err != nil) {
		panic(ErrUpdateSettings(err.Error()))
	}

	return settings, nil
}

func Inst() sconf {

	if (settings == nil) {
		fmt.Println(ErrNilSettingsMap(1))
	}

	return settings
}

func (s *sconf) Update() error {
	// update settings map from file at config_file_path
	fi, err := os.Open(config_file_path)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	var str_json string
	buff := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		str_json += string(buff[:n])
	}

	fmt.Println("str_json:")
	fmt.Println(str_json)

	settings = nil
	if err := json.Unmarshal([]byte(str_json), &settings); err != nil {
		return err
	}

	fmt.Println("settings:")
	fmt.Println(settings)

	return nil
}

func (s *sconf) Set_config_file_path(path string) {
	config_file_path = path
}

func (s *sconf) Save() bool {
	m_map, err := json.MarshalIndent(settings, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Println("writing settings map:", config_file_path)
	fmt.Println(string(m_map))

	fo, err := os.Create(config_file_path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	fo.WriteString(string(m_map))

	return true
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
//func Walk(t *tree.Tree, ch chan int) {
//	if t.Left != nil {
//		Walk(t.Left, ch)
//	}
//	fmt.Println(t.Value)
//	ch <- t.Value
//	if t.Right != nil {
//		Walk(t.Right, ch)
//	}
//}
//
//// Same determines whether the trees
//// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool {
//	c1 := make(chan int, 10)
//	c2 := make(chan int, 10)
//
//	Walk(t1, c1)
//	Walk(t2, c2)
//	close(c1)
//	close(c2)
//
//	for {
//		i1, ok1 := <-c1
//		i2, ok2 := <-c2
//
//		fmt.Println(i1, " ", ok1, " ", i2, " ", ok2)
//
//		if !ok1 && !ok2 {
//			break
//		}
//		if !ok1 || !ok2 {
//			fmt.Println("trees have different number of values")
//			return false
//		}
//		if i1 != i2 {
//			fmt.Println("values do not match")
//			return false
//		}
//	}
//	return true
//}
//
//func main() {
//
//	ftree := tree.New(1)
//	fmt.Println(ftree.String())
//
//	stree := tree.New(1)
//	fmt.Println(stree.String())
//
//	fmt.Println(Same(ftree, stree))
//}
