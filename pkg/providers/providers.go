package providers

import (
	"context"

	"github.com/valyala/fasthttp"
)

const Version = `2.1.2`

// Provider is a generic interface for all archive fetchers
type Provider interface {
	Fetch(ctx context.Context, domain string, results chan string) error
	Name() string
}

type URLScan struct {
	Host   string
	APIKey string
}

type Config struct {
	Threads           uint
	Timeout           uint
	Verbose           bool
	MaxRetries        uint
	IncludeSubdomains bool
	RemoveParameters  bool
	Client            *fasthttp.Client
	Providers         []string
	Blacklist         map[string]struct{}
	Outfile           string
	JSON              bool
	URLScan           URLScan
	Domains           []string
	OTX               string
}
