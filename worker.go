package submitter

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"time"
)

// Worker ...
type Worker struct {
	keyStack *Stack
	chainID  *big.Int
	gasPrice *big.Int
	gasLimit uint64
}

// NewWorker ...
func NewWorker(keyStack *Stack, chainID *big.Int, gasPrice *big.Int, gasLimit uint64) *Worker {
	return &Worker{
		keyStack: keyStack,
		chainID:  chainID,
		gasPrice: gasPrice,
		gasLimit: gasLimit,
	}
}

// Process ...
func (w *Worker) Process(i interface{}) interface{} {
	fn, ok := i.(ProcessFunction)
	if !ok {
		return ProcessResult{Err: errors.New("could not cast input to ProcessFunction")}
	}
	key, err := w.keyStack.Pop()
	if err != nil {
		return ProcessResult{Err: err}
	}
	defer func() {
		if err := w.keyStack.Push(key); err != nil {
		}
	}()
	txOpts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey(), w.chainID)
	if err != nil {
		return ProcessResult{Err: err}
	}
	txOpts.From = key.Address()
	txOpts.GasLimit = w.gasLimit
	txOpts.GasPrice = w.gasPrice
	tx, err := fn(txOpts)
	if err != nil {
		return ProcessResult{Err: err}
	}
	return ProcessResult{Tx: tx}
}

// BlockUntilReady ...
func (w *Worker) BlockUntilReady() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		if !w.keyStack.IsEmpty() {
			return
		}
		select {
		case <-ticker.C:
		}
	}
}

// Interrupt ...
func (w *Worker) Interrupt() {
}

// Terminate ...
func (w *Worker) Terminate() {
}
