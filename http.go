package nestsensors

import (
	"net/http"
	"strings"
)

func newClient() *http.Client {
	client := &http.Client{
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return client
}

// applyHeaders adds standard browser headers to the specified HTTP request.
func applyHeaders(req *http.Request) {
	req.Header.Set("Referer", BaseURL+"/")
	req.Header.Set("Origin", BaseURL+"/")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml,application/json;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Authority", strings.Split(BaseURL, "://")[1])
	req.Header.Set("User-Agent", UserAgent)
}
