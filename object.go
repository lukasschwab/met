package met

// ObjectOptions encapsulates arguments for the Met API Object endpoint. See
// https://metmuseum.github.io/#Object
type ObjectOptions struct {
	// ObjectID is the unique Object ID for an object. See
	// (c *Client).Objects().
	ObjectID int
}

// ObjectResult is a record for an object, containing all open access data about
// that object. See https://metmuseum.github.io/#response-1
type ObjectResult struct {
	// ObjectID is the identifying number for each artwork (unique, can be used
	// as key field).
	ObjectID int `json:"objectID"`
	// IsHighlight indicates a popular and important artwork in the collection.
	IsHighlight bool `json:"isHighlight"`
	// AccessionNumber is an identifying number for each artwork (not always
	// unique).
	AccessionNumber string `json:"accessionNumber"`
	// AccessionYear is the year the artwork was acquired.
	AccessionYear string `json:"accessionYear"`
	// IsPublicDomain indicates an artwork in the Public Domain.
	IsPublicDomain bool `json:"isPublicDomain"`
	// PrimaryImage is the URL of the primary image of an object in JPEG format.
	PrimaryImage string `json:"primaryImage"`
	// PrimaryImageSmall is the URL of the lower-res primary image of an object
	// in JPEG format
	PrimaryImageSmall string `json:"primaryImageSmall"`
	// AdditionalImages is an array of URLs of additional images of an object in
	// JPEG format.
	AdditionalImages []string `json:"additionalImages"`
	// Constituents lists the constituents associated with an object.
	Constituents []Constituent `json:"constituents"`
	// Department is the Met's curatorial department responsible for the object.
	// This is the department's name, not its numerical ID.
	Department string `json:"department"`
	// ObjectName describes the physical type of the object.
	ObjectName string `json:"objectName"`
	// Title is the title, identifying phrase, or name given to a work of art.
	Title string `json:"title"`
	// Culture is information about the culture, or people from which an object
	// was created.
	Culture string `json:"culture"`
	// Period is the time or time period when an object was created.
	Period string `json:"period"`
	// Dynasty (a succession of rulers of the same line or family) under which
	// an object was created
	Dynasty string `json:"dynasty"`
	// Reign of a monarch or ruler under which an object was created.
	Reign string `json:"reign"`
	// Portfolio identifies a set of works created as a group or published as a
	// series.
	Portfolio string `json:"portfolio"`
	// ArtistRole is the role of the artist related to the type of artwork or
	// object that was created.
	ArtistRole string `json:"artistRole"`
	// ArtistPrefix describes the extent of creation or describes an attribution
	// qualifier to the information given by ArtistRole.
	ArtistPrefix string `json:"artistPrefix"`
	// ArtistDisplayName is the artist name in the correct order for display.
	ArtistDisplayName string `json:"artistDisplayName"`
	// ArtistDisplayBio contains the nationality and life dates of an artist,
	// and the birth and death city when known.
	ArtistDisplayBio string `json:"artistDisplayBio"`
	// ArtistSuffix records complex information that qualifies the role of a
	// constituent, e.g. extent of participation by the Constituent.
	ArtistSuffix string `json:"artistSuffix"`
	// ArtistAlphaSort is the artist name used for lexographic sorting, e.g.
	// `Gogh, Vincent van`.
	ArtistAlphaSort string `json:"artistAlphaSort"`
	// ArtistNationality records the national, geopolitical, cultural, or ethnic
	// origins or affiliation of the creator or institution that made the
	// artwork.
	ArtistNationality string `json:"artistNationality"`
	// ArtistBeginDate is the year the artist was born.
	ArtistBeginDate string `json:"artistBeginDate"`
	// ArtistEndDate is the year the artist died.
	ArtistEndDate string `json:"artistEndDate"`
	// ArtistGender is the gender of the artist; currently contains 'female'
	// designations only.
	ArtistGender string `json:"artistGender"`
	// ArtistWikidataURL is the artist's Wikidata URL.
	ArtistWikidataURL string `json:"artistWikidata_URL"`
	// ArtistUlanURL is the artist's Union List of Artist Names URL.
	ArtistUlanURL string `json:"artistULAN_URL"`
	// ObjectDate is a year, a span of years, or a phrase that describes the
	// specific or approximate date when an artwork was designed or created.
	ObjectDate string `json:"objectDate"`
	// ObjectBeginDate is the year the artwork was started to be created
	ObjectBeginDate int `json:"objectBeginDate"`
	// ObjectEndDate is the year the artwork was completed (may be the same as
	// ObjectBeginDate).
	ObjectEndDate int `json:"objectEndDate"`
	// Medium refers to the materials used to create the Object.
	Medium string `json:"medium"`
	// Dimensions describes the size of the artwork or object.
	Dimensions string `json:"dimensions"`
	// Measurements is an array of measured elements of the Object.
	Measurements []Measurement `json:"measurements"`
	// CreditLine is text acknowledging the source or origin of the artwork and
	// the year the object was acquired by the museum.
	CreditLine string `json:"creditLine"`
	// GeographyType describes the relationship of the place catalogued in the
	// geography fields to the object that is being catalogued
	GeographyType string `json:"geographyType"`
	// City where the Object was created.
	City string `json:"city"`
	// State or province where the Object was created.
	State string `json:"state"`
	// County where the Object was created. Sometimes duplicates State.
	County string `json:"county"`
	// Country where the Object was created.
	Country string `json:"country"`
	// Region where the Object was created: more specific than Country but less
	// specific than Subregion.
	Region string `json:"region"`
	// Subregion where the Object was created: more specific than Region but
	// less specific than Locale. Often unavailable.
	Subregion string `json:"subregion"`
	// Locale where the Object was created: more specific than Subregion but
	// less specific than Locus. Often unavailable.
	Locale string `json:"locale"`
	// Locus where the Object was created: more specific than Locale. Often
	// unavailable.
	Locus string `json:"locus"`
	// Excavation is the name of the excavation in which the Object was
	// excavated.
	Excavation string `json:"excavation"`
	// River is a natural watercourse, usually freshwater, flowing toward an
	// ocean, a lake, a sea or another river related to the origins of an
	// artwork. Often unavailable.
	River string `json:"river"`
	// Classification is a general term describing the artwork type.
	Classification string `json:"classification"`
	// RightsAndReproduction is the credit line for artworks still under
	// copyright.
	RightsAndReproduction string `json:"rightsAndReproduction"`
	// LinkResource is the URL of the Object's page on metmuseum.org.
	LinkResource string `json:"linkResource"`
	// MetadataDate dates the last update to this object metadata.
	MetadataDate string `json:"metadataDate"`
	// Repository is `Metropolitan Museum of Art, New York, NY`.
	Repository string `json:"repository"`
	// ObjectURL is the URL of the Object's page on metmuseum.org. Duplicates
	// LinkResource.
	ObjectURL string `json:"objectURL"`
	// Tags are subject keywords associated with the Object.
	Tags []Tag `json:"tags"`
	// ObjectWikidataURL is the Wikidata URL for the Object.
	ObjectWikidataURL string `json:"objectWikidata_URL"`
	// IsTimelineWork indicates whether the object is on the Timeline of Art
	// History website.
	IsTimelineWork bool `json:"isTimelineWork"`
	// GalleryNumber is the number of the Met gallery containing the Object,
	// when available.
	GalleryNumber string `json:"GalleryNumber"`
}

// Constituent is an Object constituent.
type Constituent struct {
	// Name is the Constituent's full name.
	Name string `json:"name"`
	// Role is a Constituent's relationship to an Object.
	Role string `json:"role"`
	// UlanURL is a Union List of Artist Names URL.
	UlanURL string `json:"constituentULAN_URL"`
	// WikidataURL is the constituent's Wikidata URL.
	WikidataURL string `json:"constituentWikidata_URL"`
	// Gender is the constituent's gender, when available; currently contains
	// 'female' designations only.
	Gender string `json:"gender"`
}

// Measurement contains the measurements of an element of an Object.
type Measurement struct {
	// ElementName is a short name for this measured element.
	ElementName string `json:"elementName"`
	// ElementDescription is a human-readable description of this measured
	// element.
	ElementDescription string `json:"elementDescription,omitempty"`
	// ElementMeasurement maps attributes of the element to their measures.
	// Spatial measurements are in centimeters; weights are in kilograms.
	ElementMeasurements map[string]float64 `json:"elementMeasurements"`
}

// Tag is a subject keyword associated with an Object.
type Tag struct {
	// Term is the subject keyworkd associated with an Object.
	Term string `json:"term"`
	// AatURL is the URL of the Getty Art & Architecture Thesaurus entry for
	// Term.
	AatURL string `json:"AAT_URL"`
	// WikidataURL is the Wikidata URL for Term.
	WikidataURL string `json:"Wikidata_URL"`
}
