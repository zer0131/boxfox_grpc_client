package app

import (
	"context"
	demo "github.com/zer0131/boxfox_grpc_server/proto"
	"github.com/zer0131/toolbox/log"
	"google.golang.org/grpc"
	"gopkg.in/ini.v1"
	"strconv"
	"strings"
	"time"
)

var ConfigVal = &Config{}

func LoadConfig(ctx context.Context, path string) error {
	cfg, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, path)

	if err != nil {
		log.Errorf(ctx, "err=%s", err.Error())
		return err
	}
	ConfigVal.BaseVal.Port, err = cfg.Section("base").Key("port[int]").Int64()
	if err != nil {
		log.Errorf(ctx, "config [port[int]] err. err=%s", err.Error())
		return err
	}
	ConfigVal.BaseVal.LogLevel = cfg.Section("base").Key("log_level").String()
	ConfigVal.BaseVal.LogSize, err = cfg.Section("base").Key("log_size[int]").Int64()
	if err != nil {
		log.Errorf(ctx, "config [log_size[int]] err. err=%s", err.Error())
		return err
	}
	ConfigVal.BaseVal.Group = cfg.Section("base").Key("group").String()
	ConfigVal.BaseVal.Project = cfg.Section("base").Key("project").String()
	ConfigVal.BaseVal.DisableWebServer, err = cfg.Section("base").Key("disable_web_server[bool]").Bool()
	if err != nil {
		log.Errorf(ctx, "config [disable_web_server[bool]] err. err=%s", err.Error())
		return err
	}
	ConfigVal.BaseVal.Type = cfg.Section("base").Key("type").String()
	ConfigVal.GrpcTestVal.Addr = cfg.Section("grpc_test").Key("addr").String()
	ConfigVal.GrpcTestVal.ConnTimeout, err = cfg.Section("grpc_test").Key("conn_timeout[int]").Int64()
	if err != nil {
		log.Errorf(ctx, "config [conn_timeout[int]] err. err=%s", err.Error())
		return err
	}
	ConfigVal.GrpcTestVal.Timeout, err = cfg.Section("grpc_test").Key("timeout[int]").Int64()
	if err != nil {
		log.Errorf(ctx, "config [timeout[int]] err. err=%s", err.Error())
		return err
	}
	ConfigVal.GrpcTestVal.MaxRecvMsgSize, err = cfg.Section("grpc_test").Key("max_recv_msg_size[int]").Int64()
	if err != nil {
		log.Errorf(ctx, "config [max_recv_msg_size[int]] err. err=%s", err.Error())
		return err
	}
	ConfigVal.GrpcTestVal.Service = cfg.Section("grpc_test").Key("service").String()
	ConfigVal.GrpcTestVal.ImportPath = cfg.Section("grpc_test").Key("import_path").String()

	if err := InitProjectConfig(ctx, path); err != nil {
		return err
	}

	grpcTestConn, err := grpc.Dial(ConfigVal.GrpcTestVal.Addr, grpc.WithInsecure(), grpc.WithTimeout(time.Duration(ConfigVal.GrpcTestVal.ConnTimeout)*time.Millisecond), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(int(ConfigVal.GrpcTestVal.MaxRecvMsgSize)*1024*1024)))
	if err != nil {
		log.Errorf(ctx, "err=%s", err.Error())
		return err
	}
	ConfigVal.GrpcTestService = demo.NewBoxfoxGrpcServerServiceClient(grpcTestConn)

	return nil

}

func InitProjectConfig(ctx context.Context, path string) error {

	cfg, err := ini.LoadSources(ini.LoadOptions{
		IgnoreInlineComment: true,
	}, path)

	if err != nil {
		log.Errorf(ctx, "err=%s", err.Error())
		return err
	}
	ConfigVal.BoxfoxGrpcClientVal.A, err = cfg.Section("boxfox_grpc_client").Key("a[int]").Int64()
	if err != nil {
		log.Errorf(ctx, "config [a[int]] err. err=%s", err.Error())
		return err
	}
	tB := cfg.Section("boxfox_grpc_client").Key("b[int_array]").String()
	tBIntArr := strings.Split(tB, ",")
	for _, v := range tBIntArr {
		vInt, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Errorf(ctx, "err=%s", err.Error())
			return err
		}
		ConfigVal.BoxfoxGrpcClientVal.B = append(ConfigVal.BoxfoxGrpcClientVal.B, vInt)
	}
	ConfigVal.BoxfoxGrpcClientVal.C = cfg.Section("boxfox_grpc_client").Key("c[string]").String()
	tD := cfg.Section("boxfox_grpc_client").Key("d[string_array]").String()
	ConfigVal.BoxfoxGrpcClientVal.D = strings.Split(tD, ",")
	if err != nil {
		log.Errorf(ctx, "err=%s", err.Error())
		return err
	}
	ConfigVal.BoxfoxGrpcClientVal.E, err = cfg.Section("boxfox_grpc_client").Key("e[bool]").Bool()
	if err != nil {
		log.Errorf(ctx, "config [e[bool]] err. err=%s", err.Error())
		return err
	}

	return nil

}
