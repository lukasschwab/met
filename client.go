package met

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// defaultRoot is the public root of the Met API.
const defaultRoot = "https://collectionapi.metmuseum.org/public/collection/v1/"

// Client provides an interface for the Met API.
type Client struct {
	*http.Client
	// RootURL is the Met API root. If unspecified, Client uses defaultRoot.
	RootURL *url.URL
}

// NewClient constructs a Met API client.
func NewClient(c *http.Client) *Client {
	defaultRootURL, _ := url.Parse(defaultRoot)
	return &Client{
		Client:  c,
		RootURL: defaultRootURL,
	}
}

// copyRootURL returns a mutation-safe copy of c.RootURL.
func (c *Client) copyRootURL() *url.URL {
	u := *c.RootURL
	return &u
}

// Objects returns a listing of all valid Object IDs, pursuant to any
// restriction on metadata update dates and/or department membership
// specified in the request options.
func (c *Client) Objects(options ObjectsOptions) (*ObjectsResult, error) {
	u := c.copyRootURL()
	u.Path += "objects"
	u.RawQuery = options.toQuery().Encode()

	resp, err := c.makeRequest(u)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var out = new(ObjectsResult)
	json.Unmarshal(body, out)
	return out, nil
}

// Object returns a record for an object, containing all open access data about
// that object, including its image (if the image is available under Open
// Access).
func (c *Client) Object(options ObjectOptions) (*ObjectResult, error) {
	u := c.copyRootURL()
	u.Path += fmt.Sprintf("objects/%d", options.ObjectID)

	resp, err := c.makeRequest(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var out = new(ObjectResult)
	err = json.Unmarshal(body, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Departments returns a listing of all valid departments.
func (c *Client) Departments() (*DepartmentsResult, error) {
	u := c.copyRootURL()
	u.Path += "departments"

	resp, err := c.makeRequest(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var out = new(DepartmentsResult)
	err = json.Unmarshal(body, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Search returns a listing of all Object IDs for objects with metadata
// matching the specified options.
func (c *Client) Search(options SearchOptions) (*ObjectsResult, error) {
	u := c.copyRootURL()
	u.Path += "search"
	u.RawQuery = options.toQuery().Encode()

	resp, err := c.makeRequest(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var out = new(ObjectsResult)
	if err = json.Unmarshal(body, out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) makeRequest(u *url.URL) (*http.Response, error) {
	httpClient := c.Client
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return checkStatus(httpClient.Get(u.String()))
}

func checkStatus(res *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return res, err
	} else if res != nil && res.StatusCode != 200 {
		return res, fmt.Errorf("got non-200 response code: %d", res.StatusCode)
	}
	return res, nil
}
