FROM golang:1.19-alpine as builder
RUN apk add --no-cache gcc musl-dev linux-headers git
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn
RUN go env
COPY go.mod /go-ethereum/
COPY go.sum /go-ethereum/
# COPY cmd/recorderfile /go-ethereum/cmd/recorderfile
ADD . /go-ethereum
RUN cd /go-ethereum && go mod download
# ADD . /go-ethereum
RUN cd /go-ethereum && go run build/ci.go install -static ./cmd/geth
# Pull Geth into a second stage deploy alpine container
FROM alpine:latest
RUN apk add --no-cache ca-certificates 
RUN apk add --no-cache netcat-openbsd
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/
COPY cmd/recorderfile /go-ethereum/cmd/
COPY docker_data/ /usr/local/bin/
RUN chmod 777 /usr/local/bin/start_node1.sh
RUN chmod 777 /usr/local/bin/start_node2.sh
RUN chmod 777 /usr/local/bin/start_node3.sh
RUN chmod 777 /usr/local/bin/start_node4.sh


# admin.addPeer("enode://6b3859d6e66bdcc7c898a9b4b435499b0b5d91d695ea56125c69e3c7dd39618b95433ffc4d636a3edadc7511bf0e15ce133a37efb9425dfe22cc16aada062b21@121.36.87.222:30303")
# net.peerCount

# admin.peers

# admin.nodeInfo.enode

# admin.addPeer("enode://77b99d57e08533d4a9c18c2c57ef0a746282c2bec5ea49f0bf9896f8bc02201f15114b907bfc51c9608898da5db14becbb8a888da2636b6ac76c43f766e38416@124.70.47.171:30303")