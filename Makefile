# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: geth android ios evm all test clean

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

geth:
	$(GORUN) build/ci.go install ./cmd/geth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/geth\" to launch geth."

all:
	$(GORUN) build/ci.go install

test: all
	$(GORUN) build/ci.go test

lint: ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go install golang.org/x/tools/cmd/stringer@latest
	env GOBIN= go install github.com/fjl/gencodec@latest
	env GOBIN= go install github.com/golang/protobuf/protoc-gen-go@latest
	env GOBIN= go install ./cmd/abigen
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

NS=swh
VERSION=4_9
IMAGE=swr.cn-north-4.myhuaweicloud.com/buaa-dist-compute-ci/ethereum-recorder:${VERSION}
push:
	docker build -t ${IMAGE} -f build.Dockerfile .
	docker push ${IMAGE}

folder_path=docker_data
create:
	# if [ ! -d "$folder_path" ]; then
	# 	mkdir "$folder_path"
	# 	echo "Folder created: $folder_path"
	# else
	# 	echo "Folder already exists: $folder_path"
	# fi
	echo "20231110\n20231110" | ./build/bin/geth --datadir docker_data/node1/ account new
	echo "20231110\n20231110" | ./build/bin/geth --datadir docker_data/node2/ account new
	echo "20231110\n20231110" | ./build/bin/geth --datadir docker_data/node3/ account new
	echo "20231110\n20231110" | ./build/bin/geth --datadir docker_data/node4/ account new
	echo "20231110" > docker_data/node1/password.txt
	echo "20231110" > docker_data/node2/password.txt
	echo "20231110" > docker_data/node3/password.txt
	echo "20231110" > docker_data/node4/password.txt

clear:
	rm -rf docker_data/node*

result:
	kubectl cp ./acquire_enode.sh $(NS)/$(nodepod1):/usr/local/bin/acquire_enode.sh
	kubectl exec -it $(nodepod1) -n swh -- /bin/sh ./acquire_enode.sh
	kubectl cp $(NS)/$(nodepod1):/usr/local/bin/enode.txt ./enode.txt	# echo $(nodepod1)
	# echo $(res)

deploy: 
	# 关闭pod
	# make stop
	# 清除节点配置文件
	make clear
	# 生成节点配置文件
	make create
	# 更改节点配置、启动文件，创世区块文件
	# 处理节点1和P2P网络平均传输时延节点修改(暂未改)
	python modify.py false
	# build、push IMAGE
	make push
	# 部署节点1
	kubectl apply -f  deploy_node.yml -n $(NS)
	# 部署成功再执行后续操作
	# while true; do \
	# 	if [$(status) != "Running"]; then \
	# 		sleep 1; \
	# 	fi \
	# done
	sleep 10
	# 获取enode
	kubectl cp ./acquire_enode.sh $(NS)/$(nodepod1):/usr/local/bin/acquire_enode.sh
	kubectl exec -it $(nodepod1) -n swh -- /bin/sh ./acquire_enode.sh
	kubectl cp $(NS)/$(nodepod1):/usr/local/bin/enode.txt ./enode.txt
	# 修改配置	
	python modify.py true
	# 启动k8s
	kubectl apply -f  deploy_3node.yml -n $(NS)
backup:
	# 更新最新配置文件
	kubectl cp docker_data/start_node2.sh $(NS)/$(nodepod2):/usr/local/bin/start_node2.sh 
	kubectl cp docker_data/start_node3.sh $(NS)/$(nodepod3):/usr/local/bin/start_node3.sh 
	kubectl cp docker_data/start_node4.sh $(NS)/$(nodepod4):/usr/local/bin/start_node4.sh 
	# 初始化, 并添加节点1
	kubectl exec -it $(nodepod2) -n swh -- /bin/sh ./start_node2.sh
	kubectl exec -it $(nodepod3) -n swh -- /bin/sh ./start_node3.sh
	kubectl exec -it $(nodepod4) -n swh -- /bin/sh ./start_node4.sh

start:
	kubectl apply -f  deploy_3node.yml -n $(NS)

stop:
	kubectl delete -f  deploy_node.yml -n $(NS)
	kubectl delete -f  deploy_3node.yml -n $(NS)



output=$(shell kubectl get pod  -n ${NS})
status=$(word 8, $(output))
nodepod1=$(shell echo "$(output)" | grep -oP 'ethereum-test-node1-[\w-]+')
nodepod2=$(shell echo "$(output)" | grep -oP 'ethereum-test-node2-[\w-]+')
nodepod3=$(shell echo "$(output)" | grep -oP 'ethereum-test-node3-[\w-]+')
nodepod4=$(shell echo "$(output)" | grep -oP 'ethereum-test-node4-[\w-]+')
node1:
	kubectl exec -ti $(nodepod1) -n ${NS} /bin/sh
node2:
	kubectl exec -ti $(nodepod2) -n ${NS} /bin/sh
node3:
	kubectl exec -ti $(nodepod3) -n ${NS} /bin/sh
node4:
	kubectl exec -ti $(nodepod4) -n ${NS} /bin/sh

node:
	nohup geth --datadir node1/ --networkid 198324715 --syncmode full --unlock 0x9Ab65Ec2EC373693d50B508E5204B7C6456D7eC7 --password node1/password.txt >> system.log.1 2>&1 &
	nohup geth --datadir node2/ --networkid 198324715 --syncmode full --unlock 0x94ECbc164499BDB517A499f0DB80355dD2B39114 --password node2/password.txt >> system.log.1 2>&1 &
	nohup geth --datadir node3/ --networkid 198324715 --syncmode full --unlock 0x0ae90bb8ffc79fe785fde895341a64ff515db964 --password node3/password.txt >> system.log.1 2>&1 &
	nohup geth --datadir node4/ --networkid 198324715 --syncmode full --unlock 0xC2a46485B33cC21EEAbCfb4bED51Eb5673956288 --password node4/password.txt >> system.log.1 2>&1 &


	geth console --datadir node1/ --networkid 198324715 --syncmode full --unlock 0x9Ab65Ec2EC373693d50B508E5204B7C6456D7eC7 --password node1/password.txt
	geth console --datadir node2/ --networkid 198324715 --syncmode full --unlock 0x94ECbc164499BDB517A499f0DB80355dD2B39114 --password node2/password.txt