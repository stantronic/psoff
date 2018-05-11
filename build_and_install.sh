#!/usr/bin/env bash
echo -e "Building psoff binary"
go build
echo -e "Symlinking $(pwd)/psoff to /usr/local/bin"
ln -sf $(pwd)/psoff /usr/local/bin
