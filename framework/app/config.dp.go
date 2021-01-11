package app

import (
	demo "github.com/zer0131/boxfox_grpc_server/proto"
)

type autoBase struct {
	Port             int64
	LogLevel         string
	LogSize          int64
	Group            string
	Project          string
	DisableWebServer bool
	Type             string
}
type autoGrpcTest struct {
	Addr           string
	ConnTimeout    int64
	Timeout        int64
	MaxRecvMsgSize int64
	Service        string
	ImportPath     string
}
type autoBoxfoxGrpcClient struct {
	A int64
	B []int64
	C string
	D []string
	E bool
}
type Config struct {
	BaseVal             autoBase
	BoxfoxGrpcClientVal autoBoxfoxGrpcClient
	GrpcTestService     demo.BoxfoxGrpcServerServiceClient
	GrpcTestVal         autoGrpcTest
}
