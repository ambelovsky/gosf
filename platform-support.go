package gosf

import (
	"log"
	"runtime"
)

// SupportedPlatforms is an array of OS platform names that this framework works well on
var SupportedPlatforms []string

func init() {
	// Verify that the current server platform is supported
	SupportedPlatforms = []string{"linux", "darwin", "windows"}
	if !ArrayContainsString(SupportedPlatforms, runtime.GOOS) {
		log.Panic("Unsupported platform.")
	}
}
