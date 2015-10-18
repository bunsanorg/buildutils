package buildutils

import (
    "go/build"
)

func SrcDir(importPath string) (string, error) {
    pkg, err := build.Import(importPath, "", build.FindOnly)
    if err != nil {
        return "", err
    }
    return pkg.Dir, nil
}
