package main

import (
	"github.com/ethereum/go-ethereum/event"
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


// create cashInRequest (fiat -> blockchain)
func CashInRequest(_session *dep.DepositSession, wallet common.Address, _uuid string, _amount *big.Int) {

    
    txCashInRequest, err := _session.CashInRequest(wallet,_uuid,_amount)
    if err != nil {
        log.Printf("could not send cash IN submit to contract: %v\n", err)
    }
    fmt.Printf("CashIN Submit sent! Please wait for tx %s to be confirmed.\n", txCashInRequest.Hash().Hex())

}



// submit, that we recive funds from user and maked cash_out of this funds to user card. 
func CashOutSubmit(_session *dep.DepositSession, txid *big.Int) {

    txSubmitRequest,err := _session.CashOutSubmit(txid)
    if err != nil {
        log.Printf("could not send cash out submit to contract: %v\n", err)
        
    }
    fmt.Printf("CashOut Submit sent! Please wait for tx %s to be confirmed.\n", txSubmitRequest.Hash().Hex())
}

// revert - we have accepted cashOutSubmit from user, but can't proceed it transaction to his payment card
func CashOutRevert(_session *dep.DepositSession, txid *big.Int, errmsg string) {

    txRevertRequest,err := _session.CashOutRevert(txid, errmsg)
    if err != nil {
        log.Printf("could not send cash out submit to contract: %v\n", err)
        
    }
    fmt.Printf("CashOut Submit sent! Please wait for tx %s to be confirmed.\n", txRevertRequest.Hash().Hex())
}


func SubcribeCashOutAnonymouse(_session *dep.DepositSession, _ch chan *dep.DepositCashOutRequestEventAnonymouse) (event.Subscription, error) {

    cash_out_filter := _session.Contract.DepositFilterer
    subscription,err := cash_out_filter.WatchCashOutRequestEventAnonymouse(&bind.WatchOpts{
            Start: nil, //last block
            Context: nil,
    },_ch)
    if err != nil {
        log.Printf("error due subscription to event")
            log.Fatalln(err)
    }
    return subscription, err


}

func main(){
    loadEnv()

    ctx := context.Background()
    pk := myenv["PK"] // load private key from env

    // Connecting to network
  //  client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
	client, err := ethclient.Dial(myenv["GATEWAY"])			// load from local .env file
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
    accountAddress := common.HexToAddress("0x892eE0398C9d8C86BCA3ffa49c33b68A7b2F38d3")
    balance, _ := client.BalanceAt(ctx, accountAddress, nil)    //our balance
	fmt.Printf("Balance: %d\n",balance)
    

    // Setting up Deposit Contract
	deposit, err := dep.NewDeposit(common.HexToAddress("0xf460e21B21d294f6832aF166862c38E6Be71f423"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Deposit contract: %v", err)
    }
    
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
        GasLimit: 0,                    // 0 automatically estimates gas limit
        GasPrice: nil,                  // nil automatically suggests gas price
        Context: context.Background(),
    },
}


   /*
    Events
    */

    // Check retriving events of CashOut request

    var ch = make(chan *dep.DepositCashOutRequestEventAnonymouse)
    subscription, err := SubcribeCashOutAnonymouse(session,ch)


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
    _purce := "address karty"
    _payment_method := "karta"

    txOutRequest,err := deposit.CashOutRequest(&bind.TransactOpts{
        From: auth.From,
        Nonce: nil,           // nil uses nonce of pending state
        Signer: auth.Signer,
        Value: big.NewInt(2), // value in "Ether", we send with transaction
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
   /*
    _tx_out_id := 2
    CashOutSubmit(session,_tx_out_id) 
    */
    


    // check cash in request

    user_wallet := common.HexToAddress("0x892eE0398C9d8C86BCA3ffa49c33b68A7b2F38d3")
    uuid_in := "1"
    amount_in := big.NewInt(1)
    CashInRequest(session,user_wallet,uuid_in,amount_in)

 
   
    // Print result of events, we got during subscription
    event_result := <-ch
    fmt.Println("/n")
    fmt.Println("Destination for cash_out:", event_result.Purce)
    fmt.Println("Amount for cash_out:", event_result.Amount)
    subscription.Unsubscribe()
}



