#!/bin/bash
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I${GOPATH}/src/github.com/tvducmt/protoc-gen-copy \
  --go_out=plugins=grpc:. \
  --copy_out="lang=go:." \
  *.proto 
  

#--grpc-gateway_out=logtostderr=true:. \