package networking

import (
	"net/url"
)

type LimiterServer interface {
	Start(*url.URL)
}
