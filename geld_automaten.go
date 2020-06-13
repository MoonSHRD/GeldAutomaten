package ga

import (
	"math/big"

	dep "github.com/MoonSHRD/GeldAutomaten/artifacts/Deposit"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// CashInRequest is used to create cash-in request (fiat -> blockchain)
// Returns ID of cash-in transaction in hex representation
func CashInRequest(session *dep.DepositSession, wallet common.Address, uuid string, amount *big.Int) (string, error) {
	txCashInRequest, err := session.CashInRequest(wallet, uuid, amount)
	if err != nil {
		return "", err
	}
	return txCashInRequest.Hash().Hex(), nil
}

// CashOutSubmit is used to confirm, that we recive funds from user and did cash out of this funds to user card
// Returns ID of cash-out transaction in hex representation
func CashOutSubmit(session *dep.DepositSession, txid *big.Int) (string, error) {
	txSubmitRequest, err := session.CashOutSubmit(txid)
	if err != nil {
		return "", err
	}
	return txSubmitRequest.Hash().Hex(), nil
}

// CashOutRevert is used in the situation when we have accepted cash-out request from user, but can't proceed it transaction to his payment card
// Returns ID of revert cash-out transaction in hex representation
func CashOutRevert(_session *dep.DepositSession, txid *big.Int, errmsg string) (string, error) {
	txRevertRequest, err := _session.CashOutRevert(txid, errmsg)
	if err != nil {
		return "", err
	}
	return txRevertRequest.Hash().Hex(), nil
}

// SubcribeCashOutRequestFlow is used for subscribing on cash out request flow
func SubcribeCashOutRequestFlow(session *dep.DepositSession, listenChannel chan *dep.DepositCashOutRequestEventAnonymouse) (event.Subscription, error) {
	cashOutFilter := session.Contract.DepositFilterer
	subscription, err := cashOutFilter.WatchCashOutRequestEventAnonymouse(&bind.WatchOpts{
		Start:   nil, //last block
		Context: nil,
	}, listenChannel)
	if err != nil {
		return nil, err
	}
	return subscription, err
}
