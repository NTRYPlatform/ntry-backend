package main

import (
	"io/ioutil"
	"log"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Host string
	Port int
	Db   Database
}

type Database struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
}

var (
	conf *Config
	addr string
)

/* initConfig loads config file from the configuration folder
*  panics
 */
func initConfig() {
	//TODO: maybe the path could be a part of the command args?
	configFile, err := ioutil.ReadFile("config/ntryapp.yml")

	if err != nil {
		log.Println("Can't read config file!", err)
		panic("Panicking!")
	}

	err = yaml.Unmarshal(configFile, &conf)

	if err != nil {
		log.Println("Can't unmarshall the config properties!", err)
		panic("Panicking!")
	} else {
		log.Println("Config successfully loaded: %+v", *conf)
	}

}

func loadConfig() {
	if conf == nil {
		initConfig()
	}
}

func GetServerAddress() string {
	loadConfig()
	if addr == "" {
		addr = conf.Host + ":" + strconv.Itoa(conf.Port)
	}
	return addr
}

func GetDatabaseSettings() *Database {
	loadConfig()
	return &conf.Db
}
