// Package met provides a thin wrapper around the Metropolitan Museum of Art
// Collection API. View the full documentation: https://metmuseum.github.io
//
// All request Options here allow for the specification of a custom HTTP client:
//
//  cli := &http.Client{}
//  res, err := Search(ObjectsOptions{
//    HTTPOptions: HTTPOptions{Client: cli},
//    Q: "sunflowers"
//  })
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
	opturl "github.com/lukasschwab/optional/url"
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
	} else {
		res, err = http.Get(u.String())
	}
	return checkStatus(res, err)
}

func checkStatus(res *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return res, err
	} else if res != nil && res.StatusCode != 200 {
		return res, fmt.Errorf("got non-200 response code: %d", res.StatusCode)
	}
	return res, nil
}

type ObjectsOptions struct {
	HTTPOptions
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

// Objects returns a listing of all valid Object IDs, pursuant to any
// restriction on metadata update dates and/or department membership
// specified in the request options.
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
	Measurements          []Measurement `json:"measurements"`
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
	Tags                  []Tag      `json:"tags"`
	ObjectWikidataURL string `json:"objectWikidata_URL"`
	IsTimelineWork bool `json:"isTimelineWork"`
	GalleryNumber string `json:"GalleryNumber"`
}

type Constituent struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type Measurement struct {
	ElementName string `json:"elementName"`
	ElementDescription string `json:"elementDescription,omitempty"`
	ElementMeasurements map[string]float64 `json:"elementMeasurements"`
}

type Tag struct {
	Term string `json:"term"`
	AatURL string `json:"AAT_URL"`
	WikidataURL string `json:"Wikidata_URL"`
}

// Object returns a record for an object, containing all open access data about
// that object, including its image (if the image is available under Open
// Access).
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

// Departments returns a listing of all valid departments.
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

func (options SearchOptions) toQuery() url.Values {
	query := url.Values{}
	query.Set("q", options.Q)
	opturl.AddBoolToQuery(&query, "isHighlight", options.IsHighlight, opturl.DefaultBoolFormatter)
	opturl.AddIntToQuery(&query, "departmentId", options.DepartmentID, opturl.DefaultIntFormatter)
	opturl.AddBoolToQuery(&query, "isOnView", options.IsOnView, opturl.DefaultBoolFormatter)
	opturl.AddBoolToQuery(&query, "artistOrCulture", options.ArtistOrCulture, opturl.DefaultBoolFormatter)
	opturl.AddSliceToQuery(&query, "medium", options.Media, opturl.SeparatorFormatter("|"))
	opturl.AddBoolToQuery(&query, "hasImages", options.HasImages, opturl.DefaultBoolFormatter)
	opturl.AddSliceToQuery(&query, "geoLocations", options.GeoLocations, opturl.SeparatorFormatter("|"))
	opturl.AddIntToQuery(&query, "dateBegin", options.DateBegin, opturl.DefaultIntFormatter)
	opturl.AddIntToQuery(&query, "dateEnd", options.DateEnd, opturl.DefaultIntFormatter)
	return query
}

// Search returns a listing of all Object IDs for objects with metadata
// matching the specified options.
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
