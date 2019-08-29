ARG GO_VERSION=1.12
FROM golang:${GO_VERSION}
COPY . /trx_contract_usage
WORKDIR /trx_contract_usage
RUN go install
CMD trx_contract_usage
