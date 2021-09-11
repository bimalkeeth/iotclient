package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
)

func InstrumentingMiddleware(duration metrics.Histogram) endpoint.Middleware {

	return nil
}