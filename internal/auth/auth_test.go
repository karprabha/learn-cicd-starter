package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		wantErrText string
		want        string
	}{
		{
			name:        "no auth header",
			headers:     http.Header{},
			wantErrText: "no authorization header included",
			want:        "",
		},
		{
			name: "malformed auth header",
			headers: http.Header{
				"Authorization": []string{"Bearer 1234567890"},
			},
			wantErrText: "malformed authorization header",
			want:        "",
		},
		{
			name: "valid auth header",
			headers: http.Header{
				"Authorization": []string{"ApiKey 1234567890"},
			},
			wantErrText: "",
			want:        "1234567890",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := GetAPIKey(test.headers)

			if test.wantErrText != "" {
				if err == nil || !strings.Contains(err.Error(), test.wantErrText) {
					t.Errorf("GetAPIKey() error = %v, want error containing: %q", err, test.wantErrText)
				}
			} else if err != nil {
				t.Errorf("GetAPIKey() unexpected error = %v", err)
			}

			if got != test.want {
				t.Errorf("GetAPIKey() = %v, want = %v", got, test.want)
			}
		})
	}
}
