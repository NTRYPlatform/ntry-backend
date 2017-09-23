package config

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/Unknwon/com"
	jwt "github.com/dgrijalva/jwt-go"
	log "go.uber.org/zap"
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
	EthPvtKeyFile  string `yaml:"ethPvtKeyFile" required:"true"`
	EthPassphrase  string `yaml:"ethPassphrase" required:"true"`
	Db             Database
}

// Database struct
type Database struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
	Script   string
}

// Init loads config file from the configuration folder
func Init(path string, logger *log.Logger) (*Config, error) {
	conf := &Config{}

	configFile, err := readfile(path)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(configFile, conf); err != nil {
		return nil, err
	}

	logger.Info(fmt.Sprintf("[notary ] configuration loaded successfully: %+v", *conf))

	return conf, nil
}

func (c *Config) GetServerAddress() string {
	return c.Host + ":" + strconv.Itoa(c.Port)

}

// GetDatabaseSettings returns Database struct
func (c *Config) GetDatabaseSettings() *Database {
	return &c.Db
}

// GetDatabaseSettings returns Database struct
func (d *Database) GetDBScript() (string, error) {
	script, err := readfile(d.Script)
	if err != nil {
		return "", err
	}
	return string(script), nil
}

// GetPvtKey returns private key
func (c *Config) GetPvtKey() (*rsa.PrivateKey, error) {

	pvtBytes, err := readfile(c.PrivKeyFile)
	if err != nil {
		return nil, err
	}

	pvtKey, err := jwt.ParseRSAPrivateKeyFromPEM(pvtBytes)
	if err != nil {
		return nil, err
	}

	return pvtKey, nil
}

// GetPubKey returns public key
func (c *Config) GetPubKey() (*rsa.PublicKey, error) {

	pubBytes, err := readfile(c.PubKeyFile)
	if err != nil {
		return nil, err
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

// GetEmailInfo returns email, password and emailserver strings
func (c *Config) GetEmailInfo() (email, password, emailServer string) {
	return c.Email, c.Password, c.EmailServer
}

// GetMapperContract returns mapper contracts address
func (c *Config) GetMapperContract() string {
	return c.MapperContract
}

func (c *Config) GetEthDataDir() string {
	return c.EthDataDir
}

// GetEthIPC returns eth IPC endpoint
func (c *Config) GetEthIPC() string {
	return c.EthDataDir + "/geth.ipc"
}

//updateConfig is a util function to automatically update the configuration file
func (c *Config) updateConfig(path string) error {

	file, _ := filepath.Abs(path)
	out, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, out, 0644)
}

func (c *Config) GetEthKey() string {
	return c.EthPvtKeyFile
}

func (c *Config) GetEthPassphrase() string {
	return c.EthPassphrase
}

func (c *Config) SetMapperContractAddress(address string, path string) {
	c.MapperContract = address
	c.updateConfig(path)
}

func readfile(path string) (data []byte, err error) {
	if !com.IsFile(path) {
		return nil, fmt.Errorf("No file exist at path: %v", path)
	}

	data, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return
}
