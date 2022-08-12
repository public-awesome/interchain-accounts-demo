package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/tendermint/tendermint/libs/log"

	controllerkeeper "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/controller/keeper"
	"github.com/cosmos/interchain-accounts/x/inter-tx/types"
)

// Keeper defines the intertx module keeper.
type Keeper struct {
	cdc      codec.Codec
	storeKey storetypes.StoreKey

	controllerKeeper controllerkeeper.Keeper
	scopedKeeper     capabilitykeeper.ScopedKeeper
}

// NewKeeper creates and returns a new intertx keeper instnace.
func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey, controllerKeeper controllerkeeper.Keeper, scopedKeeper capabilitykeeper.ScopedKeeper) Keeper {
	return Keeper{
		cdc:              cdc,
		storeKey:         storeKey,
		controllerKeeper: controllerKeeper,
		scopedKeeper:     scopedKeeper,
	}
}

// Logger returns the application logger, scoped to the intertx module.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k *Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}
