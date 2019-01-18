#!/usr/bin/env bash

go test ./... -coverprofile=coverage.txt -covermode=atomic
go tool cover -html=coverage.txt