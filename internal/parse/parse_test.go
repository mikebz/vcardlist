package parse

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseOneVcf tests the sample vcf card and produce the results
func TestParseOneVcf(t *testing.T) {
	fullPath, _ := filepath.Abs("./testdata/Fred Snapps.vcf")
	nameEmail, err := ParseFile(fullPath)
	assert.NoError(t, err)
	assert.Equal(t, "Fred Snapps", nameEmail.Name)
	assert.Equal(t, "snapps@snoogle.com", nameEmail.Email)
}

// TestParseDir goes through the entire test directory and parses all
// the items in it.
func TestParseDir(t *testing.T) {
	fullPath, _ := filepath.Abs("./testdata")
	result, err := Parse(fullPath)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(*result))
	nameEmail := (*result)[0]
	assert.Equal(t, "Fred Snapps", nameEmail.Name)
	assert.Equal(t, "snapps@snoogle.com", nameEmail.Email)
}
