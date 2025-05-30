package helpers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cam-inc/viron-go/lib/constant"
)

// GenCookie cookie生成
func GenCookie(key string, value string, opts *http.Cookie) *http.Cookie {

	if opts == nil {
		opts = &http.Cookie{}
	}

	if opts.Path == "" {
		opts.Path = "/"
	}
	if opts.SameSite == http.SameSiteDefaultMode || opts.SameSite == 0 {
		opts.SameSite = http.SameSiteNoneMode
	}

	return &http.Cookie{
		Name:     key,
		Value:    value,
		Domain:   opts.Domain,
		HttpOnly: opts.HttpOnly,
		Expires:  opts.Expires,
		MaxAge:   opts.MaxAge,
		Path:     opts.Path,
		SameSite: opts.SameSite,
		Secure:   opts.Secure,
	}
}

// GenAuthorizationCookie 認証cookie生成
func GenAuthorizationCookie(token string, opts *http.Cookie) *http.Cookie {
	if opts.MaxAge == 0 && opts.Expires.IsZero() {
		opts.MaxAge = constant.DEFAULT_JWT_EXPIRATION_SEC
	}
	return GenCookie(constant.COOKIE_KEY_VIRON_AUTHORIZATION, token, opts)
}

// GenGoogleOAuth2StateCookie cookie生成
func GenGoogleOAuth2StateCookie(state string, opts *http.Cookie) *http.Cookie {
	if opts.MaxAge == 0 && opts.Expires.IsZero() {
		opts.MaxAge = constant.GOOGLE_OAUTH2_STATE_EXPIRATION_SEC
	}
	return GenCookie(constant.COOKIE_KEY_GOOGLE_OAUTH2_STATE, state, opts)
}

// GenGoogleOAuth2CodeVerifierCookie cookie生成
func GenGoogleOAuth2CodeVerifierCookie(codeVerifier string, opts *http.Cookie) *http.Cookie {
	if opts.MaxAge == 0 && opts.Expires.IsZero() {
		opts.MaxAge = constant.GOOGLE_OAUTH2_CODE_VERIFIER_EXPIRATION_SEC
	}
	return GenCookie(constant.COOKIE_KEY_GOOGLE_OAUTH2_CODE_VERIFIER, codeVerifier, opts)
}

// GenOidcStateCookie cookie生成
func GenOidcStateCookie(state string, opts *http.Cookie) *http.Cookie {
	if opts.MaxAge == 0 && opts.Expires.IsZero() {
		opts.MaxAge = constant.OIDC_STATE_EXPIRATION_SEC
	}
	return GenCookie(constant.COOKIE_KEY_OIDC_STATE, state, opts)
}

// GenOidcCodeVerifierCookie cookie生成
func GenOidcCodeVerifierCookie(codeVerifier string, opts *http.Cookie) *http.Cookie {
	if opts.MaxAge == 0 && opts.Expires.IsZero() {
		opts.MaxAge = constant.OIDC_CODE_VERIFIER_EXPIRATION_SEC
	}
	return GenCookie(constant.COOKIE_KEY_OIDC_CODE_VERIFIER, codeVerifier, opts)
}

// GetCookieToken cookieから認証token取得
func GetCookieToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie(constant.COOKIE_KEY_VIRON_AUTHORIZATION)
	if err != nil {
		return "", err
	}

	if cookie == nil {
		return "", fmt.Errorf("cookie notfound")
	}

	tokens := strings.Split(cookie.Value, " ")
	if len(tokens) != 2 || tokens[0] != constant.AUTH_SCHEME {
		return "", fmt.Errorf("token invalid")
	}

	return tokens[1], nil

}
