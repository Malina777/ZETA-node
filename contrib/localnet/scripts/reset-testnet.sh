#!/bin/bash
set -x

CHAINID="athens_101-1"
KEYRING="test"
HOSTNAME=$(hostname)

if [ $# -ne 1 ]
then
  echo "Usage: reset-testnet.sh <num of nodes>"
  exit 1
fi
NUMOFNODES=$1

# generate node list
START=1
# shellcheck disable=SC2100
END=$((NUMOFNODES - 1))
NODELIST=()
for i in $(eval echo "{$START..$END}")
do
  NODELIST+=("zetacore$i")
done

HOSTNAME=$(hostname)
if [ $HOSTNAME != "zetacore0" ]
then
  echo "You should run this only on zetacore0."
  exit 1
fi



# Init a new node to generate genesis file.
# Copy config files from existing folders which get copied via Docker Copy when building images
mkdir -p ~/.backup/config
zetacored init Zetanode-Localnet --chain-id="$CHAINID"
for NODE in "${NODELIST[@]}"
do
  ssh "$NODE" mkdir -p ~/.backup/config
  ssh "$NODE" zetacored init Zetanode-Localnet --chain-id="$CHAINID"
done

# Add two new keys for operator and hotkey and create the required json structure for os_info
source /root/os-info.sh
for NODE in "${NODELIST[@]}"; do
    ssh $NODE source /root/os-info.sh
done


# Start of genesis creation.

# 1. Accumulate all the os_info files from other nodes on zetcacore0 and create a genesis.json
ssh zetaclient0 mkdir -p ~/.zetacored/
scp ~/.zetacored/os_info/os.json zetaclient0:~/.zetacored/os.json
for NODE in "${NODELIST[@]}"; do
  INDEX=${NODE:0-1}
  ssh zetaclient"$INDEX" mkdir -p ~/.zetacored/
  scp "$NODE":~/.zetacored/os_info/os.json ~/.zetacored/os_info/os_z"$INDEX".json
  scp ~/.zetacored/os_info/os_z"$INDEX".json zetaclient"$INDEX":~/.zetacored/os.json
done

# 2. Add the observers , authorizations and required params to the genesis.json
zetacored collect-observer-info
zetacored add-observer-list
cat $HOME/.zetacored/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="azeta"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
cat $HOME/.zetacored/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="azeta"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
cat $HOME/.zetacored/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="azeta"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
cat $HOME/.zetacored/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="azeta"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
cat $HOME/.zetacored/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="azeta"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
cat $HOME/.zetacored/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json

# 3. Copy the genesis.json to all the nodes .And use it to create a gentx for every node
zetacored gentx operator 1000000000000000000000azeta --chain-id=$CHAINID --keyring-backend=$KEYRING
mkdir ~/.zetacored/config/gentx/z2gentx
for NODE in "${NODELIST[@]}"; do
    ssh $NODE rm -rf ~/.zetacored/genesis.json
    scp ~/.zetacored/config/genesis.json $NODE:~/.zetacored/config/genesis.json
    ssh $NODE zetacored gentx operator 1000000000000000000000azeta --chain-id=$CHAINID --keyring-backend=$KEYRING
    scp $NODE:~/.zetacored/config/gentx/* ~/.zetacored/config/gentx/
    scp $NODE:~/.zetacored/config/gentx/* ~/.zetacored/config/gentx/z2gentx/
done

# 4. Collect all the gentx files in zetacore0 and create the final genesis.json
zetacored collect-gentxs
zetacored validate-genesis

# 5. Copy the final genesis.json to all the nodes
for NODE in "${NODELIST[@]}"; do
    ssh $NODE rm -rf ~/.zetacored/genesis.json
    scp ~/.zetacored/config/genesis.json $NODE:~/.zetacored/config/genesis.json
done

# 6. Update Config in zetacore0 so that it has the correct persistent peer list
sleep 2
pp=$(cat $HOME/.zetacored/config/gentx/z2gentx/*.json | jq '.body.memo' )
pps=${pp:1:58}
sed -i -e "/persistent_peers =/s/=.*/= \"$pps\"/" "$HOME"/.zetacored/config/config.toml
# End of genesis creation steps . The steps below are common to all the nodes

# Misc : Copying the keyring to the client nodes so that they can sign the transactions
# We do not need to use keyring/* , as the client only needs the hotkey to sign the transactions but differentiating between the two keys would add additional logic to the script
scp ~/.zetacored/keyring-test/* "zetaclient0":~/.zetacored/keyring-test/
for NODE in "${NODELIST[@]}"; do
  INDEX=${NODE:0-1}
  ssh "zetaclient$INDEX" mkdir -p ~/.zetacored/keyring-test/
  scp ~/.zetacored/keyring-test/* "zetaclient$INDEX":~/.zetacored/keyring-test/
done