package notary

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/NTRYPlatform/ntry-backend/config"
	"github.com/NTRYPlatform/ntry-backend/eth"
	uuid "github.com/satori/go.uuid"
	log "go.uber.org/zap"
	"upper.io/db.v3/mysql"
)

// TODO:
// Handle read write locking
type Notary struct {
	id        uuid.UUID
	db        *dbServer
	conf      *config.Config
	email     *emailConf
	logger    *log.Logger
	ethClient *eth.EthClient
	ctx       context.Context
	cancel    context.CancelFunc
}

func New(args map[string]interface{}) (*Notary, error) {
	var (
		logPath       = args["--logpath"].(string)
		confPath      = args["--confpath"].(string)
		debug    bool = args["--debug"].(bool)
	)

	ntry := &Notary{}

	err := ntry.makeID()
	if err != nil {
		return nil, err
	}

	if ntry.logger, err = initLogger(&logPath, ntry.id.String(), &debug); err != nil {
		return nil, err
	}

	// Loading the congiguration
	ntry.logger.Info("[notary  ] Loading configurations from: " + confPath)
	if ntry.conf, err = config.Init(confPath, ntry.logger); err != nil {
		return nil, err
	}

	return ntry, nil
}

func (n *Notary) Init() error {

	ctx := context.Background()
	n.ctx, n.cancel = context.WithCancel(ctx)

	// TODO:
	// Conf and DB structs have publid variables
	// as well as Getters, access must be filtered
	dbConf := n.conf.GetDatabaseSettings()
	script, err := dbConf.GetDBSrcipts()
	if err != nil {
		return err
	}

	// initilizing db session
	if n.db, err =
		dbInit(script, mysql.ConnectionURL{
			Host:     dbConf.Host,
			Database: dbConf.Name,
			User:     dbConf.User,
			Password: dbConf.Password,
		}, n.logger); err != nil {
		return err
	}

	n.email = newEmail()
	n.email.from, n.email.pass, n.email.server = n.conf.GetEmailInfo()
	if !n.email.ok() {
		return fmt.Errorf("[notary  ] incomplete email configuration", nil)
	}

	// Initialize ETH client
	n.ethClient, err = eth.NewEthClient(n.conf.GetEthIPC())
	if err != nil {
		return fmt.Errorf("[notary  ] Unable to initialize ETH client", err)
	}

	if err = n.ethClient.SubscribeToMapperContract(n.conf.GetMapperContract()); err != nil {
		return fmt.Errorf("[notary  ] Unable to bind event listener", err)
	}

	return nil

}

func (n *Notary) makeID() error {
	dirName := "/var/lib/notary"
	fileName := "notary.id"
	filePath := filepath.Join(dirName, fileName)
	err := os.MkdirAll(dirName, 0644)
	if err != nil {
		return err
	}
	if f, err := os.OpenFile(filePath, os.O_RDWR, 0644); os.IsNotExist(err) {
		n.id = uuid.NewV4()
		errWrite := ioutil.WriteFile(filePath, n.id.Bytes(), 0644)
		if err != nil {
			return errWrite
		}
	} else {
		if err != nil {
			return err
		}
		buff, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		n.id, err = uuid.FromBytes(buff)
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *Notary) Start() error {

	router := n.muxServer()
	addr := n.conf.GetServerAddress()
	//TODO:generate cert and serve on TLS
	n.logger.Info("[notary  ] Server waiting for request...")
	return http.ListenAndServe(addr, router)
}

func (n *Notary) EthWatcher() {
	for {
		select {
		case ethLog := <-n.ethClient.Events:
			n.logger.Info(fmt.Sprintf("[notary  ] %+v", ethLog))
			data := hex.EncodeToString(ethLog.Data)
			address := data[24:64]
			uid := data[64:96]
			n.logger.Info(fmt.Sprintf("Address: %s, UID: %s, Tx Hash: %s", address, uid, ethLog.TxHash.String()))
			u := app.VerifyUser(uid, address, ethLog.TxHash.String())
			if err := d.UpdateUser(u, uid); err != nil {
				log.Printf("Couldn't update user email verification! %v", err.Error())
			}
			app.WriteToRegisterChannel("{\"registered\":true}")
		}
	}
}

func (n *Notary) Shutdown() error {
	n.db.CloseSession()
	return nil
}
