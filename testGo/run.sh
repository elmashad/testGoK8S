#!/usr/bin/env bash
git config --global --add url."git@github.com:".insteadOf "https://github.com/"
go get -v github.com/m-amr/gin

go get -v ./...
gin -i -p 3333 run main.go