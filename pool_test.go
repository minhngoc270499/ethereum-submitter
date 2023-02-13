package submitter

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"math/big"
	"sync"
	"testing"
)

func TestPool_Process(t *testing.T) {
	hexKeys := []string{
		// 0x39eED1b56b1Df7dA68d7C097bE6024Cc133054F1
		"8a45178ddbe19f9a62bd47db0d111a809b2dd947d7018be4d6881c3a2c5a3693",
		// 0x470fF44598A6A4890439919b353DddDBb08924B7
		"554131310bd5fd7f021bc4f699c5a43368455173835393e2ef38d03fc9af1785",
	}
	pool, err := NewPoolWithHexKeys(hexKeys, big.NewInt(1), nil, 0)
	assert.Nil(t, err)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(t *testing.T) {
			defer wg.Done()
			result := pool.Process(func(txOpts *bind.TransactOpts) (*types.Transaction, error) {
				fmt.Println(txOpts.From)
				return nil, nil
			})
			processResult, ok := result.(ProcessResult)
			assert.True(t, ok)
			assert.Nil(t, processResult.Err)
		}(t)
	}
	wg.Wait()
}
