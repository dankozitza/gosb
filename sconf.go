package sconf

import (
	"encoding/json"
	"errors"
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
	statdist.Handle(stat, false)

	return stat.Message
}

type Sconf map[string]interface{}

var settings Sconf = make(Sconf)
var config_file_path string // not being used properly
var Init_called bool = false
var stat statdist.Stat

// Init
//
// Initializes the global settings map. This must be called before calling Inst.
//
func Init(cfp string, preset Sconf) Sconf {

	if Init_called {
		stat.Status = "ERROR"
		stat.ShortStack = seestack.Short()
		stat.Message = "Init cannot be called a second time!"
		statdist.Handle(stat, false)
		return nil
	}
	Init_called = true

	config_file_path = cfp
	settings = New(cfp, preset)

	return settings
}

// Inst
//
// Returns a reference to the global settings map.
//
func Inst() Sconf {
	return settings
}

// New
//
// Creates a new Sconf object.
//
func New(cfp string, preset Sconf) Sconf {

	var local_settings Sconf = make(Sconf)

	stat.Id = statdist.GetId()
	stat.ShortStack = seestack.Short()
	stat.Status = "INIT"
	stat.Message = "object initialized, using config file " + cfp

	if preset != nil {
		for k, _ := range preset {
			local_settings[k] = preset[k]
		}
	}

	local_settings.Update(cfp)

	if stat.Status == "INIT" {
		statdist.Handle(stat, true)
	}

	return local_settings
}

// broken
func (s Sconf) Set_config_file_path(path string) {
	config_file_path = path
}

// TODO: FileWasModified
//
// Checks to see if the file at path has been modified since last update.
// Should be put in dkutils and called from main.
//

// Update
//
// Unmarshalls the JSON object in the config file path and copies those values
// to the Sconf object.
//
func (s Sconf) Update(cfp string) error {
	// update settings map from the config file path
	fi, err := os.Open(cfp)
	if err != nil {
		stat.Status = "WARN"
		stat.Message = "failed to update: " + err.Error()
		stat.ShortStack = seestack.Short()
		statdist.Handle(stat, false)
		return errors.New(stat.Message)
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
			return err
		}
		if n == 0 {
			break
		}

		str_json += string(buff[:n])
	}

	var newsettings map[string]interface{}
	if err := json.Unmarshal([]byte(str_json), &newsettings); err != nil {
		stat.Status = "WARN"
		stat.Message = "unmarshalling " + cfp + ": " + err.Error()
		stat.ShortStack = seestack.Short()
		statdist.Handle(stat, false)
		return errors.New(stat.Message)
	}

	for k, _ := range newsettings {
		s[k] = newsettings[k]
	}

	stat.Status = "PASS"
	stat.Message = "updated sconf object from file: " + cfp
	stat.ShortStack = seestack.Short()
	statdist.Handle(stat, true)

	return nil
}

// Save
//
// Saves the Sconf object as JSON.
//
func (s Sconf) Save(cfp string) error {
	m_map, err := json.MarshalIndent(s, "", "   ")
	if err != nil {
		panic(ErrSconfGeneric(err.Error()))
	}

	fo, err := os.Create(cfp)
	if err != nil {
		msg := "failed to open config file " + cfp +
			" for writing: " + err.Error()

		return ErrSconfGeneric(msg)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(ErrSconfGeneric(err.Error()))
		}
	}()

	_, err = fo.WriteString(string(m_map))
	if err != nil {
		msg := "failed to write to config file " + cfp + ": " + err.Error()
		return ErrSconfGeneric(msg)
	}

	stat.Status = "PASS"
	stat.Message = "saved config to file " + cfp
	stat.ShortStack = seestack.Short()
	statdist.Handle(stat, true)

	return nil
}

// HTTPHandler
//
// Handler used to reply to http requests. gives the Sconf object as JSON
//
// TODO: use RWMutex
//
type HTTPHandler Sconf

func (j HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	m_map, err := json.MarshalIndent(j, "", "   ")
	if err != nil {
		panic(ErrSconfGeneric("handler failed to marshal settings map: " +
			err.Error()))
	}

	fmt.Fprint(w, string(m_map))
}

// GetFilePath
//
// used to get the config file path for debug
//
func (s Sconf) GetFilePath() string {
	return config_file_path
}
