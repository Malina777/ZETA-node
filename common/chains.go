package common

// Zeta chains

func ZetaChain() Chain {
	return Chain{
		ChainName: ChainName_zeta_mainnet,
		ChainId:   7000,
	}
}

func ZetaTestnetChain() Chain {
	return Chain{
		ChainName: ChainName_zeta_testnet,
		ChainId:   7001,
	}
}

func ZetaMocknetChain() Chain {
	return Chain{
		ChainName: ChainName_zeta_mainnet,
		ChainId:   70000,
	}
}

func ZetaPrivnetChain() Chain {
	return Chain{
		ChainName: ChainName_zeta_mainnet,
		ChainId:   101,
	}
}

// Mainnet chains

func EthChain() Chain {
	return Chain{
		ChainName: ChainName_eth_mainnet,
		ChainId:   1,
	}
}

func BscMainnetChain() Chain {
	return Chain{
		ChainName: ChainName_bsc_mainnet,
		ChainId:   56,
	}
}

func BtcMainnetChain() Chain {
	return Chain{
		ChainName: ChainName_btc_mainnet,
		ChainId:   8332,
	}
}

func PolygonChain() Chain {
	return Chain{
		ChainName: ChainName_polygon_mainnet,
		ChainId:   137,
	}
}

// Testnet chains

func GoerliChain() Chain {
	return Chain{
		ChainName: ChainName_goerli_testnet,
		ChainId:   5,
	}
}

func BscTestnetChain() Chain {
	return Chain{
		ChainName: ChainName_bsc_testnet,
		ChainId:   97,
	}
}

func BtcTestNetChain() Chain {
	return Chain{
		ChainName: ChainName_btc_testnet,
		ChainId:   18332,
	}
}

func MumbaiChain() Chain {
	return Chain{
		ChainName: ChainName_mumbai_testnet,
		ChainId:   80001,
	}
}

// Privnet chains

func BtcRegtestChain() Chain {
	return Chain{
		ChainName: ChainName_btc_regtest,
		ChainId:   18444,
	}
}

func GoerliLocalnetChain() Chain {
	return Chain{
		ChainName: ChainName_goerli_localnet,
		ChainId:   1337,
	}
}

func BtcChainID() int64 {
	return BtcRegtestChain().ChainId
}

func BtcDustOffset() int64 {
	return 2000
}

// DefaultChainsList returns a list of default chains
func DefaultChainsList() []*Chain {
	chains := []Chain{
		BtcMainnetChain(),
		BscMainnetChain(),
		EthChain(),
		BtcTestNetChain(),
		MumbaiChain(),
		BscTestnetChain(),
		GoerliChain(),
		BtcRegtestChain(),
		GoerliLocalnetChain(),
		ZetaChain(),
	}
	var c []*Chain
	for i := 0; i < len(chains); i++ {
		c = append(c, &chains[i])
	}
	return c
}

// ExternalChainList returns a list chains that are not Zeta
func ExternalChainList() []*Chain {
	chains := []Chain{
		BtcMainnetChain(),
		BscMainnetChain(),
		EthChain(),
		BtcTestNetChain(),
		MumbaiChain(),
		BscTestnetChain(),
		GoerliChain(),
		BtcRegtestChain(),
		GoerliLocalnetChain(),
	}
	var c []*Chain
	for i := 0; i < len(chains); i++ {
		c = append(c, &chains[i])
	}
	return c
}