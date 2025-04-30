package zoneedit

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/mrenouf/libdns-zoneedit"
)

type Provider struct{ *zoneedit.Provider }

func init() { caddy.RegisterModule(Provider{}) }

func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.zoneedit",
		New: func() caddy.Module { return &Provider{new(zoneedit.Provider)} },
	}
}

func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Username = d.Val()
		}
		if d.NextArg() {
			p.Token = d.Val()
		}
	}
	return nil
}
