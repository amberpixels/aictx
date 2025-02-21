package aictx_test

import (
	"testing"

	"github.com/amberpixels/aictx/internal/aictx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateGitRepoName(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedURL    string
		expectedBranch string
		expectError    bool
	}{
		{
			name:        "empty input",
			input:       "",
			expectError: true,
		},
		{
			name:        "dot input",
			input:       ".",
			expectError: true,
		},
		{
			name:        "HTTPS without prefix",
			input:       "foo/bar",
			expectedURL: "https://github.com/foo/bar.git",
		},
		{
			name:           "HTTPS with branch",
			input:          "foo/bar@dev",
			expectedURL:    "https://github.com/foo/bar.git",
			expectedBranch: "dev",
		},
		{
			name:        "HTTPS with github.com prefix",
			input:       "github.com/foo/bar",
			expectedURL: "https://github.com/foo/bar.git",
		},
		{
			name:           "HTTPS with github.com prefix and branch",
			input:          "github.com/foo/bar@feature",
			expectedURL:    "https://github.com/foo/bar.git",
			expectedBranch: "feature",
		},
		{
			name:        "HTTPS with https:// prefix",
			input:       "https://github.com/foo/bar.git",
			expectedURL: "https://github.com/foo/bar.git",
		},
		{
			name:        "HTTPS with http:// prefix",
			input:       "http://github.com/foo/bar",
			expectedURL: "https://github.com/foo/bar.git",
		},
		{
			name:        "Invalid format without slash",
			input:       "foobar",
			expectError: true,
		},
		{
			name:        "SSH URL without branch",
			input:       "git@github.com:user/repo.git",
			expectedURL: "git@github.com:user/repo.git",
		},
		{
			name:           "SSH URL with branch",
			input:          "git@github.com:user/repo.git@feature-branch",
			expectedURL:    "git@github.com:user/repo.git",
			expectedBranch: "feature-branch",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			url, branch, err := aictx.ValidateGitRepoName(tc.input)
			if tc.expectError {
				require.Error(t, err, "expected error for input: %s", tc.input)
			} else {
				require.NoError(t, err, "unexpected error for input: %s", tc.input)
				assert.Equal(t, tc.expectedURL, url, "URL mismatch for input: %s", tc.input)
				assert.Equal(t, tc.expectedBranch, branch, "branch mismatch for input: %s", tc.input)
			}
		})
	}
}
