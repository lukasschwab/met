# met [![GoDoc](https://godoc.org/github.com/lukasschwab/met?status.svg)](https://godoc.org/github.com/lukasschwab/met)

Package met provides a thin wrapper around the Metropolitan Museum of Art
Collection API. View the full documentation: https://metmuseum.github.io

All request Options here allow for the specification of a custom HTTP client:

```go
cli := &http.Client{}
res, err := Search(ObjectsOptions{
  HTTPOptions: HTTPOptions{Client: cli},
  Q: "sunflowers"
})
```

## Usage

#### type Constituent

```go
type Constituent struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
```


#### type Department

```go
type Department struct {
	DepartmentID int    `json:"departmentId"`
	DisplayName  string `json:"displayName"`
}
```


#### type DepartmentsResult

```go
type DepartmentsResult struct {
	Departments []Department `json:"departments"`
}
```


#### func  Departments

```go
func Departments(options HTTPOptions) (*DepartmentsResult, error)
```
Departments returns a listing of all valid departments.

#### type HTTPOptions

```go
type HTTPOptions struct {
	Client *http.Client
}
```


#### type ObjectOptions

```go
type ObjectOptions struct {
	HTTPOptions
	ObjectID int
}
```


#### type ObjectResult

```go
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
```


#### func  Object

```go
func Object(options ObjectOptions) (*ObjectResult, error)
```
Object returns a record for an object, containing all open access data about
that object, including its image (if the image is available under Open Access).

#### type ObjectsOptions

```go
type ObjectsOptions struct {
	HTTPOptions
	// NOTE: should MetadataDate just be a string?
	MetadataDate  *time.Time
	DepartmentIDs []int
}
```


#### type ObjectsResult

```go
type ObjectsResult struct {
	Total     int   `json:"total"`
	ObjectIDs []int `json:"objectIDs"`
}
```


#### func  Objects

```go
func Objects(options ObjectsOptions) (*ObjectsResult, error)
```
Objects returns a listing of all valid Object IDs, pursuant to any restriction
on metadata update dates and/or department membership specified in the request
options.

#### func  Search

```go
func Search(options SearchOptions) (*ObjectsResult, error)
```
Search returns a listing of all Object IDs for objects with metadata matching
the specified options.

#### type SearchOptions

```go
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
```
