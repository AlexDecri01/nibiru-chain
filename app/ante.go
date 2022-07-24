package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibcante "github.com/cosmos/ibc-go/v3/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"

	"github.com/NibiruChain/nibiru/app/antedecorators"
	perpkeeper "github.com/NibiruChain/nibiru/x/perp/keeper"
	pricefeedkeeper "github.com/NibiruChain/nibiru/x/pricefeed/keeper"
)

type AnteHandlerOptions struct {
	ante.HandlerOptions
	IBCKeeper *ibckeeper.Keeper

	PricefeedKeeper *pricefeedkeeper.Keeper
	PerpKeeper      *perpkeeper.Keeper
}

/* NewAnteHandler returns and AnteHandler that checks and increments sequence
numbers, checks signatures and account numbers, and deducts fees from the first signer.
*/
func NewAnteHandler(options AnteHandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for AnteHandler")
	}
	if options.BankKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for AnteHandler")
	}
	if options.SignModeHandler == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	if options.SigGasConsumer == nil {
		options.SigGasConsumer = ante.DefaultSigVerificationGasConsumer
	}
	if options.IBCKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "ibc keeper is required for AnteHandler")
	}
	if options.PricefeedKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "pricefeed keeper is required for ante builder")
	}
	if options.PerpKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "perp keeper is required for ante builder")
	}

	memPoolDecorator := ante.NewMempoolFeeDecorator()
	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(),
		antedecorators.NewGaslessDecorator([]sdk.AnteDecorator{&memPoolDecorator}, *options.PricefeedKeeper, *options.PerpKeeper),
		ante.NewRejectExtensionOptionsDecorator(),
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewAnteDecorator(options.IBCKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}
