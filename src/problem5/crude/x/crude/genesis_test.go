package crude_test

import (
	"testing"

	keepertest "github.com/crude/testutil/keeper"
	"github.com/crude/testutil/nullify"
	"github.com/crude/x/crude"
	"github.com/crude/x/crude/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CrudeKeeper(t)
	crude.InitGenesis(ctx, *k, genesisState)
	got := crude.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
