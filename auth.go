package nestsensors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	"jaytaylor.com/nestsensors/domain"
)

var (
	// BaseURL holds the base URL for HackerNews.
	BaseURL = "https://home.nest.com"

	// UserAgent string to use for the UserAgent header value.
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.162 Safari/537.36"
)

// login returns an authenticated *http.Client (or errors out).
//
// Mimics:
//
//     curl "https://home.nest.com/session" \
//         -H "Authorization: Basic" \
//         -H "Content-Type: application/json" \
//         -H "Referer: https://home.nest.com/" \
//         --data-binary '{"email":"<Email>","password":"<Password>"}' -v
//
func login(username string, password string) (*http.Client, *domain.Auth, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, nil, fmt.Errorf("creating cookie jar: %s", err)
	}

	creds := map[string]string{
		"email":    username,
		"password": password,
	}
	postData, err := json.Marshal(creds)
	if err != nil {
		return nil, nil, err
	}

	client := newClient()
	client.Jar = jar

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/session", BaseURL), bytes.NewReader(postData))
	if err != nil {
		return nil, nil, fmt.Errorf("login: creating POST request: %s", err)
	}

	applyHeaders(req)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("login: %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode/100 != 2 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, nil, fmt.Errorf("login: expected 2xx reponse status-code but got %v (body=%v)", resp.StatusCode, string(body))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	auth := &domain.Auth{}
	if err := json.Unmarshal(data, auth); err != nil {
		return nil, nil, err
	}
	return client, auth, nil
}
