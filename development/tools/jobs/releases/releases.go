// Code generated by rendertemplates. DO NOT EDIT.

package releases

// List of currently supported releases
var (
	Release20  = mustParse("2.0")
	Release125 = mustParse("1.25")
	Release124 = mustParse("1.24")
	Release123 = mustParse("1.23")
)

// GetAllKymaReleases returns all supported kyma release branches
func GetAllKymaReleases() []*SupportedRelease {
	return []*SupportedRelease{
		Release125,
		Release124,
		Release123,
	}
}

// GetNextKymaRelease returns the version of kyma currently under development
func GetNextKymaRelease() *SupportedRelease {
	return Release20
}
