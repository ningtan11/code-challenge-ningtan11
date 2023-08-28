package cli

import (
	"strconv"
	"time"

	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/spf13/cobra"

	"github.com/Switcheo/carbon/x/market/types"
)

const (
	flagName = "name"
)

func NewCmdUpdateMarketProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-market [title] [description]",
		Args:  cobra.ExactArgs(2),
		Short: "Submit an update existing market proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			content := &types.UpdateMarketProposal{
				Title:       args[0],
				Description: args[1],
				Msg:         types.MarketParams{},
			}

			argsName, err := cmd.Flags().GetString(flagName)
			if err != nil {
				return err
			}
			if argsName != "" {
				content.Msg.Name = argsName
			}

			argsDisplayName, err := cmd.Flags().GetString(flagDisplayName)
			if err != nil {
				return err
			}
			if argsDisplayName != "" {
				content.Msg.DisplayName = &argsDisplayName
			}

			argsDescription, err := cmd.Flags().GetString(flagDescription)
			if err != nil {
				return err
			}
			if argsDescription != "" {
				content.Msg.Description = &argsDescription
			}

			argsMinQuantity, err := cmd.Flags().GetString(flagMinQuantity)
			if err != nil {
				return err
			}
			if argsMinQuantity != "" {
				m, ok := sdk.NewIntFromString(argsMinQuantity)
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid min quantity")
				}
				content.Msg.MinQuantity = &m
			}

			argsIsActive, err := cmd.Flags().GetString(flagIsActive)
			if err != nil {
				return err
			}
			if argsIsActive != "" {
				m, err := strconv.ParseBool(argsIsActive)
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid is active")
				}
				content.Msg.IsActive = &m
			}

			argsRiskStepSize, err := cmd.Flags().GetString(flagRiskStepSize)
			if err != nil {
				return err
			}
			if argsRiskStepSize != "" {
				m, ok := sdk.NewIntFromString(argsRiskStepSize)
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid risk step size")
				}
				content.Msg.RiskStepSize = &m
			}

			argsInitialMarginBase, err := cmd.Flags().GetString(flagInitialMarginBase)
			if err != nil {
				return err
			}
			if argsInitialMarginBase != "" {
				m, err := sdk.NewDecFromStr(argsInitialMarginBase)
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid initial margin base")
				}
				content.Msg.InitialMarginBase = &m
			}

			argsInitialMarginStep, err := cmd.Flags().GetString(flagInitialMarginStep)
			if err != nil {
				return err
			}
			if argsInitialMarginStep != "" {
				m, err := sdk.NewDecFromStr(argsInitialMarginStep)
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid initial margin step")
				}
				content.Msg.InitialMarginStep = &m
			}

			argsMaintenanceMarginRatio, err := cmd.Flags().GetString(flagMaintenanceMarginRatio)
			if err != nil {
				return err
			}
			if argsMaintenanceMarginRatio != "" {
				m, err := sdk.NewDecFromStr(argsMaintenanceMarginRatio)
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid maintenance margin ratio")
				}
				content.Msg.MaintenanceMarginRatio = &m
			}

			argsMaxLiquidationOrderTicket, err := cmd.Flags().GetString(flagMaxLiquidationOrderTicket)
			if err != nil {
				return err
			}
			if argsMaxLiquidationOrderTicket != "" {
				m, ok := sdk.NewIntFromString(argsMaxLiquidationOrderTicket)
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid maxLiquidation order ticket")
				}
				content.Msg.MaxLiquidationOrderTicket = &m
			}

			argsMaxLiquidationOrderDuration, err := cmd.Flags().GetString(flagMaxLiquidationOrderDuration)
			if err != nil {
				return err
			}
			if argsMaxLiquidationOrderDuration != "" {
				v, err := strconv.ParseInt(argsMaxLiquidationOrderDuration, 10, 64)
				if err != nil {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "max liquidation order duration")
				}
				m := time.Duration(v)
				content.Msg.MaxLiquidationOrderDuration = &m
			}

			argsImpactSize, err := cmd.Flags().GetString(flagImpactSize)
			if err != nil {
				return err
			}
			if argsImpactSize != "" {
				m, ok := sdk.NewIntFromString(argsImpactSize)
				if !ok {
					return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid impact size")
				}
				content.Msg.ImpactSize = &m
			}

			argsMarkPriceBand, err := cmd.Flags().GetString(flagMarkPriceBand)
			if err != nil {
				return err
			}
			if argsMarkPriceBand != "" {
				v, err := strconv.ParseUint(args[16], 10, 32)
				if err != nil {
					return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid mark price band: %s", err.Error())
				}
				m := uint32(v)
				content.Msg.MarkPriceBand = &m
			}

			argsLastPriceProtectedBand, err := cmd.Flags().GetString(flagLastPriceProtectedBand)
			if err != nil {
				return err
			}
			if argsLastPriceProtectedBand != "" {
				v, err := strconv.ParseUint(argsLastPriceProtectedBand, 10, 32)
				if err != nil {
					return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid price protected band: %s", err.Error())
				}
				m := uint32(v)
				content.Msg.MarkPriceBand = &m
			}

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := gov.NewMsgSubmitProposal(content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			if err = content.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	cmd.Flags().String(flagName, "", "identifier for market")
	cmd.Flags().String(flagDisplayName, "", "display name of market")
	cmd.Flags().String(flagDescription, "", "description of market")
	cmd.Flags().String(flagMinQuantity, "", "minimum order quantity for market")
	cmd.Flags().String(flagMakerFee, "", "maker order fee of market")
	cmd.Flags().String(flagTakerFee, "", "taker order fee of market")
	cmd.Flags().String(flagIsActive, "", "active status of market")
	cmd.Flags().String(flagRiskStepSize, "", "risk step size for market")
	cmd.Flags().String(flagInitialMarginBase, "", "initial margin base for market")
	cmd.Flags().String(flagInitialMarginStep, "", "initial margin step for market")
	cmd.Flags().String(flagMaintenanceMarginRatio, "", "maintenance margin ratio for market")
	cmd.Flags().String(flagMaxLiquidationOrderTicket, "", "max liquidation order size for market")
	cmd.Flags().String(flagMaxLiquidationOrderDuration, "", "max liquidation order duration for market")
	cmd.Flags().String(flagImpactSize, "", "impact size for market")
	cmd.Flags().String(flagMarkPriceBand, "", "mark price band for market")
	cmd.Flags().String(flagLastPriceProtectedBand, "", "last price protected band for market")

	return cmd
}

func NewCmdUpdatePerpetualsFundingIntervalProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-perpetuals-funding-interval [title] [description] [perpetuals-funding-interval]",
		Args:  cobra.ExactArgs(3),
		Short: "Submit an update existing market proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argsTitle := args[0]
			argsDescription := args[1]
			argsPerpetualsFundingIntervalSeconds, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "perpetual funding interval")
			}
			argsPerpetualsFundingInterval := time.Duration(argsPerpetualsFundingIntervalSeconds) * time.Second

			content := types.NewUpdatePerpetualsFundingIntervalProposal(argsTitle, argsDescription, argsPerpetualsFundingInterval)

			depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(depositStr)
			if err != nil {
				return err
			}

			msg, err := gov.NewMsgSubmitProposal(content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			if err = content.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")

	return cmd
}
