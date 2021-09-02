#!/bin/zsh

echo "----------- Removing deps ----------------"
rm sampleCRUD

echo "---------------------------------------"
echo "Setting env vars"
export GOSUMDB=off
export GOBIN=$PWD
export GIT_TERMINAL_PROMPT=1

echo "---------------------------------------"
echo "Building execs"
go install src/apps/SampleCRUD.go