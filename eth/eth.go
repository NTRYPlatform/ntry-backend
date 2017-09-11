package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ntryapp/auth/config"
)

var (
	once           sync.Once
	client         ethclient.Client
	mapperContract string
	ks             *keystore.KeyStore
	endPoint       string
)

func init() {
	endPoint = config.GetEthIPC()
	ks = keystore.NewKeyStore(filepath.Join(config.GetEthDataDir(), "keystore"), keystore.StandardScryptN, keystore.StandardScryptP)
	mapperContract = config.GetMapperContract()
	if len(mapperContract) < 1 {
		keyFile := filepath.Join(config.GetEthDataDir(), "keystore", config.GetEthKey())
		key, err := ioutil.ReadFile(keyFile)
		if err != nil {
			log.Fatalf("Error occurred while trying to get ntry eth key: %v", err.Error())
		}
		mapperContract = deployMapperContract(string(key), config.GetEthPassphrase())
		config.SetMapperContractAddress(mapperContract)
	}
	subscribeToMapperContract()
}

func getClient() (client *ethclient.Client, err error) {
	once.Do(func() {
		client, err = ethclient.Dial(endPoint)
		if err != nil {
			log.Fatalf("Failed to connect Ethereum client: %v", err.Error())
		}
	})
	return
}

func subscribeToMapperContract() {
	//
	// Configure the node and an ethereum full node.
	// stackConf := &node.Config{}
	// ethConf := &eth.Config{}
	// stack, err := node.New(stackConf)
	// if err != nil {
	// 	log.Printf("protocol stack: %v", err.Error())
	// }

	// Start the node. This is a bit ugly at the moment.
	// newEth := func(ctx *node.ServiceContext) (node.Service, error) {
	// 	return eth.New(ctx, ethConf)
	// }
	ch := make(chan types.Log)
	c, _ := getClient()

	if c != nil {
		_, err := c.SubscribeFilterLogs(context.TODO(), ethereum.FilterQuery{Addresses: []common.Address{common.HexToAddress(config.GetMapperContract())}}, ch)
		if err != nil {
			log.Fatalf("Can't subscribe to contract logs!")
		}
		// goroutine
		go func() {
			for true {
				// event := types.Log{}
				ethLog := <-ch
				// bytes, err := json.Marshal(ethLog)
				// json.Unmarshal(bytes, &event)
				log.Printf("%+v", ethLog)
				log.Printf("Address: %s", ethLog.Address.String())
				log.Printf("Data: %x", string(ethLog.Data))
				// log.Printf("Data: %s", ethLog.UnmarshalJSON)
			}
		}()
	}
}

//deployMapperContract deploys mapper contract to the configured ethereum network
func deployMapperContract(key, passphrase string) string {
	log.Println("Trying to deploy mapper contract...")
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := getClient()
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy a new awesome contract for the binding demo
	address, tx, _, err := DeployMapper(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	log.Printf("Contract pending deployment: 0x%x\n", address)
	log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// wait for the transaction to be mined and check
	// time.Sleep(200 * time.Millisecond)
	// getTransactionReceipt(tx.Hash().String())
	return address.String()

}

// MapSecondaryAddress maps a secondary ethereum address to a primary address
func MapSecondaryAddress(key, passphrase, primary, secondary string) (err error) {

	client, _ := getClient()
	log.Printf("Connection information: %v", client)

	contractAddress := common.HexToAddress(mapperContract)

	// create a transaction signer from an encrypted json key stream
	// and the associated passphrase.
	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	log.Printf("Auth: %+v", auth)

	// binds a generic wrapper to an already deployed contract
	contract, err := bindMapper(contractAddress, client, client)
	if err != nil {
		log.Fatalf("Failed to instansitate a mapper contract: %v", err)
		return
	}

	writerSession := &MapperTransactorSession{
		Contract: &MapperTransactor{
			contract: contract,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: big.NewInt(3141592),
		},
	}

	transaction, err := writerSession.MapAddress(common.HexToAddress(secondary))

	if err != nil {
		log.Printf("Failed to map address: %v", err)
		return
	}

	log.Printf("Transaction :  %v ", transaction.Hash().Hex())
	return

}

func getTransactionReceipt(transaction string) {
	var callOptions = []byte(`{"jsonrpc": "2.0", "method":"eth_getTransactionReceipt", "params":["` + transaction + `"], "id":"1"}`)
	req, err := http.NewRequest("post", "http://localhost:8545/", bytes.NewBuffer(callOptions))
	if err != nil {
		log.Printf("Problem with eth_getTransactionReceipt request! %v", err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}

func listAccounts() {
	var callOptions = []byte(`{"jsonrpc": "2.0", "method":"personal_listAccounts", "params":[], "id":1}`)
	req, err := http.NewRequest("post", "http://localhost:8545/", bytes.NewBuffer(callOptions))

	if err != nil {
		log.Printf("Problem with personal_listAccounts request! %v\n", err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))

}

// CreateAccount creates new ethereum account and unlocks it
func CreateAccount(password string) (string, string) {
	acc, err := ks.NewAccount(password)
	if err != nil {
		log.Printf("Problem with creating new account! %v\n", err.Error())
	}

	keyFile := acc.URL.String()[11:]
	address := acc.Address.String()
	log.Println("New address generated:", address)

	// TODO: might want to delete the file???
	keyBytes, err := ioutil.ReadFile(keyFile)
	if err := ks.Unlock(acc, password); err != nil {
		log.Printf("Problem with unlocking new account! %v\n", err.Error())

	}

	return address, string(keyBytes)
}

func unlockAccount(address string, password string) {
	var callOptions = []byte(`{"jsonrpc": "2.0", "method":"personal_unlockAccount", "params":["` + address + `","` + password + `"], "id":"1"}`)
	req, err := http.NewRequest("post", "http://localhost:8545/", bytes.NewBuffer(callOptions))

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	resMap := make(map[string]interface{})
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &resMap)
	log.Println(string(body))
	unlocked := resMap["result"].(string)
	log.Println(unlocked)
}

func sendTransaction(to string, from string, amount float32, password string) {

	tx := fmt.Sprintf(`{"from": "%s", "to":"%s", "value": web3.toWei(%v, "ether")}`, from, to, amount)
	var callOptions = []byte(`{"jsonrpc": "2.0", "method":"personal_sendTransaction", "params":["` + tx + `","` + password + `"], "id":"1"}`)

	req, err := http.NewRequest("post", "http://localhost:8545/", bytes.NewBuffer(callOptions))

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(string(body))
}
