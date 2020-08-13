package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time constant should be 24 hours")
}

func TestGetNewAcccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.False(t, at.IsExpired(), "new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new access token should not have defined token id")
	assert.True(t, at.UserID == 0, "new access token should not have associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from Now should not be expired")
}
