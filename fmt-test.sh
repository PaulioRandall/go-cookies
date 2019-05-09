#!/bin/bash

gofmt -w -s ./pkg
go test ./pkg