package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name  string
		input http.Header
		want  string
	}{
		{
			name:  "proper format",
			input: http.Header{"Authorization": []string{"ApiKey testKey"}},
			want:  "testKey",
		},
		{
			name:  "missing key",
			input: http.Header{},
			want:  "",
		},
		{
			name:  "improper format",
			input: http.Header{"Authorization": []string{"ApiKey    testKey"}},
			want:  "",
		},
		{
			name:  "improper casing",
			input: http.Header{"authorization": []string{"ApiKey testKey"}},
			want:  "testKey",
		},
		{
			name:  "multiple values",
			input: http.Header{"Authorization": []string{"ApiKey testKey", "ApiKey anotherKey"}},
			want:  "testKey",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetAPIKey(tt.input)

			if got != tt.want {
				t.Errorf("GetAPIKey() got %v, want %v", got, tt.want)
			}
		})
	}
}
