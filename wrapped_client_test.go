package submitter

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestWrappedEthereumClient_URL(t *testing.T) {
	httpClient := &http.Client{}
	client, err := NewWrappedClient(httpClient, "http://localhost:8545")
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost:8545", client.URL())
}
