#!/bin/bash

/usr/sbin/sshd

HOSTNAME=$(hostname)

cp  /root/preparams/PreParams_$HOSTNAME.json /root/preParams.json
num=$(echo $HOSTNAME | tr -dc '0-9')
node="zetacore$num"
#mv  /root/zetacored/zetacored_$node /root/.zetacored
#mv /root/tss/$HOSTNAME /root/.tss

echo "Wait for zetacore to exchange genesis file"
sleep 30
operator=$(cat $HOME/.zetacored/os.json | jq '.ObserverAddress' )
operatorAddress=$(echo "$operator" | tr -d '"')
echo "operatorAddress: $operatorAddress"
echo "Start zetaclientd"
if [ $HOSTNAME == "zetaclient0" ]
then
    rm ~/.tss/*
    export TSSPATH=~/.tss
    zetaclientd init  \
      --pre-params ~/preParams.json  --zetacore-url zetacore0 \
      --chain-id athens_101-1  --operator "$operatorAddress" --log-level 0 --hotkey=hotkey
    zetaclientd start
else
  num=$(echo $HOSTNAME | tr -dc '0-9')
  node="zetacore$num"
  SEED=$(curl --retry 10 --retry-delay 5 --retry-connrefused  -s zetaclient0:8123/p2p)
  rm ~/.tss/*
  export TSSPATH=~/.tss
  zetaclientd init  \
    --peer /ip4/172.20.0.21/tcp/6668/p2p/$SEED \
    --pre-params ~/preParams.json --zetacore-url $node \
    --chain-id athens_101-1 --operator "$operatorAddress" --log-level 0 --hotkey=hotkey
  zetaclientd start
fi
