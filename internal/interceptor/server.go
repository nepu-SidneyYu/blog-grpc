package interceptor

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TraceServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// 获取metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("miss metadata")
		}
		traceId := ""
		if val, ok := md["traceid"]; ok {
			traceId = val[0]
		}
		ctx = context.WithValue(ctx, "traceid", traceId)

		zhsId := ""
		if val, ok := md["zhsid"]; ok {
			zhsId = val[0]
		}
		ctx = context.WithValue(ctx, "zhsid", zhsId)

		userId := ""
		if val, ok := md["userid"]; ok {
			userId = val[0]
		}
		ctx = context.WithValue(ctx, "userid", userId)

		classroomId := ""
		if val, ok := md["classroomid"]; ok {
			classroomId = val[0]
		}
		ctx = context.WithValue(ctx, "classroomid", classroomId)

		userRealName := ""
		if val, ok := md["userrealname-bin"]; ok {
			userRealName = val[0]
		}
		ctx = context.WithValue(ctx, "userrealname", userRealName)

		// 继续处理请求
		return handler(ctx, req)
	}
}
