package incentive

// DO NOT EDIT - generated by aliasgen tool (github.com/rhuairahrighairidh/aliasgen)

import (
	"github.com/kava-labs/kava/x/incentive/keeper"
	"github.com/kava-labs/kava/x/incentive/types"
)

const (
	AttributeKeyClaimAmount    = types.AttributeKeyClaimAmount
	AttributeKeyClaimPeriod    = types.AttributeKeyClaimPeriod
	AttributeKeyClaimedBy      = types.AttributeKeyClaimedBy
	AttributeKeyRewardPeriod   = types.AttributeKeyRewardPeriod
	AttributeValueCategory     = types.AttributeValueCategory
	DefaultParamspace          = types.DefaultParamspace
	EventTypeClaim             = types.EventTypeClaim
	EventTypeClaimPeriod       = types.EventTypeClaimPeriod
	EventTypeClaimPeriodExpiry = types.EventTypeClaimPeriodExpiry
	EventTypeRewardPeriod      = types.EventTypeRewardPeriod
	Large                      = types.Large
	Medium                     = types.Medium
	ModuleName                 = types.ModuleName
	QuerierRoute               = types.QuerierRoute
	QueryGetClaimPeriods       = types.QueryGetClaimPeriods
	QueryGetClaims             = types.QueryGetClaims
	QueryGetParams             = types.QueryGetParams
	QueryGetRewardPeriods      = types.QueryGetRewardPeriods
	RestClaimCollateralType    = types.RestClaimCollateralType
	RestClaimOwner             = types.RestClaimOwner
	RouterKey                  = types.RouterKey
	Small                      = types.Small
	StoreKey                   = types.StoreKey
)

var (
	// function aliases
	CalculateTimeElapsed         = keeper.CalculateTimeElapsed
	NewKeeper                    = keeper.NewKeeper
	NewQuerier                   = keeper.NewQuerier
	DefaultGenesisState          = types.DefaultGenesisState
	DefaultParams                = types.DefaultParams
	GetTotalVestingPeriodLength  = types.GetTotalVestingPeriodLength
	NewGenesisAccumulationTime   = types.NewGenesisAccumulationTime
	NewGenesisState              = types.NewGenesisState
	NewMsgClaimUSDXMintingReward = types.NewMsgClaimUSDXMintingReward
	NewMultiplier                = types.NewMultiplier
	NewParams                    = types.NewParams
	NewPeriod                    = types.NewPeriod
	NewQueryClaimsParams         = types.NewQueryClaimsParams
	NewRewardIndex               = types.NewRewardIndex
	NewRewardPeriod              = types.NewRewardPeriod
	NewUSDXMintingClaim          = types.NewUSDXMintingClaim
	ParamKeyTable                = types.ParamKeyTable
	RegisterCodec                = types.RegisterCodec

	// variable aliases
	PreviousUSDXMintingRewardAccrualTimeKeyPrefix = types.PreviousUSDXMintingRewardAccrualTimeKeyPrefix
	USDXMintingClaimKeyPrefix                     = types.USDXMintingClaimKeyPrefix
	DefaultActive                                 = types.DefaultActive
	DefaultClaimEnd                               = types.DefaultClaimEnd
	DefaultClaims                                 = types.DefaultClaims
	DefaultGenesisAccumulationTimes               = types.DefaultGenesisAccumulationTimes
	DefaultMultipliers                            = types.DefaultMultipliers
	DefaultRewardPeriods                          = types.DefaultRewardPeriods
	ErrAccountNotFound                            = types.ErrAccountNotFound
	ErrClaimExpired                               = types.ErrClaimExpired
	ErrClaimNotFound                              = types.ErrClaimNotFound
	ErrInsufficientModAccountBalance              = types.ErrInsufficientModAccountBalance
	ErrInvalidAccountType                         = types.ErrInvalidAccountType
	ErrInvalidMultiplier                          = types.ErrInvalidMultiplier
	ErrNoClaimsFound                              = types.ErrNoClaimsFound
	ErrRewardPeriodNotFound                       = types.ErrRewardPeriodNotFound
	ErrZeroClaim                                  = types.ErrZeroClaim
	GovDenom                                      = types.GovDenom
	IncentiveMacc                                 = types.IncentiveMacc
	KeyClaimEnd                                   = types.KeyClaimEnd
	KeyMultipliers                                = types.KeyMultipliers
	KeyUSDXMintingRewardPeriods                   = types.KeyUSDXMintingRewardPeriods
	ModuleCdc                                     = types.ModuleCdc
	PrincipalDenom                                = types.PrincipalDenom
	USDXMintingRewardFactorKeyPrefix              = types.USDXMintingRewardFactorKeyPrefix
	USDXMintingRewardDenom                        = types.USDXMintingRewardDenom
)

type (
	Hooks                     = keeper.Hooks
	Keeper                    = keeper.Keeper
	AccountKeeper             = types.AccountKeeper
	CDPHooks                  = types.CDPHooks
	HARDHooks                 = types.HARDHooks
	CdpKeeper                 = types.CdpKeeper
	GenesisAccumulationTime   = types.GenesisAccumulationTime
	GenesisAccumulationTimes  = types.GenesisAccumulationTimes
	GenesisState              = types.GenesisState
	MsgClaimUSDXMintingReward = types.MsgClaimUSDXMintingReward
	Multiplier                = types.Multiplier
	MultiplierName            = types.MultiplierName
	Multipliers               = types.Multipliers
	Params                    = types.Params
	PostClaimReq              = types.PostClaimReq
	QueryClaimsParams         = types.QueryClaimsParams
	RewardIndex               = types.RewardIndex
	RewardIndexes             = types.RewardIndexes
	RewardPeriod              = types.RewardPeriod
	RewardPeriods             = types.RewardPeriods
	SupplyKeeper              = types.SupplyKeeper
	USDXMintingClaim          = types.USDXMintingClaim
	USDXMintingClaims         = types.USDXMintingClaims
)
