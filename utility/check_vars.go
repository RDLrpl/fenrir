package utility

import (
	"context"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

func Check_proxy(proxyAddr string, proxyLogin string, proxyPass string) (string, string, string, bool) {
	cleanProxy := strings.TrimSpace(proxyAddr)
	cleanLogin := strings.TrimSpace(proxyLogin)
	cleanPass := strings.TrimSpace(proxyPass)

	if cleanProxy == "" {
		return "", "", "", false
	}

	if !strings.Contains(cleanProxy, "://") {
		cleanProxy = "socks5://" + cleanProxy
	}

	parsedURL, err := url.Parse(cleanProxy)
	if err != nil {
		return "", "", "", false
	}

	var auth *proxy.Auth
	if cleanLogin != "" && cleanPass != "" {
		auth = &proxy.Auth{User: cleanLogin, Password: cleanPass}
	}

	dialer, err := proxy.SOCKS5("tcp", parsedURL.Host, auth, proxy.Direct)
	if err != nil {
		return "", "", "", false
	}

	contextDialer, ok := dialer.(proxy.ContextDialer)
	if !ok {
		return "", "", "", false
	}

	client := &http.Client{
		Transport: &http.Transport{DialContext: contextDialer.DialContext},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://ifconfig.me", nil)
	if err != nil {
		return "", "", "", false
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	resp, err := client.Do(req)
	if err != nil {
		return "", "", "", false
	}
	defer resp.Body.Close()

	return parsedURL.Host, cleanLogin, cleanPass, resp.StatusCode == http.StatusOK
}

func Check_CAAU() (bool, string) {
	if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
		return true, "./fpkg/CAAU/linuxamd/chrome-linux/chrome"
	}

	if runtime.GOOS == "windows" && runtime.GOARCH == "amd64" {
		return true, "./fpkg/CAAU/windowsamd/chrome-windows/chrome-windows/chrome.exe"
	} else {
		return false, ""
	}
}
