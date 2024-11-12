package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expected      string
		errorContains string
	}{
		{
			name:          "No Auth Header",
			headers:       http.Header{},
			expected:      "",
			errorContains: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name: "Malformed Authorization Header",
			headers: http.Header{
				"Authorization": []string{"notanapikey"},
			},
			errorContains: "malformed authorization header",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetAPIKey(tc.headers)
			if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s': FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s': FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s': FAIL: expected error containing: %v, got none", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - '%s': FAIL: expected API key: %s, actual: %s", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}
