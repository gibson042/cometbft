package client

import (
	"context"

	"github.com/cosmos/gogoproto/grpc"

	v1 "github.com/cometbft/cometbft/proto/tendermint/services/block_results/v1"
	ctypes "github.com/cometbft/cometbft/rpc/core/types"
)

// BlockResultsServiceClient provides the block results of a given height (or latest if none provided).
type BlockResultsServiceClient interface {
	GetBlockResults(ctx context.Context, req v1.GetBlockResultsRequest) (*ctypes.ResultBlockResults, error)
}

type blockResultServiceClient struct {
	client v1.BlockResultsServiceClient
}

// TODO: Reimplement the type ResultBlockResults
func (b blockResultServiceClient) GetBlockResults(ctx context.Context, req v1.GetBlockResultsRequest) (*ctypes.ResultBlockResults, error) {
	res, err := b.client.GetBlockResults(ctx, &v1.GetBlockResultsRequest{Height: req.Height})
	if err != nil {
		return &ctypes.ResultBlockResults{}, err
	}

	return &ctypes.ResultBlockResults{
		Height:                res.Height,
		TxsResults:            res.TxsResults,
		FinalizeBlockEvents:   res.FinalizeBlockEvents,
		ValidatorUpdates:      res.ValidatorUpdates,
		ConsensusParamUpdates: res.ConsensusParamUpdates,
		AppHash:               res.AppHash,
	}, nil
}

func newBlockResultsServiceClient(conn grpc.ClientConn) BlockResultsServiceClient {
	return &blockResultServiceClient{
		client: v1.NewBlockResultsServiceClient(conn),
	}
}

type disabledBlockResultsServiceClient struct{}

func newDisabledBlockResultsServiceClient() BlockResultsServiceClient {
	return &disabledBlockResultsServiceClient{}
}

// GetBLockResults implements BlockResultsServiceClient
func (*disabledBlockResultsServiceClient) GetBlockResults(_ context.Context, _ v1.GetBlockResultsRequest) (*ctypes.ResultBlockResults, error) {
	panic("block results service client is disabled")
}
