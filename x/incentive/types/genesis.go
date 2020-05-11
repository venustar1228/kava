package types

import (
	"bytes"
	"fmt"
	"time"
)

// GenesisClaimPeriodID stores the next claim id and its corresponding denom
type GenesisClaimPeriodID struct {
	Denom string `json:"denom" yaml:"denom"`
	ID    uint64 `json:"id" yaml:"id"`
}

// GenesisClaimPeriodIDs array of GenesisClaimPeriodID
type GenesisClaimPeriodIDs []GenesisClaimPeriodID

// GenesisState is the state that must be provided at genesis.
type GenesisState struct {
	Params             Params                `json:"params" yaml:"params"`
	PreviousBlockTime  time.Time             `json:"previous_block_time" yaml:"previous_block_time"`
	RewardPeriods      RewardPeriods         `json:"reward_periods" yaml:"reward_periods"`
	ClaimPeriods       ClaimPeriods          `json:"claim_periods" yaml:"claim_periods"`
	Claims             Claims                `json:"claims" yaml:"claims"`
	NextClaimPeriodIDs GenesisClaimPeriodIDs `json:"next_claim_period_ids" yaml:"next_claim_period_ids"`
}

// NewGenesisState returns a new genesis state
func NewGenesisState(params Params, previousBlockTime time.Time, rp RewardPeriods, cp ClaimPeriods, c Claims, ids GenesisClaimPeriodIDs) GenesisState {
	return GenesisState{
		Params:             params,
		PreviousBlockTime:  previousBlockTime,
		RewardPeriods:      rp,
		ClaimPeriods:       cp,
		Claims:             c,
		NextClaimPeriodIDs: ids,
	}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params:             DefaultParams(),
		PreviousBlockTime:  DefaultPreviousBlockTime,
		RewardPeriods:      RewardPeriods{},
		ClaimPeriods:       ClaimPeriods{},
		Claims:             Claims{},
		NextClaimPeriodIDs: GenesisClaimPeriodIDs{},
	}
}

// Validate performs basic validation of genesis data returning an
// error for any failed validation criteria.
func (gs GenesisState) Validate() error {

	if err := gs.Params.Validate(); err != nil {
		return err
	}
	if gs.PreviousBlockTime.IsZero() {
		return fmt.Errorf("previous block time not set or zero")
	}
	return nil
}

// Equal checks whether two gov GenesisState structs are equivalent
func (gs GenesisState) Equal(gs2 GenesisState) bool {
	b1 := ModuleCdc.MustMarshalBinaryBare(gs)
	b2 := ModuleCdc.MustMarshalBinaryBare(gs2)
	return bytes.Equal(b1, b2)
}

// IsEmpty returns true if a GenesisState is empty
func (gs GenesisState) IsEmpty() bool {
	return gs.Equal(GenesisState{})
}
