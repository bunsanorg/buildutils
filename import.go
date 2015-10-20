package buildutils

import (
	"fmt"
	"go/build"
	"path/filepath"
	"strings"
)

func SrcDir(importPath string) (string, error) {
	pkg, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		return "", err
	}
	return pkg.Dir, nil
}

func importPath(srcDir string, srcDirs []string) (string, error) {
	for _, srcPath := range srcDirs {
		if rel, err := filepath.Rel(srcPath, srcDir); err == nil &&
			!strings.HasPrefix(rel, "..") {

			return rel, nil
		}
	}
	return "", fmt.Errorf("unable to find import path for source %q", srcDir)
}

func ImportPath(srcDir string) (string, error) {
	return importPath(srcDir, build.Default.SrcDirs())
}
