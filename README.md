# GeldAutomaten
Bridge between fiat gateway and shard blockchain

## Basic explanation
GeldAutomat - Bankomat

Main idea of this project - launch it as a service, to be able to work with Deposit contract.

Deposit contract is entity which store user requests (to cash out money from blockchain to fiat)
Also store and proceed Administrator request (to cash in money from plastic card to blockchain)

Therefore GeldAutomat is a bridge, which allow you to exchange fiat money to your stablecoin at Ethereum network (including private networks)


## How it works:

1. service load your local .env file, where you store your admin private key and your endpoint to connect to blockchain network like this:
```
GATEWAY = 'WS://127.0.0.1:7545'
PK = "your awesome private key" 
```
**NOTE** - as we working with events - we need to connect with node by WebSocket, because HTTP doesn't support notifications.

2. Connect to network, set up private key, create auth Transactor object:
```go
// Connecting to network
// client, err := ethclient.Dial(os.Getenv("GATEWAY"))	// for global env config
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
```

3. Create instance of Deposit contract and session object:
```go
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
```

4. Subscribe to user events:
```
var ch = make(chan *dep.DepositCashOutRequestEventAnonymouse)
subscription, err := SubcribeCashOutAnonymouse(session, ch)
```

## How to use it

1. If we want to fund user address after his has made payment from plastic card and we have recived it:
```go
// check cash in request
user_wallet := common.HexToAddress("0x892eE0398C9d8C86BCA3ffa49c33b68A7b2F38d3")
uuid_in := "1"
amount_in := big.NewInt(1)
CashInRequest(session,user_wallet,uuid_in,amount_in)
```
2. If we want to payout (cashout) funds from user address to his plastic card:
```go
var ch = make(chan *dep.DepositCashOutRequestEventAnonymouse)
subscription, err := SubcribeCashOutAnonymouse(session,ch)

event_result := <-ch
...
```
After getting new events about cashout requests - we need to proceed this payout transactions through our payment provider (UnitPay at this example)


3. After proceeding transaction we need to:

	a. Submit, that we have proceeded transaction via
	```go
	CashOutSubmit(_session *dep.DepositSession, txid *big.Int)
	```
	where txid is id of requested cash_out request

	b. Revert, proceed funds back to user account:
	```go
	CashOutRevert(_session *dep.DepositSession, txid *big.Int, errmsg string)
	```







