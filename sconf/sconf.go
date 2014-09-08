package sconf

import (
	"fmt"
//	"github.com/vaughan0/go-ini"
)

type ErrNoConfigPath int
func (e ErrNoConfigPath) Error() string {
	return "sconf: no configuration file path was set"
}
type ErrNilSettingsMap int
func (e ErrNilSettingsMap) Error() string {
	return "sconf: settings map is nil!"
}
type ErrUpdateSettings string
func (e ErrUpdateSettings) Error() string {
	return "sconf: could not update settings map from ini file: [" + e + "]"
}

type sconf map[string]string

var settings sconf = make(sconf)
var config_file_path string
var update_config bool = true

func Inst(cfp string) (sconf, error) {

	if (settings == nil) {
		return nil, ErrNilSettingsMap()
	}

	if (config_file_path == nil && cfp == nil) {
		return nil, ErrNoConfigPath()
	}

	if (cfp != nil) {
		config_file_path = cfp
	}

	fmt.Println("sconf: settings map is not nil: ", settings)

	if (update_config) {
		ret := settings.Update()
		if (ret == false) {
			return nil, ErrUpdateSettings("pretend/config_file_path.ini")
		}
		update_config = false
	}

	return settings, nil

	//settings := make(sconf)
	//settings["sconf_test"] = "from_first_Inst()"
	//fmt.Print("sconf: settings map has been initialized (")
	//fmt.Println(settings, ")")

	//m_sconf := &settings
	//if (update_config == true) {
	//	m_sconf.Update()
	//}
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
