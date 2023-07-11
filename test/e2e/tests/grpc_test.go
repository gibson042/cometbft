package e2e_test

import (
	"context"
	v1 "github.com/cometbft/cometbft/proto/tendermint/services/block_results/v1"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"
	"github.com/cometbft/cometbft/version"
)

func TestGRPC_Version(t *testing.T) {
	testNode(t, func(t *testing.T, node e2e.Node) {
		if node.Mode != e2e.ModeFull && node.Mode != e2e.ModeValidator {
			return
		}

		ctx, ctxCancel := context.WithTimeout(context.Background(), time.Minute)
		defer ctxCancel()
		client, err := node.GRPCClient(ctx)
		require.NoError(t, err)

		res, err := client.GetVersion(ctx)
		require.NoError(t, err)

		require.Equal(t, version.TMCoreSemVer, res.Node)
		require.Equal(t, version.ABCIVersion, res.ABCI)
		require.Equal(t, version.P2PProtocol, res.P2P)
		require.Equal(t, version.BlockProtocol, res.Block)
	})
}

func TestGRPC_BlockResults(t *testing.T) {
	testNode(t, func(t *testing.T, node e2e.Node) {
		// TODO: What is this?
		if node.Mode != e2e.ModeFull && node.Mode != e2e.ModeValidator {
			return
		}

		nodeClient, err := node.Client()
		require.NoError(t, err)

		ctx, ctxCancel := context.WithTimeout(context.Background(), time.Minute)
		defer ctxCancel()

		client, err := node.GRPCClient(ctx)
		require.NoError(t, err)

		res, err := nodeClient.BlockResults(ctx, nil)
		require.NoError(t, err)
		t.Logf("Res: %v", res)

		grpcRes, err := client.GetBlockResults(ctx, v1.GetBlockResultsRequest{})
		require.NoError(t, err)

		require.Equal(t, res.Height, grpcRes.Height)
	})
}
