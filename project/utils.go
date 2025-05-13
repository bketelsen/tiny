package project

import (
	"fmt"
	"strings"
)

func (p *Project) computedPath(path string) string {
	if path == "" {
		return ""
	}
	if strings.HasPrefix(path, "/") {
		return path
	}
	path = strings.TrimPrefix(path, ".")
	return fmt.Sprintf("%s/%s", p.outDir, path)
}
