package submitter

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKeyWhenEmptyHexKey(t *testing.T) {
	_, err := NewKey("")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "invalid length, need 256 bits")
}

func TestNewKeyWhenInvalidHexKey(t *testing.T) {
	_, err := NewKey("8a45178ddbe19f9a62bd47db0d111a809b2dd")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "invalid hex data for private key")
}

func TestNewKeyWhenValidHexKey(t *testing.T) {
	key, err := NewKey("8a45178ddbe19f9a62bd47db0d111a809b2dd947d7018be4d6881c3a2c5a3693")
	assert.Nil(t, err)
	assert.Equal(t, common.HexToAddress("0x470fF44598A6A4890439919b353DddDBb08924B7"), key.Address())
}
