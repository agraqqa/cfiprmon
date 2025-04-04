package main

import (
	"hash/crc32"
	"fmt"
	"io"
	"net/http"
)

const (
	CloudflareIPv4URL = "https://www.cloudflare.com/ips-v4/"
	CloudflareIPv6URL = "https://www.cloudflare.com/ips-v6/"
)

// getFileChecksum returns the CRC32 checksum of the file at the specified URL
func getFileChecksum(url string) (uint32, error) {
    client := &http.Client{}

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return 0, fmt.Errorf("error creating request: %v", err)
    }

    resp, err := client.Do(req)
    if err != nil {
        return 0, fmt.Errorf("error downloading file: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    hasher := crc32.NewIEEE()
    if _, err := io.Copy(hasher, resp.Body); err != nil {
        return 0, fmt.Errorf("error calculating checksum: %v", err)
    }

    return hasher.Sum32(), nil
}

// GetCloudflareIPRangesChecksum returns the MD5 hash of the Cloudflare IP ranges file for the specified URL
func GetCloudflareIPRangesChecksum(url string) (uint32, error) {
	checksum, err := getFileChecksum(url)
	if err != nil {
		return 0, fmt.Errorf("error getting IP ranges hash from %s: %v", url, err)
	}
	return checksum, nil
}
