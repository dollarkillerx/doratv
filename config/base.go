/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 21:31 2019-09-12
 */

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type basisConf struct {
	App struct {
		Host       string `yaml:"host"`
		Debug      bool   `yaml:"debug"`
		MaxRequest int    `yaml:"max_request"`
		LogLevel   string `yaml:"log_level"`
	}
	Mysql struct {
		Dsn   string `yaml:"dsn"`
		Cache bool   `yaml:"cache"`
	}
	Pgsql struct{
		Dsn   string `yaml:"dsn"`
		MaxIdle int `yaml:"max_idle"`
		MaxOpen int `yaml:"max_open"`
		TimeOut time.Duration `yaml:"time_out"`
	}
	Redis struct {
		Maxidle     int           `yaml:"maxidle"`
		MaxActive   int           `yaml:"max_active"`
		IdleTimeout time.Duration `yaml:"idle_timeout"`
		Port        string        `yaml:"port"`
	}
}

var (
	Basis *basisConf
)

func init() {
	Basis = &basisConf{}

	bytes, e := ioutil.ReadFile("./config.yml")
	if e != nil {
		panic(e.Error())
	}

	e = yaml.Unmarshal(bytes, Basis)
	if e != nil {
		panic(e.Error())
	}

	if Basis.App.Debug {
		log.Println(Basis)
	}
}
