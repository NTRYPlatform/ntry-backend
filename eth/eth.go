package eth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClient creates an ethereum client with channel for logs subscription
type EthClient struct {
	client  *ethclient.Client
	Events  chan types.Log
	gethIPC string
}

func NewEthClient(ipc string) (*EthClient, error) {

	client, err := getClient(ipc)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect Ethereum client: %v\n", err.Error()))
	}

	fmt.Printf("Ethereum client created: %v\n", client)

	return &EthClient{client: client, Events: make(chan types.Log, 100), gethIPC: ipc}, nil
}

//TODO: singleton
func getClient(ipc string) (client *ethclient.Client, err error) {
	client, err = ethclient.Dial(ipc)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect Ethereum client: %v\n", err.Error()))
	}
	return
}

func (e *EthClient) SubscribeToMapperContract(mapperContract string) error {

	c := e.client

	if c != nil {
		_, err := c.SubscribeFilterLogs(context.TODO(), ethereum.FilterQuery{Addresses: []common.Address{common.HexToAddress(mapperContract)}}, e.Events)
		if err != nil {
			return errors.New(fmt.Sprintf("Can't subscribe to contract logs: %v\n", err.Error()))
		}
		fmt.Printf("Subscribed to mapper contract @ %s", mapperContract)
		return nil
	}
	return errors.New("Eth Client is not initialized")
}

//TODO: either remove or update

//DeployMapperContract deploys mapper contract to the configured ethereum network
func (e *EthClient) DeployMapperContract(key, passphrase string) (string, error) {
	fmt.Println("Trying to deploy mapper contract...")
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	// conn, err := getClient(e.gethIPC)

	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to create authorized transactor: %v\n", err))
	}
	// Deploy a new awesome contract for the binding demo
	address, tx, _, err := DeployMapper(auth, e.client)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to deploy new token contract: %v\n", err))
	}
	log.Printf("Contract pending deployment: 0x%x\n", address)
	log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// wait for the transaction to be mined and check
	// time.Sleep(200 * time.Millisecond)
	// getTransactionReceipt(tx.Hash().String())
	return address.String(), nil

}

/**
 * We may not need RPC interface anymore
 */

// func (e *EthClient) GetTransactionReceipt(transaction string) {
// 	var callOptions = []byte(`{"jsonrpc": "2.0", "method":"eth_getTransactionReceipt", "params":["` + transaction + `"], "id":"1"}`)
// 	req, err := http.NewRequest("post", "http://localhost:8545/", bytes.NewBuffer(callOptions))
// 	if err != nil {
// 		log.Printf("Problem with eth_getTransactionReceipt request! %v", err.Error())
// 	}

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)

// 	log.Println(string(body))
// }

// func listAccounts() {
// 	var callOptions = []byte(`{"jsonrpc": "2.0", "method":"personal_listAccounts", "params":[], "id":1}`)
// 	req, err := http.NewRequest("post", "http://localhost:8545/", bytes.NewBuffer(callOptions))

// 	if err != nil {
// 		log.Printf("Problem with personal_listAccounts request! %v\n", err.Error())
// 	}

// 	client := &http.Client{}
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)

// 	log.Println(string(body))
// }

// // CreateAccount creates new ethereum account and unlocks it
// func CreateAccount(password string) (string, string) {
// 	acc, err := ks.NewAccount(password)
// 	if err != nil {
// 		log.Printf("Problem with creating new account! %v\n", err.Error())
// 	}

// 	keyFile := acc.URL.String()[11:]
// 	address := acc.Address.String()
// 	log.Println("New address generated:", address)

// 	// TODO: might want to delete the file???
// 	keyBytes, err := ioutil.ReadFile(keyFile)
// 	if err := ks.Unlock(acc, password); err != nil {
// 		log.Printf("Problem with unlocking new account! %v\n", err.Error())

// 	}

// 	return address, string(keyBytes)
// }

// func unlockAccount(address string, password string) {
// 	var callOptions = []byte(`{"jsonrpc": "2.0", "method":"personal_unlockAccount", "params":["` + address + `","` + password + `"], "id":"1"}`)
// 	req, err := http.NewRequest("post", "http://localhost:8545/", bytes.NewBuffer(callOptions))

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	client := &http.Client{}
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()
// 	resMap := make(map[string]interface{})
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	err = json.Unmarshal(body, &resMap)
// 	log.Println(string(body))
// 	unlocked := resMap["result"].(string)
// 	log.Println(unlocked)
// }

// func sendTransaction(to string, from string, amount float32, password string) {

// 	tx := fmt.Sprintf(`{"from": "%s", "to":"%s", "value": web3.toWei(%v, "ether")}`, from, to, amount)
// 	var callOptions = []byte(`{"jsonrpc": "2.0", "method":"personal_sendTransaction", "params":["` + tx + `","` + password + `"], "id":"1"}`)

// 	req, err := http.NewRequest("post", "http://localhost:8545/", bytes.NewBuffer(callOptions))

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	client := &http.Client{}
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)

// 	log.Println(string(body))
// }
