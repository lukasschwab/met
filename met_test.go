package met

import (
	"fmt"
	"testing"
	"time"
)

// TODO: test that custom client is used.
// TODO: more comprehensive tests of each method.
// TODO: test invalid input.

func TestObjects(t *testing.T) {
	all, _ := Objects(ObjectsOptions{})
	checkObjectsLengthsAgree(t, all)

	yesterday := time.Now().AddDate(0, 0, -1)
	recent, _ := Objects(ObjectsOptions{
		MetadataDate: &yesterday,
	})
	checkObjectsLengthsAgree(t, recent)

	if recent.Total >= all.Total {
		t.Errorf("New (%d) should be fewer objects than All (%d).", recent.Total, all.Total)
	}

	deps, _ := Objects(ObjectsOptions{
		DepartmentIDs: []int{1},
	})
	checkObjectsLengthsAgree(t, deps)
}

func ExampleObjects_all() {
	// Get all objects.
	allObjects, err := Objects(ObjectsOptions{})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("All %d object IDs: %v\n", allObjects.Total, allObjects.ObjectIDs)
}

func ExampleObjects_date() {
	// Get all objects updated in the last 20 years.
	twentyYearsAgo := time.Now().AddDate(-20, 0, 0)
	recentObjects, err := Objects(ObjectsOptions{
		MetadataDate: &twentyYearsAgo,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("All %d object IDs: %v\n", recentObjects.Total, recentObjects.ObjectIDs)
}

func ExampleObjects_department() {
	// Get all objects updated in the last 20 years in Department 1.
	twentyYearsAgo := time.Now().AddDate(-20, 0, 0)
	d1Objects, err := Objects(ObjectsOptions{
		MetadataDate:  &twentyYearsAgo,
		DepartmentIDs: []int{1},
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("All %d object IDs: %v\n", d1Objects.Total, d1Objects.ObjectIDs)
}

func TestObject(t *testing.T) {
	targetObject := 436535
	o, err := Object(ObjectOptions{ObjectID: targetObject})
	if err != nil {
		t.Errorf("Valid fetch got error: %s", err)
	} else if o.ObjectID != targetObject {
		t.Errorf("Object ID does not match target.")
	}

	_, err = Object(ObjectOptions{ObjectID: -1})
	if err == nil {
		t.Errorf("Invalid ID should produce 404 status error.")
	}
}

func ExampleObject() {
	obj, err := Object(ObjectOptions{ObjectID: 436535})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("Object %d is titled '%s.'\n", obj.ObjectID, obj.Title)
}

func TestDepartments(t *testing.T) {
	o, err := Departments(HTTPOptions{})
	if err != nil {
		t.Errorf("Valid fetch got error: %s", err)
	} else if len(o.Departments) <= 0 {
		t.Errorf("No departments listed: %d", len(o.Departments))
	} else if o.Departments[0].DepartmentID != 1 {
		t.Errorf("Departments listed in incorrect order: first element has ID %d", o.Departments[0].DepartmentID)
	}
}

func ExampleDepartments() {
	depts, err := Departments(HTTPOptions{})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d departments: %v\n", len(depts.Departments), depts.Departments)
}

func TestSearch(t *testing.T) {
	o, err := Search(SearchOptions{Q: "sunflowers"})
	if err != nil {
		t.Errorf("Valid fetch got error: %s", err)
	}
	checkObjectsLengthsAgree(t, o)
	o, err = Search(SearchOptions{Q: "sunflowers", IsHighlight: true})
	if err != nil {
		t.Errorf("Valid fetch got error: %s", err)
	}
	checkObjectsLengthsAgree(t, o)
}

func ExampleSearch_query() {
	results, err := Search(SearchOptions{Q: "sunflower"})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleSearch_highlights() {
	results, err := Search(SearchOptions{
		Q:           "sunflower",
		IsHighlight: true,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleSearch_department() {
	results, err := Search(SearchOptions{
		Q:            "cat",
		DepartmentID: 6,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleSearch_view() {
	results, err := Search(SearchOptions{
		Q:        "sunflower",
		IsOnView: true,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleSearch_culture() {
	results, err := Search(SearchOptions{
		Q:               "french",
		ArtistOrCulture: true,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleSearch_media() {
	results, err := Search(SearchOptions{
		Q:     "quilt",
		Media: []string{"Quilts", "Silk", "Bedcovers"},
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleSearch_images() {
	results, err := Search(SearchOptions{
		Q:         "Auguste Renoir",
		HasImages: true,
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleSearch_geolocation() {
	results, err := Search(SearchOptions{
		Q:            "flowers",
		GeoLocations: []string{"France"},
	})
	if err != nil {
		// Handle error.
	}
	fmt.Printf("There are %d results.\n", results.Total)
}

func ExampleSearch_dates() {
	results, err := Search(SearchOptions{
		Q:         "African",
		DateBegin: 1700,
		DateEnd:   1800,
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
