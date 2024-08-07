#!/bin/bash

export CGO_CFLAGS="-I/home/nikhilsamrat/Downloads/rocksdb/include" 
export CGO_LDFLAGS="-L/home/nikhilsamrat/Downloads/rocksdb/build -lrocksdb -lstdc++ -lm -lz -lsnappy -llz4 -lzstd" 
export LD_LIBRARY_PATH=/home/nikhilsamrat/Downloads/rocksdb/build:$LD_LIBRARY_PATH
go build
