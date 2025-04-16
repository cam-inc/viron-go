package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWTConfig(t *testing.T) {
	jwt := JWT{
		Secret:        "test-secret",
		ExpirationSec: 3600,
		Issuer:        "test-issuer",
		Audience:      []string{"test-audience"},
	}

	assert.Equal(t, "test-secret", jwt.Secret)
	assert.Equal(t, 3600, jwt.ExpirationSec)
	assert.Equal(t, "test-issuer", jwt.Issuer)
	assert.Equal(t, []string{"test-audience"}, jwt.Audience)
}

func TestOIDCConfig(t *testing.T) {
	oidc := OIDC{
		ClientID:    "test-client-id",
		IssuerURL:   "https://example.com",
		RedirectURL: "https://example.com/callback",
	}

	assert.Equal(t, "test-client-id", oidc.ClientID)
	assert.Equal(t, "https://example.com", oidc.IssuerURL)
	assert.Equal(t, "https://example.com/callback", oidc.RedirectURL)
}
