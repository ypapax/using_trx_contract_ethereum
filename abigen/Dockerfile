ARG GO_VERSION=1.12
FROM golang:${GO_VERSION}
RUN go get github.com/ethereum/go-ethereum/...
WORKDIR $GOPATH/src/github.com/ethereum/go-ethereum
RUN make
RUN make devtools
RUN apt-get update
#RUN apt-get install software-properties-common -y
#RUN add-apt-repository ppa:ethereum/ethereum
#RUN apt-get update
RUN apt-get install wget -y
WORKDIR $GOPATH/bin
RUN wget -O solc https://github.com/ethereum/solidity/releases/download/v0.4.25/solc-static-linux
#RUN wget -O solc https://github.com/ethereum/solidity/releases/download/v0.5.11/solc-static-linux
RUN chmod +x solc
COPY TronToken.sol /root/TronToken.sol
COPY entrypoint.sh /root/entrypoint.sh
RUN chmod +x /root/entrypoint.sh
WORKDIR /root
ENTRYPOINT ["/root/entrypoint.sh"]