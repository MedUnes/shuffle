#!/bin/bash
COMPILE_DIR=./.shuffle/data/build
go build -ldflags="-s -w" -o "${COMPILE_DIR}/shuffle"