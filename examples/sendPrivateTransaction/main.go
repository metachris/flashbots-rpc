package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

var privateKey, _ = crypto.GenerateKey() // creating a new private key for testing. you probably want to use an existing key.

func main() {
	rpc := flashbotsrpc.New("https://relay.flashbots.net")
	rpc.Debug = true

	sendPrivTxArgs := flashbotsrpc.FlashbotsSendPrivateTransactionRequest{
		Tx: "0x02f9019d011e843b9aca008477359400830247fa94def1c0ded9bec7f1a1670819833240f027b25eff88016345785d8a0000b90128d9627aa40000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000016345785d8a000000000000000000000000000000000000000000000000001394b63b2cbaea253a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee0000000000000000000000006b175474e89094c44da98b954eedeac495271d0f869584cd00000000000000000000000086003b044f70dac0abc80ac8957305b6370893ed0000000000000000000000000000000000000000000000d3443651ba615f6cd6c001a011a9f58ebe30aa679783b31793f897fdb603dd2ea086845723a22dae85ab2864a0090cf1fcce0f6e85da54f4eccf32a485d71a7d39bc0b43a53a9e64901c656230",
	}

	txHash, err := rpc.FlashbotsSendPrivateTransaction(privateKey, sendPrivTxArgs)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	// Print txHash
	fmt.Printf("txHash: %s\n", txHash)
}