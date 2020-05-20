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
	deposit, err := dep.NewDeposit(common.HexToAddress("0xf460e21B21d294f6832aF166862c38E6Be71f423"), client)
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


   /*
    Events
    */

    // Check retriving (past) events of CashOut request

    var ch = make(chan *dep.DepositCashOutRequestEventAnonymouse)
    cash_out_filter := session.Contract.DepositFilterer
   // cash_out_filter.WatchCashOutRequestEventAnonymouse(nil,ch)
    subscription,err := cash_out_filter.WatchCashOutRequestEventAnonymouse(&bind.WatchOpts{
            Start: nil,
            Context: nil,
    },ch)
    if err != nil {
        log.Printf("error due subscription to event")
            log.Fatalln(err)
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
   // _tx_out_id := 2
   /*
    txSubmitRequest,err := session.CashOutSubmit(big.NewInt(2))
    if err != nil {
        log.Printf("could not send cash out submit to contract: %v\n", err)
        
    }
    fmt.Printf("CashOut Submit sent! Please wait for tx %s to be confirmed.\n", txSubmitRequest.Hash().Hex())
    */


    // check cash in request

    user_wallet := common.HexToAddress("0x892eE0398C9d8C86BCA3ffa49c33b68A7b2F38d3")
    uuid_in := "1"
    amount_in := big.NewInt(1)


    txCashInRequest, err := session.CashInRequest(user_wallet,uuid_in,amount_in)
    if err != nil {
        log.Printf("could not send cash IN submit to contract: %v\n", err)
    }
    fmt.Printf("CashOut Submit sent! Please wait for tx %s to be confirmed.\n", txCashInRequest.Hash().Hex())


 
    //cash_out_filter := session.Contract.WatchCashOutRequestEventAnonymouse(nil,)
    
   // fmt.Printf("subscription:")
   // fmt.Printf(subscription)

    event_result := <-ch

    fmt.Println("/n")
    fmt.Println("Destination for cash_out:", event_result.Purce)
    fmt.Println("Amount for cash_out:", event_result.Amount)
    subscription.Unsubscribe()

    /*
    for {
        select {
        case err := <-subscription.Err():
            log.Printf("error due subscription to event")
            log.Fatalln(err)
        
          
    case event_log := <-ch :
        fmt.Println("Destination: ")

        }

    }
    */

    /*
    for {
		select {
		case err := <-subscription.Err():
			log.Fatal(err)
		case log := <-ch:
			var greetEvent struct {
				Name  string
				Count *big.Int
			}

			err = greeterAbi.Unpack(&greetEvent, "_Greet", log.Data)

			if err != nil {
				fmt.Println("Failed to unpack:", err)
			}

			fmt.Println("Contract:", log.Address.Hex())
			fmt.Println("Name:", greetEvent.Name)
			fmt.Println("Count:", greetEvent.Count)
		}
	}
    */



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

