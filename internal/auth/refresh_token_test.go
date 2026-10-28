package auth

import (
	"encoding/hex"
	"testing"
)

func TestMakeRefreshToken(t *testing.T) {
	token := MakeRefreshToken()

	decoded, err := hex.DecodeString(token)
	if err != nil {
		t.Fatalf("MakeRefreshToken() returned non-hex string %q: %v", token, err)
	}
	if len(decoded) != 32 {
		t.Errorf("MakeRefreshToken() decoded length = %d bytes, want 32", len(decoded))
	}
}

func TestMakeRefreshTokenIsUnique(t *testing.T) {
	const tokenCount = 1000
	seen := make(map[string]bool, tokenCount)

	for range tokenCount {
		token := MakeRefreshToken()
		if seen[token] {
			t.Fatalf("MakeRefreshToken() produced duplicate token %q", token)
		}
		seen[token] = true
	}
}
