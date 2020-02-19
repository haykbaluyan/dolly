package resolve_test

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-phorce/dolly/algorithms/guid"
	"github.com/go-phorce/dolly/fileutil/resolve"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ResolveDirectory(t *testing.T) {
	tmpDir := path.Join(os.TempDir(), "resolve-test", guid.MustCreate())
	testData := []struct {
		dir     string
		baseDir string
		create  bool
		err     string
	}{
		{
			dir:     "a1/a2",
			baseDir: tmpDir,
			create:  false,
			err:     "no such file or directory",
		},
		{
			dir:     "a1/a2",
			baseDir: tmpDir,
			create:  true,
			err:     "",
		},
		{
			dir:     "a1/a2",
			baseDir: tmpDir,
			create:  false,
			err:     "",
		},
	}

	// Run test
	for idx, v := range testData {
		t.Run(fmt.Sprintf("[%d] %s", idx, v.dir), func(t *testing.T) {
			d, err := resolve.Directory(v.dir, v.baseDir, v.create)
			if v.err != "" {
				require.Error(t, err)
				assert.True(t, strings.Contains(err.Error(), v.err))
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, d)
				assert.True(t, strings.HasSuffix(d, v.dir))
			}
		})
	}
}

func Test_File(t *testing.T) {
	f, err := resolve.File("", ".")
	assert.NoError(t, err)
	assert.Empty(t, f)

	f, err = resolve.File("etc/dev/softhsm_unittest.json", "../..")
	assert.NoError(t, err)
	assert.NotEmpty(t, f)

	// now f is relative to current folder
	f2, err := resolve.File(f, ".")
	assert.NoError(t, err)
	assert.Equal(t, f, f2)

	fabs, err := filepath.Abs(f2)
	require.NoError(t, err)

	f3, err := resolve.File(fabs, "/does/not/matter")
	assert.NoError(t, err)
	assert.Equal(t, fabs, f3)

	_, err = resolve.File(fabs+".junk", "/does/not/matter")
	assert.Error(t, err)
}