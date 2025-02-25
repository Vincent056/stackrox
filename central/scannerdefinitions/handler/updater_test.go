package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stackrox/rox/central/scannerdefinitions/file"
	"github.com/stackrox/rox/pkg/fileutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	defURL = "https://definitions.stackrox.io/e799c68a-671f-44db-9682-f24248cd0ffe/diff.zip"
)

var (
	nov23 = time.Date(2019, time.November, 23, 0, 0, 0, 0, time.Local)
)

func assertOnFileExistence(t *testing.T, path string, shouldExist bool) {
	exists, err := fileutils.Exists(path)
	require.NoError(t, err)
	assert.Equal(t, shouldExist, exists)
}

func TestUpdate(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "dump.zip")
	u := newUpdater(file.New(filePath), &http.Client{Timeout: 30 * time.Second}, defURL, 1*time.Hour)
	// Should fetch first time.
	require.NoError(t, u.doUpdate())
	assertOnFileExistence(t, filePath, true)

	lastUpdatedTime := time.Now().Add(time.Hour)
	mustSetModTime(t, filePath, lastUpdatedTime)
	// Should not fetch since it can't be updated in a time in the future.
	require.NoError(t, u.doUpdate())
	assert.Equal(t, lastUpdatedTime.UTC(), mustGetModTime(t, filePath))
	assertOnFileExistence(t, filePath, true)

	// Should definitely fetch.
	mustSetModTime(t, filePath, nov23)
	require.NoError(t, u.doUpdate())
	assert.True(t, lastUpdatedTime.UTC().After(mustGetModTime(t, filePath)))
	assert.True(t, mustGetModTime(t, filePath).After(nov23.UTC()))
	assertOnFileExistence(t, filePath, true)
}

func mustGetModTime(t *testing.T, path string) time.Time {
	fi, err := os.Stat(path)
	require.NoError(t, err)
	return fi.ModTime().UTC()
}

func mustSetModTime(t *testing.T, path string, modTime time.Time) {
	require.NoError(t, os.Chtimes(path, time.Now(), modTime))
}
