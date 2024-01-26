package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/google/uuid"
)

type SetCitizenDecorator struct {
	ak AccountKeeper
}

func NewSetCitizenDecorator(ak AccountKeeper) SetCitizenDecorator {
	return SetCitizenDecorator{
		ak: ak,
	}
}

func (scd SetCitizenDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
	}

	signers := sigTx.GetSigners()
	pubkeys, err := sigTx.GetPubKeys()
	if err != nil {
		return ctx, err
	}
	for _, signer := range signers {
		acc, err := GetSignerAcc(ctx, scd.ak, signer)
		if err != nil {
			return ctx, err
		}
		// account already has pubkey set,no need to reset
		if len(acc.GetCitizen()) != 0 {
			continue
		}
		citizenId := uuid.New().String()
		citizen := types.Citizen{
			CitizenId: citizenId,
			CitizenName: namesgenerator.GetRandomName(1),
			EthAddress:
		}
		if err != nil {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, err.Error())
		}
		scd.ak.SetAccount(ctx, acc)
	}
	return next(ctx, tx, simulate)
}
