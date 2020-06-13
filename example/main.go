package example

import (
	"math/big"

	ga "github.com/MoonSHRD/GeldAutomaten"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"context"
	"fmt"
	"log"

	dep "github.com/MoonSHRD/GeldAutomaten/artifacts/Deposit"
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

func main() {
	loadEnv()

	ctx := context.Background()
	pk := myenv["PK"] // load private key from env

	// Connecting to network
	//  client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
	client, err := ethclient.Dial(myenv["GATEWAY"]) // load from local .env file
	if err != nil {
		log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
	}
	defer client.Close()

	// setting up private key in proper format
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}

	// Creating an auth transactor
	auth := bind.NewKeyedTransactor(privateKey)

	// check calls
	// check balance
	accountAddress := common.HexToAddress("0x3214Db97Bf87a057C039a39594E91CD31f5d2a2C")
	balance, _ := client.BalanceAt(ctx, accountAddress, nil) //our balance
	fmt.Printf("Balance: %d\n", balance)

	// Setting up Deposit Contract
	deposit, err := dep.NewDeposit(common.HexToAddress("0x56D17255fFFab24677ba0e6760d9F6ac86313493"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Deposit contract: %v", err)
	}

	// Wrap the Deposit contract instance into a session
	session := &dep.DepositSession{
		Contract: deposit,
		CallOpts: bind.CallOpts{
			Pending: true,
			From:    auth.From,
			Context: context.Background(),
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: 0,   // 0 automatically estimates gas limit
			GasPrice: nil, // nil automatically suggests gas price
			Context:  context.Background(),
		},
	}

	/*
	   Events
	*/

	// Check retriving events of CashOut request

	var ch = make(chan *dep.DepositCashOutRequestEventAnonymouse)
	subscription, err := ga.SubcribeCashOutRequestFlow(session, ch)

	// check call for getting owner inside session
	/*
	   ownr,err := session.Owner()
	   if err != nil {
	   log.Fatalf("Failed to owner of contract: %v", err)
	   }
	   oh:= ownr.Hex()
	   fmt.Println(oh)
	*/

	// check Cash out request
	purce := "address karty"
	paymentMethod := "karta"

	txOutRequest, err := deposit.CashOutRequest(
		&bind.TransactOpts{
			From:     auth.From,
			Nonce:    nil, // nil uses nonce of pending state
			Signer:   auth.Signer,
			Value:    big.NewInt(2), // value in "Ether", we send with transaction
			GasPrice: nil,           // nil automatically suggests gas price
			GasLimit: 0,             // 0 automatically estimates gas limit
			Context:  context.Background(),
		}, purce, paymentMethod,
	)

	if err != nil {
		log.Printf("could not send cash out request to contract: %v\n", err)

	}
	fmt.Printf("CashOut sent! Please wait for tx %s to be confirmed.\n", txOutRequest.Hash().Hex())

	// check Cash Out Submit
	/*
	   _tx_out_id := 2
	   ga.CashOutSubmit(session,_tx_out_id)
	*/

	// check cash in request
	/*
	   user_wallet := common.HexToAddress("0x892eE0398C9d8C86BCA3ffa49c33b68A7b2F38d3")
	   uuid_in := "3"
	   amount_in := big.NewInt(1)
	   ga.CashInRequest(session,user_wallet,uuid_in,amount_in)
	*/

	// Print result of events, we got during subscription
	eventResult := <-ch
	fmt.Println("/n")
	fmt.Println("Destination for cash_out:", eventResult.Purce)
	fmt.Println("Amount for cash_out:", eventResult.Amount)
	subscription.Unsubscribe()
}
