package rest

import (
	"net/url"
	"path"
)

func DefaultServerURL(host, apiPath string, groupVersion string) (*url.URL, string, error) {
	hostURL, _ := url.Parse(host)

	versionedAPIPath := path.Join("/", apiPath, groupVersion)

	return hostURL, versionedAPIPath, nil
}

func defaultServerURLFor(config *Config) (*url.URL, string, error) {
	return DefaultServerURL(config.Host, config.APIPath, config.GroupVersion)
}
