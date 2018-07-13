package paychan

import (
	"testing"
	//"github.com/stretchr/testify/assert"
)

// example from x/bank
func TestKeeper(t *testing.T) {
	// ms, authKey := setupMultiStore()

	// cdc := wire.NewCodec()
	// auth.RegisterBaseAccount(cdc)

	// ctx := sdk.NewContext(ms, abci.Header{}, false, nil, log.NewNopLogger())
	// accountMapper := auth.NewAccountMapper(cdc, authKey, &auth.BaseAccount{})
	// coinKeeper := NewKeeper(accountMapper)

	// addr := sdk.Address([]byte("addr1"))
	// addr2 := sdk.Address([]byte("addr2"))
	// addr3 := sdk.Address([]byte("addr3"))
	// acc := accountMapper.NewAccountWithAddress(ctx, addr)

	// // Test GetCoins/SetCoins
	// accountMapper.SetAccount(ctx, acc)
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{}))

	// coinKeeper.SetCoins(ctx, addr, sdk.Coins{{"foocoin", 10}})
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{{"foocoin", 10}}))

	// // Test HasCoins
	// assert.True(t, coinKeeper.HasCoins(ctx, addr, sdk.Coins{{"foocoin", 10}}))
	// assert.True(t, coinKeeper.HasCoins(ctx, addr, sdk.Coins{{"foocoin", 5}}))
	// assert.False(t, coinKeeper.HasCoins(ctx, addr, sdk.Coins{{"foocoin", 15}}))
	// assert.False(t, coinKeeper.HasCoins(ctx, addr, sdk.Coins{{"barcoin", 5}}))

	// // Test AddCoins
	// coinKeeper.AddCoins(ctx, addr, sdk.Coins{{"foocoin", 15}})
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{{"foocoin", 25}}))

	// coinKeeper.AddCoins(ctx, addr, sdk.Coins{{"barcoin", 15}})
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{{"barcoin", 15}, {"foocoin", 25}}))

	// // Test SubtractCoins
	// coinKeeper.SubtractCoins(ctx, addr, sdk.Coins{{"foocoin", 10}})
	// coinKeeper.SubtractCoins(ctx, addr, sdk.Coins{{"barcoin", 5}})
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{{"barcoin", 10}, {"foocoin", 15}}))

	// coinKeeper.SubtractCoins(ctx, addr, sdk.Coins{{"barcoin", 11}})
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{{"barcoin", 10}, {"foocoin", 15}}))

	// coinKeeper.SubtractCoins(ctx, addr, sdk.Coins{{"barcoin", 10}})
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{{"foocoin", 15}}))
	// assert.False(t, coinKeeper.HasCoins(ctx, addr, sdk.Coins{{"barcoin", 1}}))

	// // Test SendCoins
	// coinKeeper.SendCoins(ctx, addr, addr2, sdk.Coins{{"foocoin", 5}})
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{{"foocoin", 10}}))
	// assert.True(t, coinKeeper.GetCoins(ctx, addr2).IsEqual(sdk.Coins{{"foocoin", 5}}))

	// _, err2 := coinKeeper.SendCoins(ctx, addr, addr2, sdk.Coins{{"foocoin", 50}})
	// assert.Implements(t, (*sdk.Error)(nil), err2)
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{{"foocoin", 10}}))
	// assert.True(t, coinKeeper.GetCoins(ctx, addr2).IsEqual(sdk.Coins{{"foocoin", 5}}))

	// coinKeeper.AddCoins(ctx, addr, sdk.Coins{{"barcoin", 30}})
	// coinKeeper.SendCoins(ctx, addr, addr2, sdk.Coins{{"barcoin", 10}, {"foocoin", 5}})
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{{"barcoin", 20}, {"foocoin", 5}}))
	// assert.True(t, coinKeeper.GetCoins(ctx, addr2).IsEqual(sdk.Coins{{"barcoin", 10}, {"foocoin", 10}}))

	// // Test InputOutputCoins
	// input1 := NewInput(addr2, sdk.Coins{{"foocoin", 2}})
	// output1 := NewOutput(addr, sdk.Coins{{"foocoin", 2}})
	// coinKeeper.InputOutputCoins(ctx, []Input{input1}, []Output{output1})
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{{"barcoin", 20}, {"foocoin", 7}}))
	// assert.True(t, coinKeeper.GetCoins(ctx, addr2).IsEqual(sdk.Coins{{"barcoin", 10}, {"foocoin", 8}}))

	// inputs := []Input{
	// 	NewInput(addr, sdk.Coins{{"foocoin", 3}}),
	// 	NewInput(addr2, sdk.Coins{{"barcoin", 3}, {"foocoin", 2}}),
	// }

	// outputs := []Output{
	// 	NewOutput(addr, sdk.Coins{{"barcoin", 1}}),
	// 	NewOutput(addr3, sdk.Coins{{"barcoin", 2}, {"foocoin", 5}}),
	// }
	// coinKeeper.InputOutputCoins(ctx, inputs, outputs)
	// assert.True(t, coinKeeper.GetCoins(ctx, addr).IsEqual(sdk.Coins{{"barcoin", 21}, {"foocoin", 4}}))
	// assert.True(t, coinKeeper.GetCoins(ctx, addr2).IsEqual(sdk.Coins{{"barcoin", 7}, {"foocoin", 6}}))
	// assert.True(t, coinKeeper.GetCoins(ctx, addr3).IsEqual(sdk.Coins{{"barcoin", 2}, {"foocoin", 5}}))

}
