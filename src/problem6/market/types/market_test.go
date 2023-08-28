package types_test

import (
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/Switcheo/carbon/x/market/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ = Describe("Market", func() {
	var (
		addr = sdk.AccAddress(strings.Repeat("1", 20)).String()
	)

	Describe("ValidateBasic", func() {
		It("should pass validation when all attributes are valid", func() {
			msg := types.NewMsgCreateMarket(
				addr,   // creator
				"spot", // market type
				"eth",  // base
				"btc",  // quote
				sdk.NewDec(2500),
				sdk.NewDec(2500),
				"",              // indexOracleID,
				time.Unix(0, 0), // expiryTime
			)
			err := msg.ValidateBasic()
			Expect(err).ToNot(BeNil())
		})

		It("should pass name validation with .", func() {
			msg := types.NewMsgCreateMarket(
				addr,   // creator
				"spot", // market type
				"eth",  // base
				"btc",  // quote
				sdk.NewDec(2500),
				sdk.NewDec(2500),
				"",              // indexOracleID,
				time.Unix(0, 0), // expiryTime
			)
			err := msg.ValidateBasic()
			Expect(err).To(BeNil())
		})

		It("should fail validation when name is too long", func() {
			msg := types.NewMsgCreateMarket(
				addr,   // creator
				"spot", // market type
				"eth",  // base
				"btc",  // quote
				sdk.NewDec(2500),
				sdk.NewDec(2500),
				"",              // indexOracleID,
				time.Unix(0, 0), // expiryTime
			)
			err := msg.ValidateBasic()
			Expect(err).ToNot(BeNil())
			Expect(types.ErrInvalidMarket.Is(err)).To(BeTrue())
		})

		It("should fail validaiton when display name is too long", func() {
			msg := types.NewMsgCreateMarket(
				addr,   // creator
				"spot", // market type
				"eth",  // base
				"btc",  // quote
				sdk.NewDec(2500),
				sdk.NewDec(2500),
				"",              // indexOracleID,
				time.Unix(0, 0), // expiryTime
			)
			err := msg.ValidateBasic()
			Expect(types.ErrInvalidMarket.Is(err)).To(BeTrue())
		})

		It("should fail a-zA-Z_. display name validation when there is a * symbol", func() {
			msg := types.NewMsgCreateMarket(
				addr,   // creator
				"spot", // market type
				"eth",  // base
				"btc",  // quote
				sdk.NewDec(2500),
				sdk.NewDec(2500),
				"",              // indexOracleID,
				time.Unix(0, 0), // expiryTime
			)
			err := msg.ValidateBasic()
			Expect(types.ErrInvalidMarket.Is(err)).To(BeTrue())
		})

		It("should fail a-zA-Z_. display name validation there is a space", func() {
			msg := types.NewMsgCreateMarket(
				addr,   // creator
				"spot", // market type
				"eth",  // base
				"btc",  // quote
				sdk.NewDec(2500),
				sdk.NewDec(2500),
				"",              // indexOracleID,
				time.Unix(0, 0), // expiryTime
			)
			err := msg.ValidateBasic()
			Expect(err).To(BeNil())
		})
	})
})
