package main

import (
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

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

    pk := myenv["PK"] // load private key from env

  //  client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
	client, err := ethclient.Dial(myenv["GATEWAY"])			// load from local .env file
    if err != nil {
        log.Fatalf("could not connect to Ethereum gateway: %v\n", err)
    }
    defer client.Close()


    // check calls 
    // check balance
    accountAddress := common.HexToAddress("0x892eE0398C9d8C86BCA3ffa49c33b68A7b2F38d3")
    balance, _ := client.BalanceAt(ctx, accountAddress, nil)
	fmt.Printf("Balance: %d\n",balance)
    
    // Setting up Deposit Contract
	deposit, err := dep.NewDeposit(common.HexToAddress("0xdD5A73bfE8f907D6592197cF4Ba5945b1bFc608f"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Deposit contract: %v", err)
	}

    // Check view calls (owner address)
	owner,err := deposit.Owner(nil)
	if err != nil {
		log.Fatalf("Failed to owner of contract: %v", err)
	}

	owner_human := owner.Hex()

    fmt.Println(owner_human)
    
    // setting up private key in proper format
    privateKey, err := crypto.HexToECDSA(pk)
    if err != nil {
        log.Fatal(err)
    }

    // Creating an auth transactor
    auth := bind.NewKeyedTransactor(privateKey)
    
    // Wrap the Deposit contract instance into a session
      session := &dep.DepositSession{
	Contract: deposit,
	CallOpts: bind.CallOpts{
        Pending: true,
        From: auth.From,
        Context: context.Background(),
	},
	TransactOpts: bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
        GasLimit: uint64(5000000),
        GasPrice: big.NewInt(1),
        Context: context.Background(),
	},
}

// check call for getting owner inside session
    ownr,err := session.Owner()
    if err != nil {
    log.Fatalf("Failed to owner of contract: %v", err)
    }
    oh:= ownr.Hex()
    fmt.Println(oh)


    // check Cash out request
    _purce := "address karty"
    _payment_method := "karta"
    /*
    txOutRequest,err := session.CashOutRequest(_purce,_payment_method),big.NewInt(1)
    */

    txOutRequest,err := deposit.CashOutRequest(&bind.TransactOpts{
        From: auth.From,
        Nonce: nil,           // nil uses nonce of pending state
        Signer: auth.Signer,
        Value: big.NewInt(2),
        GasPrice: nil,        // nil automatically suggests gas price
        GasLimit: 0,          // 0 automatically estimates gas limit
        Context: context.Background(),
    },_purce,_payment_method,
    )

    if err != nil {
        log.Printf("could not send cash out request to contract: %v\n", err)
        
    }
    fmt.Printf("CashOut sent! Please wait for tx %s to be confirmed.\n", txOutRequest.Hash().Hex())



    // check Cash Out Submit
   // _tx_out_id := 2
    txSubmitRequest,err := session.CashOutSubmit(big.NewInt(2))
    if err != nil {
        log.Printf("could not send cash out submit to contract: %v\n", err)
        
    }
    fmt.Printf("CashOut Submit sent! Please wait for tx %s to be confirmed.\n", txSubmitRequest.Hash().Hex())




}

/*
// cash out request (for test purposes)
func OutRequest(session &dep.DepositSession, string _purce, string _paymentMethod) {
    
    txOutRequest, err := session.
}
*/


/*
// cash out request (for test purposes)
func OutRequest(&dep.DepositSession session, string _purce, string _paymentMethod) {
    
    txOutRequest, err := session.
}
*/

