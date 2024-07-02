package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// 客户端拦截器
func TraceClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		m := make(map[string]string)
		traceId := ctx.Value("traceid")
		if traceId != nil {
			m["traceid"] = traceId.(string)
		}
		zhsId := ctx.Value("zhsid")
		if zhsId != nil {
			m["zhsid"] = zhsId.(string)
		}
		userId := ctx.Value("userid")
		if userId != nil {
			m["userid"] = userId.(string)
		}
		classroomId := ctx.Value("classroomid")
		if classroomId != nil {
			m["classroomid"] = classroomId.(string)
		}
		userRealName := ctx.Value("userrealname")
		if userRealName != nil {
			m["userrealname-bin"] = userRealName.(string)
		}
		if len(m) > 0 {
			ctx = metadata.NewOutgoingContext(ctx, metadata.New(m))
		}

		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}
