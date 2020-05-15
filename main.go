package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
//	"os"
    "context"
    "log"
    "fmt"

   // "github.com/ethereum/go-ethereum"
    "github.com/joho/godotenv"
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
//	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
    if err != nil {
        log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
    }
    defer client.Close()

    accountAddress := common.HexToAddress("0x892eE0398C9d8C86BCA3ffa49c33b68A7b2F38d3")
    balance, _ := client.BalanceAt(ctx, accountAddress, nil)
    fmt.Printf("Balance: %d\n",balance)
}