package blockresultservice

import (
	"context"

	v1 "github.com/cometbft/cometbft/proto/tendermint/services/block_results/v1"
	"github.com/cometbft/cometbft/rpc/core"
	rpctypes "github.com/cometbft/cometbft/rpc/jsonrpc/types"
)

type blockResultsService struct {
	nodeEnv *core.Environment
}

// New creates a new CometBFT block results service server.
func New(env *core.Environment) v1.BlockResultsServiceServer {
	return &blockResultsService{nodeEnv: env}
}

// GetBlockResults returns the block results of the requested height.
// If no height is given, the block results for the latest height are returned.
func (s *blockResultsService) GetBlockResults(_ context.Context, req *v1.GetBlockResultsRequest) (*v1.GetBlockResultsResponse, error) {
	res, err := s.nodeEnv.BlockResults(&rpctypes.Context{}, &req.Height)
	if err != nil {
		return &v1.GetBlockResultsResponse{}, err
	}

	return &v1.GetBlockResultsResponse{
		Height:              req.Height,
		TxsResults:          res.TxsResults,
		FinalizeBlockEvents: res.FinalizeBlockEvents,
		ValidatorUpdates:    res.ValidatorUpdates,
		AppHash:             res.AppHash,
	}, nil
}
