package transaport

import (
	_ "github.com/go-kit/kit/circuitbreaker"
	_ "github.com/go-kit/kit/log"
	_ "github.com/go-kit/kit/ratelimit"
	_ "github.com/go-kit/kit/tracing/opentracing"
	_ "github.com/go-kit/kit/tracing/zipkin"
	_ "github.com/go-kit/kit/transport"
	"github.com/go-kit/kit/transport/grpc"
	_ "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	Add grpc.Handler
}