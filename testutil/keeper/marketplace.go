package keeper

import (
	"fmt"
	"testing"

	"github.com/CudoVentures/cudos-node/simapp"
	"github.com/CudoVentures/cudos-node/x/marketplace/keeper"
	"github.com/CudoVentures/cudos-node/x/marketplace/types"
	nftkeeper "github.com/CudoVentures/cudos-node/x/nft/keeper"
	nfttypes "github.com/CudoVentures/cudos-node/x/nft/types"
	tmdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
)

func MarketplaceKeeper(t testing.TB) (*keeper.Keeper, *nftkeeper.Keeper, *bankkeeper.BaseKeeper, sdk.Context) {
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	maccPerms := simapp.GetMaccPerms()
	appCodec := simapp.MakeTestEncodingConfig().Codec

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	authModuleStore, err := setupModuleStore(t, cdc, db, stateStore, authtypes.StoreKey)
	require.NoError(t, err)

	maccPerms[types.ModuleName] = []string{authtypes.Minter}
	authKeeper := authkeeper.NewAccountKeeper(appCodec, authModuleStore.storeKey, authtypes.ProtoBaseAccount, maccPerms, sdk.Bech32MainPrefix, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	bankModuleStore, err := setupModuleStore(t, cdc, db, stateStore, banktypes.StoreKey)
	require.NoError(t, err)

	// bankKeeper := bankkeeper.NewBaseKeeper(appCodec, bankModuleStore.storeKey, authKeeper, bankModuleStore.paramSubspace, nil)
	blockedAddr := make(map[string]bool)
	bankKeeper := bankkeeper.NewBaseKeeper(appCodec, bankModuleStore.storeKey, authKeeper, blockedAddr, authtypes.NewModuleAddress(govtypes.ModuleName).String())

	nftModuleStore, err := setupModuleStore(t, cdc, db, stateStore, nfttypes.StoreKey)
	require.NoError(t, err)

	nftKeeper := nftkeeper.NewKeeper(cdc, nftModuleStore.storeKey, nftModuleStore.memStoreKey)

	moduleStore, err := setupModuleStore(t, cdc, db, stateStore, types.StoreKey)
	require.NoError(t, err)

	k := keeper.NewKeeper(cdc, moduleStore.storeKey, moduleStore.memStoreKey, moduleStore.paramSubspace, bankKeeper, nftKeeper)

	require.NoError(t, stateStore.LoadLatestVersion())

	authKeeper.SetParams(ctx, authtypes.DefaultParams())
	bankKeeper.SetParams(ctx, banktypes.DefaultParams())
	k.SetParams(ctx, types.DefaultParams())

	return k, nftKeeper, &bankKeeper, ctx
}

func setupModuleStore(t testing.TB, cdc *codec.ProtoCodec, db *tmdb.MemDB, stateStore storetypes.CommitMultiStore, storeKeyName string) (moduleStore, error) {
	storeKey := sdk.NewKVStoreKey(storeKeyName)
	memStoreKey := storetypes.NewMemoryStoreKey(fmt.Sprintf("mem_%s", storeKeyName))
	paramsSubspace := typesparams.NewSubspace(cdc, types.Amino, storeKey, memStoreKey, fmt.Sprintf("%sParams", storeKeyName))

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)

	return moduleStore{
		storeKey:      storeKey,
		paramSubspace: paramsSubspace,
	}, nil
}

type moduleStore struct {
	storeKey      *storetypes.KVStoreKey
	memStoreKey   *storetypes.MemoryStoreKey
	paramSubspace typesparams.Subspace
}
