package wallet

import (
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// CreateWallet create wallet from private key
func CreateWallet() (string, string) {
	getPrivateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Println(err)
	}

	getPublicKey := crypto.FromECDSA(getPrivateKey)
	thePublicKey := hexutil.Encode(getPublicKey)

	thePublicAddress := crypto.PubkeyToAddress(getPrivateKey.PublicKey).Hex()
	return thePublicAddress, thePublicKey
}

// ImportWallet import wallet from private key
func ImportWallet(privateKey string) (string, string) {
	getPrivateKey, err := crypto.HexToECDSA(privateKey)

	if err != nil {
		log.Println(err)
	}

	getPublicKey := crypto.FromECDSA(getPrivateKey)
	thePublicKey := hexutil.Encode(getPublicKey)

	thePublicAddress := crypto.PubkeyToAddress(getPrivateKey.PublicKey).Hex()
	return thePublicAddress, thePublicKey
}
