package submitter

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Key ...
type Key struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

// NewKey ...
func NewKey(hexKey string) (*Key, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, err
	}
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("could not cast to *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKey)
	key := &Key{
		address:    address,
		privateKey: privateKey,
	}
	return key, nil
}

// Address ...
func (k *Key) Address() common.Address {
	return k.address
}

// PrivateKey ...
func (k *Key) PrivateKey() *ecdsa.PrivateKey {
	return k.privateKey
}
