package env

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ysugimoto/go-cliargs"
	"github.com/ysugimoto/signalingo/static"
	"os"
)

type Env struct {
	Server  ServerEnv  `toml:"server"`
	Hook    HookEnv    `toml:"webhook"`
	Log     LogEnv     `toml:"log"`
	Storage StorageEnv `toml:"storage"`
	Redis   RedisEnv   `toml:"redis"`
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

type StorageEnv struct {
	Type string `toml:"type"`
}

type RedisEnv struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

func mergeEnv(env *Env) Env {
	args := cliarg.NewArguments()
	args.Alias("h", "host", env.Server.Host)
	args.Alias("p", "port", env.Server.Port)
	args.Alias("e", "endpoint", env.Server.Endpoint)
	args.Parse()

	host, _ := args.GetOptionAsString("host")
	port, _ := args.GetOptionAsInt("port")
	endpoint, _ := args.GetOptionAsString("endpoint")

	env.Server.Host = host
	env.Server.Port = port
	env.Server.Endpoint = endpoint

	return *env
}

func InitEnv(path string) Env {
	var env Env
	defaultConf, _ := static.Asset("etc/env.conf")
	toml.Decode(string(defaultConf), &env)

	if _, err := os.Stat(path); err == nil {
		if _, err := toml.DecodeFile(path, &env); err != nil {
			panic(fmt.Sprintf("%v", err))
		}
	}

	return mergeEnv(&env)
}
