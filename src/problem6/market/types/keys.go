package types

const (
	// ModuleName defines the module name
	ModuleName = "market"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// StoreKey defines the transient store key
	TStoreKey = "transient_" + ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_market"

	// this line is used by starport scaffolding # ibc/keys/name
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	MarketKey           = "Market-value-"
	ControlledParamsKey = "ControlledParams-value-"
	FeeTierKey          = "FeeTier-value-"
	StakeEquivalenceKey = "StakeEquivalence-value-"

	// Transient store
	IncomingDisableSpotMarketTKey = "IncomingDisableSpotMarket-value-"
	StakeEquivalenceTKey          = "StakeEquivalence-value-"
)

var (
	// Market Store
	MarketNameSequenceSuffixKey = []byte{0x00}

	// ControlledParams Store
	PerpetualsFundingIntervalKey = []byte{0x00}
)
