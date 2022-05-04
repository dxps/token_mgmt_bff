#!/bin/sh

## -------------------------------------------------------------------------------
## This script restarts the main on code changes using the reflex tool.
## If you don't have it already, install it using one of the options listed below.
## While being outside of this (and any other Go Modules based) project directory,
## run either:
##   - go install github.com/cespare/reflex@latest
##     if you have Go 1.16 or newer
##   - go get github.com/cespare/reflex 
##     if you have Go 1.15 or older
## Also, include $HOME/go/bin (or whatever your GOBIN env var
## is explicitly defined) to your PATH env var.
## -------------------------------------------------------------------------------

reflex -d none -r '\.go' -s -t 8000ms -- sh -c "go run ./cmd/token_mgmt_bff/main.go"
