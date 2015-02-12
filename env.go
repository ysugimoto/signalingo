package signaling

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Env struct {
	Server ServerEnv `toml:"server"`
	Hook   HookEnv   `toml:"webhook"`
	Log    LogEnv    `toml:"log"`
}

type ServerEnv struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Endpoint string `toml:"endpoint"`
	Tls      bool   `toml:"tls"`
	Cert     string `toml:"certfile"`
	Key      string `toml:"keyfile"`
}

type HookEnv struct {
	Url string `toml:"url"`
}

type LogEnv struct {
	Type     string `toml:"type"`
	Filepath string `toml:"filepath"`
	Level    string `toml:"level"`
}

func InitEnv(path string) Env {
	var env Env
	defaultConf, _ := Asset("conf/env.conf")
	toml.Decode(string(defaultConf), &env)

	if _, err := os.Stat(path); err == nil {
		if _, err := toml.DecodeFile(path, &env); err != nil {
			panic(fmt.Sprintf("%v", err))
		}
	}

	//fmt.Printf("%v\n", env)
	return env
}
