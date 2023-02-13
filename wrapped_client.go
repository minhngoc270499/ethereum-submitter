package submitter

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"net/http"
)

// WrappedClient ...
type WrappedClient struct {
	url    string
	client *ethclient.Client
}

// NewWrappedClient ...
func NewWrappedClient(httpClient *http.Client, url string) (*WrappedClient, error) {
	client, err := rpc.DialHTTPWithClient(url, httpClient)
	if err != nil {
		return nil, err
	}
	wrappedClient := &WrappedClient{
		url:    url,
		client: ethclient.NewClient(client),
	}
	return wrappedClient, nil
}

// URL ...
func (c WrappedClient) URL() string {
	return c.url
}

// Client ...
func (c WrappedClient) Client() *ethclient.Client {
	return c.client
}
