package zetaclient

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types/query"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
	zetaObserverTypes "github.com/zeta-chain/zetacore/x/observer/types"
	"google.golang.org/grpc"
)

func (b *ZetaCoreBridge) GetInboundPermissions() (types.PermissionFlags, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.PermissionFlags(context.Background(), &types.QueryGetPermissionFlagsRequest{})
	if err != nil {
		return types.PermissionFlags{}, err
	}
	return resp.PermissionFlags, nil

}

func (b *ZetaCoreBridge) GetCoreParamsForChainID(externalChainID int64) (*zetaObserverTypes.CoreParams, error) {
	client := zetaObserverTypes.NewQueryClient(b.grpcConn)
	resp, err := client.GetCoreParamsForChain(context.Background(), &zetaObserverTypes.QueryGetCoreParamsForChainRequest{ChainID: externalChainID})
	if err != nil {
		return &zetaObserverTypes.CoreParams{}, err
	}
	return resp.CoreParams, nil
}

func (b *ZetaCoreBridge) GetCoreParams() ([]*zetaObserverTypes.CoreParams, error) {
	client := zetaObserverTypes.NewQueryClient(b.grpcConn)
	resp, err := client.GetCoreParams(context.Background(), &zetaObserverTypes.QueryGetCoreParamsRequest{})
	if err != nil {
		return nil, err
	}
	return resp.CoreParams.CoreParams, nil
}

func (b *ZetaCoreBridge) GetObserverParams() (zetaObserverTypes.Params, error) {
	client := zetaObserverTypes.NewQueryClient(b.grpcConn)
	resp, err := client.Params(context.Background(), &zetaObserverTypes.QueryParamsRequest{})
	if err != nil {
		return zetaObserverTypes.Params{}, err
	}
	return resp.Params, nil
}

//func (b *ZetaCoreBridge) GetAccountDetails(address string) (string, error) {
//	client := authtypes.NewQueryClient(b.grpcConn)
//	resp, err := client.Account(context.Background(), &authtypes.QueryAccountRequest{
//		Address: address,
//	})
//	if err != nil {
//		b.logger.Error().Err(err).Msg("Query account failed")
//		return "", err
//	}
//
//	err := resp.UnpackInterfaces
//	return resp.Account.GetTypeUrl(), nil
//
//}

func (b *ZetaCoreBridge) GetAllCctx() ([]*types.CrossChainTx, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.CctxAll(context.Background(), &types.QueryAllCctxRequest{})
	if err != nil {
		return nil, err
	}
	return resp.CrossChainTx, nil
}

func (b *ZetaCoreBridge) GetCctxByHash(sendHash string) (*types.CrossChainTx, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.Cctx(context.Background(), &types.QueryGetCctxRequest{Index: sendHash})
	if err != nil {
		return nil, err
	}
	return resp.CrossChainTx, nil
}

func (b *ZetaCoreBridge) GetObserverList(chain common.Chain) ([]string, error) {
	client := zetaObserverTypes.NewQueryClient(b.grpcConn)
	resp, err := client.ObserversByChain(context.Background(), &zetaObserverTypes.QueryObserversByChainRequest{
		ObservationChain: chain.ChainName.String(),
	})
	if err != nil {
		return nil, err
	}
	return resp.Observers, nil
}

func (b *ZetaCoreBridge) GetAllPendingCctx() ([]*types.CrossChainTx, error) {
	client := types.NewQueryClient(b.grpcConn)
	maxSizeOption := grpc.MaxCallRecvMsgSize(32 * 1024 * 1024)
	resp, err := client.CctxAllPending(context.Background(), &types.QueryAllCctxPendingRequest{}, maxSizeOption)
	if err != nil {
		return nil, err
	}
	return resp.CrossChainTx, nil
}

func (b *ZetaCoreBridge) GetLastBlockHeight() ([]*types.LastBlockHeight, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.LastBlockHeightAll(context.Background(), &types.QueryAllLastBlockHeightRequest{})
	if err != nil {
		b.logger.Error().Err(err).Msg("query GetBlockHeight error")
		return nil, err
	}
	return resp.LastBlockHeight, nil
}

func (b *ZetaCoreBridge) GetLatestZetaBlock() (*tmtypes.Block, error) {
	client := tmservice.NewServiceClient(b.grpcConn)
	res, err := client.GetLatestBlock(context.Background(), &tmservice.GetLatestBlockRequest{})
	if err != nil {
		return nil, err
	}
	return res.Block, nil
}

func (b *ZetaCoreBridge) GetLastBlockHeightByChain(chain common.Chain) (*types.LastBlockHeight, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.LastBlockHeight(context.Background(), &types.QueryGetLastBlockHeightRequest{Index: chain.ChainName.String()})
	if err != nil {
		return nil, err
	}
	return resp.LastBlockHeight, nil
}

func (b *ZetaCoreBridge) GetZetaBlockHeight() (int64, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.LastZetaHeight(context.Background(), &types.QueryLastZetaHeightRequest{})
	if err != nil {
		return 0, err
	}
	return resp.Height, nil
}

func (b *ZetaCoreBridge) GetNonceByChain(chain common.Chain) (*types.ChainNonces, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.ChainNonces(context.Background(), &types.QueryGetChainNoncesRequest{Index: chain.ChainName.String()})
	if err != nil {
		return nil, err
	}
	return resp.ChainNonces, nil
}

func (b *ZetaCoreBridge) GetAllNodeAccounts() ([]*types.NodeAccount, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.NodeAccountAll(context.Background(), &types.QueryAllNodeAccountRequest{})
	if err != nil {
		return nil, err
	}
	b.logger.Debug().Msgf("GetAllNodeAccounts: %d", len(resp.NodeAccount))
	return resp.NodeAccount, nil
}

func (b *ZetaCoreBridge) GetKeyGen() (*types.Keygen, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.Keygen(context.Background(), &types.QueryGetKeygenRequest{})
	if err != nil {
		//log.Error().Err(err).Msg("query GetKeyGen error")
		return nil, err
	}
	return resp.Keygen, nil
}

func (b *ZetaCoreBridge) GetOutTxTracker(chain common.Chain, nonce uint64) (*types.OutTxTracker, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.OutTxTracker(context.Background(), &types.QueryGetOutTxTrackerRequest{
		ChainID: chain.ChainId,
		Nonce:   nonce,
	})
	if err != nil {
		return nil, err
	}
	return &resp.OutTxTracker, nil
}

func (b *ZetaCoreBridge) GetAllOutTxTrackerByChain(chain common.Chain) ([]types.OutTxTracker, error) {
	client := types.NewQueryClient(b.grpcConn)
	resp, err := client.OutTxTrackerAllByChain(context.Background(), &types.QueryAllOutTxTrackerByChainRequest{
		Chain: chain.ChainId,
		Pagination: &query.PageRequest{
			Key:        nil,
			Offset:     0,
			Limit:      300,
			CountTotal: false,
			Reverse:    false,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp.OutTxTracker, nil
}

func (b *ZetaCoreBridge) GetClientParams(chainID int64) (zetaObserverTypes.QueryGetCoreParamsForChainResponse, error) {
	client := zetaObserverTypes.NewQueryClient(b.grpcConn)
	resp, err := client.GetCoreParamsForChain(context.Background(), &zetaObserverTypes.QueryGetCoreParamsForChainRequest{ChainID: chainID})
	if err != nil {
		return zetaObserverTypes.QueryGetCoreParamsForChainResponse{}, err
	}
	return *resp, nil
}
