package submitter

import (
	"net/http"
)

// NewClientBalancer ...
func NewClientBalancer(clients ...*WrappedClient) *Balancer[*WrappedClient] {
	return NewBalancer(clients...)
}

// NewClientBalancerWithHttpClient ...
func NewClientBalancerWithHttpClient(httpClient *http.Client, urls []string) (*Balancer[*WrappedClient], error) {
	clients := make([]*WrappedClient, 0, len(urls))
	for _, url := range urls {
		wrappedClient, err := NewWrappedClient(httpClient, url)
		if err != nil {
			return nil, err
		}
		clients = append(clients, wrappedClient)
	}
	return NewBalancer(clients...), nil
}

// NewClientBalancerWithURLs ...
func NewClientBalancerWithURLs(fn func() *http.Client, urls []string) (*Balancer[*WrappedClient], error) {
	clients := make([]*WrappedClient, 0, len(urls))
	for _, url := range urls {
		wrappedClient, err := NewWrappedClient(fn(), url)
		if err != nil {
			return nil, err
		}
		clients = append(clients, wrappedClient)
	}
	return NewBalancer(clients...), nil
}
