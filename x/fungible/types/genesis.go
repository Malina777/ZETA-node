package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ForeignCoinsList: []ForeignCoins{},
		SystemContract:   nil,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in foreignCoins
	foreignCoinsIndexMap := make(map[string]struct{})

	for _, elem := range gs.ForeignCoinsList {
		index := string(ForeignCoinsKey(elem.Index))
		if _, ok := foreignCoinsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for foreignCoins")
		}
		foreignCoinsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
