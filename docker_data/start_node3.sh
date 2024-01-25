geth init --datadir node3/ genesis.json
nohup geth --datadir node3/ --networkid 198324715 --syncmode full --unlock 0xbfa8a5a4476ef8b92decdbf952dc8c3f2a1c1943 --password node3/password.txt >> system.log.1 2>&1 &
sleep 5
/usr/local/bin/geth attach --datadir node3 <<EOF
admin.addPeer("enode://3ece85ba254beebff873de70b72d4d17bb1f2ca7bb828a22207cf4f34a5303ef9973a73840235c098fa5ea248bf9468d4f69f0eb3f44ee195bc5ed68ffd18cc6@124.70.47.171:30303")
EOF