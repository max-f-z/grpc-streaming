//go:generate protoc -I . --go_out=plugins=grpc:. srvstream.proto

package srvstream
