package thebug

import (
	"context"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
)

type thebug struct {
	Next plugin.Handler
}
var _ plugin.Handler = (*thebug)(nil)

// Name
func (*thebug) Name() string { return "thebug" }

// ServeDNS
func (t *thebug) ServeDNS(c context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	return plugin.NextOrFailure(t.Name(), t.Next, c, w, r)
}
