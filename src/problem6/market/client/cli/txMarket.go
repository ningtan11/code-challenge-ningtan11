package cli

import (
	"strconv"
	"time"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"

	"github.com/Switcheo/carbon/x/market/types"
)

func CmdCreateMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-market [market-type] [base] [quote] [base-usd-price] [quote-usd-price] " +
			"[index-oracle-id] [expiry-time]",
		Short: "Creates a new market",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsMarketType := args[0]
			argsBase := args[1]
			argsQuote := args[2]
			argsCurrentBasePrice := args[3]
			argsCurrentQuotePrice := args[4]
			argsIndexOracleID := args[5]
			argsExpiryTime, err := strconv.ParseInt(args[6], 10, 64)
			if err != nil {
				return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid expiry time: %s", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMarket(
				clientCtx.GetFromAddress().String(),
				argsMarketType,
				argsBase,
				argsQuote,
				sdk.MustNewDecFromStr(argsCurrentBasePrice),
				sdk.MustNewDecFromStr(argsCurrentQuotePrice),
				argsIndexOracleID,
				time.Unix(argsExpiryTime, 0),
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
