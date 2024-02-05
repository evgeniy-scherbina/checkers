package keeper

import (
	"context"

	"github.com/alice/checkers/x/monobank/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BalanceAll(c context.Context, req *types.QueryAllBalanceRequest) (*types.QueryAllBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var balances []types.Balance
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	balanceStore := prefix.NewStore(store, types.KeyPrefix(types.BalanceKeyPrefix))

	pageRes, err := query.Paginate(balanceStore, req.Pagination, func(key []byte, value []byte) error {
		var balance types.Balance
		if err := k.cdc.Unmarshal(value, &balance); err != nil {
			return err
		}

		balances = append(balances, balance)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBalanceResponse{Balance: balances, Pagination: pageRes}, nil
}

func (k Keeper) Balance(c context.Context, req *types.QueryGetBalanceRequest) (*types.QueryGetBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBalance(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetBalanceResponse{Balance: val}, nil
}
