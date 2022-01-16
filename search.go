package met

import (
	"net/url"

	opt "github.com/lukasschwab/optional/pkg/optional"
	opturl "github.com/lukasschwab/optional/pkg/url"
)

// SearchOptions encapsulates arguments for the Met API Search endpoint. See
// https://metmuseum.github.io/#search
type SearchOptions struct {
	// Q is a search term, e.g. `sunflowers`.
	Q string
	// IsHighlight restricts results to objects designated as highlights.
	// Highlights are selected works of art from The Met Museum’s permanent collection representing
	// different cultures and time periods.
	IsHighlight opt.Bool
	// DepartmentID restricts results to objects in a department.
	DepartmentID opt.Int
	// IsOnView restricts results to objects on view in the Met.
	IsOnView opt.Bool
	// ArtistOrCulture restricts results to objects with artist name and culture
	// fields that match Q.
	ArtistOrCulture opt.Bool
	// Media restricts results to objects matching one or more medium or object
	// type, e.g. "Ceramics", "Furniture", "Paintings", "Sculpture", "Textiles."
	Media []string
	// HasImages restricts results to objects with images.
	HasImages opt.Bool
	// GeoLocations restricts results to objects matching one or more
	// locations, e.g. "Europe", "France", "Paris", "China", "New York."
	GeoLocations []string
	// YearRange restricts results to a range of years. See NewYearRange.
	YearRange *yearRange
}

type yearRange struct {
	yearBegin int
	yearEnd   int
}

// NewYearRange returns a yearRange (from yearBegin to yearEnd) for use in
// SearchOptions.
func NewYearRange(yearBegin, yearEnd int) *yearRange {
	return &yearRange{yearBegin, yearEnd}
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
	if options.YearRange != nil {
		opturl.AddIntToQuery(&query, "dateBegin", options.YearRange.yearBegin, opturl.DefaultIntFormatter)
		opturl.AddIntToQuery(&query, "dateEnd", options.YearRange.yearEnd, opturl.DefaultIntFormatter)
	}
	return query
}
