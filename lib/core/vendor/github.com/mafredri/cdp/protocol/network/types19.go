// +build go1.9

// Code generated by cdpgen. DO NOT EDIT.

package network

import (
	"github.com/mafredri/cdp/protocol/internal"
)

// CachedResource Information about the cached resource.
type CachedResource struct {
	URL      string                    `json:"url"`                // Resource URL. This is the url of the original network request.
	Type     internal.PageResourceType `json:"type"`               // Type of this resource.
	Response *Response                 `json:"response,omitempty"` // Cached response data.
	BodySize float64                   `json:"bodySize"`           // Cached response body size.
}
