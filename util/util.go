package util

import (
	"crypto/tls"
	"github.com/lucas-clemente/quic-go/http3"
	"net/http"
	"net/url"
)

// Generate HTTP client
func GetHttpClient(insecure bool, usesHttp3 bool) *http.Client {
	// Set insecure or not
	tlsClientConfig := &tls.Config{InsecureSkipVerify: insecure}
	var tr http.RoundTripper
	if usesHttp3 {
		tr = &http3.RoundTripper{TLSClientConfig: tlsClientConfig}
	} else {
		tr = &http.Transport{
			TLSClientConfig: tlsClientConfig,
		}
	}
	return &http.Client{Transport: tr}
}

// Fill "https" scheme when the given url has no scheme
func FillHttps(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" {
		u.Scheme = "https"
	}
	return u.String(), nil
}
