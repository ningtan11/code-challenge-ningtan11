package cli

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Switcheo/carbon/x/admin/client/cli"
	"github.com/Switcheo/carbon/x/market/types"
)

var (
	flagMarketType       = "market-type"
	flagMarketId         = "market-id"
	flagReqStake         = "req-stake"
	flagWhitelistAddress = "whitelist-address"
	flagMakerFee         = "maker-fee"
	flagTakerFee         = "taker-fee"
)

func CmdAddFeeTier() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-fee-tier",
		Short: "Adds a fee tier to a market by id or type",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var marketType string
			if viper.IsSet(flagMarketType) {
				v := viper.GetString(flagMarketType)
				marketType = v
			}
			var marketId string
			if viper.IsSet(flagMarketId) {
				v := viper.GetString(flagMarketId)
				marketId = v
			}
			var whitelistAddress string
			if viper.IsSet(flagWhitelistAddress) {
				v := viper.GetString(flagWhitelistAddress)
				whitelistAddress = v
			}
			feeCategory := types.FeeCategory{
				MarketId:           marketId,
				MarketType:         marketType,
				WhitelistedAddress: whitelistAddress,
			}

			var reqStake sdkmath.Int
			if viper.IsSet(flagReqStake) {
				v, ok := sdk.NewIntFromString(viper.GetString(flagReqStake))
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid stake qty")
				}
				reqStake = v
			}
			var makerFee sdk.Dec
			if viper.IsSet(flagMakerFee) {
				v, err := sdk.NewDecFromStr(viper.GetString(flagMakerFee))
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid maker fee")
				}
				makerFee = v
			}
			var takerFee sdk.Dec
			if viper.IsSet(flagTakerFee) {
				v, err := sdk.NewDecFromStr(viper.GetString(flagTakerFee))
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid taker fee")
				}
				if v.IsNegative() {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "taker fee cannot be negative")
				}
				takerFee = v
			}
			feeTier := types.FeeTier{
				RequiredStake: reqStake,
				TradingFees: types.TradingFees{
					MakerFee: makerFee,
					TakerFee: takerFee,
				},
			}

			msg := types.NewMsgAddFeeTier(
				cli.GetFromOrGroupAddress(clientCtx),
				feeCategory,
				feeTier,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return cli.GenerateOrBroadcastTx(clientCtx, cmd, msg)
		},
	}
	cmd.Flags().String(flagMakerFee, "", "Maker Fee")
	cmd.Flags().String(flagTakerFee, "", "Taker Fee")
	cmd.Flags().String(flagMarketId, "", "Market Id")
	cmd.Flags().String(flagWhitelistAddress, "", "Whitelist Address")
	cmd.Flags().String(flagMarketType, "", "Market Type")
	cmd.Flags().String(flagReqStake, "", "Required Stake")
	if err := cmd.MarkFlagRequired(flagTakerFee); err != nil {
		panic(err)
	}
	if err := cmd.MarkFlagRequired(flagMakerFee); err != nil {
		panic(err)
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateFeeTier() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-fee-tier",
		Short: "Updates a fee tier to a market by id or type",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var marketType string
			if viper.IsSet(flagMarketType) {
				v := viper.GetString(flagMarketType)
				marketType = v
			}
			var marketId string
			if viper.IsSet(flagMarketId) {
				v := viper.GetString(flagMarketId)
				marketId = v
			}
			var whitelistAddress string
			if viper.IsSet(flagWhitelistAddress) {
				v := viper.GetString(flagWhitelistAddress)
				whitelistAddress = v
			}
			feeCategory := types.FeeCategory{
				MarketId:           marketId,
				MarketType:         marketType,
				WhitelistedAddress: whitelistAddress,
			}

			var reqStake sdkmath.Int
			if viper.IsSet(flagReqStake) {
				v, ok := sdk.NewIntFromString(viper.GetString(flagReqStake))
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid stake qty")
				}
				reqStake = v
			}
			var makerFee *sdk.Dec
			if viper.IsSet(flagMakerFee) {
				v, err := sdk.NewDecFromStr(viper.GetString(flagMakerFee))
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid maker fee")
				}
				makerFee = &v
			}
			var takerFee *sdk.Dec
			if viper.IsSet(flagTakerFee) {
				v, err := sdk.NewDecFromStr(viper.GetString(flagTakerFee))
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid taker fee")
				}
				if v.IsNegative() {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "taker fee cannot be negative")
				}
				takerFee = &v
			}

			msg := types.NewMsgUpdateFeeTier(
				cli.GetFromOrGroupAddress(clientCtx),
				feeCategory,
				reqStake,
				makerFee,
				takerFee,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return cli.GenerateOrBroadcastTx(clientCtx, cmd, msg)
		},
	}
	cmd.Flags().String(flagReqStake, "", "Required Stake")
	cmd.Flags().String(flagMakerFee, "", "Maker Fee")
	cmd.Flags().String(flagTakerFee, "", "Taker Fee")
	cmd.Flags().String(flagMarketId, "", "Market Id")
	cmd.Flags().String(flagMarketType, "", "Market Type")
	cmd.Flags().String(flagWhitelistAddress, "", "Whitelist Address")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveFeeTier() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-fee-tier",
		Short: "Removes a fee tier from a market by id or type",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var marketType string
			if viper.IsSet(flagMarketType) {
				v := viper.GetString(flagMarketType)
				marketType = v
			}
			var marketId string
			if viper.IsSet(flagMarketId) {
				v := viper.GetString(flagMarketId)
				marketId = v
			}
			var whitelistAddress string
			if viper.IsSet(flagWhitelistAddress) {
				v := viper.GetString(flagWhitelistAddress)
				whitelistAddress = v
			}
			feeCategory := types.FeeCategory{
				MarketId:           marketId,
				MarketType:         marketType,
				WhitelistedAddress: whitelistAddress,
			}

			var reqStake sdkmath.Int
			if viper.IsSet(flagReqStake) {
				v, ok := sdk.NewIntFromString(viper.GetString(flagReqStake))
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid stake qty")
				}
				reqStake = v
			}
			msg := types.NewMsgRemoveFeeTier(cli.GetFromOrGroupAddress(clientCtx), feeCategory, reqStake)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return cli.GenerateOrBroadcastTx(clientCtx, cmd, msg)
		},
	}
	cmd.Flags().String(flagMarketId, "", "Market Id")
	cmd.Flags().String(flagMarketType, "", "Market Type")
	cmd.Flags().String(flagWhitelistAddress, "", "Whitelist Address")
	cmd.Flags().String(flagReqStake, "", "Required Stake")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSetStakeEquivalence() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-stake-eq [denom] [ratio]",
		Short: "Sets a stake equivalence for a denom and ratio",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			denom := args[0]
			ratio, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid ratio qty")
			}

			msg := types.NewMsgSetStakeEquivalence(cli.GetFromOrGroupAddress(clientCtx), denom, ratio)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return cli.GenerateOrBroadcastTx(clientCtx, cmd, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateAllPoolTradingFees() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update-all-pool-trading-fees [market-type]",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			marketType := args[0]

			var makerFee *sdk.Dec
			if viper.IsSet(flagMakerFee) {
				v, err := sdk.NewDecFromStr(viper.GetString(flagMakerFee))
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid maker fee")
				}
				makerFee = &v
			}

			var takerFee *sdk.Dec
			if viper.IsSet(flagTakerFee) {
				v, err := sdk.NewDecFromStr(viper.GetString(flagTakerFee))
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid taker fee")
				}
				takerFee = &v
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAllPoolTradingFees(cli.GetFromOrGroupAddress(clientCtx), marketType, *makerFee, *takerFee)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return cli.GenerateOrBroadcastTx(clientCtx, cmd, msg)
		},
	}

	cmd.Flags().String(flagMakerFee, "", "Maker Fee")
	cmd.Flags().String(flagTakerFee, "", "Taker Fee")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
