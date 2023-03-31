# rpc-interface
rpc服务接口

go mod init github.com/xlizy/rpc-interface

protoc --go_out=. --go-triple_out=. ./MsgPushMessage.proto