package metrics

import (
	"github.com/coredns/coredns/plugin"

	"golang.org/x/net/context"
)

// WithServer returns the current server handling the request. It returns the
// server listening address: <scheme>://[<bind>]:<port> Normally this is
// something like "dns://:53", but if the bind plugin is used, i.e. "bind
// 127.0.0.53", it will be "dns://127.0.0.53:53", etc. If not address is found
// the empty string is returned.
//
// Basic usage with a metric:
//
// <metric>.WithLabelValues(metrics.WithServer(ctx), labels..).Add(1)
func WithServer(ctx context.Context) string {
	srv := ctx.Value(plugin.ServerCtx{})
	if srv == nil {
		return ""
	}
	return srv.(string)
}
