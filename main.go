package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/SmartContractWithGolang/config"
	"github.com/SmartContractWithGolang/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// load config
	config.LoadConfig()

	// address of etherum env
	client, err := ethclient.Dial(config.Config.PROVIDER_RPC)
	if err != nil {
		panic(err)
	}

	contract, err := contract.NewMain(common.HexToAddress("0x563585FBc6256BB01B93cBAB406bCFCc4414F7a8"), client)
	if err != nil {
		panic(err)
	}

	value, err := contract.Balance(nil)
	if err != nil {
		panic(err)
	}

	println("balance before: ", value.Uint64())

	tx, err := contract.Deposite(getAccountAuth(client, config.Config.PRIVATE_KEY), big.NewInt(2))
	if err != nil {
		log.Fatal("error: ", err)
		panic(err)
	}

	fmt.Printf("tx sent: %s \n", tx.Hash().Hex())

	value2, err := contract.Balance(nil)
	if err != nil {
		panic(err)
	}

	println("balance After: ", value2.Uint64())

}

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(accountAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//fetch the last use nonce of account
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce = ", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	gas, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Println(err)
	}

	auth.GasPrice = gas

	return auth
}
