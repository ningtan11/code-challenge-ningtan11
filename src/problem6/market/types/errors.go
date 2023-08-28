package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/market module sentinel errors
var (
	ErrInvalidMarket             = errorsmod.Register(ModuleName, 2, "invalid market")
	ErrMarketNotFound            = errorsmod.Register(ModuleName, 3, "market not found")
	ErrDuplicateName             = errorsmod.Register(ModuleName, 4, "duplicate market name")
	ErrDuplicateDisplayName      = errorsmod.Register(ModuleName, 5, "duplicate market display name")
	ErrInvalidExpiry             = errorsmod.Register(ModuleName, 6, "market expiry time must be at least 1 hr in the future")
	ErrInvalidCandlesticksLength = errorsmod.Register(ModuleName, 7, "invalid candlesticks length")
	ErrInvalidExpiryDay          = errorsmod.Register(ModuleName, 8, "market expiry time must be on a friday 8am utc")
	ErrMaxActiveMarketsExceeded  = errorsmod.Register(ModuleName, 9, "maximum number of active markets exceeded")
	ErrInvalidLotSize            = errorsmod.Register(ModuleName, 10, "lot size should not be less than 1")
	ErrInvalidTickSize           = errorsmod.Register(ModuleName, 11, "tick size should not be less than 1")
	ErrSameBaseAndQuoteToken     = errorsmod.Register(ModuleName, 12, "quote and base token can't be the same")
	ErrMarketNotUnique           = errorsmod.Register(ModuleName, 13, "market already exists for this base/quote denom, marketType and expiryTime")
	ErrHasOutstandingPositions   = errorsmod.Register(ModuleName, 14, "tick size cannot be changed when outstanding positions exist")
	ErrInvalidFeeTier            = errorsmod.Register(ModuleName, 15, "fee tier is not valid")

	// this line is used by starport scaffolding # ibc/errors
)
