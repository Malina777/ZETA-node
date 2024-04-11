package observer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zeta-chain/zetacore/testutil/keeper"
	"github.com/zeta-chain/zetacore/testutil/nullify"
	"github.com/zeta-chain/zetacore/testutil/sample"
	"github.com/zeta-chain/zetacore/x/observer"
	"github.com/zeta-chain/zetacore/x/observer/types"
)

func TestGenesis(t *testing.T) {
	t.Run("genState fields defined", func(t *testing.T) {
		params := types.DefaultParams()
		tss := sample.Tss()
		genesisState := types.GenesisState{
			Params:    &params,
			Tss:       &tss,
			BlameList: sample.BlameRecordsList(t, 10),
			Ballots: []*types.Ballot{
				sample.Ballot(t, "0"),
				sample.Ballot(t, "1"),
				sample.Ballot(t, "2"),
			},
			Observers: sample.ObserverSet(3),
			NodeAccountList: []*types.NodeAccount{
				sample.NodeAccount(),
				sample.NodeAccount(),
				sample.NodeAccount(),
			},
			CrosschainFlags:   types.DefaultCrosschainFlags(),
			Keygen:            sample.Keygen(t),
			ChainParamsList:   sample.ChainParamsList(),
			LastObserverCount: sample.LastObserverCount(10),
			TssFundMigrators:  []types.TssFundMigratorInfo{sample.TssFundsMigrator(1), sample.TssFundsMigrator(2)},
			ChainNonces: []types.ChainNonces{
				sample.ChainNonces(t, "0"),
				sample.ChainNonces(t, "1"),
				sample.ChainNonces(t, "2"),
			},
			PendingNonces: sample.PendingNoncesList(t, "sample", 20),
			NonceToCctx:   sample.NonceToCctxList(t, "sample", 20),
			TssHistory:    []types.TSS{sample.Tss()},
		}

		// Init and export
		k, ctx, _, _ := keepertest.ObserverKeeper(t)
		observer.InitGenesis(ctx, *k, genesisState)
		got := observer.ExportGenesis(ctx, *k)
		require.NotNil(t, got)

		// Compare genesis after init and export
		nullify.Fill(&genesisState)
		nullify.Fill(got)
		require.Equal(t, genesisState, *got)
	})

	t.Run("genState fields not defined", func(t *testing.T) {
		genesisState := types.GenesisState{}

		k, ctx, _, _ := keepertest.ObserverKeeper(t)
		observer.InitGenesis(ctx, *k, genesisState)
		got := observer.ExportGenesis(ctx, *k)
		require.NotNil(t, got)

		defaultParams := types.DefaultParams()
		btcChainParams := types.GetDefaultBtcRegtestChainParams()
		btcChainParams.IsSupported = true
		goerliChainParams := types.GetDefaultGoerliLocalnetChainParams()
		goerliChainParams.IsSupported = true
		zetaPrivnetChainParams := types.GetDefaultZetaPrivnetChainParams()
		zetaPrivnetChainParams.IsSupported = true
		localnetChainParams := types.ChainParamsList{
			ChainParams: []*types.ChainParams{
				btcChainParams,
				goerliChainParams,
				zetaPrivnetChainParams,
			},
		}
		expectedGenesisState := types.GenesisState{
			Params:            &defaultParams,
			CrosschainFlags:   types.DefaultCrosschainFlags(),
			ChainParamsList:   localnetChainParams,
			Tss:               &types.TSS{},
			Keygen:            &types.Keygen{},
			LastObserverCount: &types.LastObserverCount{},
			NodeAccountList:   []*types.NodeAccount{},
		}

		require.Equal(t, expectedGenesisState, *got)
	})

	t.Run("genState fields not defined except tss", func(t *testing.T) {
		tss := sample.Tss()
		genesisState := types.GenesisState{
			Tss: &tss,
		}

		k, ctx, _, _ := keepertest.ObserverKeeper(t)
		observer.InitGenesis(ctx, *k, genesisState)
		got := observer.ExportGenesis(ctx, *k)
		require.NotNil(t, got)

		defaultParams := types.DefaultParams()
		btcChainParams := types.GetDefaultBtcRegtestChainParams()
		btcChainParams.IsSupported = true
		goerliChainParams := types.GetDefaultGoerliLocalnetChainParams()
		goerliChainParams.IsSupported = true
		zetaPrivnetChainParams := types.GetDefaultZetaPrivnetChainParams()
		zetaPrivnetChainParams.IsSupported = true
		localnetChainParams := types.ChainParamsList{
			ChainParams: []*types.ChainParams{
				btcChainParams,
				goerliChainParams,
				zetaPrivnetChainParams,
			},
		}
		pendingNonces, err := k.GetAllPendingNonces(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, pendingNonces)
		expectedGenesisState := types.GenesisState{
			Params:            &defaultParams,
			CrosschainFlags:   types.DefaultCrosschainFlags(),
			ChainParamsList:   localnetChainParams,
			Tss:               &tss,
			Keygen:            &types.Keygen{},
			LastObserverCount: &types.LastObserverCount{},
			NodeAccountList:   []*types.NodeAccount{},
			PendingNonces:     pendingNonces,
		}

		require.Equal(t, expectedGenesisState, *got)
	})

	t.Run("export without init", func(t *testing.T) {
		k, ctx, _, _ := keepertest.ObserverKeeper(t)

		got := observer.ExportGenesis(ctx, *k)
		require.NotNil(t, got)

		params := k.GetParams(ctx)
		expectedGenesisState := types.GenesisState{
			Params:            &params,
			CrosschainFlags:   types.DefaultCrosschainFlags(),
			ChainParamsList:   types.ChainParamsList{},
			Tss:               &types.TSS{},
			Keygen:            &types.Keygen{},
			LastObserverCount: &types.LastObserverCount{},
			NodeAccountList:   []*types.NodeAccount{},
			Ballots:           k.GetAllBallots(ctx),
			TssHistory:        k.GetAllTSS(ctx),
			TssFundMigrators:  k.GetAllTssFundMigrators(ctx),
			BlameList:         k.GetAllBlame(ctx),
			ChainNonces:       k.GetAllChainNonces(ctx),
			NonceToCctx:       k.GetAllNonceToCctx(ctx),
		}

		require.Equal(t, expectedGenesisState, *got)
	})
}
