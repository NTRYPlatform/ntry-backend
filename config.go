package main

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Host        string
	Port        int
	PubKeyFile  string
	PrivKeyFile string
	Db          Database
}

type Database struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
}

var (
	conf   *Config
	addr   string
	pubKey *rsa.PublicKey
	pvtKey *rsa.PrivateKey
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

func LoadConfig() {
	if conf == nil {
		initConfig()
		// _ = GetServerAddress()
		// _ = GetDatabaseSettings()
		// _ = GetPubKey()
		// _ = GetPvtKey()
	}
}

func GetServerAddress() string {
	LoadConfig()
	if addr == "" {
		addr = conf.Host + ":" + strconv.Itoa(conf.Port)
	}
	return addr
}

func GetDatabaseSettings() *Database {
	LoadConfig()
	return &conf.Db
}

func GetPvtKey() *rsa.PrivateKey {
	LoadConfig()
	if pvtKey == nil {
		pvtBytes, err := ioutil.ReadFile(conf.PrivKeyFile)
		if err != nil {
			log.Fatal("Can't read key!", err)
		}
		pvtKey, err = jwt.ParseRSAPrivateKeyFromPEM(pvtBytes)
		if err != nil {
			log.Fatal("Can't read key!", err)
		}
	}
	return pvtKey
}

func GetPubKey() *rsa.PublicKey {
	LoadConfig()
	if pubKey == nil {
		pubBytes, err := ioutil.ReadFile(conf.PubKeyFile)
		if err != nil {
			log.Fatal("Can't read key!", err)
		}
		pubKey, err = jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	}
	return pubKey
}
