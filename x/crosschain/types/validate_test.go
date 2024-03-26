package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/zetacore/common"
	"github.com/zeta-chain/zetacore/testutil/sample"
	"github.com/zeta-chain/zetacore/x/crosschain/types"
)

func TestValidateAddressForChain(t *testing.T) {
	require.Error(t, types.ValidateAddressForChain("0x123", common.GoerliChain().ChainId))
	require.Error(t, types.ValidateAddressForChain("", common.GoerliChain().ChainId))
	require.Error(t, types.ValidateAddressForChain("%%%%", common.GoerliChain().ChainId))
	require.NoError(t, types.ValidateAddressForChain("0x792c127Fa3AC1D52F904056Baf1D9257391e7D78", common.GoerliChain().ChainId))
	require.Error(t, types.ValidateAddressForChain("1EYVvXLusCxtVuEwoYvWRyN5EZTXwPVvo3", common.BtcMainnetChain().ChainId))
	require.Error(t, types.ValidateAddressForChain("bcrt1qs758ursh4q9z627kt3pp5yysm78ddny6txaqgw", common.BtcMainnetChain().ChainId))
	require.Error(t, types.ValidateAddressForChain("", common.BtcRegtestChain().ChainId))
	require.NoError(t, types.ValidateAddressForChain("bc1qysd4sp9q8my59ul9wsf5rvs9p387hf8vfwatzu", common.BtcMainnetChain().ChainId))
	require.NoError(t, types.ValidateAddressForChain("bcrt1qs758ursh4q9z627kt3pp5yysm78ddny6txaqgw", common.BtcRegtestChain().ChainId))
	require.NoError(t, types.ValidateAddressForChain("bcrt1qs758ursh4q9z627kt3pp5yysm78ddny6txaqgw", common.ZetaChainMainnet().ChainId))
	require.NoError(t, types.ValidateAddressForChain("0x792c127Fa3AC1D52F904056Baf1D9257391e7D78", common.ZetaChainMainnet().ChainId))
}

func TestValidateZetaIndex(t *testing.T) {
	require.NoError(t, types.ValidateZetaIndex("0x84bd5c9922b63c52d8a9ca686e0a57ff978150b71be0583514d01c27aa341910"))
	require.NoError(t, types.ValidateZetaIndex(sample.ZetaIndex(t)))
	require.Error(t, types.ValidateZetaIndex("0"))
	require.Error(t, types.ValidateZetaIndex("0x70e967acFcC17c3941E87562161406d41676FD83"))
}

func TestValidateHashForChain(t *testing.T) {
	require.NoError(t, types.ValidateHashForChain("0x84bd5c9922b63c52d8a9ca686e0a57ff978150b71be0583514d01c27aa341910", common.GoerliChain().ChainId))
	require.Error(t, types.ValidateHashForChain("", common.GoerliChain().ChainId))
	require.Error(t, types.ValidateHashForChain("a0fa5a82f106fb192e4c503bfa8d54b2de20a821e09338094ab825cc9b275059", common.GoerliChain().ChainId))
	require.NoError(t, types.ValidateHashForChain("15b7880f5d236e857a5e8f043ce9d56f5ef01e1c3f2a786baf740fc0bb7a22a3", common.BtcMainnetChain().ChainId))
	require.NoError(t, types.ValidateHashForChain("a0fa5a82f106fb192e4c503bfa8d54b2de20a821e09338094ab825cc9b275059", common.BtcTestNetChain().ChainId))
	require.Error(t, types.ValidateHashForChain("0x84bd5c9922b63c52d8a9ca686e0a57ff978150b71be0583514d01c27aa341910", common.BtcMainnetChain().ChainId))
}