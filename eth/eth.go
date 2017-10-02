package eth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClient creates an ethereum client with channel for logs subscription
type EthClient struct {
	client         *ethclient.Client
	Events         chan types.Log
	gethIPC        string
	carContract    common.Address
	notaryContract common.Address
	key            string
	passphrase     string
}

func NewEthClient(ipc, carContract, notaryContract, key, passphrase string) (*EthClient, error) {
	client, err := getClient(ipc)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect Ethereum client: %v\n", err.Error()))
	}

	fmt.Printf("Ethereum client created: %v\n", client)

	return &EthClient{client: client, Events: make(chan types.Log, 100), gethIPC: ipc, carContract: common.HexToAddress(carContract), notaryContract: common.HexToAddress(notaryContract), key: key, passphrase: passphrase}, nil
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

func (e *EthClient) CarDeal(hash, buyerAddress, sellerAddress string, cid int64) error {
	fmt.Println("Trying to make a car deal...")
	// auth
	auth, err := bind.NewTransactor(strings.NewReader(e.key), e.passphrase)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create authorized transactor: %v\n", err))
	}
	contract, err := NewCarContractFunctions(e.carContract, e.client)
	opts := bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(3141592),
	}

	c := big.NewInt(cid)
	bytes := []byte(hash)
	var arr [32]byte
	copy(arr[:], bytes)
	tx, err := contract.CarDeal(&opts, c, common.HexToAddress(sellerAddress), common.HexToAddress(buyerAddress), arr)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create car deal: %v\n", err))
	}
	log.Printf("Contract pending deployment: %v\n", tx)
	log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// wait for the transaction to be mined and check
	// time.Sleep(200 * time.Millisecond)
	// getTransactionReceipt(tx.Hash().String())
	return nil

}

func (e *EthClient) NotaryBalance(address string) (*big.Int, error) {
	fmt.Println("Trying to get balance...")
	a := common.HexToAddress(address)
	contract, err := NewTokenContractFunctions(e.notaryContract, e.client)
	opts := bind.CallOpts{
		From: a,
	}

	b, err := contract.BalanceOf(&opts, a)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to get balance: %v\n", err))
	}
	fmt.Printf("Balance... %v\n", b)
	return b, nil

}
