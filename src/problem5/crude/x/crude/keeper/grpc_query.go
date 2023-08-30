package keeper

import (
	"github.com/crude/x/crude/types"
)

var _ types.QueryServer = Keeper{}
