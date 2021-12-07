package ant_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertsExpr(t *testing.T) {

	regexpPatternStr, specificity, err := convertToRegex("/foo*/bar/tep**.jsp")
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, 4, specificity)
	assert.Equal(t, "^/foo[^/]*/bar/tep.*\\.jsp$", regexpPatternStr)
}

func TestMatchesRoot(t *testing.T) {

	antPattern := MustCompile("/**")

	assert.True(t, antPattern.Matches("/"))
	assert.Equal(t, "/**", antPattern.String())
}

func TestFindStringSubmatch(t *testing.T) {

	antPattern := MustCompile("/foo(/**)")

	result := antPattern.FindStringSubmatch("/foo/bar")

	assert.Equal(t, "/bar", result[1])
}
