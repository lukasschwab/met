package met

import (
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
		DepartmentIds: []int{1},
	})
	checkObjectsLengthsAgree(t, deps)
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

// Utilities.

func checkObjectsLengthsAgree(t *testing.T, o *ObjectsResult) {
	if o.Total != len(o.ObjectIDs) {
		t.Errorf("Lengths don't match: Total=%d, Length=%d", o.Total, len(o.ObjectIDs))
	}
}