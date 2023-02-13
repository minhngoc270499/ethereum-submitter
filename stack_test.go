package submitter

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_WhenInitNilKeys(t *testing.T) {
	stack := NewStack(nil)
	key, _ := NewKey("8a45178ddbe19f9a62bd47db0d111a809b2dd947d7018be4d6881c3a2c5a3693")
	err := stack.Push(key)
	assert.Nil(t, err)
	assert.False(t, stack.IsEmpty())
	popKey, err := stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, common.HexToAddress("0x470fF44598A6A4890439919b353DddDBb08924B7"), popKey.Address())
	assert.True(t, stack.IsEmpty())
}

func TestStack_WhenPushNilKey(t *testing.T) {
	stack := NewStack(nil)
	err := stack.Push(nil)
	assert.Equal(t, "nil key", err.Error())
}

func TestStack_WhenPopFromNilKeyS(t *testing.T) {
	stack := NewStack(nil)
	_, err := stack.Pop()
	assert.Equal(t, "empty keys", err.Error())
}

func TestStack_WhenPopMultipleKeys(t *testing.T) {
	hexKeys := []string{
		// 0x39eED1b56b1Df7dA68d7C097bE6024Cc133054F1
		"8a45178ddbe19f9a62bd47db0d111a809b2dd947d7018be4d6881c3a2c5a3693",
		// 0x470fF44598A6A4890439919b353DddDBb08924B7
		"554131310bd5fd7f021bc4f699c5a43368455173835393e2ef38d03fc9af1785",
	}
	stack := NewStack(nil)
	for _, hexKey := range hexKeys {
		key, _ := NewKey(hexKey)
		err := stack.Push(key)
		assert.Nil(t, err)
	}
	firstPopKey, err := stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, common.HexToAddress("0x39eED1b56b1Df7dA68d7C097bE6024Cc133054F1"), firstPopKey.Address())
	secondPopKey, err := stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, common.HexToAddress("0x470fF44598A6A4890439919b353DddDBb08924B7"), secondPopKey.Address())
	assert.True(t, stack.IsEmpty())
}

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack(nil)
	assert.True(t, stack.IsEmpty())
}
