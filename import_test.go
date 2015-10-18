package buildutils

import (
    "os"
    "path"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

const (
    CurrentImportPath = "github.com/bunsanorg/buildutils"
    CurrentFileName   = "import_test.go"
)

func TestSrcDir(t *testing.T) {
    dir, err := SrcDir(CurrentImportPath)
    require.NoError(t, err)
    _, err = os.Stat(path.Join(dir, CurrentFileName))
    assert.NoError(t, err, "unable to find current file")
}

func TestImportPath(t *testing.T) {
    dir, err := SrcDir(CurrentImportPath)
    require.NoError(t, err)
    importPath, err := ImportPath(dir)
    require.NoError(t, err)
    assert.Equal(t, CurrentImportPath, importPath)
}

func TestImportPathRaw(t *testing.T) {
    testData := []struct {
        srcDir     string
        srcDirs    []string
        importPath string
        errString  string
    }{
        {
            srcDir:     "/root/1/import/path",
            srcDirs:    []string{"/root/1", "/root/2", "/root/3"},
            importPath: "import/path",
            errString:  "",
        },
        {
            srcDir:     "/root/2/import/path",
            srcDirs:    []string{"/root/1", "/root/2", "/root/3"},
            importPath: "import/path",
            errString:  "",
        },
        {
            srcDir:     "/root/3/import/path",
            srcDirs:    []string{"/root/1", "/root/2", "/root/3"},
            importPath: "import/path",
            errString:  "",
        },
        {
            srcDir:     "/root/4/import/path",
            srcDirs:    []string{"/root/1", "/root/2", "/root/3"},
            importPath: "",
            errString:  "unable to find import path for source \"/root/4/import/path\"",
        },
    }
    for _, tt := range testData {
        impPath, err := importPath(tt.srcDir, tt.srcDirs)
        if tt.errString == "" {
            if assert.NoError(t, err) {
                assert.Equal(t, tt.importPath, impPath)
            }
        } else {
            assert.EqualError(t, err, tt.errString)
        }
    }
}
