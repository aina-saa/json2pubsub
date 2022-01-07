#!/usr/bin/env bash
version="0.0.3"
time=$(date)
hash=$(git log -n 1 --pretty=format:"%H")
hostname=$(hostname)
go build -ldflags="\
    -X 'github.com/aina-saa/json2pubsub/version.BuildVersion=$version' \
    -X 'github.com/aina-saa/json2pubsub/version.BuildTime=$time' \
    -X 'github.com/aina-saa/json2pubsub/version.BuildHost=$hostname' \
    -X 'github.com/aina-saa/json2pubsub/version.BuildSha=$hash'" .
./json2pubsub --version

# eof