package util

import (
	"crypto/tls"
	"net/http"
)

// Generate HTTP client
func GetHttpClient(insecure bool) *http.Client {
	// Set insecure or not
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{ InsecureSkipVerify: insecure },
	}
	return &http.Client{Transport: tr}

}
