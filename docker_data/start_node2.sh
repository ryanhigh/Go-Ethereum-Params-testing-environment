#!/bin/sh
geth init --datadir node2/ genesis.json
nohup geth --datadir node2/ --networkid 198324715 --syncmode full --unlock 0xab9605ed5b7dc5c8283fc4f6d83a34d583e37ec5 --password node2/password.txt >> system.log.1 2>&1 &
sleep 5
/usr/local/bin/geth attach --datadir node2 <<EOF
admin.addPeer("enode://3ece85ba254beebff873de70b72d4d17bb1f2ca7bb828a22207cf4f34a5303ef9973a73840235c098fa5ea248bf9468d4f69f0eb3f44ee195bc5ed68ffd18cc6@124.70.47.171:30303")
EOF