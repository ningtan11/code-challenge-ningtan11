package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Switcheo/carbon/x/market/types"
)

func CmdGetTradingFees() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trading-fee [marketName] [userAddress]",
		Short: "Query trading fee for a given market and user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			marketName := args[0]
			userAddress := args[1]

			params := &types.QueryGetTradingFeesRequest{
				MarketName:  marketName,
				UserAddress: userAddress,
			}

			res, err := queryClient.TradingFees(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdFeeTiersMarketId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fee-tiers-id [marketId]",
		Short: "Query all fee tiers for a given market id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			marketName := args[0]

			var whitelistAddress string
			if viper.IsSet(flagWhitelistAddress) {
				v := viper.GetString(flagWhitelistAddress)
				whitelistAddress = v
			}

			params := &types.QueryGetFeeTiersRequest{
				MarketName:  marketName,
				UserAddress: whitelistAddress,
			}

			res, err := queryClient.FeeTiers(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagWhitelistAddress, "", "Whitelist Address")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdFeeTiersMarketType() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fee-tiers-type [marketType]",
		Short: "Query all fee tiers for a given market type",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			marketType := args[0]

			var whitelistAddress string
			if viper.IsSet(flagWhitelistAddress) {
				v := viper.GetString(flagWhitelistAddress)
				whitelistAddress = v
			}

			params := &types.QueryGetFeeTiersRequest{
				MarketType:  marketType,
				UserAddress: whitelistAddress,
			}

			res, err := queryClient.FeeTiers(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(flagWhitelistAddress, "", "Whitelist Address")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdStakeEquivalenceAll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake-eqs",
		Short: "Query for all stake equivalences that were set",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllStakeEquivalenceRequest{}

			res, err := queryClient.StakeEquivalenceAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdFeeStructureAll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fee-structures",
		Short: "Query for all fee structures",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllFeeStructuresRequest{}

			res, err := queryClient.FeeStructuresAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
