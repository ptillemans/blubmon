#!/usr/bin/env bash

export GOOS=linux
export GOARCH=arm
export GOARM=5
go build ./main.go
scp main debian@blub.snamellit.com:blubmon
ssh debian@blub.snamellit.com ./blubmon
