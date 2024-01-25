#!/bin/sh

# ...... enode ......
enode=$(geth --exec "admin.nodeInfo.enode" attach --datadir node1)

# ...... enode
echo $enode > enode.txt
