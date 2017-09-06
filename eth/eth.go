package eth

import (
	"log"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ntryapp/auth/config"
)

var (
	once           sync.Once
	client         ethclient.Client
	mapperContract string
)

func getClient() (client *ethclient.Client, err error) {
	once.Do(func() {
		mapperContract = config.GetMapperContract()
		endPoint := config.GetEthIPC()
		client, err = ethclient.Dial(endPoint)
		if err != nil {
			log.Fatalf("Failed to connect Ethereum client: %v", err.Error())
		}
	})
	return
}

func MapSecondaryAddress(key, passphrase, primary, secondary string) (err error) {

	client, _ := getClient()
	log.Printf("Connection information: %v", client)

	contractAddress := common.HexToAddress(mapperContract)

	// Account information of Writer/Reader
	// key := `{"address":"0xce114b931dbd8cfcf517c6a72be5f4f8b58944bb","crypto":{"cipher":"aes-128-ctr","ciphertext":"252d251ad0c3800da78186c62f427fd6cfb6711b519a1b723c6897438ef986e2","cipherparams":{"iv":"5574b60debba87d5b54112feb600d3ed"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"89e61af14b23d8021ae477f6e5e201d021ca166c47c9121b06fce3d722dfc2af"},"mac":"9e8dbac0aa6f0ebb4cf50c3b049dd4c212c0c9a242738ad69506c87a0749bd8b"},"id":"83386624-e582-4434-8f96-76e0ceef72d9","version":3}`
	// Passphrase to authenticate the user account on ethereum
	// passphrase := "testnet"

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
