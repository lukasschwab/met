package met

// DepartmentsResult is a listing of all departments. See
// https://metmuseum.github.io/#departments
type DepartmentsResult struct {
	Departments []Department `json:"departments"`
}

// Department is a department listing.
type Department struct {
	// DepartmentID is the fully-qualified ID for this department for use with
	// the Objects and Search APIs.
	DepartmentID int `json:"departmentId"`
	// DisplayName is this department's name.
	DisplayName string `json:"displayName"`
}
