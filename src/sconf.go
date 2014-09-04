package main

import ()
//	"fmt"
//	"strings"

type sconf map[string]string

var settings sconf
var config_file_path string
var update_config bool

func main() {
	m_sconf := &settings
	if (update_config == true) {
		m_sconf.Update()
	}
}

func (s *sconf) Update() bool {
	// update settings map from file at config_file_path
	return false
}

func (s *sconf) Set_config_file_path(path string) {
	config_file_path = path
}

func (s *sconf) Defined(key string) bool {
	return false
}

func (s *sconf) Get(key string) string {
	return "value"
}

func (s *sconf) Set(key string) bool {
	return false
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
