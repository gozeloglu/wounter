package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expCount int
	}{
		{
			name:     "Empty string",
			s:        "",
			expCount: 0,
		},
		{
			name:     "String with only spaces",
			s:        "    ",
			expCount: 0,
		},
		{
			name:     "One letter without space",
			s:        "word",
			expCount: 1,
		},
		{
			name:     "Strings with spaces",
			s:        "  This is      example text   with spaces",
			expCount: 6,
		},
		{
			name:     "Strings with some special characters",
			s:        "  This is,      example. text/   -with spaces",
			expCount: 6,
		},
		{
			name:     "Strings with seperated special characters",
			s:        "  This is   , . /   example text  - with spaces 124 ",
			expCount: 11,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			count := Counter(tc.s)
			assert.Equal(t, tc.expCount, count)
		})
	}
}
