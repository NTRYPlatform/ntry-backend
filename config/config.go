package config

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	yaml "gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	Host           string `yaml:"host" required:"true"`
	Port           int    `yaml:"port" required:"true"`
	Email          string `yaml:"email" required:"true"`
	Password       string `yaml:"password" required:"true"`
	EmailServer    string `yaml:"emailServer" required:"true"`
	PubKeyFile     string `yaml:"pubKeyFile" required:"true"`
	PrivKeyFile    string `yaml:"privKeyFile" required:"true"`
	MapperContract string `yaml:"mapperContract" required:"true"`
	EthDataDir     string `yaml:"ethDataDir" required:"true"`
	Db             Database
}

// Database struct
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
*
 */
func initConfig() {
	//TODO: maybe the path could be a part of the command args?
	file, err := filepath.Abs(".notaryconf/ntryapp.yml")
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("Can't read config file! %v", err.Error())
		panic("Panicking!")
	}

	err = yaml.Unmarshal(configFile, &conf)

	if err != nil {
		log.Printf("Can't unmarshall the config properties!\n%v", err)
		panic("Panicking!")
	} else {
		log.Printf("Config successfully loaded: %+v", *conf)
	}

}

// LoadConfig initializes config
// TODO: add once sync
// TODO: use the init function for initializing config
func LoadConfig() {
	if conf == nil {
		initConfig()
	}
}

// GetServerAddress returns server address
func GetServerAddress() string {
	LoadConfig()
	if addr == "" {
		addr = conf.Host + ":" + strconv.Itoa(conf.Port)
	}
	return addr
}

// GetDatabaseSettings returns Database struct
func GetDatabaseSettings() *Database {
	LoadConfig()
	return &conf.Db
}

// GetPvtKey returns private key
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

// GetPubKey returns public key
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

// GetEmailInfo returns email, password and emailserver strings
func GetEmailInfo() (email, password, emailServer string) {
	LoadConfig()
	return conf.Email, conf.Password, conf.EmailServer
}

// GetMapperContract returns mapper contracts address
func GetMapperContract() string {
	return conf.MapperContract
}

func GetEthDataDir() string {
	return conf.EthDataDir
}

// GetEthIPC returns eth IPC endpoint
func GetEthIPC() string {
	return conf.EthDataDir + "/geth.ipc"
}
