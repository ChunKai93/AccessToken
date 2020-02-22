package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"wcjs/lib/client/mysql"
)

type jsonConfig struct {
	mu sync.RWMutex
	mysqlconfig []*mysql.MultiMysqlOptions
}

func (jc *jsonConfig) reloadConfig(filename string) error{
	workPath, _ := os.Getwd()
	conf_file_name := filepath.Join(workPath, "conf", filename)
	//fmt.Print(conf_file_name)
	if contents,err := ioutil.ReadFile(conf_file_name);err == nil {
		jc.mu.Lock()
		e := json.Unmarshal(contents,&jc.mysqlconfig)
		jc.mu.Unlock()
		if e != nil {
			return e
		} else {
			return nil
		}
	} else {
		return errors.New("无法打开mysql配置文件")
	}
	return nil
}


var jsonconfig *jsonConfig
var chemm_jsonconfig *jsonConfig
func InitJsonConfig() error {
	jsonconfig = &jsonConfig{
		mu:sync.RWMutex{},
		mysqlconfig:make([]*mysql.MultiMysqlOptions,0),
	}
	err := jsonconfig.reloadConfig("mysql.json")
	if err != nil {
		return err
	}
	return nil
}

func JsonMysqlGet() []*mysql.MultiMysqlOptions{
	if jsonconfig == nil {
		err := InitJsonConfig()
		if err != nil {
			fmt.Print(err)
		}
	}
	jsonconfig.mu.RLock()
	defer jsonconfig.mu.RUnlock()
	return jsonconfig.mysqlconfig
}

func InitChemmJsonConfig() error {
	chemm_jsonconfig = &jsonConfig{
		mu:sync.RWMutex{},
		mysqlconfig:make([]*mysql.MultiMysqlOptions,0),
	}
	err := chemm_jsonconfig.reloadConfig("chemm_mysql.json")
	if err != nil {
		return err
	}
	return nil
}

func JsonChemmMysqlGet() []*mysql.MultiMysqlOptions{
	if chemm_jsonconfig == nil {
		err := InitChemmJsonConfig()
		if err != nil {
			fmt.Print(err)
		}
	}
	chemm_jsonconfig.mu.RLock()
	defer chemm_jsonconfig.mu.RUnlock()
	return chemm_jsonconfig.mysqlconfig
}