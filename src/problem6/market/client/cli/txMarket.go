package cli

import (
	"strconv"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Switcheo/carbon/x/admin/client/cli"
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

var (
	flagDisplayName                 = "display-name"
	flagDescription                 = "description"
	flagLotSize                     = "lot-size"
	flagTicksize                    = "tick-size"
	flagMinQuantity                 = "min-quantity"
	flagIsActive                    = "is-active"
	flagRiskStepSize                = "risk-step-size"
	flagInitialMarginBase           = "initial-margin-base"
	flagInitialMarginStep           = "initial-margin-step"
	flagMaintenanceMarginRatio      = "maintenance-margin-ratio"
	flagMaxLiquidationOrderTicket   = "max-liquidation-order-ticket"
	flagMaxLiquidationOrderDuration = "max-liquidation-order-duration"
	flagImpactSize                  = "impact-size"
	flagMarkPriceBand               = "mark-price-band"
	flagLastPriceProtectedBand      = "last-price-protected-band"
	flagTradingBandwidth            = "trading-bandwidth"
	flagMarketExpiry                = "market-expiry"
)

func CmdUpdateMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-market [name] ",
		Short: "Update an existing market",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var displayName *string
			if viper.IsSet(flagDisplayName) {
				v := viper.GetString(flagDisplayName)
				displayName = &v
			}

			var description *string
			if viper.IsSet(flagDescription) {
				v := viper.GetString(flagDescription)
				description = &v
			}

			var lotSize *sdkmath.Int
			if viper.IsSet(flagLotSize) {
				v, ok := sdk.NewIntFromString(viper.GetString(flagLotSize))
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid lot size")
				}
				lotSize = &v
			}

			var tickSize *sdk.Dec
			if viper.IsSet(flagTicksize) {
				v, err := sdk.NewDecFromStr(viper.GetString(flagTicksize))
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid tick size")
				}
				tickSize = &v
			}

			var minQuantity *sdkmath.Int
			if viper.IsSet(flagMinQuantity) {
				v, ok := sdk.NewIntFromString(viper.GetString(flagMinQuantity))
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid min quantity")
				}
				minQuantity = &v
			}

			var riskStepSize *sdkmath.Int
			if viper.IsSet(flagRiskStepSize) {
				v, ok := sdk.NewIntFromString(viper.GetString(flagRiskStepSize))
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid risk step size")
				}
				riskStepSize = &v
			}

			var initialMarginBase *sdk.Dec
			if viper.IsSet(flagInitialMarginBase) {
				v, err := sdk.NewDecFromStr(viper.GetString(flagInitialMarginBase))
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid initial margin base")
				}
				initialMarginBase = &v
			}

			var initialMarginStep *sdk.Dec
			if viper.IsSet(flagInitialMarginStep) {
				v, err := sdk.NewDecFromStr(viper.GetString(flagInitialMarginStep))
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid initial margin step")
				}
				initialMarginStep = &v
			}

			var maintenanceMarginRatio *sdk.Dec
			if viper.IsSet(flagMaintenanceMarginRatio) {
				v, err := sdk.NewDecFromStr(viper.GetString(flagMaintenanceMarginRatio))
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid maintenance margin ratio")
				}
				maintenanceMarginRatio = &v
			}

			var maxLiquidationOrderTicket *sdkmath.Int
			if viper.IsSet(flagMaxLiquidationOrderTicket) {
				v, ok := sdk.NewIntFromString(viper.GetString(flagMaxLiquidationOrderTicket))
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid max liquidation order ticket")
				}
				maxLiquidationOrderTicket = &v
			}

			var maxLiquidationOrderDuration *time.Duration
			if viper.IsSet(flagMaxLiquidationOrderDuration) {
				v := viper.GetDuration(flagMaxLiquidationOrderDuration)
				maxLiquidationOrderDuration = &v
			}

			var impactSize *sdkmath.Int
			if viper.IsSet(flagImpactSize) {
				v, ok := sdk.NewIntFromString(viper.GetString(flagImpactSize))
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid impact size")
				}
				impactSize = &v
			}

			var markPriceBand *uint32
			if viper.IsSet(flagMarkPriceBand) {
				v := viper.GetUint32(flagMarkPriceBand)
				markPriceBand = &v
			}

			var lastPriceProtectedBand *uint32
			if viper.IsSet(flagLastPriceProtectedBand) {
				v := viper.GetUint32(flagLastPriceProtectedBand)
				lastPriceProtectedBand = &v
			}

			var tradingBandwidth *uint32
			if viper.IsSet(flagTradingBandwidth) {
				v := viper.GetUint32(flagTradingBandwidth)
				tradingBandwidth = &v
			}

			var isActive *bool
			if viper.IsSet(flagIsActive) {
				v := viper.GetBool(flagIsActive)
				isActive = &v
			}

			var expiryTime *time.Time

			if viper.IsSet(flagMarketExpiry) {
				v := viper.GetString(flagMarketExpiry)
				argsExpiryTime, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid expiry time: %s", err)
				}
				t := time.Unix(argsExpiryTime, 0)
				expiryTime = &t
			}

			msg := types.NewMsgUpdateMarket(
				cli.GetFromOrGroupAddress(clientCtx),
				types.MarketParams{
					Name:                        args[0],
					DisplayName:                 displayName,
					Description:                 description,
					LotSize:                     lotSize,
					TickSize:                    tickSize,
					MinQuantity:                 minQuantity,
					RiskStepSize:                riskStepSize,
					InitialMarginBase:           initialMarginBase,
					InitialMarginStep:           initialMarginStep,
					MaintenanceMarginRatio:      maintenanceMarginRatio,
					MaxLiquidationOrderTicket:   maxLiquidationOrderTicket,
					MaxLiquidationOrderDuration: maxLiquidationOrderDuration,
					ImpactSize:                  impactSize,
					MarkPriceBand:               markPriceBand,
					LastPriceProtectedBand:      lastPriceProtectedBand,
					TradingBandwidth:            tradingBandwidth,
					IsActive:                    isActive,
					ExpiryTime:                  expiryTime,
				},
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return cli.GenerateOrBroadcastTx(clientCtx, cmd, msg)
		},
	}

	cmd.Flags().String(flagDisplayName, "", "Market Display Name")
	cmd.Flags().String(flagDescription, "", "Market Description")
	cmd.Flags().String(flagLotSize, "", "Lot Size")
	cmd.Flags().String(flagTicksize, "", "Tick Size")
	cmd.Flags().String(flagMinQuantity, "", "Minimum Order Quantity")
	cmd.Flags().String(flagIsActive, "", "Is Active")
	cmd.Flags().String(flagRiskStepSize, "", "Risk Step Size")
	cmd.Flags().String(flagInitialMarginBase, "", "Initial Margin Base")
	cmd.Flags().String(flagInitialMarginStep, "", "Initial Margin Step")
	cmd.Flags().String(flagMaintenanceMarginRatio, "", "Maintenance Margin Ratio")
	cmd.Flags().String(flagMaxLiquidationOrderTicket, "", "Maximum Size for a Liquidation Order")
	cmd.Flags().Duration(flagMaxLiquidationOrderDuration, 0, "Maximum Duration for a Liquidation Order")
	cmd.Flags().String(flagImpactSize, "", "Impact Size")
	cmd.Flags().Uint32(flagMarkPriceBand, 0, "Mark Price Band")
	cmd.Flags().Uint32(flagLastPriceProtectedBand, 0, "Last Price Protected Band")
	cmd.Flags().Uint32(flagTradingBandwidth, 0, "Allowed Trading Bandwidth")
	cmd.Flags().String(flagMarketExpiry, "", "Market Expiry")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSetPerpetualsFundingInterval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-perp-funding-interval [interval]",
		Short: "Set perpetuals funding interval (interval is in seconds)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsInterval, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			argsIntervalInDuration := time.Duration(argsInterval) * time.Second

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.NewMsgUpdatePerpetualsFundingInterval(cli.GetFromOrGroupAddress(clientCtx), argsIntervalInDuration)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return cli.GenerateOrBroadcastTx(clientCtx, cmd, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDisableSpotMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable-spot-market [market]",
		Short: "Broadcast message DisableSpotMarket",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argsMarketName := args[0]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDisableSpotMarket(
				cli.GetFromOrGroupAddress(clientCtx),
				argsMarketName,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return cli.GenerateOrBroadcastTx(clientCtx, cmd, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
