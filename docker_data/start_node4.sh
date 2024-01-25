geth init --datadir node4/ genesis.json
nohup geth --datadir node4/ --networkid 198324715 --syncmode full --unlock 0xe050845fa06d45b84112df25fef592d813afbd13 --password node4/password.txt >> system.log.1 2>&1 &
sleep 5
/usr/local/bin/geth attach --datadir node4 <<EOF
admin.addPeer("enode://3ece85ba254beebff873de70b72d4d17bb1f2ca7bb828a22207cf4f34a5303ef9973a73840235c098fa5ea248bf9468d4f69f0eb3f44ee195bc5ed68ffd18cc6@124.70.47.171:30303")
EOF