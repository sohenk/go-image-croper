package server

import (
	v1 "imgcropper/api/imgcropper/service"
	"imgcropper/internal/conf"
	"imgcropper/internal/service"

	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.ImgcropperService, logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			ratelimit.Server(),
			validate.Validator(),
			tracing.Server(),
		),
		http.ResponseEncoder(responseEncoder),
		http.ErrorEncoder(responseErrorEncoder),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterImgcropperHTTPServer(srv, greeter)
	return srv
}
