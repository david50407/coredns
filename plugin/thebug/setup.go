package thebug

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
)

var log = clog.NewWithPlugin("thebug")

func init() { plugin.Register("thebug", setup) }

func setup(c *caddy.Controller) error {
	instance := thebugParse(c)

	c.OnStartup(func () error {
		log.Infof("Listing Handlers, this should match the order in plugin.cfg:")
		for _, h := range dnsserver.GetConfig(c).Handlers() {
			log.Infof("> Handler: %s", h.Name())
		}
		return nil
	})

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		instance.Next = next
		return instance
	})

	return nil
}

func thebugParse(*caddy.Controller) *thebug {
	return &thebug{}
}
