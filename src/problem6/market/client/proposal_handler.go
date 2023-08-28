package client

import (
	"github.com/Switcheo/carbon/x/market/client/cli"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

// ProposalHandlers are the proposal handlers for this module.
var ProposalHandlers = []govclient.ProposalHandler{
	govclient.NewProposalHandler(cli.NewCmdUpdateMarketProposal),
	govclient.NewProposalHandler(cli.NewCmdUpdateMarketProposal),
}
