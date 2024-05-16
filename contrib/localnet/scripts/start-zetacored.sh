#!/bin/bash

# This script is used to start the zetacored nodes
# It initializes the nodes and creates the genesis.json file
# It also starts the nodes
# The number of nodes is passed as an first argument to the script

/usr/sbin/sshd

# This function add authz observer authorizations for inbound/outbound votes and tracker messages
# These messages have been renamed for v17: https://github.com/zeta-chain/node/blob/refactor/rename-outbound-inbound/docs/releases/v17_breaking_changes.md#intx-and-outtx-renaming
# There if the genesis is generated with a v16 binary for the upgrade tests, it will not contains authorizations for new messages
# This function will add the missing authorizations to the genesis file
# TODO: Remove this function when v17 is released
#
add_v17_message_authorizations() {
    # Path to the JSON file
    json_file="/root/.zetacored/config/genesis.json"

    # Using jq to parse JSON, create new entries, and append them to the authorization array
    jq '
        # Store the nodeAccountList array
        .app_state.observer.nodeAccountList as $list |

        # Iterate over the stored list to construct new objects and append to the authorization array
        .app_state.authz.authorization += [
            $list[] |
            {
                "granter": .operator,
                "grantee": .granteeAddress,
                "authorization": {
                    "@type": "/cosmos.authz.v1beta1.GenericAuthorization",
                    "msg": "/zetachain.zetacore.crosschain.MsgVoteInbound"
                },
                "expiration": null
            },
            {
                "granter": .operator,
                "grantee": .granteeAddress,
                "authorization": {
                    "@type": "/cosmos.authz.v1beta1.GenericAuthorization",
                    "msg": "/zetachain.zetacore.crosschain.MsgVoteOutbound"
                },
                "expiration": null
            },
            {
                "granter": .operator,
                "grantee": .granteeAddress,
                "authorization": {
                    "@type": "/cosmos.authz.v1beta1.GenericAuthorization",
                    "msg": "/zetachain.zetacore.crosschain.MsgAddOutboundTracker"
                },
                "expiration": null
            },
            {
                "granter": .operator,
                "grantee": .granteeAddress,
                "authorization": {
                    "@type": "/cosmos.authz.v1beta1.GenericAuthorization",
                    "msg": "/zetachain.zetacore.crosschain.Msg
                    AddInboundTracker"
                },
                "expiration": null
            }
        ]
    ' $json_file > temp.json && mv temp.json $json_file
}

if [ $# -lt 1 ]
then
  echo "Usage: genesis.sh <num of nodes> [option]"
  exit 1
fi
NUMOFNODES=$1

# create keys
CHAINID="athens_101-1"
KEYRING="test"
HOSTNAME=$(hostname)
INDEX=${HOSTNAME:0-1}

# Environment variables used for upgrade testing
export DAEMON_HOME=$HOME/.zetacored
export DAEMON_NAME=zetacored
export DAEMON_ALLOW_DOWNLOAD_BINARIES=true
export DAEMON_RESTART_AFTER_UPGRADE=true
export CLIENT_DAEMON_NAME=zetaclientd
export CLIENT_DAEMON_ARGS="-enable-chains,GOERLI,-val,operator"
export DAEMON_DATA_BACKUP_DIR=$DAEMON_HOME
export CLIENT_SKIP_UPGRADE=true
export CLIENT_START_PROCESS=false
export UNSAFE_SKIP_BACKUP=true

# generate node list
START=1
# shellcheck disable=SC2100
END=$((NUMOFNODES - 1))
NODELIST=()
for i in $(eval echo "{$START..$END}")
do
  NODELIST+=("zetacore$i")
done

echo "HOSTNAME: $HOSTNAME"

# init ssh keys
# we generate keys at runtime to ensure that keys are never pushed to
# a docker registry
if [ $HOSTNAME == "zetacore0" ]; then
  if [[ ! -f ~/.ssh/id_rsa ]]; then
    ssh-keygen -t rsa -q -N "" -f ~/.ssh/id_rsa
    cp ~/.ssh/id_rsa.pub ~/.ssh/authorized_keys
    # keep localtest.pem for compatibility
    cp ~/.ssh/id_rsa ~/.ssh/localtest.pem
    chmod 600 ~/.ssh/*
  fi
fi

# Wait for authorized_keys file to exist (zetacore1+)
while [ ! -f ~/.ssh/authorized_keys ]; do
    echo "Waiting for authorized_keys file to exist..."
    sleep 1
done

# Init a new node to generate genesis file .
# Copy config files from existing folders which get copied via Docker Copy when building images
mkdir -p ~/.backup/config
zetacored init Zetanode-Localnet --chain-id=$CHAINID
rm -rf ~/.zetacored/config/app.toml
rm -rf ~/.zetacored/config/client.toml
rm -rf ~/.zetacored/config/config.toml
cp -r ~/zetacored/common/app.toml ~/.zetacored/config/
cp -r ~/zetacored/common/client.toml ~/.zetacored/config/
cp -r ~/zetacored/common/config.toml ~/.zetacored/config/
sed -i -e "/moniker =/s/=.*/= \"$HOSTNAME\"/" "$HOME"/.zetacored/config/config.toml

# Add two new keys for operator and hotkey and create the required json structure for os_info
source ~/add-keys.sh

# Pause other nodes so that the primary can node can do the genesis creation
if [ $HOSTNAME != "zetacore0" ]
then
  while [ ! -f ~/.zetacored/config/genesis.json ]; do
    echo "Waiting for genesis.json file to exist..."
    sleep 1
  done
  # need to wait for zetacore0 to be up otherwise you get
  # 
  while ! curl -s -o /dev/null zetacore0:26657/status ; do
    echo "Waiting for zetacore0 rpc"
    sleep 1
done
fi

# Genesis creation following steps
# 1. Accumulate all the os_info files from other nodes on zetcacore0 and create a genesis.json
# 2. Add the observers , authorizations and required params to the genesis.json
# 3. Copy the genesis.json to all the nodes .And use it to create a gentx for every node
# 4. Collect all the gentx files in zetacore0 and create the final genesis.json
# 5. Copy the final genesis.json to all the nodes and start the nodes
# 6. Update Config in zetacore0 so that it has the correct persistent peer list
# 7. Start the nodes

# Start of genesis creation . This is done only on zetacore0
if [ $HOSTNAME == "zetacore0" ]
then
  # Misc : Copying the keyring to the client nodes so that they can sign the transactions
  ssh zetaclient0 mkdir -p ~/.zetacored/keyring-test/
  scp ~/.zetacored/keyring-test/* zetaclient0:~/.zetacored/keyring-test/
  ssh zetaclient0 mkdir -p ~/.zetacored/keyring-file/
  scp ~/.zetacored/keyring-file/* zetaclient0:~/.zetacored/keyring-file/

# 1. Accumulate all the os_info files from other nodes on zetcacore0 and create a genesis.json
  for NODE in "${NODELIST[@]}"; do
    INDEX=${NODE:0-1}
    ssh zetaclient"$INDEX" mkdir -p ~/.zetacored/
    while ! scp "$NODE":~/.zetacored/os_info/os.json ~/.zetacored/os_info/os_z"$INDEX".json; do
      echo "Waiting for os_info.json from node $NODE"
      sleep 1
    done
    scp ~/.zetacored/os_info/os_z"$INDEX".json zetaclient"$INDEX":~/.zetacored/os.json
  done

  ssh zetaclient0 mkdir -p ~/.zetacored/
  scp ~/.zetacored/os_info/os.json zetaclient0:/root/.zetacored/os.json

# 2. Add the observers, authorizations, required params and accounts to the genesis.json
  zetacored collect-observer-info
  zetacored add-observer-list --keygen-block 55

  # Check for the existence of "AddToOutTxTracker" string in the genesis file
  # If this message is found in the genesis, it means add-observer-list has been run with the v16 binary for upgrade tests
  # In this case, we need to add authorizations for the new v17 messages to the genesis file
  if jq -e 'tostring | contains("AddToOutTxTracker")' "/root/.zetacored/config/genesis.json" > /dev/null; then
    add_v17_message_authorizations
  fi

  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="azeta"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="azeta"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="azeta"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="azeta"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="azeta"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  cat $HOME/.zetacored/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="500000000"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["gov"]["voting_params"]["voting_period"]="100s"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["feemarket"]["params"]["min_gas_price"]="10000000000.0000"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json

# set admin account
  zetacored add-genesis-account zeta1srsq755t654agc0grpxj4y3w0znktrpr9tcdgk 100000000000000000000000000azeta
  zetacored add-genesis-account zeta1n0rn6sne54hv7w2uu93fl48ncyqz97d3kty6sh 100000000000000000000000000azeta # Funds the localnet_gov_admin account
  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["authority"]["policies"]["items"][0]["address"]="zeta1srsq755t654agc0grpxj4y3w0znktrpr9tcdgk"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["authority"]["policies"]["items"][1]["address"]="zeta1srsq755t654agc0grpxj4y3w0znktrpr9tcdgk"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["authority"]["policies"]["items"][2]["address"]="zeta1srsq755t654agc0grpxj4y3w0znktrpr9tcdgk"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["observer"]["params"]["admin_policy"][0]["address"]="zeta1srsq755t654agc0grpxj4y3w0znktrpr9tcdgk"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  cat $HOME/.zetacored/config/genesis.json | jq '.app_state["observer"]["params"]["admin_policy"][1]["address"]="zeta1srsq755t654agc0grpxj4y3w0znktrpr9tcdgk"' > $HOME/.zetacored/config/tmp_genesis.json && mv $HOME/.zetacored/config/tmp_genesis.json $HOME/.zetacored/config/genesis.json
  
# give balance to runner accounts to deploy contracts directly on zEVM
# deployer
  zetacored add-genesis-account zeta1uhznv7uzyjq84s3q056suc8pkme85lkvhrz3dd 100000000000000000000000000azeta
# erc20 tester
  zetacored add-genesis-account zeta1datate7xmwm4uk032f9rmcu0cwy7ch7kg6y6zv 100000000000000000000000000azeta
# zeta tester
  zetacored add-genesis-account zeta1tnp0hvsq4y5mxuhrq9h3jfwulxywpq0ads0rer 100000000000000000000000000azeta
# bitcoin tester
  zetacored add-genesis-account zeta19q7czqysah6qg0n4y3l2a08gfzqxydla492v80 100000000000000000000000000azeta
# ethers tester
  zetacored add-genesis-account zeta134rakuus43xn63yucgxhn88ywj8ewcv6ezn2ga 100000000000000000000000000azeta

# 3. Copy the genesis.json to all the nodes .And use it to create a gentx for every node
  zetacored gentx operator 1000000000000000000000azeta --chain-id=$CHAINID --keyring-backend=$KEYRING --gas-prices 20000000000azeta
  # Copy host gentx to other nodes
  for NODE in "${NODELIST[@]}"; do
    ssh $NODE mkdir -p ~/.zetacored/config/gentx/peer/
    scp ~/.zetacored/config/gentx/* $NODE:~/.zetacored/config/gentx/peer/
  done
  # Create gentx files on other nodes and copy them to host node
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
   pp=$(cat $HOME/.zetacored/config/gentx/z2gentx/*.json | jq '.body.memo' )
   pps=${pp:1:58}
   sed -i -e 's/^persistent_peers =.*/persistent_peers = "'$pps'"/' "$HOME"/.zetacored/config/config.toml
fi
# End of genesis creation steps . The steps below are common to all the nodes

# Update persistent peers
if [ $HOSTNAME != "zetacore0" ]
then
  # Misc : Copying the keyring to the client nodes so that they can sign the transactions
  ssh zetaclient"$INDEX" mkdir -p ~/.zetacored/keyring-test/
  scp ~/.zetacored/keyring-test/* "zetaclient$INDEX":~/.zetacored/keyring-test/
  ssh zetaclient"$INDEX" mkdir -p ~/.zetacored/keyring-file/
  scp ~/.zetacored/keyring-file/* "zetaclient$INDEX":~/.zetacored/keyring-file/

  pp=$(cat $HOME/.zetacored/config/gentx/peer/*.json | jq '.body.memo' )
  pps=${pp:1:58}
  sed -i -e "/persistent_peers =/s/=.*/= \"$pps\"/" "$HOME"/.zetacored/config/config.toml
fi

cosmovisor run start --pruning=nothing --minimum-gas-prices=0.0001azeta --json-rpc.api eth,txpool,personal,net,debug,web3,miner --api.enable --home /root/.zetacored