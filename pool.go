package submitter

import (
	"context"
	"errors"
	"github.com/Jeffail/tunny"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"time"
)

// Pool ...
type Pool struct {
	tunnyPool *tunny.Pool
}

// ProcessFunction ...
type ProcessFunction func(txOpts *bind.TransactOpts) (*types.Transaction, error)

// ProcessResult ...
type ProcessResult struct {
	Tx  *types.Transaction
	Err error
}

// NewPool ...
func NewPool(keyStack *Stack, chainID *big.Int, gasPrice *big.Int, gasLimit uint64) (*Pool, error) {
	if keyStack.IsEmpty() {
		return nil, errors.New("empty keys stack")
	}
	tunnyPool := tunny.New(keyStack.Len(), func() tunny.Worker {
		return NewWorker(keyStack, chainID, gasPrice, gasLimit)
	})
	pool := &Pool{
		tunnyPool: tunnyPool,
	}
	return pool, nil
}

// NewPoolWithHexKeys ...
func NewPoolWithHexKeys(hexKeys []string, chainID *big.Int, gasPrice *big.Int, gasLimit uint64) (*Pool, error) {
	keys := make([]*Key, 0, len(hexKeys))
	stack := NewStack(keys)
	for _, hexKey := range hexKeys {
		key, err := NewKey(hexKey)
		if err != nil {
			return nil, err
		}
		if err := stack.Push(key); err != nil {
			return nil, err
		}
	}
	return NewPool(stack, chainID, gasPrice, gasLimit)
}

// Process ...
func (p *Pool) Process(fn ProcessFunction) interface{} {
	return p.tunnyPool.Process(fn)
}

// ProcessCtx ...
func (p *Pool) ProcessCtx(ctx context.Context, fn ProcessFunction) (interface{}, error) {
	return p.tunnyPool.ProcessCtx(ctx, fn)
}

// ProcessTimed ...
func (p *Pool) ProcessTimed(timeout time.Duration, fn ProcessFunction) (interface{}, error) {
	return p.tunnyPool.ProcessTimed(fn, timeout)
}

// Close ...
func (p *Pool) Close() {
	p.tunnyPool.Close()
}
