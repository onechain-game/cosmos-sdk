#!/bin/sh
rm -rf ~/.simapp

echo "=====  Create .simapp ====="
echo ""
simd init simapp --chain-id my-test-chain
echo "=====  Create my_validator ====="
echo ""
simd keys add my_validator --keyring-backend test
echo "=====  Add genesis account ====="
echo ""
MY_VALIDATOR=$(simd keys show my_validator --keyring-backend test --address)
simd genesis add-genesis-account $MY_VALIDATOR 100000000stake
echo "=====  Add gentx ====="
echo ""
simd genesis gentx my_validator 60000000stake --chain-id my-test-chain --keyring-backend test
echo "=====  collect-gentxs ====="
echo ""
simd genesis collect-gentxs
cd Documents/cosmos/chain/cosmos-sdk && make install

