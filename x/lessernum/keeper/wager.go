package keeper

import (
	"github.com/alice/checkers/x/lessernum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CollectWager(ctx sdk.Context, k *Keeper, storedGame *types.StoredGame) error {
	if storedGame.PlayerToMove == 1 {
		addr1, err := GetPlayer1Addr(storedGame)
		if err != nil {
			return err
		}

		coin := GetWagerCoin(storedGame)
		err = k.bank.SendCoinsFromAccountToModule(ctx, addr1, types.ModuleName, sdk.NewCoins(coin))
		if err != nil {
			return err
		}

	} else if storedGame.PlayerToMove == 2 {
		addr2, err := GetPlayer2Addr(storedGame)
		if err != nil {
			return err
		}

		coin := GetWagerCoin(storedGame)
		err = k.bank.SendCoinsFromAccountToModule(ctx, addr2, types.ModuleName, sdk.NewCoins(coin))
		if err != nil {
			return err
		}
	}

	return nil
}

func GetPlayer1Addr(storedGame *types.StoredGame) (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(storedGame.Player1)
}

func GetPlayer2Addr(storedGame *types.StoredGame) (sdk.AccAddress, error) {
	return sdk.AccAddressFromBech32(storedGame.Player2)
}

func GetWagerCoin(storedGame *types.StoredGame) (wager sdk.Coin) {
	return sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(storedGame.Wager)))
}
