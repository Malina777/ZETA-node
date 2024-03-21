package types

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
)

// DefaultGenesis returns the default authority genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Policies: DefaultPolicies(),
	}
}

// Validate performs basic genesis state validation returning an error upon any failure
func (gs GenesisState) Validate() error {
	return gs.Policies.Validate()
}

func GetGenesisStateFromAppState(marshaler codec.JSONCodec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState
	if appState[ModuleName] != nil {
		err := marshaler.UnmarshalJSON(appState[ModuleName], &genesisState)
		if err != nil {
			panic(fmt.Sprintf("Failed to get genesis state from app state: %s", err.Error()))
		}
	}
	return genesisState
}
