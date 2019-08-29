#!/usr/bin/env bash
set -ex
abigen --sol /root/TronToken.sol --pkg main --out=token.go
ls -la