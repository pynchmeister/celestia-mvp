package simulation

import (
	"math/rand"

	"blog/x/blog/keeper"
	"blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgUploadW3S(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUploadW3S{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UploadW3S simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UploadW3S simulation not implemented"), nil, nil
	}
}
