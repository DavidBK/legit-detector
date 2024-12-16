package github

import (
	"testing"
)

func TestIsValidateSignature(t *testing.T) {
	tests := []struct {
		name       string
		payload    []byte
		secret     string
		headerHash string
		want       bool
	}{
		{
			name:       "Valid signature",
			payload:    []byte("Hello, World!"),
			secret:     "It's a Secret to Everybody",
			headerHash: "sha256=757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
			want:       true,
		},
		{
			name:       "Invalid signature",
			payload:    []byte("Hello, World!"),
			secret:     "It's a Secret to Everybody",
			headerHash: "sha256=invalidhashvalue",
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidateSignature(tt.payload, tt.secret, tt.headerHash)
			if got != tt.want {
				t.Errorf("isValidateSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
