package sconf

import (
	"fmt"
	"encoding/json"
	"io"
	"os"
	"net/http"
	//"github.com/dankozitza/seestack"
	//"github.com/dankozitza/statshare"
)

type ErrSconfGeneric string
func (e ErrSconfGeneric) Error() string {
	stat.Status = "ERROR"
	stat.Message = "error: " + e.Error()

	return stat.Message
}

type ErrUpdateSettings string
func (e ErrUpdateSettings) Error() string {
	return "Sconf could not update: " + string(e)
}

type Sconf map[string]string

var settings Sconf = make(Sconf)
var config_file_path string
var New_called bool = false
//var stat statshare.Statshare = statshare.New("test")
var stat SconfStatshare = SconfStatshare{"PASS", "sconf package initialized"}

func New(cfp string) (Sconf, error) {
	stat.Status = "PASS"
	stat.Message = "sconf object initialized"

	if (New_called) {
		//return nil, stat.Err("New() cannot be called a second time!")
		return nil, ErrSconfGeneric("New() cannot be called a second time!")
		//panic(ErrSconfGeneric("New() cannot be called a second time!"))
	}
	New_called = true

	if (settings == nil) {
		//return nil, stat.Err("settings map cannot be nil!")
		return nil, ErrSconfGeneric("settings map cannot be nil!")
		//panic(ErrSconfGeneric("settings map cannot be nil!"))
	}

	config_file_path = cfp

	err := settings.Update()
	if (err != nil) {
		//stat.PanicErr("failed to update settings", err)
		panic(ErrSconfGeneric("failed to update settings: " + err.Error()))
	}

	return settings, nil
}

func Inst() Sconf {
	if (settings == nil) {
		panic(ErrSconfGeneric("settings map cannot be nil!"))
	}
	return settings
}

func (s *Sconf) Set_config_file_path(path string) {
	config_file_path = path
}

func (s *Sconf) Update() error {
	// update settings map from file at config_file_path
	fi, err := os.Open(config_file_path)
	if err != nil {
		//stat.Warn("failed to open config file: " + err.Error())
		stat.Status = "WARN"
		stat.Message = "failed to open config file: " + err.Error()
		return nil
	}
	defer func() {
		if err := fi.Close(); err != nil {
			//stat.PanicErr("filehande failed to close", err)
			panic(ErrSconfGeneric("filehande failed to close: " + err.Error()))
		}
	}()

	var str_json string
	buff := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buff)
		if err != nil && err != io.EOF {
			return ErrSconfGeneric(err.Error())
		}
		if n == 0 {
			break
		}

		str_json += string(buff[:n])
	}

	settings = nil
	if err := json.Unmarshal([]byte(str_json), &settings); err != nil {
		//stat.Status = "ERROR"
		//stat.Message = "failed to unmarshal config file: " +
		//	config_file_path + ": " + err.Error()

		return ErrSconfGeneric("failed to unmarshal config file: " +
			config_file_path + ": " + err.Error())
	}

	return nil
}

func (s *Sconf) Save() error {
	m_map, err := json.MarshalIndent(settings, "", "   ")
	if err != nil {
		panic(err)
	}

	fo, err := os.Create(config_file_path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = fo.WriteString(string(m_map))

	return err
}

// JSONSconfshareMap
//
// Handler used to reply to http requests. gives the settings map as JSON
//
// TODO: use RWMutex
//
type JSONSconfshareMap string
func (j JSONSconfshareMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	m_map, err := json.MarshalIndent(settings, "", "   ")
	if err != nil {
		//stat.PanicErr("handler failed to marshal settings map", err)
		panic(ErrSconfGeneric("handler failed to marshal settings map: " +
			err.Error()))
	}

	//fmt.Println(seestack.Short(), "r:")
	//fmt.Println(r)

	fmt.Fprint(w, string(m_map))
}

// GetFilePath
//
// used to get the config file path for debug
//
func (s Sconf) GetFilePath() string {
	return config_file_path
}

// CustomStatshare interface used by statshare
type SconfStatshare struct {
	Status string
	Message string
}

//func (css SconfStatshare) GiveStat() (string, string) {
//	return css.Status, css.Message
//}

//type SconfStatshare string
func (s SconfStatshare) GiveStat() (string, string) {
	return "TEST", "this is a test in sconf"
}
