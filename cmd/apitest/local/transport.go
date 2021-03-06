// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package local

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/moov-io/base/http/bind"
)

// Transport intercepts HTTP requests and re-writes them according to bind.HTTP's local port binds.
// This is done to provide an shared http.RoundTripper usable by clients wishing for local dev with Moov.
//
// The underlying http.RoundTripper is required to enforce timeouts and other config be non-default
// and so we don't hack into http.DefaultClient (which has no timeouts).
type Transport struct {
	Underlying http.RoundTripper

	Debug bool
}

// RoundTrip modifies the incoming request to reshape a Moov API production URL to a local dev URL.
// The GoDoc for http.RoundTripper state that the request SHOULD not be modified, not MUST. If this
// ends up causing problems we'll have to figure out another solution.
//
// This means:
//  - Dropping /v1/$app routing prefix
//  - Changing the local port used (each app runs on its own port now)
//    - Adjusting the scheme if needed.
func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	origURL := r.URL.String()

	// Each route looks like /v1/$app/... so we need to trim off the v1 and $app segments
	// while looking up $app's port mapping.
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 { // parts splits into: "", v1, $app, (rest of url)
		// Pass through whatever this request is.
		return t.Underlying.RoundTrip(r)
	}

	r.URL.Scheme = "http"
	r.URL.Host = "localhost"
	r.URL.Path = "/"

	// Routing logic, should match Ingress routes
	switch strings.ToLower(parts[2]) {
	case "ach":
		switch strings.ToLower(parts[3]) {
		case "depositories", "gateways", "originators", "receivers", "transfers":
			r.URL.Host += bind.HTTP("paygate")
		default:
			r.URL.Host += bind.HTTP("ach")
		}
	case "auth":
		r.URL.Host += bind.HTTP("auth")
	case "fed":
		r.URL.Host += bind.HTTP("fed")
		if len(parts) > 3 && parts[3] != "ping" {
			r.URL.Path = "/fed/" // fed expects /fed/ as a prefix on routes on non-ping routes
		}
	case "accounts", "gl":
		r.URL.Host += bind.HTTP("accounts")
		r.URL.Path = "/accounts"
		if len(parts) > 3 {
			if parts[3] == "ping" {
				r.URL.Path = "/"
			} else {
				r.URL.Path += "/"
			}
		}
	case "customers":
		if len(parts) >= 5 && strings.EqualFold(parts[4], "accounts") {
			r.URL.Host += bind.HTTP("accounts")
		} else {
			r.URL.Host += bind.HTTP("customers")
		}
		r.URL.Path = "/customers"
		if len(parts) > 3 {
			r.URL.Path += "/"
		}
	case "paygate":
		r.URL.Host += bind.HTTP("paygate")
	case "oauth2":
		r.URL.Host += bind.HTTP("auth")
		r.URL.Path = "/oauth2/"
	case "ofac", "watchman":
		r.URL.Host += bind.HTTP("watchman")
	case "users":
		r.URL.Host += bind.HTTP("auth")
		r.URL.Path = "/users/"
	}

	r.URL.Path += strings.Join(parts[3:], "/") // everything after $app

	if t.Debug {
		log.Printf("%v %v request URL (Original: %v)", r.Method, r.URL.String(), origURL)
	}

	if t.Underlying == nil {
		return nil, errors.New("nil underlying Transport")
	}
	return t.Underlying.RoundTrip(r)
}
