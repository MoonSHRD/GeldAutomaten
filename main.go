package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"

//	"os"
    "context"
    "log"
    "fmt"

   // "github.com/ethereum/go-ethereum"
	"github.com/joho/godotenv"
	dep	"github.com/MoonSHRD/GeldAutomaten/artifacts/Deposit"

)

var myenv map[string]string

const envLoc = ".env"

func loadEnv() {
    var err error
    if myenv, err = godotenv.Read(envLoc); err != nil {
        log.Printf("could not load env from %s: %v", envLoc, err)
    }
}

func main(){
    loadEnv()

    ctx := context.Background()

//	gw := myenv["GATEWAY"]
//	fmt.Println(gw)

  //  client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
	client, err := ethclient.Dial(myenv["GATEWAY"])			// load from local .env file
    if err != nil {
        log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
    }
    defer client.Close()

    accountAddress := common.HexToAddress("0x892eE0398C9d8C86BCA3ffa49c33b68A7b2F38d3")
    balance, _ := client.BalanceAt(ctx, accountAddress, nil)
	fmt.Printf("Balance: %d\n",balance)
	
	deposit, err := dep.NewDeposit(common.HexToAddress("0xdD5A73bfE8f907D6592197cF4Ba5945b1bFc608f"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Deposit contract: %v", err)
	}

	owner,err := deposit.Owner(nil)
	if err != nil {
		log.Fatalf("Failed to owner of contract: %v", err)
	}

	owner_human := owner.Hex()

	fmt.Println(owner_human)

}