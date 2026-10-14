package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"
	expiresIn := time.Hour

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("making JWT: %v", err)
	}

	validatedUserID, err := ValidateJWT(tokenString, secret)
	if err != nil {
		t.Fatalf("validate JWT: %v", err)
	}

	if validatedUserID != userID {
		t.Errorf("got %+v, want %+v", validatedUserID, userID)
	}
}

func TestRejectExpiredToken(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"
	expiresIn := -time.Hour

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("making JWT: %v", err)
	}

	_, err = ValidateJWT(tokenString, secret)
	if err == nil {
		t.Fatalf("expected expired token to be rejected")
	}
}

func TestRejectWrongSecret(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"
	wrongSecret := "wrong-secret"
	expiresIn := time.Hour

	tokenString, err := MakeJWT(userID, secret, expiresIn)
	if err != nil {
		t.Fatalf("making JWT: %v", err)
	}

	_, err = ValidateJWT(tokenString, wrongSecret)
	if err == nil {
		t.Fatalf("expected wrong secret to be rejected")
	}
}
