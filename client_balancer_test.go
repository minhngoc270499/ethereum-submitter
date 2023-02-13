package submitter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEthereumClientBalancer_Next(t *testing.T) {
	urls := []string{
		"http://localhost:8540",
		"http://localhost:8541",
		"http://localhost:8542",
		"http://localhost:8543",
		"http://localhost:8544",
		"http://localhost:8545",
		"http://localhost:8546",
		"http://localhost:8547",
		"http://localhost:8548",
		"http://localhost:8549",
	}
	balancer, err := NewClientBalancerWithHttpClient(nil, urls)
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost:8540", balancer.Next().URL())
	assert.Equal(t, "http://localhost:8541", balancer.Next().URL())
	assert.Equal(t, "http://localhost:8542", balancer.Next().URL())
	assert.Equal(t, "http://localhost:8543", balancer.Next().URL())
	assert.Equal(t, "http://localhost:8544", balancer.Next().URL())
	assert.Equal(t, "http://localhost:8545", balancer.Next().URL())
	assert.Equal(t, "http://localhost:8546", balancer.Next().URL())
	assert.Equal(t, "http://localhost:8547", balancer.Next().URL())
	assert.Equal(t, "http://localhost:8548", balancer.Next().URL())
	assert.Equal(t, "http://localhost:8549", balancer.Next().URL())
}
