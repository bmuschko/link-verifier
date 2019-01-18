#!/usr/bin/env bash

go test ./... -coverprofile=coverage.txt -covermode=count
go tool cover -html=coverage.txt