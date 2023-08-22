package cmdline

import (
	// nolint:gosec
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertExpectedHashFile(t *testing.T, expHash, fName string) {
	// nolint:gosec
	fileContent, err := os.ReadFile(fName)
	assert.NoError(t, err)
	assertExpectedHashStr(t, expHash, string(fileContent))
}

func assertExpectedHashStr(t *testing.T, expHash, actual string) {
	// nolint:gosec
	h := md5.New()
	_, err := io.WriteString(h, actual)
	assert.NoError(t, err)
	// assert.Equal(t, expHash, fmt.Sprintf("%x", h.Sum(nil)), "%s", actual)
	assert.Equal(t, expHash, fmt.Sprintf("%x", h.Sum(nil)), "%s", actual[0:int(math.Min(100, float64(len(actual))))])
}
