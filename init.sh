#!/bin/sh

# set variables for the chain
VALIDATOR_NAME=validator1
CHAIN_ID=raffle
KEY_NAME=raffle-key-1
KEY_2_NAME=raffle-key-2
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000urfl"
STAKING_AMOUNT="1000000000urfl"
STAKING_BOND_DENOM="urfl"

# create a random Namespace ID for your rollup to post blocks to
NAMESPACE_ID=$(openssl rand -hex 8)
echo $NAMESPACE_ID

# query the DA Layer start height, in this case we are querying
# our local devnet at port 26657, the RPC. The RPC endpoint is
# to allow users to interact with Celestia's nodes by querying
# the node's state and broadcasting transactions on the Celestia
# network. The default port is 26657.
DA_BLOCK_HEIGHT=$(curl https://rpc-celestia-testnet-blockspacerace.keplr.app/block | jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT

# build the Raffle chain with Rollkit
ignite chain build

# reset any existing genesis/chain data
raffled tendermint unsafe-reset-all

# initialize the validator with the chain ID you set
raffled init $VALIDATOR_NAME --chain-id $CHAIN_ID --staking-bond-denom $STAKING_BOND_DENOM

# add keys for key 1 and key 2 to keyring-backend test
raffled keys add $KEY_NAME --keyring-backend test
raffled keys add $KEY_2_NAME --keyring-backend test

# add these as genesis accounts
raffled add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test
raffled add-genesis-account $KEY_2_NAME $TOKEN_AMOUNT --keyring-backend test

# set the staking amounts in the genesis transaction
raffled gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test

# collect genesis transactions
raffled collect-gentxs

# start the chain
raffled start --rollkit.aggregator true --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT --grpc.address 0.0.0.0:9098 --grpc-web.address 0.0.0.0:9099 --p2p.laddr tcp://0.0.0.0:27656 --rpc.laddr tcp://0.0.0.0:27657

# uncomment the next command if you are using lazy aggregation
# raffled start --rollkit.aggregator true --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT --rollkit.lazy_aggregator true
