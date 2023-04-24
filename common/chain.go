package common

import (
	"strings"
)

var (
	SigningAlgoSecp256k1 = SigninAlgo("secp256k1")
	SigningAlgoEd25519   = SigninAlgo("ed25519")
)

// return the ChainName from a string
// if no such name exists, returns the empty chain name: ChainName_empty
func ParseChainName(chain string) ChainName {
	c := ChainName_value[chain]
	return ChainName(c)
}

type SigninAlgo string

// Chain is an alias of string , represent a block chain
//type Chain string

// Chains represent a slice of Chain
type Chains []Chain

// Equals compare two chain to see whether they represent the same chain
func (chain Chain) IsEqual(c Chain) bool {
	if chain.ChainName == c.ChainName && chain.ChainId == c.ChainId {
		return true
	}
	return false
}

func (chain Chain) IsZetaChain() bool {
	return chain.IsEqual(ZetaChain())
}
func (chain Chain) IsExternalChain() bool {
	return !chain.IsEqual(ZetaChain())
}

func (chain Chain) IsEVMChain() bool {
	return chain.ChainId == 1 || // Ethereum
		chain.ChainId == 56 || // BSC
		chain.ChainId == 137 || // Polygon
		chain.ChainId == 5 || // Goerli
		chain.ChainId == 80001 || // Polygon mumbai
		chain.ChainId == 97 || // BSC testnet
		chain.ChainId == 1001 || // klaytn baobab
		chain.ChainId == 1337 // eth privnet
}

func (chain Chain) IsKlaytnChain() bool {
	return chain.ChainId == 1001
}

func (chain Chain) IsBitcoinChain() bool {
	return chain.ChainId == 18444 || // regtest
		chain.ChainId == 18332 || //testnet
		chain.ChainId == 8332 // mainnet
}

// IsEmpty is to determinate whether the chain is empty
func (chain Chain) IsEmpty() bool {
	return strings.TrimSpace(chain.String()) == ""
}

// Has check whether chain c is in the list
func (chains Chains) Has(c Chain) bool {
	for _, ch := range chains {
		if ch.IsEqual(c) {
			return true
		}
	}
	return false
}

// Distinct return a distinct set of chains, no duplicates
func (chains Chains) Distinct() Chains {
	var newChains Chains
	for _, chain := range chains {
		if !newChains.Has(chain) {
			newChains = append(newChains, chain)
		}
	}
	return newChains
}

func (chains Chains) Strings() []string {
	strings := make([]string, len(chains))
	for i, c := range chains {
		strings[i] = c.String()
	}
	return strings
}

func GetChainFromChainName(chainName ChainName) *Chain {
	chains := DefaultChainsList()
	for _, chain := range chains {
		if chainName == chain.ChainName {
			return chain
		}
	}
	return nil
}

func GetChainFromChainID(chainID int64) *Chain {
	chains := DefaultChainsList()
	for _, chain := range chains {
		if chainID == chain.ChainId {
			return chain
		}
	}
	return nil
}
