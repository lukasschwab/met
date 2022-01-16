package met

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestCustomClient(t *testing.T) {
	cli := NewClient(&http.Client{
		// Too short to reasonably succeed.
		Timeout: 1 * time.Nanosecond,
	})
	_, err := cli.Objects(ObjectsOptions{})
	if err == nil {
		t.Errorf("Custom client with infinitessimal timeout should always error.")
	}
}

func TestObjects(t *testing.T) {
	c := NewClient(&http.Client{})
	all, _ := c.Objects(ObjectsOptions{})
	checkObjectsLengthsAgree(t, all)

	yesterday := time.Now().AddDate(0, 0, -1)
	recent, _ := c.Objects(ObjectsOptions{
		MetadataDate: &yesterday,
	})
	checkObjectsLengthsAgree(t, recent)

	if recent.Total >= all.Total {
		t.Errorf("New (%d) should be fewer objects than All (%d).", recent.Total, all.Total)
	}

	deps, _ := c.Objects(ObjectsOptions{
		DepartmentIDs: []int{1},
	})
	checkObjectsLengthsAgree(t, deps)
}

func ExampleClient_Objects_all() {
	c := NewClient(&http.Client{})
	// Get all objects.
	allObjects, err := c.Objects(ObjectsOptions{})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("All %d object IDs: %v\n", allObjects.Total, allObjects.ObjectIDs)
}

func ExampleClient_Objects_date() {
	c := NewClient(&http.Client{})
	// Get all objects updated in the last 20 years.
	twentyYearsAgo := time.Now().AddDate(-20, 0, 0)
	recentObjects, err := c.Objects(ObjectsOptions{
		MetadataDate: &twentyYearsAgo,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("All %d object IDs: %v\n", recentObjects.Total, recentObjects.ObjectIDs)
}

func ExampleClient_Objects_department() {
	c := NewClient(&http.Client{})
	// Get all objects updated in the last 20 years in Department 1.
	twentyYearsAgo := time.Now().AddDate(-20, 0, 0)
	d1Objects, err := c.Objects(ObjectsOptions{
		MetadataDate:  &twentyYearsAgo,
		DepartmentIDs: []int{1},
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("All %d object IDs: %v\n", d1Objects.Total, d1Objects.ObjectIDs)
}

func TestObject(t *testing.T) {
	c := NewClient(&http.Client{})
	targetObject := 436535
	o, err := c.Object(ObjectOptions{ObjectID: targetObject})
	if err != nil {
		t.Errorf("Valid fetch got error: %s", err)
	} else if o.ObjectID != targetObject {
		t.Errorf("Object ID does not match target.")
	}

	_, err = c.Object(ObjectOptions{ObjectID: -1})
	if err == nil {
		t.Errorf("Invalid ID should produce 404 status error.")
	}
}

func ExampleClient_Object() {
	c := NewClient(&http.Client{})
	obj, err := c.Object(ObjectOptions{ObjectID: 436535})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("Object %d is titled '%s.'\n", obj.ObjectID, obj.Title)
}

func TestDepartments(t *testing.T) {
	c := NewClient(&http.Client{})
	o, err := c.Departments()
	if err != nil {
		t.Errorf("Valid fetch got error: %s", err)
	} else if len(o.Departments) <= 0 {
		t.Errorf("No departments listed: %d", len(o.Departments))
	} else if o.Departments[0].DepartmentID != 1 {
		t.Errorf("Departments listed in incorrect order: first element has ID %d", o.Departments[0].DepartmentID)
	}
}

func ExampleClient_Departments() {
	c := NewClient(&http.Client{})
	depts, err := c.Departments()
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d departments: %v\n", len(depts.Departments), depts.Departments)
}

func TestSearch(t *testing.T) {
	c := NewClient(&http.Client{})
	o, err := c.Search(SearchOptions{Q: "sunflowers"})
	if err != nil {
		t.Errorf("Valid fetch got error: %s", err)
	}
	checkObjectsLengthsAgree(t, o)
	o, err = c.Search(SearchOptions{Q: "sunflowers", IsHighlight: true})
	if err != nil {
		t.Errorf("Valid fetch got error: %s", err)
	}
	checkObjectsLengthsAgree(t, o)
}

func ExampleClient_Search_query() {
	c := NewClient(&http.Client{})
	results, err := c.Search(SearchOptions{Q: "sunflower"})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleClient_Search_highlights() {
	c := NewClient(&http.Client{})
	results, err := c.Search(SearchOptions{
		Q:           "sunflower",
		IsHighlight: true,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleClient_Search_department() {
	c := NewClient(&http.Client{})
	results, err := c.Search(SearchOptions{
		Q:            "cat",
		DepartmentID: 6,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleClient_Search_view() {
	c := NewClient(&http.Client{})
	results, err := c.Search(SearchOptions{
		Q:        "sunflower",
		IsOnView: true,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleClient_Search_culture() {
	c := NewClient(&http.Client{})
	results, err := c.Search(SearchOptions{
		Q:               "french",
		ArtistOrCulture: true,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleClient_Search_media() {
	c := NewClient(&http.Client{})
	results, err := c.Search(SearchOptions{
		Q:     "quilt",
		Media: []string{"Quilts", "Silk", "Bedcovers"},
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleClient_Search_images() {
	c := NewClient(&http.Client{})
	results, err := c.Search(SearchOptions{
		Q:         "Auguste Renoir",
		HasImages: true,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleClient_Search_geolocation() {
	c := NewClient(&http.Client{})
	results, err := c.Search(SearchOptions{
		Q:            "flowers",
		GeoLocations: []string{"France"},
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleClient_Search_dates() {
	c := NewClient(&http.Client{})
	results, err := c.Search(SearchOptions{
		Q:         "African",
		YearRange: NewYearRange(1700, 1800),
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

// Utilities.

func checkObjectsLengthsAgree(t *testing.T, o *ObjectsResult) {
	if o.Total != len(o.ObjectIDs) {
		t.Errorf("Lengths don't match: Total=%d, Length=%d", o.Total, len(o.ObjectIDs))
	}
}
