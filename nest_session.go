package nestsensors

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"jaytaylor.com/nestsensors/domain"
)

// NestSession holds the API session context.
type NestSession struct {
	User     string
	Password string
	Client   *http.Client
	Auth     *domain.Auth
}

// New returns a new nest sensors session.
func New(user string, password string) *NestSession {
	ns := &NestSession{
		User:     user,
		Password: password,
	}
	return ns
}

// Login logs in to home.nest.com and initiates a new session.
func (ns *NestSession) Login() error {
	client, auth, err := login(ns.User, ns.Password)
	if err != nil {
		return err
	}
	ns.Client = client
	ns.Auth = auth
	return nil
}

func (ns *NestSession) Raw() error {
	u := fmt.Sprintf("%v/api/0.1/user/%v/app_launch", BaseURL, ns.Auth.UserID)
	// TODO: this could be exposed via domain as the keys of the `keys` var.
	s := `{"known_bucket_types":["buckets","delayed_topaz","demand_response","device","device_alert_dialog","geofence_info","kryptonite","link","message","message_center","metadata","occupancy","quartz","safety","rate_plan","rcs_settings","safety_summary","schedule","shared","structure","structure_history","structure_metadata","topaz","topaz_resource","tou","track","trip","tuneups","user","user_alert_dialog","user_settings","where","widget_track"],"known_bucket_versions":[]}`

	body, err := ns.checkedPostJSON(u, strings.NewReader(s))
	if err != nil {
		return err
	}
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	// m := map[string]interface{}{}
	state := &domain.NestState{}
	if err := json.Unmarshal(data, state); err != nil {
		return err
	}

	tss, err := state.TemperatureSensors()
	if err != nil {
		return err
	}

	for _, ts := range tss {
		name, err := state.WhereToName(ts.Value.WhereID)
		if err != nil {
			return err
		}
		fmt.Printf("%v %+v\n", name, ts.Value)
	}

	// bs, err := json.MarshalIndent(state, "", "    ")
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%v\n", string(bs))

	return nil
}

// checkedGet requires an already authenticated *http.Client and retrieves
// content from the specified url.
//
// Automatically injects the "Authorization: Basic <access-token>" header.
//
// Caller is responsible for closing the returned response body.
func (ns *NestSession) checkedGet(url string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating logged-in GET request: %s", err)
	}
	applyHeaders(req)
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", ns.Auth.AccessToken))

	resp, err := ns.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getting %v: %s", url, err)
	}

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("expected 2xx response status-code from %v but got %v", url, resp.StatusCode)
	}

	return resp.Body, nil
}

func (ns *NestSession) checkedPostJSON(url string, body io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("creating logged-in POST request: %s", err)
	}
	applyHeaders(req)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", ns.Auth.AccessToken))

	resp, err := ns.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getting %v: %s", url, err)
	}

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("expected 2xx response status-code from %v but got %v", url, resp.StatusCode)
	}

	return resp.Body, nil
}

/*

 */
