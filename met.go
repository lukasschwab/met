// Package met provides a thin wrapper around the Metropolitan Museum of Art
// Collection API: https://metmuseum.github.io
package met

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	opt "github.com/lukasschwab/optional"
)

const apiRoot = "https://collectionapi.metmuseum.org/public/collection/v1/"

type HTTPOptions struct {
	Client *http.Client
}

func (options HTTPOptions) makeRequest(u *url.URL) (*http.Response, error) {
	var res *http.Response
	var err error
	// Prefer the client specified in options.
	if options.Client != nil {
		res, err = options.Client.Get(u.String())
	}
	res, err = http.Get(u.String())
	return checkStatus(res, err)
}

func checkStatus(res *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return res, err
	} else if res != nil && res.StatusCode != 200 {
		return res, fmt.Errorf("Got non-200 response code: %d", res.StatusCode)
	}
	return res, nil
}

type ObjectsOptions struct {
	HTTPOptions
	// NOTE: should MetadataDate just be a string?
	MetadataDate  *time.Time
	DepartmentIDs []int
}

func (options ObjectsOptions) toQuery() url.Values {
	query := url.Values{}
	if options.MetadataDate != nil {
		t := *options.MetadataDate
		query.Set("metadataDate", t.Format("2006-01-02"))
	}
	if options.DepartmentIDs != nil && len(options.DepartmentIDs) > 0 {
		asStrings := make([]string, len(options.DepartmentIDs))
		for i := range options.DepartmentIDs {
			asStrings[i] = strconv.Itoa(options.DepartmentIDs[i])
		}
		query.Set("departmentIds", strings.Join(asStrings, "|"))
	}
	return query
}

type ObjectsResult struct {
	Total     int   `json:"total"`
	ObjectIDs []int `json:"objectIDs"`
}

func Objects(options ObjectsOptions) (*ObjectsResult, error) {
	u, _ := url.Parse(apiRoot)
	u.Path += "objects"
	u.RawQuery = options.toQuery().Encode()

	resp, err := options.makeRequest(u)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var out = new(ObjectsResult)
	json.Unmarshal(body, out)
	return out, nil
}

type ObjectOptions struct {
	HTTPOptions
	ObjectID int
}

type ObjectResult struct {
	ObjectID              int           `json:"objectID"`
	IsHighlight           bool          `json:"isHighlight"`
	AccessionNumber       string        `json:"accessionNumber"`
	IsPublicDomain        bool          `json:"isPublicDomain"`
	PrimaryImage          string        `json:"primaryImage"`
	PrimaryImageSmall     string        `json:"primaryImageSmall"`
	AdditionalImages      []string      `json:"additionalImages"`
	Constituents          []Constituent `json:"constituents"`
	Department            string        `json:"department"`
	ObjectName            string        `json:"objectName"`
	Title                 string        `json:"title"`
	Culture               string        `json:"culture"`
	Period                string        `json:"period"`
	Dynasty               string        `json:"dynasty"`
	Reign                 string        `json:"reign"`
	Portfolio             string        `json:"portfolio"`
	ArtistRole            string        `json:"artistRole"`
	ArtistPrefix          string        `json:"artistPrefix"`
	ArtistDisplayName     string        `json:"artistDisplayName"`
	ArtistDisplayBio      string        `json:"artistDisplayBio"`
	ArtistSuffix          string        `json:"artistSuffix"`
	ArtistAlphaSort       string        `json:"artistAlphaSort"`
	ArtistNationality     string        `json:"artistNationality"`
	ArtistBeginDate       string        `json:"artistBeginDate"`
	ArtistEndDate         string        `json:"artistEndDate"`
	ObjectDate            string        `json:"objectDate"`
	ObjectBeginDate       int           `json:"objectBeginDate"`
	ObjectEndDate         int           `json:"objectEndDate"`
	Medium                string        `json:"medium"`
	Dimensions            string        `json:"dimensions"`
	CreditLine            string        `json:"creditLine"`
	GeographyType         string        `json:"geographyType"`
	City                  string        `json:"city"`
	State                 string        `json:"state"`
	County                string        `json:"county"`
	Country               string        `json:"country"`
	Region                string        `json:"region"`
	Subregion             string        `json:"subregion"`
	Locale                string        `json:"locale"`
	Locus                 string        `json:"locus"`
	Excavation            string        `json:"excavation"`
	River                 string        `json:"river"`
	Classification        string        `json:"classification"`
	RightsAndReproduction string        `json:"rightsAndReproduction"`
	LinkResource          string        `json:"linkResource"`
	MetadataDate          string        `json:"metadataDate"`
	Repository            string        `json:"repository"`
	ObjectURL             string        `json:"objectURL"`
	Tags                  []string      `json:"tags"`
}

type Constituent struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

func Object(options ObjectOptions) (*ObjectResult, error) {
	u, _ := url.Parse(apiRoot)
	u.Path += fmt.Sprintf("objects/%d", options.ObjectID)

	resp, err := options.makeRequest(u)
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

type DepartmentsResult struct {
	Departments []Department `json:"departments"`
}

type Department struct {
	DepartmentID int    `json:"departmentId"`
	DisplayName  string `json:"displayName"`
}

func Departments(options HTTPOptions) (*DepartmentsResult, error) {
	u, _ := url.Parse(apiRoot)
	u.Path += "departments"

	resp, err := options.makeRequest(u)
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

type SearchOptions struct {
	HTTPOptions
	Q               string
	IsHighlight     opt.Bool
	DepartmentID    opt.Int
	IsOnView        opt.Bool
	ArtistOrCulture opt.Bool
	// Media are concatenated to form the "medium" option.
	Media        []string
	HasImages    opt.Bool
	GeoLocations []string
	// If DateBegin is non-nil, DateEnd must be non-nil.
	DateBegin opt.Int
	// If DateEnd is non-nil, DateBegin must be non-nil.
	DateEnd opt.Int
}

func (options SearchOptions) validate() error {
	if options.DateBegin != nil && options.DateEnd == nil {
		return errors.New("DateBegin is defined, but DateEnd is not")
	} else if options.DateBegin == nil && options.DateEnd != nil {
		return errors.New("DateEnd is defined, but DateBegin is not")
	}
	return nil
}

func maybeAddBoolToQuery(q *url.Values, name string, b opt.Bool) {
	if b != nil {
		q.Set(name, strconv.FormatBool(opt.ToBool(b)))
	}
}

func maybeAddIntToQuery(q *url.Values, name string, i opt.Int) {
	if i != nil {
		q.Set(name, strconv.Itoa(opt.ToInt(i)))
	}
}

func maybeAddSliceToQuery(q *url.Values, name string, s []string) {
	if s != nil && len(s) > 0 {
		q.Set(name, strings.Join(s, "|"))
	}
}

func (options SearchOptions) toQuery() url.Values {
	query := url.Values{}
	query.Set("q", options.Q)
	maybeAddBoolToQuery(&query, "isHighlight", options.IsHighlight)
	maybeAddIntToQuery(&query, "departmentId", options.DepartmentID)
	maybeAddBoolToQuery(&query, "isOnView", options.IsOnView)
	maybeAddBoolToQuery(&query, "artistOrCulture", options.ArtistOrCulture)
	maybeAddSliceToQuery(&query, "medium", options.Media)
	maybeAddBoolToQuery(&query, "hasImages", options.HasImages)
	maybeAddSliceToQuery(&query, "geoLocations", options.GeoLocations)
	maybeAddIntToQuery(&query, "dateBegin", options.DateBegin)
	maybeAddIntToQuery(&query, "dateEnd", options.DateEnd)
	return query
}

func Search(options SearchOptions) (*ObjectsResult, error) {
	err := options.validate()
	if err != nil {
		return nil, err
	}

	u, _ := url.Parse(apiRoot)
	u.Path += "search"
	u.RawQuery = options.toQuery().Encode()

	resp, err := options.makeRequest(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var out = new(ObjectsResult)
	err = json.Unmarshal(body, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
