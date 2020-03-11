// nolint
// DO NOT EDIT - generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen)
package committee

import (
	"github.com/kava-labs/kava/x/committee/keeper"
	"github.com/kava-labs/kava/x/committee/types"
)

const (
	DefaultNextProposalID = types.DefaultNextProposalID
	ModuleName            = types.ModuleName
	StoreKey              = types.StoreKey
)

var (
	// function aliases
	NewKeeper           = keeper.NewKeeper
	DefaultGenesisState = types.DefaultGenesisState
	GetKeyFromID        = types.GetKeyFromID
	GetVoteKey          = types.GetVoteKey
	NewGenesisState     = types.NewGenesisState
	RegisterCodec       = types.RegisterCodec
	Uint64FromBytes     = types.Uint64FromBytes

	// variable aliases
	CommitteeKeyPrefix = types.CommitteeKeyPrefix
	ModuleCdc          = types.ModuleCdc
	NextProposalIDKey  = types.NextProposalIDKey
	ProposalKeyPrefix  = types.ProposalKeyPrefix
	VoteKeyPrefix      = types.VoteKeyPrefix
	VoteThreshold      = types.VoteThreshold
)

type (
	Keeper                        = keeper.Keeper
	Committee                     = types.Committee
	GeneralShutdownPermission     = types.GeneralShutdownPermission
	GenesisState                  = types.GenesisState
	GodPermission                 = types.GodPermission
	GroupChangeProposal           = types.GroupChangeProposal
	InflationRateChangePermission = types.InflationRateChangePermission
	MsgSubmitProposal             = types.MsgSubmitProposal
	MsgVote                       = types.MsgVote
	Permission                    = types.Permission
	Proposal                      = types.Proposal
	PubProposal                   = types.PubProposal
	ShutdownCDPDepsitPermission   = types.ShutdownCDPDepsitPermission
	Vote                          = types.Vote
)
