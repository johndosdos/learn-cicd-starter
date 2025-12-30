package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		input      http.Header
		wantAPIKey string
		wantErr    error
	}{
		{
			name:       "valid key",
			input:      http.Header{"Authorization": {"ApiKey 123"}},
			wantAPIKey: "123",
			wantErr:    nil,
		},
		{
			name:       "no header",
			input:      http.Header{},
			wantAPIKey: "",
			wantErr:    ErrNoAuthHeaderIncluded,
		},
		{
			name:       "malformed: with auth scheme but no value",
			input:      http.Header{"Authorization": {"ApiKey"}},
			wantAPIKey: "",
			wantErr:    errors.New("malformed authorization header"),
		},
		{
			name:       "malformed: wrong auth scheme",
			input:      http.Header{"Authorization": {"Bearer 123"}},
			wantAPIKey: "",
			wantErr:    errors.New("malformed authorization header"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotAPIKey, err := GetAPIKey(test.input)

			if test.wantErr != nil {
				if err == nil {
					t.Fatalf("error: want [%+v], got [%+v]", test.wantErr, err)
				}

				if test.wantErr.Error() != err.Error() {
					t.Fatalf("error: want [%+v], got [%+v]", test.wantErr, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}

			if test.wantAPIKey != gotAPIKey {
				t.Errorf("error: want [%+v], got [%+v]", test.wantAPIKey, gotAPIKey)
			}
		})
	}
}
