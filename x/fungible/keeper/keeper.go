package keeper

import (
	"fmt"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	"github.com/zeta-chain/zetacore/x/fungible/types"
)

type (
	Keeper struct {
		cdc                codec.BinaryCodec
		storeKey           storetypes.StoreKey
		memKey             storetypes.StoreKey
		paramstore         paramtypes.Subspace
		authKeeper         types.AccountKeeper
		evmKeeper          evmkeeper.Keeper
		bankKeeper         types.BankKeeper
		zetaobserverKeeper types.ZetaObserverKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	authKeeper types.AccountKeeper,
	evmKeeper evmkeeper.Keeper,
	bankKeeper types.BankKeeper,
	zetacobservKeeper types.ZetaObserverKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:                cdc,
		storeKey:           storeKey,
		memKey:             memKey,
		paramstore:         ps,
		authKeeper:         authKeeper,
		evmKeeper:          evmKeeper,
		bankKeeper:         bankKeeper,
		zetaobserverKeeper: zetacobservKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
