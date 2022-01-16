package met

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

// ObjectsOptions encapsulates arguments for the Met API Objects endpoint. See
// https://metmuseum.github.io/#objects
type ObjectsOptions struct {
	// MetadataDate restricts Options results to objects updated after the day
	// of the specified time.
	MetadataDate *time.Time
	// DepartmentIDs restricts Options results to objects in the specified
	// departments. See (c *Client).Departments().
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

// ObjectsResult is a listing of object IDs. See
// https://metmuseum.github.io/#response
type ObjectsResult struct {
	// Total is the total number of publicly-available objects.
	Total int `json:"total"`
	// ObjectIDs is an array containing the object ID of each publicly-available
	// object.
	ObjectIDs []int `json:"objectIDs"`
}
