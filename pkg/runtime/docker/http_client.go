package docker

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

const defaultTimeout = 30 * time.Second

// HTTPClient create new http client to connect to the docker daemon
func HTTPClient(daemonURL string, tlsConfig *tls.Config) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newHTTPClient(address *url.URL, tlsConfig *tls.Config, timeout time.Duration) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:wrapcheck // pass dialer error as-is

//nolint:wrapcheck // pass dialer error as-is

// Override the main URL object so the HTTP lib won't complain
