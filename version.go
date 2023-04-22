package argparser

import "fmt"

const MAJOR_VERSION = 0

const MINOR_VERSION = 2

const PATCH_VERSION = 0

const PRERELEASE_ID = "dev"

func MakeVersionString() string {
	var suffix string
	if PRERELEASE_ID == "" {
		suffix = ""
	} else {
		suffix = "-" + PRERELEASE_ID
	}

	return fmt.Sprintf("%d.%d.%d%s", MAJOR_VERSION, MINOR_VERSION, PATCH_VERSION, suffix)
}
