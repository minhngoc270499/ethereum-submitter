package submitter

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/sync/errgroup"
	"strings"
	"time"
)

// SendTransactionWithBalancer ...
func SendTransactionWithBalancer(ctx context.Context, balancer *Balancer[*WrappedClient], tx *types.Transaction) error {
	g := new(errgroup.Group)
	for _, client := range balancer.All() {
		client := client
		g.Go(func() error {
			err := client.Client().SendTransaction(ctx, tx)
			if err == nil {
				return nil
			}
			if err.Error() != "known transaction" || strings.Contains(err.Error(), "already known") {
				return nil
			}
			return err
		})
	}
	return g.Wait()
}

// SendAndWaitTransactionWithBalancer ...
func SendAndWaitTransactionWithBalancer(ctx context.Context, balancer *Balancer[*WrappedClient], tx *types.Transaction) (*types.Receipt, error) {
	if err := SendTransactionWithBalancer(ctx, balancer, tx); err != nil {
		return nil, err
	}
	return WaitMinedWithBalancer(ctx, balancer, tx)
}

// WaitMinedWithBalancer ...
func WaitMinedWithBalancer(ctx context.Context, balancer *Balancer[*WrappedClient], tx *types.Transaction) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	for {
		b := balancer.Next().Client()
		receipt, err := b.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			return receipt, nil
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}
