package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"wcjs/access_token/common"
	"wcjs/lib/config"
	"wcjs/lib/log"
)

type iniConfig struct {
	mu sync.RWMutex
	baseconfig *BaseConfig
	sessionconfig *common.SessionConfig
	logconfig *log.Loggerconfig
	collectconfig *common.CollectConfig
}

func (ic *iniConfig) reloadConfig() error{
	ini_parser := config.IniParser{}
	workPath, _ := os.Getwd()
	conf_file_name := filepath.Join(workPath, "conf", "app.conf")
	if err := ini_parser.Load(conf_file_name); err != nil {
		fmt.Printf("try load config file[%s] error[%s]\n", conf_file_name, err.Error())
		return err
	}

	ic.mu.Lock()
	ic.baseconfig.Host = ini_parser.GetString("base", "host")
	ic.baseconfig.Port = ini_parser.GetInt64("base", "port")

	ic.sessionconfig.Host = ini_parser.GetString("session", "host")
	ic.sessionconfig.Select = int(ini_parser.GetInt64("session", "select"))
	ic.sessionconfig.Authstring = ini_parser.GetString("session", "authstring")

	ic.logconfig.DebugEnabled = ini_parser.GetBool("log", "debug")
	ic.logconfig.Dir = ini_parser.GetString("log", "path")
	ic.logconfig.Loglevel = ini_parser.GetInt64("log", "level")
	ic.logconfig.DistingType = ini_parser.GetInt64("log", "distingtype")

	ic.collectconfig.FilePath = ini_parser.GetString("collect", "filepath")
	ic.mu.Unlock()
	return nil
}


var iniconfig *iniConfig
func InitIniConfig() error {
	iniconfig = &iniConfig{
		mu:sync.RWMutex{},
		baseconfig:&BaseConfig{},
		sessionconfig:&common.SessionConfig{},
		logconfig:&log.Loggerconfig{},
		collectconfig:&common.CollectConfig{
			FilePath:"",
			//MysqlConfig:JsonMysqlGet(),
			ChemmMysqlConfig:JsonChemmMysqlGet(),
		},
	}
	err := iniconfig.reloadConfig()
	if err != nil {
		return err
	}
	return nil
}

func IniBaseGet() *BaseConfig{
	if iniconfig == nil{
		err := InitIniConfig()
		if err != nil {
			fmt.Print(err)
		}
	}
	iniconfig.mu.RLock()
	defer iniconfig.mu.RUnlock()
	return iniconfig.baseconfig
}

func IniSessionGet() *common.SessionConfig{
	if iniconfig == nil {
		err := InitIniConfig()
		if err != nil {
			fmt.Print(err)
		}
	}
	iniconfig.mu.RLock()
	defer iniconfig.mu.RUnlock()
	return iniconfig.sessionconfig
}

func IniLogGet() *log.Loggerconfig{
	if iniconfig == nil {
		err := InitIniConfig()
		if err != nil {
			fmt.Print(err)
		}
	}
	iniconfig.mu.RLock()
	defer iniconfig.mu.RUnlock()
	return iniconfig.logconfig
}

func IniCollectGet() *common.CollectConfig{
	if iniconfig == nil {
		err := InitIniConfig()
		if err != nil {
			fmt.Print(err)
		}
	}
	iniconfig.mu.RLock()
	defer iniconfig.mu.RUnlock()
	return iniconfig.collectconfig
}