package sconf

import (
	"encoding/json"
	"fmt"
	"github.com/dankozitza/seestack"
	"github.com/dankozitza/statdist"
	"io"
	"net/http"
	"os"
)

type ErrSconfGeneric string

func (e ErrSconfGeneric) Error() string {
	stat.Status = "ERROR"
	stat.Message = string(e)
	stat.ShortStack = seestack.Short()
	statdist.Handle(stat)

	return stat.Message
}

type ErrUpdateSettings string

func (e ErrUpdateSettings) Error() string {
	return "Sconf could not update: " + string(e)
}

type Sconf map[string]interface{}

var settings Sconf = make(Sconf)
var config_file_path string
var New_called bool = false
var stat statdist.Stat

func New(cfp string) (Sconf, error) {

	config_file_path = cfp

	stat.Id = statdist.GetId()
	stat.ShortStack = seestack.Short()
	stat.Status = "PASS"
	stat.Message = "object initialized"

	if New_called {
		return nil, ErrSconfGeneric("New() cannot be called a second time!")
	}
	New_called = true

	if settings == nil {
		return nil, ErrSconfGeneric("settings map cannot be nil!")
	}

	stat.Message += ", using config file " + cfp

	err := settings.Update()
	if err != nil {
		panic(ErrSconfGeneric("failed to update settings: " + err.Error()))
	}

	statdist.Handle(stat)

	return settings, nil
}

func Inst() Sconf {
	if settings == nil {
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
		stat.ShortStack = seestack.Short()
		stat.Status = "WARN"
		stat.Message = "failed read config file: " + err.Error()
		statdist.Handle(stat)
		return nil
	}
	defer func() {
		if err := fi.Close(); err != nil {
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

	var newsettings map[string]interface{}
	if err := json.Unmarshal([]byte(str_json), &newsettings); err != nil {

		return ErrSconfGeneric("failed to unmarshal config file: " +
			config_file_path + ": " + err.Error())
	}

	for k, _ := range newsettings {
		settings[k] = newsettings[k]
	}

	return nil
}

func (s *Sconf) Save() error {
	m_map, err := json.MarshalIndent(settings, "", "   ")
	if err != nil {
		panic(ErrSconfGeneric(err.Error()))
	}

	fo, err := os.Create(config_file_path)
	if err != nil {
		msg := "failed to open config file " +
			config_file_path + " for writing: " + err.Error()

		return ErrSconfGeneric(msg)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(ErrSconfGeneric(err.Error()))
		}
	}()

	_, err = fo.WriteString(string(m_map))

	if err != nil {
		msg := "failed to write to config file " +
			config_file_path + ": " + err.Error()
		return ErrSconfGeneric(msg)
	}

	stat.Status = "PASS"
	stat.Message = "saved config to file " + config_file_path
	stat.ShortStack = seestack.Short()
	statdist.Handle(stat)

	return nil
}

// JSONSettingsMap
//
// Handler used to reply to http requests. gives the settings map as JSON
//
// TODO: use RWMutex
//
type JSONSettingsMap string

func (j JSONSettingsMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	m_map, err := json.MarshalIndent(settings, "", "   ")
	if err != nil {
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
