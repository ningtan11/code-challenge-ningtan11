package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/Switcheo/carbon/x/market/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group market queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdListMarket())
	cmd.AddCommand(CmdShowMarket())
	cmd.AddCommand(CmdGetTradingFees())
	cmd.AddCommand(CmdFeeTiersMarketId())
	cmd.AddCommand(CmdFeeTiersMarketType())
	cmd.AddCommand(CmdStakeEquivalenceAll())
	cmd.AddCommand(CmdFeeStructureAll())

	return cmd
}
