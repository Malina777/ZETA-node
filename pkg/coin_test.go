package pkg_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/zetacore/pkg"
)

func Test_GetAzetaDecFromAmountInZeta(t *testing.T) {
	tt := []struct {
		name        string
		zetaAmount  string
		err         require.ErrorAssertionFunc
		azetaAmount sdk.Dec
	}{
		{
			name:        "valid zeta amount",
			zetaAmount:  "210000000",
			err:         require.NoError,
			azetaAmount: sdk.MustNewDecFromStr("210000000000000000000000000"),
		},
		{
			name:        "very high zeta amount",
			zetaAmount:  "21000000000000000000",
			err:         require.NoError,
			azetaAmount: sdk.MustNewDecFromStr("21000000000000000000000000000000000000"),
		},
		{
			name:        "very low zeta amount",
			zetaAmount:  "1",
			err:         require.NoError,
			azetaAmount: sdk.MustNewDecFromStr("1000000000000000000"),
		},
		{
			name:        "zero zeta amount",
			zetaAmount:  "0",
			err:         require.NoError,
			azetaAmount: sdk.MustNewDecFromStr("0"),
		},
		{
			name:        "decimal zeta amount",
			zetaAmount:  "0.1",
			err:         require.NoError,
			azetaAmount: sdk.MustNewDecFromStr("100000000000000000"),
		},
		{
			name:        "invalid zeta amount",
			zetaAmount:  "%%%%%$#",
			err:         require.Error,
			azetaAmount: sdk.MustNewDecFromStr("0"),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			azeta, err := pkg.GetAzetaDecFromAmountInZeta(tc.zetaAmount)
			tc.err(t, err)
			if err == nil {
				require.Equal(t, tc.azetaAmount, azeta)
			}
		})
	}

}

func TestGetCoinType(t *testing.T) {
	tests := []struct {
		name    string
		coin    string
		want    pkg.CoinType
		wantErr bool
	}{
		{
			name:    "valid coin type 0",
			coin:    "0",
			want:    pkg.CoinType(0),
			wantErr: false,
		},
		{
			name:    "valid coin type 1",
			coin:    "1",
			want:    pkg.CoinType(1),
			wantErr: false,
		},
		{
			name:    "valid coin type 2",
			coin:    "2",
			want:    pkg.CoinType(2),
			wantErr: false,
		},
		{
			name:    "valid coin type 3",
			coin:    "3",
			want:    pkg.CoinType(3),
			wantErr: false,
		},
		{
			name:    "invalid coin type negative",
			coin:    "-1",
			want:    pkg.CoinType_Cmd,
			wantErr: true,
		},
		{
			name:    "invalid coin type large number",
			coin:    "4",
			want:    pkg.CoinType_Cmd,
			wantErr: true,
		},
		{
			name:    "invalid coin type non-integer",
			coin:    "abc",
			want:    pkg.CoinType_Cmd,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pkg.GetCoinType(tt.coin)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}