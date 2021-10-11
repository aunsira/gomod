package actions

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/aunsira/gomod/config"
)

const (
	ModerationPath = "/api/v1/moderations"
)

// Example:
//
//  imgData, get := &posmoni.GetModeration{}, &actions.GetModeration{
//      ID: "5a546e916e11571f570c1533",
//  }
//
//  if err := client.Call(imgData, get); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Moderation: %#v\n", imgData)
//
type GetModeration struct {
	ID string
}

// Example:
//
//  list, get := &posmoni.GetModerations{}, &actions.GetModerations{
//      ID: "5a546e916e11571f570c1533",
//      Page: 1,
//      Item: 20,
//  }
//
//  if err := client.Call(list, get); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Moderation: %#v\n", list)
//  fmt.Printf("First element: %#v\n", list.Data.Images[0])
//
type GetModerations struct {
	ID   string
	Page string
	Item string
}

// Example:
//
//  imgData, post := &posmoni.PostModeration{}, &actions.PostModeration{
//		Data: TestImageDataURL,
//  }
//
//  if err := client.Call(imgData, post); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Moderation: %#v\n", imgData)
//
type PostModeration struct {
	Data           string
	PostbackURL    string
	PostbackMethod string
	CustomID       string
}

// Endpoint returns Posmoni's request url, verb and endpoint for calling GET Moderation API.
func (g *GetModeration) Endpoint() (string, string, string) {
	return config.PosmoniAPIURL, "GET", ModerationPath
}

// Endpoint returns Posmoni's request url, verb and endpoint for calling Get list of
// Moderation API.
func (g *GetModerations) Endpoint() (string, string, string) {
	return config.PosmoniAPIURL, "GET", ModerationPath
}

// Endpoint returns Posmoni's request url, verb and endpoint for calling Create
// Moderation API.
func (p *PostModeration) Endpoint() (string, string, string) {
	return config.PosmoniAPIURL, "POST", ModerationPath
}

// Payload creates request's payload for Get Moderation API. Returns http.Request
// object which contains required query parameters.
func (g *GetModeration) Payload(endpoint, method, path string) (*http.Request, error) {
	if g.ID == "" {
		return nil, errors.New("ID is required ")
	}

	req, err := http.NewRequest(method, string(endpoint)+path+"/"+g.ID, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Payload creates request's payload for Get list Moderation API. Returns
// http.Request which contains required query parameters.
func (g *GetModerations) Payload(endpoint, method, path string) (*http.Request, error) {
	req, err := http.NewRequest(method, string(endpoint)+path, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if g.ID != "" {
		q.Add("id", g.ID)
	}
	if g.Page != "" {
		q.Add("page", g.Page)
	}
	if g.Item != "" {
		q.Add("per_page", g.Item)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

// Payload creates request's payload for Create Moderation API. Returns
// http.Request which contains required query parameters.
func (p *PostModeration) Payload(endpoint, method, path string) (*http.Request, error) {
	values := url.Values{}
	if p.Data != "" {
		values.Set("data", p.Data)
	}
	if p.PostbackURL != "" {
		values.Set("postback_url", p.PostbackURL)
	}
	if p.PostbackMethod != "" {
		values.Set("postback_method", p.PostbackMethod)
	}
	if p.CustomID != "" {
		values.Set("custom_id", p.CustomID)
	}

	body := strings.NewReader(values.Encode())
	req, err := http.NewRequest(method, string(endpoint)+path, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}
