package version

import "fmt"

var (
	RELEASE = ""

	API = "v1"

	Long = fmt.Sprintf("Permastar release: %s, API version: %s", RELEASE, API)
)
