// nolint
// DO NOT EDIT - generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen)
package committee

import (
	"github.com/kava-labs/kava/x/committee/client"
	"github.com/kava-labs/kava/x/committee/keeper"
	"github.com/kava-labs/kava/x/committee/types"
)

const (
	AttributeKeyCommitteeID         = types.AttributeKeyCommitteeID
	AttributeKeyProposalCloseStatus = types.AttributeKeyProposalCloseStatus
	AttributeKeyProposalID          = types.AttributeKeyProposalID
	AttributeValueCategory          = types.AttributeValueCategory
	AttributeValueProposalFailed    = types.AttributeValueProposalFailed
	AttributeValueProposalPassed    = types.AttributeValueProposalPassed
	AttributeValueProposalTimeout   = types.AttributeValueProposalTimeout
	DefaultNextProposalID           = types.DefaultNextProposalID
	DefaultParamspace               = types.DefaultParamspace
	EventTypeProposalClose          = types.EventTypeProposalClose
	EventTypeProposalSubmit         = types.EventTypeProposalSubmit
	EventTypeProposalVote           = types.EventTypeProposalVote
	MaxCommitteeDescriptionLength   = types.MaxCommitteeDescriptionLength
	ModuleName                      = types.ModuleName
	ProposalTypeCommitteeChange     = types.ProposalTypeCommitteeChange
	ProposalTypeCommitteeDelete     = types.ProposalTypeCommitteeDelete
	QuerierRoute                    = types.QuerierRoute
	QueryCommittee                  = types.QueryCommittee
	QueryCommittees                 = types.QueryCommittees
	QueryProposal                   = types.QueryProposal
	QueryProposals                  = types.QueryProposals
	QueryTally                      = types.QueryTally
	QueryVote                       = types.QueryVote
	QueryVotes                      = types.QueryVotes
	RouterKey                       = types.RouterKey
	StoreKey                        = types.StoreKey
	TypeMsgSubmitProposal           = types.TypeMsgSubmitProposal
	TypeMsgVote                     = types.TypeMsgVote
)

var (
	// function aliases
	NewKeeper                   = keeper.NewKeeper
	NewQuerier                  = keeper.NewQuerier
	DefaultGenesisState         = types.DefaultGenesisState
	GetKeyFromID                = types.GetKeyFromID
	GetVoteKey                  = types.GetVoteKey
	NewCommittee                = types.NewCommittee
	NewCommitteeChangeProposal  = types.NewCommitteeChangeProposal
	NewCommitteeDeleteProposal  = types.NewCommitteeDeleteProposal
	NewGenesisState             = types.NewGenesisState
	NewMsgSubmitProposal        = types.NewMsgSubmitProposal
	NewMsgVote                  = types.NewMsgVote
	NewProposal                 = types.NewProposal
	NewQueryCommitteeParams     = types.NewQueryCommitteeParams
	NewQueryProposalParams      = types.NewQueryProposalParams
	NewQueryVoteParams          = types.NewQueryVoteParams
	RegisterCodec               = types.RegisterCodec
	RegisterPermissionTypeCodec = types.RegisterPermissionTypeCodec
	RegisterProposalTypeCodec   = types.RegisterProposalTypeCodec
	Uint64FromBytes             = types.Uint64FromBytes

	// variable aliases
	ProposalHandler            = client.ProposalHandler
	CommitteeKeyPrefix         = types.CommitteeKeyPrefix
	ErrInvalidCommittee        = types.ErrInvalidCommittee
	ErrInvalidGenesis          = types.ErrInvalidGenesis
	ErrInvalidPubProposal      = types.ErrInvalidPubProposal
	ErrNoProposalHandlerExists = types.ErrNoProposalHandlerExists
	ErrProposalExpired         = types.ErrProposalExpired
	ErrUnknownCommittee        = types.ErrUnknownCommittee
	ErrUnknownProposal         = types.ErrUnknownProposal
	ErrUnknownVote             = types.ErrUnknownVote
	ModuleCdc                  = types.ModuleCdc
	NextProposalIDKey          = types.NextProposalIDKey
	ProposalKeyPrefix          = types.ProposalKeyPrefix
	VoteKeyPrefix              = types.VoteKeyPrefix
)

type (
	Keeper                  = keeper.Keeper
	AllowedParam            = types.AllowedParam
	AllowedParams           = types.AllowedParams
	Committee               = types.Committee
	CommitteeChangeProposal = types.CommitteeChangeProposal
	CommitteeDeleteProposal = types.CommitteeDeleteProposal
	GenesisState            = types.GenesisState
	GodPermission           = types.GodPermission
	MsgSubmitProposal       = types.MsgSubmitProposal
	MsgVote                 = types.MsgVote
	ParamChangePermission   = types.ParamChangePermission
	Permission              = types.Permission
	Proposal                = types.Proposal
	PubProposal             = types.PubProposal
	QueryCommitteeParams    = types.QueryCommitteeParams
	QueryProposalParams     = types.QueryProposalParams
	QueryVoteParams         = types.QueryVoteParams
	TextPermission          = types.TextPermission
	Vote                    = types.Vote
)
