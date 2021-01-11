package handler

import (
	"boxfox_grpc_client/framework/app"
	pb "boxfox_grpc_client/proto"
	"context"
	pbs "github.com/zer0131/boxfox_grpc_server/proto"
)

type BoxfoxGrpcClientServiceHandler struct{}

func (s *BoxfoxGrpcClientServiceHandler) Monitor(ctx context.Context, in *pb.MonitorReqNew) (*pb.MonitorRespNew, error) {
	if err := in.Validate(); err != nil {
		// 返回err，直接交给框架处理
		// 1. http: 调用方收到500，和错误消息体
		// 2. grpc-go: ...
		return nil, err
	}

	ret, _ := app.ConfigVal.GrpcTestService.Monitor(ctx, &pbs.MonitorReq{Name: "haha", Age: 12})
	return &pb.MonitorRespNew{
		Errno:  1,
		Errmsg: ret.Errmsg,
	}, nil
}
