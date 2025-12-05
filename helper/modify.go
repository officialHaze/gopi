package helper

import (
	"fmt"
	"io"
	"os"
	"projinit/settings"
	"strings"
)

func ReplaceProjectNamePlaceholder(filepath, projname string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %v", filepath, err)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %v", filepath, err)
	}

	// replace the project name placeholder with actual name
	modified := strings.ReplaceAll(string(b), "<PROJECT_NAME>", projname)

	return []byte(modified), nil
}

func ReplaceGOVersionPlaceholder(b []byte, version string) []byte {
	content := string(b)

	if version == "" { // empty
		version = settings.MySettings.DEF_GO_VERSION
	}

	modified := strings.ReplaceAll(content, "<GO_VERSION>", version)

	return []byte(modified)
}
