package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/draios/shared-go/pkg/sdauth"
	pkghttp "github.com/draios/secure-backend/pkg/http"
)

func getAuthClient(authServiceEndpoint string) (*sdauth.Client, error) {
	httpClient, err := getMonitorBackendHTTPClient()
	if err != nil {
		return nil, errors.New("cannot create http client for legacy services")
	}

	beURL, err := url.Parse(authServiceEndpoint)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot create parse endpoint '%s', %v", authServiceEndpoint, err))
	}

	config := &sdauth.Config{
		Endpoint:   beURL,
		HTTPClient: httpClient,
	}

	return sdauth.New(config), nil
}

func getMonitorBackendHTTPClient() (*http.Client, error) {
	tlsSkipCheck := os.Getenv("MONITOR_BACKEND_TLS_SKIP_CHECK") != ""
	transport, err := pkghttp.BuildTransport(tlsSkipCheck)
	if err != nil {
		return nil, err
	}

	return &http.Client{
		Transport: transport,
	}, nil
}

