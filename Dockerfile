ARG GO_VERSION=1.12
FROM golang:${GO_VERSION}
COPY . /transaction_info
WORKDIR /transaction_info
RUN go install
CMD transcation_info_etherium
