package sconf

import (
	"fmt"
	"encoding/json"
	"io"
	"os"
)

type ErrSconfGeneric string
func (e ErrSconfGeneric) Error() string {
	return "sconf: error: " + string(e)
}

type ErrUpdateSettings string
func (e ErrUpdateSettings) Error() string {
	return "sconf: could not update: " + string(e)
}

type sconf map[string]string

var settings sconf = make(sconf)
var config_file_path string
var New_called bool = false

func New(cfp string) (sconf, error) {

	if (New_called) {
		panic(ErrSconfGeneric("New() cannot be called a second time!"))
	}
	New_called = true

	if (settings == nil) {
		panic(ErrSconfGeneric("settings map cannot be nil!"))
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
		panic(ErrSconfGeneric("settings map cannot be nil!"))
	}
	return settings
}

func (s *sconf) Set_config_file_path(path string) {
	config_file_path = path
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

	settings = nil
	if err := json.Unmarshal([]byte(str_json), &settings); err != nil {
		return err
	}

	return nil
}

func (s *sconf) Save() error {
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
