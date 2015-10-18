package buildutils

import (
    "os"
    "path"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestSrcDir(t *testing.T) {
    dir, err := SrcDir("github.com/bunsanorg/buildutils")
    require.NoError(t, err)
    _, err = os.Stat(path.Join(dir, "import_test.go"))
    assert.NoError(t, err, "unable to find current file")
}
