// Code generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen) DO NOT EDIT.

package swap

import (
	"github.com/kava-labs/kava/x/swap/keeper"
	"github.com/kava-labs/kava/x/swap/types"
)

const (
	AttributeKeyDepositor      = types.AttributeKeyDepositor
	AttributeKeyExactDirection = types.AttributeKeyExactDirection
	AttributeKeyFeePaid        = types.AttributeKeyFeePaid
	AttributeKeyOwner          = types.AttributeKeyOwner
	AttributeKeyPoolID         = types.AttributeKeyPoolID
	AttributeKeyRequester      = types.AttributeKeyRequester
	AttributeKeyShares         = types.AttributeKeyShares
	AttributeKeySwapInput      = types.AttributeKeySwapInput
	AttributeKeySwapOutput     = types.AttributeKeySwapOutput
	AttributeValueCategory     = types.AttributeValueCategory
	DefaultParamspace          = types.DefaultParamspace
	EventTypeSwapDeposit       = types.EventTypeSwapDeposit
	EventTypeSwapTrade         = types.EventTypeSwapTrade
	EventTypeSwapWithdraw      = types.EventTypeSwapWithdraw
	ModuleAccountName          = types.ModuleAccountName
	ModuleName                 = types.ModuleName
	PoolIDSep                  = types.PoolIDSep
	QuerierRoute               = types.QuerierRoute
	QueryGetDeposits           = types.QueryGetDeposits
	QueryGetParams             = types.QueryGetParams
	QueryGetPool               = types.QueryGetPool
	QueryGetPools              = types.QueryGetPools
	RouterKey                  = types.RouterKey
	StoreKey                   = types.StoreKey
)

var (
	// function aliases
	AllInvariants                        = keeper.AllInvariants
	NewKeeper                            = keeper.NewKeeper
	NewQuerier                           = keeper.NewQuerier
	PoolRecordsInvariant                 = keeper.PoolRecordsInvariant
	PoolReservesInvariant                = keeper.PoolReservesInvariant
	PoolSharesInvariant                  = keeper.PoolSharesInvariant
	RegisterInvariants                   = keeper.RegisterInvariants
	ShareRecordsInvariant                = keeper.ShareRecordsInvariant
	DefaultGenesisState                  = types.DefaultGenesisState
	DefaultParams                        = types.DefaultParams
	DepositorPoolSharesKey               = types.DepositorPoolSharesKey
	NewAllowedPool                       = types.NewAllowedPool
	NewAllowedPools                      = types.NewAllowedPools
	NewBasePool                          = types.NewBasePool
	NewBasePoolWithExistingShares        = types.NewBasePoolWithExistingShares
	NewDenominatedPool                   = types.NewDenominatedPool
	NewDenominatedPoolWithExistingShares = types.NewDenominatedPoolWithExistingShares
	NewDepositsQueryResult               = types.NewDepositsQueryResult
	NewGenesisState                      = types.NewGenesisState
	NewMsgDeposit                        = types.NewMsgDeposit
	NewMsgSwapExactForTokens             = types.NewMsgSwapExactForTokens
	NewMsgSwapForExactTokens             = types.NewMsgSwapForExactTokens
	NewMsgWithdraw                       = types.NewMsgWithdraw
	NewParams                            = types.NewParams
	NewPoolRecord                        = types.NewPoolRecord
	NewPoolRecordFromPool                = types.NewPoolRecordFromPool
	NewPoolStatsQueryResult              = types.NewPoolStatsQueryResult
	NewQueryDepositsParams               = types.NewQueryDepositsParams
	NewQueryPoolParams                   = types.NewQueryPoolParams
	NewShareRecord                       = types.NewShareRecord
	ParamKeyTable                        = types.ParamKeyTable
	PoolID                               = types.PoolID
	PoolIDFromCoins                      = types.PoolIDFromCoins
	PoolKey                              = types.PoolKey
	RegisterCodec                        = types.RegisterCodec

	// variable aliases
	DefaultAllowedPools       = types.DefaultAllowedPools
	DefaultPoolRecords        = types.DefaultPoolRecords
	DefaultShareRecords       = types.DefaultShareRecords
	DefaultSwapFee            = types.DefaultSwapFee
	DepositorPoolSharesPrefix = types.DepositorPoolSharesPrefix
	ErrDeadlineExceeded       = types.ErrDeadlineExceeded
	ErrDepositNotFound        = types.ErrDepositNotFound
	ErrInsufficientLiquidity  = types.ErrInsufficientLiquidity
	ErrInvalidCoin            = types.ErrInvalidCoin
	ErrInvalidDeadline        = types.ErrInvalidDeadline
	ErrInvalidPool            = types.ErrInvalidPool
	ErrInvalidShares          = types.ErrInvalidShares
	ErrInvalidSlippage        = types.ErrInvalidSlippage
	ErrNotAllowed             = types.ErrNotAllowed
	ErrNotImplemented         = types.ErrNotImplemented
	ErrSlippageExceeded       = types.ErrSlippageExceeded
	KeyAllowedPools           = types.KeyAllowedPools
	KeySwapFee                = types.KeySwapFee
	MaxSwapFee                = types.MaxSwapFee
	ModuleCdc                 = types.ModuleCdc
	PoolKeyPrefix             = types.PoolKeyPrefix
)

type (
	Keeper                = keeper.Keeper
	AccountKeeper         = types.AccountKeeper
	AllowedPool           = types.AllowedPool
	AllowedPools          = types.AllowedPools
	BasePool              = types.BasePool
	DenominatedPool       = types.DenominatedPool
	DepositsQueryResult   = types.DepositsQueryResult
	DepositsQueryResults  = types.DepositsQueryResults
	GenesisState          = types.GenesisState
	MsgDeposit            = types.MsgDeposit
	MsgSwapExactForTokens = types.MsgSwapExactForTokens
	MsgSwapForExactTokens = types.MsgSwapForExactTokens
	MsgWithDeadline       = types.MsgWithDeadline
	MsgWithdraw           = types.MsgWithdraw
	Params                = types.Params
	PoolRecord            = types.PoolRecord
	PoolRecords           = types.PoolRecords
	PoolStatsQueryResult  = types.PoolStatsQueryResult
	PoolStatsQueryResults = types.PoolStatsQueryResults
	QueryDepositsParams   = types.QueryDepositsParams
	QueryPoolParams       = types.QueryPoolParams
	ShareRecord           = types.ShareRecord
	ShareRecords          = types.ShareRecords
	SupplyKeeper          = types.SupplyKeeper
	SwapHooks             = types.SwapHooks
)
