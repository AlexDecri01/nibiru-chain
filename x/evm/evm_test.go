// Copyright (c) 2023-2024 Nibi, Inc.
package evm_test

import (
	"strconv"
	"testing"

	"github.com/NibiruChain/nibiru/eth"
	"github.com/NibiruChain/nibiru/x/evm"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestSuite_RunAll(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestFunToken() {
	for testIdx, tc := range []struct {
		bankDenom string
		erc20Addr string
		wantErr   string
	}{
		{
			bankDenom: "",
			erc20Addr: "",
			wantErr:   "FunTokenError",
		},
		{
			bankDenom: "unibi",
			erc20Addr: "5aaeb6053f3e94c9b9a09f33669435e7ef1beaed",
			wantErr:   "not encoded as expected",
		},
		{
			bankDenom: "unibi",
			erc20Addr: eth.NewHexAddrFromStr("5aaeb6053f3e94c9b9a09f33669435e7ef1beaed").String(),
			wantErr:   "",
		},
		{
			bankDenom: "ibc/AAA/BBB",
			erc20Addr: "0xE1aA1500b962528cBB42F05bD6d8A6032a85602f",
			wantErr:   "",
		},
		{
			bankDenom: "tf/contract-addr/subdenom",
			erc20Addr: "0x6B2e60f1030aFa69F584829f1d700b47eE5Fc74a",
			wantErr:   "",
		},
	} {
		s.Run(strconv.Itoa(testIdx), func() {
			funtoken := evm.FunToken{
				Erc20Addr: tc.erc20Addr,
				BankDenom: tc.bankDenom,
			}
			err := funtoken.Validate()
			if tc.wantErr != "" {
				s.Require().Error(err)
				return
			}
			s.Require().NoError(err)
		})
	}

	for _, tc := range []struct {
		name string
		A    string
		B    string
	}{
		{
			name: "capital and lowercase match",
			A:    "5aaeb6053f3e94c9b9a09f33669435e7ef1beaed",
			B:    "5AAEB6053F3E94C9B9A09F33669435E7EF1BEAED",
		},
		{
			name: "0x prefix and no prefix match",
			A:    "5aaeb6053f3e94c9b9a09f33669435e7ef1beaed",
			B:    "0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed",
		},
		{
			name: "0x prefix and no prefix match",
			A:    "5aaeb6053f3e94c9b9a09f33669435e7ef1beaed",
			B:    "0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed",
		},
		{
			name: "mixed case compatibility",
			A:    "0x5Bdb32670a05Daa22Cb2E279B80044c37dc85e61",
			B:    "0x5BDB32670A05DAA22CB2E279B80044C37DC85E61",
		},
	} {
		s.Run(tc.name, func() {
			funA := evm.FunToken{Erc20Addr: tc.A}
			funB := evm.FunToken{Erc20Addr: tc.B}

			s.EqualValues(funA.ERC20Addr(), funB.ERC20Addr())
		})
	}
}
