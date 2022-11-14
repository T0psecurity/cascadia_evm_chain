#!/bin/bash

KEY="mykey"
CHAINID="evmos_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t evmos-datadir.XXXXX)

echo "create and add new keys"
./cascadiad keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Evmos with moniker=$MONIKER and chain-id=$CHAINID"
./cascadiad init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./cascadiad add-genesis-account \
"$(./cascadiad keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000aevmos,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./cascadiad gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./cascadiad collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./cascadiad validate-genesis --home $DATA_DIR

echo "starting evmos node $i in background ..."
./cascadiad start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started evmos node"
tail -f /dev/null