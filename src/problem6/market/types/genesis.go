package types

import (
	"time"

	"github.com/Switcheo/carbon/constants"
)

var ShortenPerpetualsFundingInterval string

// NewGenesisState creates a new GenesisState object
func NewGenesisState(markets []Market, params Params, controlledParams ControlledParams, stakeEq []StakeEquivalence, feeStructures []FeeStructure) *GenesisState {
	return &GenesisState{
		Params:            params,
		Markets:           markets,
		ControlledParams:  controlledParams,
		StakeEquivalences: stakeEq,
		FeeStructures:     feeStructures,
	}
}

func DefaultControlledParams() ControlledParams {
	perpetualsFundingInterval := 60 * time.Minute
	if ShortenPerpetualsFundingInterval == constants.True {
		perpetualsFundingInterval = 1 * time.Minute
	}
	return ControlledParams{
		PerpetualsFundingInterval: perpetualsFundingInterval,
	}
}

// DefaultGenesisState defines the default market genesis state
func DefaultGenesisState() *GenesisState {
	if constants.UseTestGenesis != constants.True {
		return &GenesisState{
			Params:             DefaultParams(),
			ControlledParams:   DefaultControlledParams(),
			MarketNameSequence: 1,
			Markets:            []Market{},
			StakeEquivalences:  []StakeEquivalence{},
			FeeStructures:      []FeeStructure{},
		}
	}

	return &GenesisState{
		Params:             DefaultParams(),
		ControlledParams:   DefaultControlledParams(),
		MarketNameSequence: 1,
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		Markets:           GenesisMarkets,
		StakeEquivalences: GenesisStakeEq,
		FeeStructures:     GenesisFeeStructures,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	for _, m := range gs.Markets {
		if err := m.ValidateBasic(); err != nil {
			return err
		}
	}
	for _, fs := range gs.FeeStructures {
		if err := FeeTiers(fs.FeeTiers).ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}
