package submitter

import (
	"fmt"
	"net/http"
)

// Proxy ...
type Proxy struct {
	balancerMap map[uint64]*Balancer[*WrappedClient]
}

// NewProxy ...
func NewProxy(balancerMap map[uint64]*Balancer[*WrappedClient]) *Proxy {
	return &Proxy{
		balancerMap: balancerMap,
	}
}

// NewProxyWithUrls ...
func NewProxyWithUrls(httpClient *http.Client, urlMap map[uint64][]string) (*Proxy, error) {
	balancerMap := make(map[uint64]*Balancer[*WrappedClient], 0)
	for chainID, urls := range urlMap {
		balancer, err := NewClientBalancerWithHttpClient(httpClient, urls)
		if err != nil {
			return nil, err
		}
		balancerMap[chainID] = balancer
	}
	proxy := &Proxy{
		balancerMap: balancerMap,
	}
	return proxy, nil
}

// GetBalancerByChainID ...
func (p Proxy) GetBalancerByChainID(chainID uint64) (*Balancer[*WrappedClient], error) {
	balancer, ok := p.balancerMap[chainID]
	if !ok {
		return nil, fmt.Errorf("balancer with chain id %v not found", chainID)
	}
	return balancer, nil
}
