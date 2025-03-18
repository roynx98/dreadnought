package framework

import (
	"net/url"
)

type LimiterServer interface {
	Start(*url.URL)
}
