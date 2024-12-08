package utils

import (
	"net/url"
	"errors"
	"net"
)

// ValidateURL checks if a URL is valid
func ValidateURL(rawURL string) error {
	// Parse the URL
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return errors.New("invalid URL format")
	}

	// Check for allowed protocols
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("URL must start with http:// or https://")
	}

	// Optional: Check if the domain exists
	hostname := parsedURL.Hostname()
	_, err = net.LookupIP(hostname)
	if err != nil {
		return errors.New("domain does not exist")
	}

	return nil
}