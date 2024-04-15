package types_test

import (
	"testing"

	"github.com/NibiruChain/collections"
	ethtypes "github.com/NibiruChain/nibiru/eth/types"
	cmtdb "github.com/cometbft/cometbft-db"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	sdkcodec "github.com/cosmos/cosmos-sdk/codec"
	sdkcodectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkstore "github.com/cosmos/cosmos-sdk/store"
	sdkstoretypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// encoderDeps: Initializes a database and KV store useful for testing
// abstractions over merklized storage like the `collections.Map` and
// `collections.Item`.
func encoderDeps() (sdkstoretypes.StoreKey, sdk.Context, sdkcodec.BinaryCodec) {
	sk := sdk.NewKVStoreKey("mock")
	dbm := cmtdb.NewMemDB()
	ms := sdkstore.NewCommitMultiStore(dbm)
	ms.MountStoreWithDB(sk, sdkstoretypes.StoreTypeIAVL, dbm)
	if err := ms.LoadLatestVersion(); err != nil {
		panic(err)
	}

	return sk,
		sdk.Context{}.
			WithMultiStore(ms).
			WithGasMeter(sdk.NewGasMeter(1_000_000_000)),
		sdkcodec.NewProtoCodec(sdkcodectypes.NewInterfaceRegistry())
}

func assertBijectiveKey[T any](t *testing.T, encoder collections.KeyEncoder[T], key T) {
	encodedKey := encoder.Encode(key)
	readLen, decodedKey := encoder.Decode(encodedKey)
	require.Equal(t, len(encodedKey), readLen, "encoded key and read bytes must have same size")
	require.Equal(t, key, decodedKey, "encoding and decoding produces different keys")
	wantStr := encoder.Stringify(key)
	gotStr := encoder.Stringify(decodedKey)
	require.Equal(t, wantStr, gotStr,
		"encoding and decoding produce different string representations")
}

func assertBijectiveValue[T any](t *testing.T, encoder collections.ValueEncoder[T], value T) {
	encodedValue := encoder.Encode(value)
	decodedValue := encoder.Decode(encodedValue)
	require.Equal(t, value, decodedValue, "encoding and decoding produces different values")

	wantStr := encoder.Stringify(value)
	gotStr := encoder.Stringify(decodedValue)
	require.Equal(t, wantStr, gotStr,
		"encoding and decoding produce different string representations")
	require.NotEmpty(t, encoder.Name())
}

type SuiteEncoders struct {
	suite.Suite
}

func TestSuiteEncoders_RunAll(t *testing.T) {
	suite.Run(t, new(SuiteEncoders))
}

func (s *SuiteEncoders) TestEncoderBytes() {
	testCases := []struct {
		name  string
		value string
	}{
		{"dec-like number", "-1000.5858"},
		{"Nibiru bech32 addr", "nibi1rlvdjfmxkyfj4tzu73p8m4g2h4y89xccf9622l"},
		{"Nibiru EVM addr", "0xA52c829E935C30F4C7dcD66739Cf91BF79dD9253"},
		{"normal text with special symbols", "abc123日本語!!??foobar"},
	}
	for _, tc := range testCases {
		s.Run("bijectivity: []byte encoders "+tc.name, func() {
			given := []byte(tc.value)
			assertBijectiveKey(s.T(), ethtypes.KeyEncoderBytes, given)
			assertBijectiveValue(s.T(), ethtypes.ValueEncoderBytes, given)
		})
	}
}

func (s *SuiteEncoders) TestEncoderEthAddr() {
	testCases := []struct {
		name      string
		given     ethtypes.EthAddr
		wantPanic bool
	}{
		{
			name:  "Nibiru EVM addr",
			given: ethcommon.BytesToAddress([]byte("0xA52c829E935C30F4C7dcD66739Cf91BF79dD9253")),
		},
		{
			name:  "Nibiru EVM addr length above 20 bytes",
			given: ethcommon.BytesToAddress([]byte("0xA52c829E935C30F4C7dcD66739Cf91BF79dD92532456BF123421")),
		},
		{
			name:  "Nibiru Bech 32 addr (hypothetically)",
			given: ethtypes.EthAddr([]byte("nibi1rlvdjfmxkyfj4tzu73p8m4g2h4y89xccf9622l")),
		},
	}
	for _, tc := range testCases {
		s.Run("bijectivity: []byte encoders "+tc.name, func() {
			given := tc.given
			runTest := func() {
				assertBijectiveKey(s.T(), ethtypes.KeyEncoderEthAddr, given)
				assertBijectiveValue(s.T(), ethtypes.ValueEncoderEthAddr, given)
			}
			if tc.wantPanic {
				s.Require().Panics(runTest)
			} else {
				s.Require().NotPanics(runTest)
			}
		})
	}
}
