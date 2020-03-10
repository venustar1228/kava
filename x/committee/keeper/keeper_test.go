package keeper_test

import (
	"testing"

	"github.com/kava-labs/kava/x/committee/types"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/kava-labs/kava/app"
	"github.com/kava-labs/kava/x/committee/keeper"
)

type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper
	app    app.TestApp
	ctx    sdk.Context

	addresses []sdk.AccAddress
}

func (suite *KeeperTestSuite) SetupTest() {

	suite.app = app.NewTestApp()
	suite.keeper = suite.app.GetCommitteeKeeper()
	suite.ctx = suite.app.NewContext(true, abci.Header{})
	_, suite.addresses = app.GeneratePrivKeyAddressPairs(2)
}

func (suite *KeeperTestSuite) TestGetSetCommittee() {
	// test setup
	com := types.Committee{
		ID: 12,
		// TODO other fields
	}

	// write and read from store
	suite.keeper.SetCommittee(suite.ctx, com)
	readCommittee, found := suite.keeper.GetCommittee(suite.ctx, com.ID)

	// check before and after match
	suite.True(found)
	suite.Equal(com, readCommittee)
}

func (suite *KeeperTestSuite) TestGetSetProposal() {
	// test setup
	prop := types.Proposal{
		ID: 12,
		// TODO other fields
	}

	// write and read from store
	suite.keeper.SetProposal(suite.ctx, prop)
	readProposal, found := suite.keeper.GetProposal(suite.ctx, prop.ID)

	// check before and after match
	suite.True(found)
	suite.Equal(prop, readProposal)
}

func (suite *KeeperTestSuite) TestGetSetVote() {
	// test setup
	vote := types.Vote{
		ProposalID: 12,
		Voter:      suite.addresses[0],
		// TODO other fields
	}

	// write and read from store
	suite.keeper.SetVote(suite.ctx, vote)
	readVote, found := suite.keeper.GetVote(suite.ctx, vote.ProposalID, vote.Voter)

	// check before and after match
	suite.True(found)
	suite.Equal(vote, readVote)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
