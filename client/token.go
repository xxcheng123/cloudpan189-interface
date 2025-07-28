package client

import "time"

type AuthToken interface {
	IsExpired() bool
	ExpireTime() string
	ExpireDuration() time.Duration
	AccessToken() string
}

type authToken struct {
	accessToken string
	expires     int64
}

func NewAuthToken(accessToken string, expires int64) AuthToken {
	return &authToken{
		accessToken: accessToken,
		expires:     expires,
	}
}

func (t *authToken) IsExpired() bool {
	return t.expires < time.Now().Unix()*1e3
}

func (t *authToken) ExpireTime() string {
	return time.Unix(t.expires/1e3, 0).Format("2006-01-02 15:04:05")
}

// ExpireDuration 剩余过期时间
func (t *authToken) ExpireDuration() time.Duration {
	return time.Duration(t.expires-time.Now().Unix()*1e3) * time.Millisecond
}

func (t *authToken) AccessToken() string {
	return t.accessToken
}
